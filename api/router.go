package api

import "github.com/julienschmidt/httprouter"

func Init(rt *httprouter.Router) {
	rt.POST("/send_gift", SendGiftHandler)
	rt.GET("/recv_gifts_ranking", RecvGiftRankingHandler)
	rt.GET("/recv_gifts_record", RecvGiftRecordHandler)
}
