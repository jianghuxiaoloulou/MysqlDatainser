package main

import (
	"WowjoyProject/MysqlDataInser/global"
	"WowjoyProject/MysqlDataInser/internal/model"
	"time"
)

// @title 数据库历史数据插入
// @version 1.0.0.1
// @description 数据库历史数据插入
// @termsOfService
func main() {
	for i := 0; i < 268100001; i += 10000 {
		start := i
		end := start + 10000 - 1
		time.Sleep(time.Second * 1)
		global.Logger.Info("***开始插入数据")
		model.AutoUpdate(start, end)
	}
}
