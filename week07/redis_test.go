package week07

import (
	"context"
	"fmt"
	"log"
	"strings"
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

func GetUsedMemory(ctx context.Context, cli *redis.Client) string {
	infoMem := cli.Info(ctx, "memory")
	result, _ := infoMem.Result()
	for _, str := range strings.Split(result, "\n") {
		if strings.HasPrefix(str, "used_memory_human") {
			return str
		}
	}
	return ""
}

func GenValue(size int) string {
	b := strings.Builder{}
	for i := 0; i < size; i++ {
		b.WriteString("0")
	}
	return b.String()
}

func Test_InfoMemory(t *testing.T) {
	cli := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	ctx := context.Background()
	cli.Set(ctx, "hello", "world", 0)
	cmd := cli.Get(ctx, "hello")
	log.Printf("cmd: %+v\n", cmd)

	genData := func(size int, num int) {
		prefix := "test_info_memory"
		for i := 0; i < num; i++ {
			key := fmt.Sprintf("%s_%d", prefix, i)
			val := GenValue(size)
			cli.Set(ctx, key, val, 0)
		}
	}
	log.Printf("infoMem: %+v\n", GetUsedMemory(ctx, cli))
	size := 1024
	genData(size, 10000)
	log.Printf("size:%d infoMem: %+v\n", size, GetUsedMemory(ctx, cli))
	genData(size, 20000)
	log.Printf("size:%d infoMem: %+v\n", size, GetUsedMemory(ctx, cli))
	genData(size, 30000)
	log.Printf("size:%d infoMem: %+v\n", size, GetUsedMemory(ctx, cli))
}

/*
=== RUN   Test_InfoMemory
2022/07/02 20:29:15 cmd: get hello: world
2022/07/02 20:29:15 infoMem: used_memory_human:1000.29K
2022/07/02 20:29:26 size:1024 infoMem: used_memory_human:13.90M
2022/07/02 20:29:43 size:1024 infoMem: used_memory_human:26.85M
2022/07/02 20:30:07 size:1024 infoMem: used_memory_human:39.66M


val大小为1k时，平均key占1.3k字节
*/
