package grpcplugin

import (
	"context"
	"grpc-exampl/infra"
	number "grpc-exampl/proto"
	"net"
	"strconv"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	PLUGIN_NAME = "grpc-plugin"
)

type GrpcPlugin struct {
	name   string
	logger *logrus.Entry
	number.UnimplementedParityServer
}

func New() *GrpcPlugin {
	return &GrpcPlugin{
		name:   PLUGIN_NAME,
		logger: infra.DefaultLogger.WithField("logger", PLUGIN_NAME),
	}
}

var DefaultGrpcPlugin = New()

// impl parity service
func (p *GrpcPlugin) CheckParity(ctx context.Context, n *number.NumberRequest) (*number.NumberResponse, error) {
	// parse string to int
	p.logger.Infof("got number request with value : %s", n.Number)
	num, err := strconv.Atoi(n.Number)
	if err != nil {
		p.logger.Error(err)
		return nil, err
	}
	if num%2 == 0 {
		p.logger.Infof("number : %d is pair", num)
		return &number.NumberResponse{
			Response: number.NumberResponse_PAIR,
		}, nil
	} else {
		p.logger.Infof("number : %d is impair", num)
		return &number.NumberResponse{
			Response: number.NumberResponse_IMPAIR,
		}, nil
	}

}

// run the grpc server
func (p *GrpcPlugin) Run() error {
	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
		p.logger.Errorf("failed to listen: %v", err)
		return err
	}
	s := grpc.NewServer()
	number.RegisterParityServer(s, p)
	p.logger.Infof("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		p.logger.Errorf("failed to serve: %v", err)
				return err
	}
	return nil
}
