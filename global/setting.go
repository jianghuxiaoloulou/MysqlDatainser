package global

import (
	"WowjoyProject/MysqlDataInser/pkg/logger"
	"WowjoyProject/MysqlDataInser/pkg/setting"
)

var (
	GeneralSetting  *setting.GeneralSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
