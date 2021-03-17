package services

import (
	"douyu-new/dao"
	"douyu-new/dao/cache"
	"fmt"
	"time"
)

func GetRedis() *cache.RedisRDB {
	return dao.RedisDao
}

func GetMongo() *cache.MongoDBDAO {
	return dao.MongoDao
}

func SendGift(reqItem *SendGiftReq) (respCode int) {
	//TODO 需要根据gift_id去查询gift表获取对应的价格，此处省略，直接根据传的价格
	formatParams := FormatParams4DB(reqItem)
	err := GetMongo().SendGiftDBItem(formatParams)
	if err != nil {
		respCode = 1
		fmt.Println("insert SendGift error:", err.Error())
		return
	}

	err = GetRedis().AddSendGiftCache(reqItem.FromUid, reqItem.ToUid, reqItem.Amount*reqItem.Fee)
	if err != nil {
		respCode = 1
		fmt.Println("AddSendGiftCache err:", err.Error())
		return
	}
	return
}

func FormatParams4DB(reqItem *SendGiftReq) (formatParams *cache.SendGiftItem) {
	now := time.Now()
	formatParams = &cache.SendGiftItem{}
	formatParams.FromUid = reqItem.FromUid
	formatParams.ToUid = reqItem.ToUid
	formatParams.Amount = reqItem.Amount
	formatParams.RoomId = reqItem.RoomId
	formatParams.Fee = reqItem.Fee
	formatParams.GiftId = reqItem.GiftId
	formatParams.CreateTime = now
	formatParams.ModTime = now
	return
}

func RecvGiftRanking(toUid uint64) (respData []*RecvGiftRankingResponse, respCode int) {
	data, err := GetRedis().GetSendGiftCache(toUid)
	if err != nil {
		respCode = 1
		fmt.Errorf("AddSendGiftCache err:%s", err.Error())
	}

	if data != nil {
		for _, item := range data {
			temp := new(RecvGiftRankingResponse)
			temp.FromUid = item.Uid
			temp.Score = item.Score
			respData = append(respData, temp)
		}
	}
	return
}

func RecvGiftRecord(toUid uint64, limit, offset int) (respData []*cache.SendGiftItem, respCode int) {
	data, err := GetMongo().GetSendGiftRecordDBItem(toUid, limit, offset)
	if err != nil {
		respCode = 1
		fmt.Errorf("AddSendGiftCache err:%s", err.Error())
	}

	if data != nil {
		respData = data
	}
	return
}
