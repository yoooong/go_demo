package api

import (
	"douyu-new/common"
	"douyu-new/services"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func SendGiftHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	res := new(common.CommonHttpBaseRes)

	defer func() {
		if err := common.WriteHttpJsonRes(w, res); err != nil {
			//err log
			fmt.Println()
		}
	}()

	reqItem := &services.SendGiftReq{}
	decoder := json.NewDecoder(r.Body)
	decoder.UseNumber()
	err := decoder.Decode(&reqItem)
	if err != nil {
		//log
		fmt.Println("参数格式错误")
		return
	}
	if reqItem.FromUid == 0 || reqItem.ToUid == 0 || reqItem.Fee == 0 || reqItem.Amount == 0 || reqItem.GiftId == "" || reqItem.RoomId == 0 {
		res.Code = 3 //参数错误
		return
	}
	res.Code = services.SendGift(reqItem)
}

func RecvGiftRankingHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res := new(common.CommonHttpBaseRes)
	defer func() {
		if err := common.WriteHttpJsonRes(w, res); err != nil {
		}
	}()
	toUid, _ := common.Uint64(r.URL.Query().Get("uid"))
	if toUid == 0 {
		res.Code = 3 //参数错误
		return
	}
	res.Data, res.Code = services.RecvGiftRanking(toUid)
}

func RecvGiftRecordHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res := new(common.CommonHttpBaseRes)
	defer func() {
		if err := common.WriteHttpJsonRes(w, res); err != nil {
		}
	}()
	toUid, _ := common.Uint64(r.URL.Query().Get("uid"))
	limit, _ := common.Int(r.URL.Query().Get("limit"))
	offset, _ := common.Int(r.URL.Query().Get("offset"))
	if toUid == 0 {
		res.Code = 3 //参数错误
		return
	}
	if limit == 0 {
		limit = 30
	}
	res.Data, res.Code = services.RecvGiftRecord(toUid, limit, offset)
}
