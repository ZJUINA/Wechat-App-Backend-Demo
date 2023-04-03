package mdDef

import "gorm.io/gorm"

type CommentMusicReq struct {
	Name   string `json:"name" gorm:"not null; index; column:name"`
	Limit  string `json:"limit" gorm:"index; column:limit"`
	Offset string `json:"offset" gorm:"index; column:offset"`
	Before string `json:"before" gorm:"index; column:before"`
	Tags   string `gorm:"column:tags" json:"tags"`
}

type CommentMusicBasic struct {
	gorm.Model
	Name    string `json:"name" gorm:"not null; index; column:name"`
	Content string `gorm:"index; column:content" json:"content" `
	Tags    string `gorm:"column:tags" json:"tags"`
}
