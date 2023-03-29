package mDAO

import (
	"WangYiYunDemo/music/mDAO/mdDef"
	"log"
)

func ListSongs(req *mdDef.ListSongsReq) ([]*mdDef.SongBasic, error) {
	cond := make(map[string]interface{})
	if len(req.Name) > 0 {
		cond["Name"] = req.Name
	}
	if len(req.Tags) > 0 {
		cond["Tags"] = req.Tags
	}
	if len(req.Singer) > 0 {
		cond["Singer"] = req.Singer
	}
	resp := make([]*mdDef.SongBasic, 0)
	result := MysqlDB.Where(cond).Find(resp)
	if result.Error != nil {
		log.Println("mysql inner Error : " + result.Error.Error())
		return nil, result.Error
	}
	return resp, nil
}
