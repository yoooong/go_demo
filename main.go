package main

import (
	"douyu-new/api"
	"douyu-new/dao"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()

	api.Init(router) //初始化路由

	dao.InitMongoDBDAO() //初始化mongo

	dao.RedisInit() //初始化Redis

	log.Fatal(http.ListenAndServe(":8080", router))
}
