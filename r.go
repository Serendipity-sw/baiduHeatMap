package main

// 启动优先级：500
// 连接redis数据库
// 从redis中恢复数据
//
import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/guotie/config"
	"github.com/guotie/deferinit"
	"github.com/smtc/glog"
)

var (
	rpool           *redis.Pool
	cacheBssSeconds int = 86400
)

func init() {
	deferinit.AddInit(connectRedis, disconnectRedis, 500) // 优先级最高
	deferinit.AddInit(func() {
		cacheBssSeconds = config.GetIntDefault("cacheBssSeconds", 86400)
	}, nil, 50)
}

func connectRedis() {
	proto := config.GetStringDefault("redisProto", "tcp")
	addr := config.GetStringDefault("redisAddr", "127.0.0.1:6379")
	dbindex := config.GetIntDefault("redisDatabase", 0)
	openRedis(proto, addr, dbindex)

	glog.Info("open redis successfully.\n")
}

func disconnectRedis() {
	closeRedis()
}

// OpenRedis: 建立于redis服务器的连接池
//
// @proto: 协议, 通常为tcp, unix
// @addr: 地址, 如果proto为tcp, addr为IP地址; 如果proto为unix, addr为文件
func openRedis(proto, addr string, idx int) {
	rpool = &redis.Pool{
		MaxIdle:     100,
		IdleTimeout: 600 * time.Second,
		Dial: func() (redis.Conn, error) {
			do := redis.DialDatabase(idx)
			c, err := redis.Dial(proto, addr, do)
			if err != nil {
				panic(err)
				return nil, err
			}

			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	conn := rpool.Get()
	defer conn.Close()
	_, err := conn.Do("PING")
	if err != nil {
		panic(err.Error())
	}

	return
}

// CloseRedis: 关闭redis连接池
//
func closeRedis() {
	rpool.Close()
}

/**
设置redis缓存
创建人:邵炜
输入参数: uuid键值  ps存入字符串
输出参数: 错误对象
*/
func setRedisCachePs(uuid string, ps string) error {

	c := rpool.Get()
	defer c.Close()
	_, err := c.Do("SETEX", "axonVnav-"+uuid, cacheBssSeconds, ps)
	if err != nil {
		glog.Error("set redis cache failed: string %s error=%v\n",
			uuid, ps, err)
	}
	return err
}

/**
获取redis缓存 根据键值获取
创建人:邵炜
创建时间:2016年3月7日12:01:55
输入参数: key 键值
输出参数: redis获取值 错误对象
*/
func getRedisCachePs(key string) (*string, error) {

	c := rpool.Get()
	r, err := c.Do("GET", "axonVnav-"+key)
	c.Close()
	if err != nil {
		return nil, err
	}
	if r == nil {
		err = fmt.Errorf("getRedisCache: key %s is nil", "axonVnav-"+key)
		return nil, err
	}

	ps := r.(string)

	return &ps, nil
}
