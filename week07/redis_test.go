package week07

import (
	"context"
	"log"
	"testing"

	"github.com/go-redis/redis/v8"
)

// redis benchmark test
/*
	1. redis-benchmark -p 6379 -n 10000 -d 1024 -t get,set

				1k		 5k       20
		get: 12531.33  11210.76  12919.90
		set: 13679.89  9727.63   12642.22

*/

func Test_InfoMemory(t *testing.T) {
	cli := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	ctx := context.Background()
	cli.Set(ctx, "hello", "world", 0)
	cmd := cli.Get(ctx, "hello")
	log.Printf("cmd: %+v\n", cmd)
}
