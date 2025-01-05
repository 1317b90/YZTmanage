package API

import (
	"YZT/Data"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 数据库表
type Message struct {
	gorm.Model
	Type           string `gorm:"size:30"`          // 消息所属类型（组）
	Sender         string `gorm:"size:30;not null"` // 发送者
	NewMessage     string `gorm:"type:text"`        // 最新消息
	HistoryMessage string `gorm:"type:text"`        // 历史消息
	ReplyMessage   string `gorm:"type:text"`        // 回复的消息
}

// 接收消息结构体
type MessageInput struct {
	Type           string `form:"type" json:"type" binding:"required"`
	Sender         string `form:"sender" json:"sender" binding:"required"`
	NewMessage     string `form:"new_message" json:"new_message" binding:"required"`
	HistoryMessage string `form:"history_message" json:"history_message" default:""`
	ReplyMessage   string `form:"reply_message" json:"reply_message" default:""`
}

// 查询所有消息记录
func Get_messages(c *gin.Context) {
	var messages []Message
	result := Data.DB.Find(&messages)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "查询消息失败", "error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": messages})
}

// 新增消息记录接口
func Add_message(c *gin.Context) {
	var inputData MessageInput
	if err := c.ShouldBind(&inputData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	result := Data.DB.Create(&inputData)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "添加消息到数据库失败", "error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "消息添加成功"})
}

// 删除消息记录
func Del_message(c *gin.Context) {
	id := c.DefaultQuery("id", "")

	result := Data.DB.Delete(&Message{}, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "删除消息失败", "error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "消息删除成功"})
}

// 查询单个消息
func Get_message(c *gin.Context) {
	id := c.Param("id")
	var message Message
	result := Data.DB.First(&message, "id = ?", id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "查询消息失败", "error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": message})
}
