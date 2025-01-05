package main

import (
	"YZT/API"
	"YZT/Data"
	"YZT/Server"
	"os"
)

func main() {
	// 全局设置时区
	os.Setenv("TZ", "Asia/Shanghai")
	// 开启定时器
	Server.Timer.Start()
	// 初始化Gin
	Server.Gin_init()

	// 迁移数据表
	_ = Data.DB.AutoMigrate(&API.Task{})
	_ = Data.DB.AutoMigrate(&API.Log{})
	_ = Data.DB.AutoMigrate(&API.User{})
	_ = Data.DB.AutoMigrate(&API.Message{})

	// 打个招呼

	// --------- 任务类 --------- 任务类 --------- 任务类 --------- 任务类 --------- 任务类 --------- 任务类 --------- 任务类
	// 添加任务
	Server.R.POST("/task", API.Create_task)

	// 查询任务 从数据库中
	Server.R.GET("/task/db", API.Get_task_db)

	// 更新任务状态为ing
	Server.R.GET("/task/ing", API.Ing_task)

	// 任务完成后，修改数据库中任务状态
	Server.R.PUT("/task/done", API.Done_task)

	// 删除任务，从数据库
	Server.R.DELETE("/task", API.Del_task)

	// 统计任务数据
	Server.R.GET("/task/count", API.Count_task)

	// --------- 定时任务 --------- 定时任务 --------- 定时任务 --------- 定时任务 --------- 定时任务 --------- 定时任务 --------- 定时任务 --------- 定时任务

	// 新增定时任务
	Server.R.POST("/cron", API.Create_cron)

	// 查询定时任务
	Server.R.GET("/cron", API.Get_cron)

	// 删除定时任务
	Server.R.DELETE("/cron", API.Del_cron)

	// -------- Redis --------- Redis --------- Redis --------- Redis --------- Redis --------- Redis --------- Redis --------- Redis --------- Redis

	// 根据key获取value
	Server.R.GET("/redis/value", API.Get_redis_value)

	// 根据key设置value
	Server.R.POST("/redis/value", API.Set_redis_value)

	// 新增消息到Redis队列
	Server.R.POST("/redis/message", API.Add_redis_message)

	// 查询Redis中的任务
	Server.R.GET("/redis", API.Get_redis)

	// 删除Redis中的任务
	Server.R.DELETE("/redis", API.Del_redis)

	// --------- 用户表 --------- 用户表 --------- 用户表 --------- 用户表 --------- 用户表 --------- 用户表 --------- 用户表 --------- 用户表

	// 查询所有用户
	Server.R.GET("/user", API.Get_user)

	// 根据id查询用户
	Server.R.GET("/user/:phone", API.Get_user_by_phone)

	// 新增或更新用户
	Server.R.PUT("/user", API.Put_user)

	// 删除用户
	Server.R.DELETE("/user", API.Del_user)

	// --------- 消息记录 --------- 消息记录 --------- 消息记录 --------- 消息记录 --------- 消息记录 --------- 消息记录 --------- 消息记录 --------- 消息记录
	// 查询所有消息记录
	Server.R.GET("/message", API.Get_messages)

	// 新增消息记录
	Server.R.POST("/message", API.Add_message)

	// 根据id查询消息记录
	Server.R.GET("/message/:id", API.Get_message)

	// 删除消息记录
	Server.R.DELETE("/message", API.Del_message)

	// --------- 文件 --------- 文件 --------- 文件 --------- 文件 --------- 文件 --------- 文件 --------- 文件 --------- 文件
	// 设置文件目录
	Server.R.Static("/file", "./file")
	// 上传文件
	Server.R.POST("/file", API.Up_file)

	// ---------- 日志 ---------- 日志 ---------- 日志 ---------- 日志 ---------- 日志 ---------- 日志 ---------- 日志 ---------- 日志
	// 查询所有日志
	Server.R.GET("/log", API.Get_log)

	// 新增日志
	Server.R.POST("/log", API.Add_log)

	// 删除日志
	Server.R.DELETE("/log", API.Del_log)

	// 启动Gin
	_ = Server.R.Run()
}
