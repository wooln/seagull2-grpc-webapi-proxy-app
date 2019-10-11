package main

import (	
	proxyLib "github.com/wooln/seagull2-grpc-webapi-proxy"
	"log"
	gw "Foo_Contracts"  // Update
)

func main()  {
	actions := []proxyLib.RegisterAction {
		proxyLib.RegisterAction{
			gw.RegisterGreeterHandlerFromEndpoint,
			"localhost:8080",
		},		
		proxyLib.RegisterAction{
			gw.RegisterGreeterNewHandlerFromEndpoint,
			"localhost:8080",
		},		
	}
	log.Printf(proxyLib.Msg)
	proxyLib.Proxy(actions)
}