package redis

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

func Connect() {
	c, err := redis.Dial("tcp", "111.231.228.108:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()
	// 永不过期
	set(c)
	// 带有过期时间
	setExpire(c)
	// 检查key可以是否过期
	existsKey(c)
	// 删除key
	delKey(c)
}
func set(c redis.Conn) {
	_, err := c.Do("SET", "username", "nick")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	username, err := redis.String(c.Do("GET", "username"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Got username %v \n", username)
	}
}
func setExpire(c redis.Conn) {
	_, err := c.Do("SET", "mykey", "superWang", "EX", "5")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}
	time.Sleep(8 * time.Second)
	username, err = redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}
}
func existsKey(c redis.Conn) {
	_, err := c.Do("SET", "mykey", "superWang")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	is_key_exit, err := redis.Bool(c.Do("EXISTS", "mykey1"))
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("exists or not: %v \n", is_key_exit)
	}
}

func delKey(c redis.Conn) {
	_, err := c.Do("SET", "mykey", "superWang")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}

	_, err = c.Do("DEL", "mykey")
	if err != nil {
		fmt.Println("redis delelte failed:", err)
	}

	username, err = redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}
}
