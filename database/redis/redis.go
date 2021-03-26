package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// Redis 数据库
func Redis() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "IP:端口",
		Password: "密码（没有则为空）",
		DB:       6, // 数据库编号
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err := rdb.Ping(ctx).Err()

	if err != nil {
		fmt.Println("ping redis failed", err)
		return
	}

	// // set/get示例(过期事件为两个时间段之间的差值)
	err = rdb.Set(ctx, "name", "卧槽什么鬼啊啊啊啊啊啊!", time.Second*60).Err()

	if err != nil {
		fmt.Println("set name failed", err)
		return
	}

	val, err := rdb.Get(ctx, "name").Result()

	if err == redis.Nil {
		fmt.Println("name doesn't !", err)
	} else if err != nil {
		fmt.Println("get name failed , err :", err)
		return
	} else {
		fmt.Println("name:", val)
	}

	// // zset示例
	// zsetkey := "language_rank"

	// languages := []*redis.Z{
	// 	&redis.Z{Score: 90.0, Member: "Golang"},
	// 	&redis.Z{Score: 98.0, Member: "Java"},
	// 	&redis.Z{Score: 95.0, Member: "Python"},
	// 	&redis.Z{Score: 97.0, Member: "Javascript"},
	// 	&redis.Z{Score: 99.0, Member: "C/C++"},
	// }

	// // ZADD
	// fmt.Println("新增zset(ZAdd)：")
	// num, err := rdb.ZAdd(ctx, zsetkey, languages...).Result()

	// if err != nil {
	// 	fmt.Println("zadd failed,err:", err)
	// 	return
	// }

	// fmt.Println("Zadd succ:", num)

	// // 把golang的分数加10
	// fmt.Println("Golang自增(ZIncBy)：")
	// newScore, err := rdb.ZIncrBy(ctx, zsetkey, 10.0, "Golang").Result()

	// if err != nil {
	// 	fmt.Println("zincrby failed,err:", err)
	// 	return
	// }

	// fmt.Println("Golang's score is :", newScore)

	// // 取分数最高的三个
	// fmt.Println("分数最高的三个(ZRevRangeWithScores)：")
	// ret, err := rdb.ZRevRangeWithScores(ctx, zsetkey, 0, 2).Result()

	// if err != nil {
	// 	fmt.Println("zrevrange failed:", err)
	// 	return
	// }

	// for _, z := range ret {
	// 	fmt.Println(z.Member, z.Score)
	// }

	// // 取95-100分的
	// fmt.Println("95-100的语言(ZRangeByScoreWithScores)：")
	// op := &redis.ZRangeBy{
	// 	Min: "95",
	// 	Max: "100",
	// }

	// ret, err = rdb.ZRangeByScoreWithScores(ctx, zsetkey, op).Result()

	// if err != nil {
	// 	fmt.Println("ZRangeByScoreWithScores failed,err:", err)
	// 	return
	// }

	// for _, z := range ret {
	// 	fmt.Println(z.Member, z.Score)
	// }

}
