package common

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// 获取redis 配置
func GetRedisFromConsul(vip *viper.Viper) (red *redis.Client, err error) {
	red = redis.NewClient(&redis.Options{
		Addr:         vip.GetString("addr"),
		Password:     vip.GetString("password"),
		DB:           vip.GetInt("DB"),
		PoolSize:     vip.GetInt("poolSize"),
		MinIdleConns: vip.GetInt("minIdleConn"),
	})
	//集群
	clusterClients := redis.NewClusterClient(
		&redis.ClusterOptions{
			Addrs: []string{"192.168.100.131:6380", "192.168.100.131:6381", "192.168.100.131:6382"},
		})
	fmt.Println(clusterClients)

	return red, nil
}

// 设置用户登录信息
func SetUserToken(red *redis.Client, key string, val []byte, timeTTL time.Duration) {
	red.Set(context.Background(), key, val, timeTTL)
}

// 获取用户登录信息
func GetUserToken(red *redis.Client, key string) string {
	res, err := red.Get(context.Background(), key).Result()
	if err != nil {
		log.Print("GetUserToken  err  ", err)
	}
	return res
}

func GetConsulConfig(url string, fileKey string) (*viper.Viper, error) {
	conf := viper.New()
	err := conf.AddRemoteProvider("consul", url, fileKey)

	if err != nil {
		log.Println("viper conf err:", err)
	}
	conf.SetConfigType("json")

	err = conf.ReadRemoteConfig()
	if err != nil {
		log.Println("viper conf err:", err)
	}

	return conf, nil
}
func GetMysqlFromConsul(vip *viper.Viper) (db *gorm.DB, err error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		})

	str := vip.GetString("user") + ":" + vip.GetString("pwd") +
		"@tcp(" + vip.GetString("host") + ":" + vip.GetString("port") +
		")/" + vip.GetString("database") + "?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(str), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Println("mysql err:", err)
	}
	return db, nil
}
