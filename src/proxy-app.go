package main

import (	
	proxyLib "github.com/wooln/seagull2-grpc-webapi-proxy"
	gw "Foo_Contracts"
)

func main()  {
	actions := []proxyLib.RegisterAction {
		proxyLib.RegisterAction{
			gw.RegisterGreeterHandlerFromEndpoint,
			"Greeter",
		},		
		proxyLib.RegisterAction{
			gw.RegisterGreeterNewHandlerFromEndpoint,
			"GreeterNew",
		},		
	}
	proxyLib.Proxy(actions)
}