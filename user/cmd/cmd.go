package usercmd

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/spf13/cobra"
	"github.com/wheelergeo/g-otter-gen/user/userservice"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/mysql"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/redis"
	"github.com/wheelergeo/g-otter-micro-user/biz/handler"
	"github.com/wheelergeo/g-otter-micro-user/conf"
	"github.com/wheelergeo/g-otter-micro-user/pkg/auth"
	"github.com/wheelergeo/g-otter-micro-user/pkg/logger"
	"github.com/wheelergeo/g-otter-micro-user/pkg/token"
)

var once sync.Once
var c *cobra.Command

func Command() *cobra.Command {
	once.Do(func() {
		c = &cobra.Command{
			Use:   "user",
			Short: "start user micro server",
			Run: func(cmd *cobra.Command, args []string) {
				dal.Init()
				auth.NewServer(
					conf.GetConf().Casbin.ModelName,
					mysql.DB,
					conf.GetConf().Casbin.PolicyTable,
					conf.GetConf().Casbin.PolicyRedis,
				)
				token.Init(
					redis.RedisClient,
					conf.GetConf().Paseto.CacheKey,
					nil,
				)
				logger.InitKlogWithLogrus(logger.Config{
					Mode:   logger.StdOut,
					Format: logger.Text,
					Level:  hlog.Level(conf.LogLevel()),
					FileCfg: logger.FileConfig{
						FileName:      conf.GetConf().Kitex.LogFileName,
						MaxSize:       conf.GetConf().Kitex.LogMaxSize,
						MaxBackups:    conf.GetConf().Kitex.LogMaxBackups,
						MaxAge:        conf.GetConf().Kitex.LogMaxAge,
						FlushInterval: time.Minute,
					},
				})

				k := microService()

				klog.Infof("[%s] Server start", conf.GetEnv())
				err := k.Run()
				if err != nil {
					klog.Error(err.Error())
				}
			},
		}
	})

	return c
}

func microService() (k server.Server) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}

	// rpc info
	rpcInfo := rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}

	opts := []server.Option{
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcInfo),
		server.WithMetaHandler(transmeta.ServerTTHeaderHandler),
	}

	// otel
	if conf.GetConf().Kitex.EnableOtel &&
		conf.GetConf().Otel.Endpoint != "" {
		p := provider.NewOpenTelemetryProvider(
			provider.WithServiceName(conf.GetConf().Kitex.Service),
			provider.WithExportEndpoint(conf.GetConf().Otel.Endpoint),
			provider.WithInsecure(),
		)

		opts = append(opts, server.WithSuite(tracing.NewServerSuite()))
		server.RegisterShutdownHook(func() {
			p.Shutdown(context.Background())
		})
	}

	k = userservice.NewServer(new(handler.UserServiceImpl), opts...)
	return
}
