package cache

import (
	"context"
	"douyu-new/common"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type RedisRDB struct {
	RDB *redis.Client
	Ctx context.Context
}

const (
	CharmRankingPrefix = "charm_ranking"
)

type UserScoreValue struct {
	Uid   uint64 `json:"uid"`
	Score uint64 `json:"score"`
}

func (r *RedisRDB) AddSendGiftCache(fromUid, toUid, total uint64) (err error) {
	key := fmt.Sprintf("%s:%d", CharmRankingPrefix, toUid)
	cmd := r.RDB.ZIncrBy(r.Ctx, key, float64(total), common.String(fromUid))
	err = cmd.Err()
	return
}

func (r *RedisRDB) GetSendGiftCache(toUid uint64) (scoreValues []*UserScoreValue, err error) {
	key := fmt.Sprintf("%s:%d", CharmRankingPrefix, toUid)
	cmd := r.RDB.ZRevRangeWithScores(r.Ctx, key, 0, 99)
	err = cmd.Err()
	if err != nil {
		//err log
		return
	}

	for _, item := range cmd.Val() {
		scoreValueItem := new(UserScoreValue)
		scoreValueItem.Uid, _ = common.Uint64(item.Member)
		scoreValueItem.Score = uint64(item.Score)
		scoreValues = append(scoreValues, scoreValueItem)
	}
	return
}
