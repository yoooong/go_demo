package cache

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	DATABASE           = "douyu"
	SendGiftCollection = "send_gift"
)

type SendGiftItem struct {
	FromUid    uint64    `bson:"from_uid" json:"from_uid"`
	ToUid      uint64    `bson:"to_uid" json:"to_uid"`
	RoomId     uint64    `bson:"room_id" json:"room_id"`
	GiftId     string    `bson:"gift_id" json:"gift_id"`
	Amount     uint64    `bson:"Amount" json:"amount"`
	Fee        uint64    `bson:"fee" json:"fee"`
	CreateTime time.Time `bson:"create_time" json:"create_time"`
	ModTime    time.Time `bson:"mod_time" json:"mod_time"`
}

type MongoDBDAO struct {
	Sess *mgo.Session
}

func (this *MongoDBDAO) SendGiftDBItem(info *SendGiftItem) (err error) {
	sess := this.Sess
	if sess == nil {
		fmt.Print("init mongodb err")
		return
	}
	sess = sess.Copy()
	defer sess.Close()
	c := sess.DB(DATABASE).C(SendGiftCollection)

	err = c.Insert(info)
	if err == mgo.ErrNotFound {
		err = nil
		return
	}
	return
}

func (this *MongoDBDAO) GetSendGiftRecordDBItem(uid uint64, limit, offset int) (ret []*SendGiftItem, err error) {
	sess := this.Sess
	if sess == nil {
		fmt.Print("init mongodb err")
		return
	}
	sess = sess.Copy()
	defer sess.Close()
	c := sess.DB(DATABASE).C(SendGiftCollection)
	cond := bson.M{}
	if uid > 0 {
		cond["to_uid"] = uid
	}
	query := c.Find(cond).Sort("-create_time")
	query = query.Skip(offset).Limit(limit)

	err = query.All(&ret)
	if err == mgo.ErrNotFound {
		err = nil
		return
	}
	return
}
