package redis

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

// Exists return true if key is present
func Exists(rp *redis.Pool, key string) (bool, error) {
	red := rp.Get()

	defer func() {
		if err := red.Close(); err != nil {
			log.Fatalf("Exists(%v) Close: %v", key, err)
		}
	}()

	exists, err := redis.Bool(red.Do("EXISTS", key))
	if err != nil {
		return false, fmt.Errorf("Exists GET(%s): %v", key, err)
	}

	return exists, nil
}
