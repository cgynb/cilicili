package redis

import (
	"github.com/go-redis/redis"
	"strconv"
)

func AddClick(videoId float64) (float64, bool) {
	val, err := DB.ZIncrBy("click", 1, strconv.Itoa(int(videoId))).Result()
	return val, err == nil
}

type tmp struct {
	VideoId       int
	VideoClickNum float64
}

func GetClickList() (v []tmp, ok bool) {
	val, err := DB.ZRevRangeByScoreWithScores("click", redis.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Offset: 0,
		Count:  10,
	}).Result()
	for i := range val {
		m, _ := strconv.Atoi(val[i].Member.(string))
		v = append(v, tmp{
			m,
			val[i].Score,
		})
	}
	return v, err == nil
}
