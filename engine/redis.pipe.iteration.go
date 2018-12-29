package engine

import (
	"strconv"

	"github.com/go-redis/redis"
)

func getRedisIteration(tx redis.Pipeliner, pipeKey string) (i int, err error) {
	iStr, err := tx.HGet(pipeKey, "iteration").Result()
	if err != nil {
		return -1, err
	}
	return strconv.Atoi(iStr)
}

func setRedisIteration(tx redis.Pipeliner, pipeKey string, iter int) (err error) {
	_, err = tx.HSet(pipeKey, "iteration", iter).Result()
	return err
}