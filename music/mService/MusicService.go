package mService

import (
	"WangYiYunDemo/music/mDAO"
	"WangYiYunDemo/music/mDAO/mdDef"
	"WangYiYunDemo/music/mService/msDef"
	"log"
)

func ListSongs(req *mdDef.ListSongsReq) (*msDef.ListSongsResp, error) {
	log.Println("not implemented")
	daoResp, err := mDAO.ListSongs(req)
	if err != nil {
		log.Println("Dao Error : " + err.Error())
		return nil, err
	}
	resp := new(msDef.ListSongsResp)
	resp.SongsData = daoResp
	resp.Length = len(daoResp)
	return resp, nil
}
