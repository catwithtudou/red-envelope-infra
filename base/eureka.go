package base

import (
	infra "github.com/catwithtudou/red-envelope-infra"
	"github.com/kataras/iris"
	"github.com/tietang/go-eureka-client/eureka"
	"time"
)

/**
 *@Author tudou
 *@Date 2020/8/30
 **/

type EurekaStarter struct {
	infra.BaseStarter
	client *eureka.Client
}

func (e *EurekaStarter) Init(ctx infra.StarterContext) {
	e.client = eureka.NewClient(ctx.Props())
	e.client.Start()
}

func (e *EurekaStarter) SetUp(ctx infra.StarterContext) {
	info := make(map[string]interface{})
	info["startTime"] = time.Now()
	info["appName"] = ctx.Props().GetDefault("app.name", "red-envelope")
	Iris().Get("/info", func(context iris.Context) {
		context.JSON(info)
	})
	Iris().Get("/health", func(context iris.Context) {
		health := eureka.Health{
			Details: make(map[string]interface{}),
		}
		health.Status = eureka.StatusUp

		context.JSON(health)
	})
}
