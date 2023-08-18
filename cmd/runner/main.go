package main

import (
	"context"
	"grpc-exampl/infra"
	"grpc-exampl/plugins/redisplugin"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

const NAME = "task-runner"

type Runner struct {
	name   string
	logger *logrus.Entry
	*redis.Client
}

func newRunner() *Runner {
	return &Runner{
		name:   NAME,
		logger: infra.DefaultLogger.WithField("logger", NAME),
		Client: redisplugin.DefaultRedisClient}
}
func gennumber() int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return r1.Intn(1000)
}

func main() {

	runner := newRunner()
	runner.logger.Infof("starting runner")
	for {
		r := gennumber()
		result := runner.Client.Publish(context.Background(), "number", r)
		if result.Err() != nil {
			runner.logger.Error(result.Err())
			continue
		}
		runner.logger.Infof("publishing number with value = %d", r)

		time.Sleep(10 * time.Second)

	}

}
