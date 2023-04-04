package mService

import (
	"WangYiYunDemo/music/mDAO"
	"WangYiYunDemo/music/mDAO/mdDef"
	"WangYiYunDemo/music/mService/msDef"
	"log"
)

func WriteComment(req *mdDef.UploadCommentReq) (*msDef.WriteCommentResp, error) {
	//log.Println("not implemented")
	daoResp, err := mDAO.WriteComment(req)
	if err != nil {
		log.Println("Dao Error : " + err.Error())
		return nil, err
	}
	resp := new(msDef.WriteCommentResp)
	resp.IsSuccess = daoResp
	return resp, nil
}
