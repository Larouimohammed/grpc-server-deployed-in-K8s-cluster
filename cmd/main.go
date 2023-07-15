package main

import (
	"grpc-exampl/agent"
)

func main() {
	a := agent.DefaultAgent // A0
	if err := a.Start(); err != nil {
		a.Logger.Fatal(err)
	}

}
