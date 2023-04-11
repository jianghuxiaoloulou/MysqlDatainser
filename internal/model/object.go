package model

import (
	"WowjoyProject/MysqlDataInser/global"
)

// 自动获取需要删除的数据
func AutoInserData(start, end int) {
	global.Logger.Info("***开始在数据库中插入处理的数据***", "from: ", start, "  to: ", end)
	sql := `insert into file_remote (instance_key,dcm_update_time_retrieve,dcm_file_name_remote,img_file_name_remote)
	select ins.instance_key,im.image_dttm,im.dcm_file_name_remote,im.img_file_name_remote
	from instance ins
	left join image im on ins.instance_key = im.instance_key
	where ins.instance_key between ? and ?;`
	global.DBEngine.Exec(sql, start, end)
}

func AutoUpdateDCMData(start, end int) {
	global.Logger.Info("***开始在数据库中插入DCM处理的数据***", "from: ", start, "  to: ", end)
	sql := `update file_remote fr 
	left join image im on im.instance_key = fr.instance_key
	set fr.dcm_file_exist = 1
	where im.dcm_file_upload_status = 0 and im.instance_key between ? and ?;`
	global.DBEngine.Exec(sql, start, end)
}

func AutoUpdateJPGData(start, end int) {
	global.Logger.Info("***开始在数据库中插入JPG处理的数据***", "from: ", start, "  to: ", end)
	sql := `update file_remote fr 
	left join image im on im.instance_key = fr.instance_key
	set fr.img_file_exist = 1
	where im.img_file_upload_status = 0 and im.instance_key between ? and ?;`
	global.DBEngine.Exec(sql, start, end)
}

func AutoUpdate(start, end int) {
	global.Logger.Info("***开始在数据库中更新数据***", "from: ", start, "  to: ", end)
	sql := `update file_remote fr 
	left join instance ins on ins.instance_key = fr.instance_key
	set fr.dcm_file_exist = 1
	where fr.dcm_file_exist = 0 and ins.FileExist = 1 and ins.instance_key between ? and ?;`
	global.DBEngine.Exec(sql, start, end)
}
