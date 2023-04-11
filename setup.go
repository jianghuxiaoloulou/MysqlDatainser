package main

import (
	"WowjoyProject/MysqlDataInser/global"
	"WowjoyProject/MysqlDataInser/internal/model"
	"WowjoyProject/MysqlDataInser/pkg/logger"
	"WowjoyProject/MysqlDataInser/pkg/setting"
	"log"

	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	readSetup()
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("General", &global.GeneralSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.GeneralSetting.LogSavePath + "/" + global.GeneralSetting.LogFileName + global.GeneralSetting.LogFileExt,
		MaxSize:   global.GeneralSetting.LogMaxSize,
		MaxAge:    global.GeneralSetting.LogMaxAge,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func readSetup() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
}
