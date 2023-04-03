package msDef

import "WangYiYunDemo/music/mDAO/mdDef"

type ListCommentMusicResp struct {
	Length           int                        `json:"length"`
	CommentMusicData []*mdDef.CommentMusicBasic `json:"commentMusic_data"`
}
