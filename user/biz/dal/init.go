package dal

import (
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/mysql"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
