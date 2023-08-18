package agent

import (
	"grpc-exampl/infra"
	"grpc-exampl/plugins/grpcplugin"
	"grpc-exampl/plugins/redisplugin"
	"sync"

	"github.com/sirupsen/logrus"
)

const (
	AGENT_NAME = "laroui-agent"
)

type Agent struct {
	name         string
	Logger       *logrus.Entry
	AgentPlugins []AgentAPI
}

var defaultPlugins = []AgentAPI{
	redisplugin.DefaultRedisPlugin, grpcplugin.DefaultGrpcPlugin,
}

func (a *Agent) Start(wg *sync.WaitGroup) error {
	a.Logger.Info("agent start")
	for _, plugin := range a.AgentPlugins {
		wg.Add(1)
		go func(p AgentAPI) {
			if err := p.Run(); err != nil {
				a.Logger.Error(err)
				wg.Done()

			}

		}(plugin)

	}
	return nil
}
func New() *Agent {
	return &Agent{
		name:         AGENT_NAME,
		Logger:       infra.DefaultLogger.WithField("logger", AGENT_NAME),
		AgentPlugins: defaultPlugins,
	}
}

var DefaultAgent = New()
