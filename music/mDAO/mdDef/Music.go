package mdDef

import "gorm.io/gorm"

type SongBasic struct {
	gorm.Model
	Name        string `gorm:"not null; index; column:name" json:"name"`
	Singer      string `gorm:"index; column:singer" json:"singer"`
	Tags        string `gorm:"column:tags" json:"tags"` //初代版本 ： 只支持一个标签
	PlayedTimes int    `gorm:"column:played_times" json:"played_times"`
}

type ListSongsReq struct {
	IDs     []uint   `json:"ids"`
	Names   []string `json:"names"`
	Singers []string `json:"singers"`
	Tags    []string `json:"tags"`
}

type PlayList struct {
	Length int          `json:"length"`
	Songs  []*SongBasic `json:"songs"`
}
