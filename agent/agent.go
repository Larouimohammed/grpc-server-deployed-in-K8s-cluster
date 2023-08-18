package agent

import (
	"grpc-exampl/infra"
	"grpc-exampl/plugins/redisplugin"
	"sync"
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

func (a *Agent) Start(wg *sync.WaitGroup) error {
	a.Logger.Info("agent start")
	for _, plugin := range a.AgentPlugins {
		wg.Add(1)
		go func(p AgentAPI){
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
		LogUtil:      infra.New(AGENT_NAME),
		AgentPlugins: defaultPlugins,
	}
}

var DefaultAgent = New()
