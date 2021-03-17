package dao

import (
	"douyu-new/dao/cache"
	"fmt"
	"gopkg.in/mgo.v2"
)

var Sess *mgo.Session
var MongoDao *cache.MongoDBDAO

func InitMongoDBDAO() {

	var err error
	session := &mgo.DialInfo{Addrs: []string{fmt.Sprint("127.0.0.1", ":", "27017")},
		Source: "admin", Username: "root", Password: "123456", PoolLimit: 10}

	if Sess, err = mgo.DialWithInfo(session); err != nil {
		return
	}
	fmt.Printf("mongo error:%v", err)
	MongoDao = NewMongoDao(Sess)
	Sess.SetMode(mgo.SecondaryPreferred, true)

}

func NewMongoDao(sess *mgo.Session) (mongoDao *cache.MongoDBDAO) {
	mongoDao = &cache.MongoDBDAO{
		Sess: sess,
	}
	return
}
