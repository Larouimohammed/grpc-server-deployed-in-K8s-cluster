package redisplugin

import (
	"context"
	"grpc-exampl/infra"
	number "grpc-exampl/proto"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	//redisClient.ConfigSet(bgCtx, "notify-keyspace-events", "KA")
	if _, err := redisClient.Ping(bgCtx).Result(); err != nil {
		log.Fatalf("Failed to connect redis client : %s \n", err)
	}
	return redisClient
}

var DefaultRedisClient = NewRedisClient()

// redsi plugin
type Redisplugin struct {
	name   string
	logger *logrus.Entry
	*redis.Client
}

func New() *Redisplugin {
	return &Redisplugin{
		name:   PLUGIN_NAME,
		Client: DefaultRedisClient,
		logger: infra.DefaultLogger.WithField("logger", PLUGIN_NAME),
	}
}

var DefaultRedisPlugin = New()

func prepareGrpcClinet() (number.ParityClient, error) {

	conn, err := grpc.Dial("localhost:3333", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := number.NewParityClient(conn)

	return client, nil
}

func (r *Redisplugin) Run() error {
	r.logger.Info("starting subscribtion on topic name : number")
	client, err := prepareGrpcClinet()
	if err != nil {
		r.logger.Error(err)
		return err
	}

	sub := r.Subscribe(context.Background(), "number")
	for message := range sub.Channel() {
		r.logger.Infof("got message with value = %s ", message.Payload)
		res, err := client.CheckParity(r.Context(), &number.NumberRequest{Number: message.Payload})
		if err != nil {
			r.logger.Error(err)
			continue
		} else {
			r.logger.Info(res.Response)
		}

	}
	return nil

}
