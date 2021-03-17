package services

import "time"

type SendGiftReq struct {
	FromUid    uint64    `json:"from_uid"`
	ToUid      uint64    `json:"to_uid"`
	Amount     uint64    `json:"amount"`
	Fee        uint64    `json:"fee"`
	RoomId     uint64    `json:"room_id"`
	GiftId     string    `json:"gift_id"`
	CreateTime time.Time `json:"create_time"`
	ModTime    time.Time `json:"mod_time"`
}

type RecvGiftRankingResponse struct {
	FromUid uint64 `json:"from_uid"`
	Score   uint64 `json:"score"`
}
