package API

import (
	"YZT/Data"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 根据key获取value
func Get_redis_value(c *gin.Context) {
	key := c.Query("key")

	value, err := Data.Rdb.Get(Data.CTX, key).Result()
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": value})
}

// 根据key设置value
func Set_redis_value(c *gin.Context) {
	type inputS struct {
		Key   string `form:"key" json:"key" binding:"required"`
		Value string `form:"value" json:"value" binding:"required"`
	}

	var inputData inputS

	if err := c.ShouldBind(&inputData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := Data.Rdb.Set(Data.CTX, inputData.Key, inputData.Value, 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "设置Redis值失败", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "设置Redis值成功"})
}

// 获取所有redis数据
func Get_redis(c *gin.Context) {
	// 获取所有Redis键
	keys, err := Data.Rdb.Keys(Data.CTX, "*").Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "获取Redis键失败", "error": err.Error()})
		return
	}

	// 获取每个键对应的值
	var redisData = make(map[string]interface{})
	for _, key := range keys {
		value, err := Data.Rdb.Get(Data.CTX, key).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "获取Redis值失败", "error": err.Error()})
			return
		}
		redisData[key] = value
	}

	c.JSON(http.StatusOK, gin.H{"message": "获取Redis数据成功", "data": redisData})
}

// 删除指定redis数据
func Del_redis(c *gin.Context) {
	key := c.Query("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "缺少键参数"})
		return
	}

	err := Data.Rdb.Del(Data.CTX, key).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "删除Redis数据失败", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除Redis数据成功"})
}

// 新增消息到Redis队列
func Add_redis_message(c *gin.Context) {
	type inputS struct {
		Key       string `form:"key" json:"key" binding:"required"`
		Recipient string `form:"recipient" json:"recipient" binding:"required"` // 接收人
		Type      string `form:"type" json:"type" binding:"required"`           // 消息类型
		Content   string `form:"content" json:"content" binding:"required"`     // 消息内容
	}

	var inputData inputS
	if err := c.ShouldBind(&inputData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	inputDataStr, err := json.Marshal(inputData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "将输入数据转化为字符串失败", "error": err.Error()})
		return
	}

	Data.Rdb.LPush(Data.CTX, inputData.Key, inputDataStr)

	c.JSON(http.StatusOK, gin.H{"message": "消息添加成功"})
}
