package grpcplugin

import (
	"context"
	"grpc-exampl/infra"
	number "grpc-exampl/proto"
	"strconv"
)

const (
	PLUGIN_NAME = "grpc-plugin"
)

type GrpcPlugin struct {
	name string
	*infra.LogUtil
	number.UnimplementedParityServer
}

func New() *GrpcPlugin {
	return &GrpcPlugin{
		name:    PLUGIN_NAME,
		LogUtil: infra.New(PLUGIN_NAME),
	}
}

var DefaultGrpcPlugin = New()

// impl parity service
func (p *GrpcPlugin) CheckParity(ctx context.Context, n *number.NumberRequest) (*number.NumberResponse, error) {
	// parse string to int
	p.Logger.Infof("got number request with value : %s", n.Number)
	num, err := strconv.Atoi(n.Number)
	if err != nil {
		p.Logger.Error(err)
		return nil, err
	}
	if num%2 == 0 {
		p.Logger.Infof("number : %d is pair", num)
		return &number.NumberResponse{
			Response: number.NumberResponse_PAIR,
		}, nil
	} else {
		p.Logger.Infof("number : %d is impair", num)
		return &number.NumberResponse{
			Response: number.NumberResponse_IMPAIR,
		}, nil
	}

}

// run the grpc server
func (p *GrpcPlugin) Run() error {
	return nil
}
