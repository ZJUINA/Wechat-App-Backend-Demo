package msDef

import "WangYiYunDemo/music/mDAO/mdDef"

type ListSongsResp struct {
	Length    int                `json:"length"`
	SongsData []*mdDef.SongBasic `json:"songs_data"`
}
