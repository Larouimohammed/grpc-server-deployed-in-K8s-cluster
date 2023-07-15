package agent

import (
	"grpc-exampl/infra"
	"grpc-exampl/plugins/redisplugin"
)

const (
	AGENT_NAME = "grpc-agent"
)

type Agent struct {
	name string
	*infra.LogUtil
	AgentPlugins []AgentAPI
}

var defaultPlugins = []AgentAPI{
	redisplugin.DefaultRedisPlugin,
}

func (a *Agent) Start() error {
	a.Logger.Info("agent start")
	for _, plugin := range a.AgentPlugins {
		if err := plugin.Run(); err != nil {
			a.Logger.Error(err)
			return err
		}
	}
	return nil
}
func New() *Agent {
	return &Agent{
		name:         AGENT_NAME,
		LogUtil:      infra.New(AGENT_NAME),
		AgentPlugins: defaultPlugins,
	}
}

var DefaultAgent = New()
