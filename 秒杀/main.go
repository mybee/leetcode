package main

import (
	"github.com/garyburd/redigo/redis"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const (
	REDIS_HOST = "127.0.0.1:6379"
	REDIS_PASSWORD = ""
)
var mutex sync.RWMutex
var pool *redis.Pool
var logChan = make(chan []byte,1024)

func main()  {
	// 日志线程
	go output()
	// http 客户端线程
	go qiang()
	// redis 连接池
	pool = &redis.Pool{
		MaxIdle: 80 ,
		MaxActive:100,
		IdleTimeout:time.Second * 180 ,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp",REDIS_HOST,redis.DialPassword(REDIS_PASSWORD))
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_ , err := c.Do("PING")
			return err
		},
	}
	c := pool.Get()
	c.Do("SET","total",100)
	c.Close()

	// 主线程
	// 太过依赖redis , 如果redis挂了，可考虑根据redis性能调整连接数量. 保证redis连接可用性。
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		c := pool.Get()
		defer c.Close()
		c.Do("WATCH","total")
		resp,err := c.Do("GET","total")
		if err != nil {
			writer.Write([]byte(err.Error()))
			return
		}
		total , err := redis.Int(resp,err)
		if err != nil {
			writer.Write([]byte(err.Error()))
			return
		}
		if total == 0 {
			writer.Write([]byte("抢光了"))
			return
		}
		c.Send("MULTI")
		c.Send("INCRBY","total",-1) // 这里是原子性的
		resp , err = c.Do("EXEC")
		if resp == nil {
			writer.Write([]byte("抢购失败"))
			return
		}
		res , err := redis.Values(resp,err)
		if err != nil {
			// 错误
			log.Print("err",err)
			return
		}
		// 写入成功应该是如下这种格式，第三个返回值才是total
		// 1) (integer) 78
		for _, v := range res {
			switch v.(type) {
			case int:
				total = v.(int)
			}
		}
		// 将用户信息写入redis 队列. 总数应该是和商品数量相等.
		// 也可以是 订阅/发布模式，在另一台服务器或另一个线程里，订阅一个话题，将商品信息和用户id写入数据库，
		// 数据库没有压力，压力 httpserver > redis > db
		// 连接过大直接拒绝连接 获取启用多台服务器负载均衡。
		c.Do("LPUSH","users",rand.Int())
		io.WriteString(writer,strconv.Itoa(total))
	})
	http.ListenAndServe(":8090",nil)
}

func qiang()  {
	for j := 0 ; j < 100 ; j ++ {
		time.Sleep(4 * time.Second)
		for i := 0 ; i < 1000 ; i ++{
			go req()
		}
	}
}

func req()  {
	defer func() {
		if err := recover(); err != nil {
			log.Print("panic:",err)
		}
	}()
	resp, err := http.Get("http://localhost:8090")
	if err != nil {
		log.Print("err",err)
		return
	}
	defer resp.Body.Close()
	all , err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("err",err)
	}
	logChan <- all
}

func output()  {
	for{
		select {
		case s := <- logChan :
			log.Print(string(s))
		}
	}
}