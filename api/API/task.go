package API

import (
	"YZT/Data"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 任务表
type Task struct {
	gorm.Model
	Type   string          `gorm:"size:255"` // 任务的类型，可以是make_invoice（开发票）
	Input  json.RawMessage // 任务的输入
	Output json.RawMessage // 任务的输出
	// 可以是waiting（等待执行） running（执行中）success（执行成功） error（执行失败）
	State   string `gorm:"type:varchar(10);default:'waiting'"`
	Message string `gorm:"type:text"`
	UserID  string `gorm:"type:varchar(30);default:'0'"`
}

// 创建任务
func Create_task(c *gin.Context) {
	type inputS struct {
		Type  string          `form:"type" json:"type" binding:"required"`
		Input json.RawMessage `form:"input" json:"input" binding:"required"`
	}

	var inputData inputS

	if err := c.ShouldBind(&inputData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// 将input转换为Map
	var inputMap map[string]interface{}
	if err := json.Unmarshal(inputData.Input, &inputMap); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "输入数据非JSON，请检查后重试"})
		return
	}

	// 创建任务到数据库
	task := &Task{
		Type:  inputData.Type,
		Input: inputData.Input,
	}

	result := Data.DB.Create(task)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "创建任务到数据库失败", "error": result.Error.Error()})
		return
	}

	// 将创建完毕的task的id添加到inputData.Input中，键名为task_id
	inputMap["task_id"] = task.ID

	inputMapStr, err := json.Marshal(inputMap)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "将inputMap转为字符串失败", "error": err.Error()})
		return
	}
	// 创建任务到Redis
	err = Data.Rdb.LPush(Data.CTX, inputData.Type, inputMapStr).Err()
	if err != nil {
		fmt.Println("创建任务到Redis失败", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "创建任务到Redis失败", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "任务创建成功", "data": task})
}

// 查询任务，从数据库
func Get_task_db(c *gin.Context) {
	taskID := c.Query("id")
	// 如果id为空，从内存，返回全部任务
	if taskID == "" {
		state := c.Query("state")
		var tasks []Task
		var result *gorm.DB
		if state == "" {
			result = Data.DB.Find(&tasks)
		} else {
			result = Data.DB.Where("state = ?", state).Find(&tasks)
		}
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "获取任务列表失败",
				"error":   result.Error,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"data":    tasks,
		})
	} else {
		var task Task

		result := Data.DB.First(&task, taskID)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "任务不存在",
				"error":   result.Error,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"data":    task,
		})
	}
}

// 任务开始执行，修改数据状态
func Ing_task(c *gin.Context) {
	taskID := c.Query("id")
	// 将taskID转换为uint
	id, err := strconv.ParseUint(taskID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "无效的任务ID",
			"error":   err.Error(),
		})
		return
	}

	// 在数据库中更新任务状态
	result := Data.DB.Model(&Task{}).Where("id = ?", id).Update("state", "running")
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "更新任务状态失败",
			"error":   result.Error.Error(),
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "未找到指定ID的任务",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "任务状态已更新为进行中",
	})
}

// 完成任务时，修改数据库状态
func Done_task(c *gin.Context) {
	type inputType struct {
		TaskID  uint64 `form:"id" json:"id" binding:"required"`
		Succeed bool   `form:"succeed" json:"succeed" default:"false"`
		Msg     string `form:"msg" json:"msg" `
		Data    string `form:"data" json:"data" `
	}
	var inputData inputType

	if err := c.ShouldBind(&inputData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// 更新任务的Output字段和状态
	state := "error"
	if inputData.Succeed {
		state = "success"
	}

	result := Data.DB.Model(&Task{}).Where("id = ?", inputData.TaskID).Updates(map[string]interface{}{
		"output":  json.RawMessage(inputData.Data),
		"state":   state,
		"message": inputData.Msg,
	})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "更新数据库状态失败",
			"error":   result.Error.Error(),
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "未找到指定ID的任务",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "任务输出已成功更新",
	})
}

// 删除任务，从数据库
func Del_task(c *gin.Context) {
	taskID := c.Query("id")
	// 将taskID转换为uint
	id, err := strconv.ParseUint(taskID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "无效的任务ID",
			"error":   err.Error(),
		})
		return
	}

	// 从数据库中删除任务
	result := Data.DB.Unscoped().Delete(&Task{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "删除失败，" + result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

// 统计任务数据
func Count_task(c *gin.Context) {
	var totalCount int64
	var waitingCount int64
	var runningCount int64
	var successCount int64
	var errorCount int64

	// 统计总任务数
	Data.DB.Model(&Task{}).Count(&totalCount)

	// 统计不同状态的任务数
	Data.DB.Model(&Task{}).Where("state = ?", "waiting").Count(&waitingCount)
	Data.DB.Model(&Task{}).Where("state = ?", "running").Count(&runningCount)
	Data.DB.Model(&Task{}).Where("state = ?", "success").Count(&successCount)
	Data.DB.Model(&Task{}).Where("state = ?", "error").Count(&errorCount)

	c.JSON(http.StatusOK, gin.H{
		"message": "任务统计数据",
		"data": gin.H{
			"all":     totalCount,
			"waiting": waitingCount,
			"running": runningCount,
			"success": successCount,
			"error":   errorCount,
		},
	})
}

// 定时清理任务记录 本地执行
// 为什么不做成接口：做成接口谁来读取？
func Clean_task_func() (string, error) {
	// 删除超过24小时且状态为waiting的任务记录
	deleteResult := Data.DB.Where("state = ? AND created_at < ?", "waiting", time.Now().Add(-24*time.Hour)).Delete(&Task{})
	deletedCount := deleteResult.RowsAffected

	// 将超过24小时且状态为running的任务记录更新为error状态
	updateResult := Data.DB.Model(&Task{}).
		Where("state = ? AND created_at < ?", "running", time.Now().Add(-24*time.Hour)).
		Updates(map[string]interface{}{
			"state":  "error",
			"output": json.RawMessage(`{"Code":500,"Msg":"任务执行超时","Data":null}`),
		})
	updatedCount := updateResult.RowsAffected

	// 记录日志
	logMessage := fmt.Sprintf("清理任务记录：删除了 %d 条等待状态的过期任务，更新了 %d 条执行中的超时任务", deletedCount, updatedCount)
	return logMessage, nil
}
