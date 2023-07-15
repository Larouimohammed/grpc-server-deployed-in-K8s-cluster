package redisplugin

import (
	"context"
	"grpc-exampl/infra"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

const (
	PLUGIN_NAME = "redis-plugin"
)

var (
	bgCtx       = context.Background()
	defaulHost  = "localhost"
	defaultPort = "6379"
)

func NewRedisClient() *redis.Client {

	if os.Getenv("REDIS_HOST") != "" {
		defaulHost = os.Getenv("REDIS_HOST")
	}
	if string(os.Getenv("REDIS_PORT")) != "" {
		defaultPort = string(os.Getenv("REDIS_PORT"))
	}
	redisClient := redis.NewClient(
		&redis.Options{
			Addr:       defaulHost + ":" + defaultPort,
			DB:         0,
			MaxRetries: 10,
		},
	)
	redisClient.ConfigSet(bgCtx, "notify-keyspace-events", "KA")
	if _, err := redisClient.Ping(bgCtx).Result(); err != nil {
		log.Fatalf("Failed to connect redis client : %s \n", err)
	}
	return redisClient
}

var DefaultRedisClient = NewRedisClient()

// / redsi plugin
type Redisplugin struct {
	name string
	*infra.LogUtil
	*redis.Client
}

func New() *Redisplugin {
	return &Redisplugin{
		name:    PLUGIN_NAME,
		Client:  DefaultRedisClient,
		LogUtil: infra.New(PLUGIN_NAME),
	}
}

var DefaultRedisPlugin = New()

func (r *Redisplugin) Run() error {
	r.Logger.Info("starting subscribtion on topic name : number")
	sub := r.Subscribe(context.Background(), "number")
	for message := range sub.Channel() {
		r.Logger.Infof("got message with value = %s ", message.Payload)
	}
	return nil
}
