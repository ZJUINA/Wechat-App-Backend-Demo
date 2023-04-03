package mService

import (
	"WangYiYunDemo/music/mDAO"
	"WangYiYunDemo/music/mDAO/mdDef"
	"WangYiYunDemo/music/mService/msDef"
	"log"
)

func ListCommentMusic(req *mdDef.CommentMusicReq) (*msDef.ListCommentMusicResp, error) {
	//log.Println("not implemented")
	daoResp, err := mDAO.ListCommentMusic(req)
	if err != nil {
		log.Println("Dao Error : " + err.Error())
		return nil, err
	}
	resp := new(msDef.ListCommentMusicResp)
	resp.CommentMusicData = daoResp
	resp.Length = len(daoResp)
	return resp, nil
}
