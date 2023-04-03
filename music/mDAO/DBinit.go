package mDAO

import (
	"WangYiYunDemo/music/mDAO/mdDef"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var MysqlDB *gorm.DB

func DBInit() {
	var err error
	MysqlDB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:423319@tcp(127.0.0.1:3306)/wyyDemo?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         256,                                                                             // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                            // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                            // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                            // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                           // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		log.Println("initialize gorm DB error : " + err.Error())
	} else {
		migrator := MysqlDB.Migrator()
		if !migrator.HasTable(&mdDef.CommentMusicBasic{}) {
			migrator.CreateTable(&mdDef.CommentMusicBasic{})
		}
		newItem := mdDef.CommentMusicBasic{Name: "一万次悲伤", Content: "逃跑计划-世界", Tags: "test"}
		MysqlDB.Create(&newItem)
		newItem = mdDef.CommentMusicBasic{Name: "一万次悲伤", Content: "每一颗眼泪是一万道光", Tags: "test"}
		MysqlDB.Create(&newItem)
		newItem = mdDef.CommentMusicBasic{Name: "一万次悲伤", Content: "我一直在最后的地方等你", Tags: "test"}
		MysqlDB.Create(&newItem)
		migrator.RenameTable(&mdDef.CommentMusicBasic{}, "music_comments")
	}

}
