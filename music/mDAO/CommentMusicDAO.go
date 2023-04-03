package mDAO

import (
	"WangYiYunDemo/music/mDAO/mdDef"
	"log"
	"strconv"
)

func ListCommentMusic(req *mdDef.CommentMusicReq) ([]*mdDef.CommentMusicBasic, error) {
	cond := make(map[string]interface{})
	if len(req.Name) > 0 {
		cond["Name"] = req.Name
	}
	if len(req.Tags) > 0 {
		cond["Tags"] = req.Tags
	}
	if req.Limit == "" {
		cond["Limit"] = 20
	} else {
		limit, _ := strconv.Atoi(req.Limit)
		cond["Limit"] = limit
	}
	if req.Before == "" {
		cond["BeforeTime"] = "0"
	} else {
		cond["BeforeTime"] = req.Before
	}
	if req.Offset == "" {
		cond["Offset"] = 0
	} else {
		offset, _ := strconv.Atoi(req.Offset)
		cond["Offset"] = offset
	}
	resp := make([]*mdDef.CommentMusicBasic, 0)
	result := MysqlDB.Table("music_comments").Where("Name = ? And Tags = ?", cond["Name"], cond["Tags"]).Offset(cond["Offset"].(int)).Limit(cond["Limit"].(int)).Scan(&resp)
	if nil != result.Error {
		log.Println("mysql inner Error : " + result.Error.Error())
		return nil, result.Error
	}
	return resp, nil
}
