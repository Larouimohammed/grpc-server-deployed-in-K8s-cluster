package main

import (
	"grpc-exampl/agent"
	"sync"
)

func main() {
	 wg:= &sync.WaitGroup{}
	
	a:= agent.DefaultAgent // A0
	if err := a.Start(wg); err != nil {
		a.Logger.Warn(err)
	   
	}

wg.Wait()
}
