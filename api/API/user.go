package API

import (
	"YZT/Data"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	Phone       string    `gorm:"size:11;primaryKey"`
	Enable      bool      `gorm:"default:true"`
	Uscid       string    `gorm:"size:30;not null"`
	DsjUsername string    `gorm:"size:30;not null"`
	DsjPassword string    `gorm:"size:30;not null"`
	CompanyName string    `gorm:"size:30;not null"`
	BankName    string    `gorm:"size:30;not null"`
	BankID      string    `gorm:"size:30;not null"`
	PuppetID    string    `gorm:"size:30"`
	CreateTime  time.Time `gorm:"autoCreateTime"`
	UpdateTime  time.Time `gorm:"autoUpdateTime"`
}

// 查询用户
func Get_user(c *gin.Context) {
	var users []User
	query := Data.DB

	queryParams := map[string]string{
		"Phone":       "phone",
		"CompanyName": "company_name",
		"Uscid":       "uscid",
		"DsjUsername": "dsj_username",
		"DsjPassword": "dsj_password",
		"BankName":    "bank_name",
		"BankID":      "bank_id",
		"Enable":      "enable",
	}

	for param, column := range queryParams {
		if value := c.Query(param); value != "" {
			if param == "Enable" {
				query = query.Where(column+" = ?", value == "true")
			} else {
				query = query.Where(column+" = ?", value)
			}
		}
	}

	result := query.Find(&users)

	if result.Error != nil {
		c.JSON(500, gin.H{"message": "查询用户失败", "error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{"data": users})
}

// 根据phone查询指定用户
func Get_user_by_phone(c *gin.Context) {
	phone := c.Param("phone")
	var user User

	result := Data.DB.First(&user, "phone = ?", phone)

	if result.Error != nil {
		c.JSON(500, gin.H{"message": "查询用户失败", "error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{"data": user})
}

// 新增或更新用户
func Put_user(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"message": "请求参数错误", "error": err.Error()})
		return
	}

	result := Data.DB.Save(&user)
	if result.Error != nil {
		c.JSON(500, gin.H{"message": "保存用户失败", "error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "用户保存成功"})
}

// 删除用户
func Del_user(c *gin.Context) {
	phone := c.DefaultQuery("phone", "")

	fmt.Println(phone)

	result := Data.DB.Delete(&User{}, "phone = ?", phone)

	if result.Error != nil {
		c.JSON(500, gin.H{"message": "删除用户失败", "error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "用户删除成功"})
}
