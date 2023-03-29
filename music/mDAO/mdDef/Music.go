package mdDef

import "gorm.io/gorm"

type SongBasic struct {
	gorm.Model
	Name   string `gorm:"not null; index; column:name" json:"name"`
	Singer string `gorm:"singer; index" json:"singer"`
	Tags   string `gorm:"tags" json:"tags"` //初代版本 ： 只支持一个标签
}

type ListSongsReq struct {
	Name   string `gorm:"not null; column:name" json:"name"`
	Singer string `gorm:"singer" json:"singer"`
	Tags   string `gorm:"tags" json:"tags"`
}
