package data

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8" // 导入 go-redis 库
	"github.com/google/wire"
	"verify-code/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewCustomerData, NewStudentRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	Rdb *redis.Client
}

// NewData .初始化
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	data := &Data{}
	//初始化rdb
	//连接Redis的配置
	redisURL := fmt.Sprintf("redis://%s/1?dial_timeout=%d", c.Redis.Addr, 1)
	options, err := redis.ParseURL(redisURL)
	if err != nil {
		data.Rdb = nil
	}
	data.Rdb = redis.NewClient(options)
	cleanup := func() {
		//清理redis连接
		_ = data.Rdb.Close()
		log.NewHelper(logger).Info("closing the data resources")
	}
	return data, cleanup, nil
}
