package main

import (	
	proxyLib "github.com/wooln/seagull2-grpc-webapi-proxy"
	gw "Foo_Contracts"
)

func main()  {
	actions := []proxyLib.RegisterAction {
		proxyLib.RegisterAction{
			Action: gw.RegisterGreeterHandlerFromEndpoint,
			EndpointKey: "Greeter",
		},		
		proxyLib.RegisterAction{
			Action: gw.RegisterGreeterNewHandlerFromEndpoint,
			EndpointKey: "GreeterNew",
		},		
	}
	proxyLib.Proxy(actions)
}