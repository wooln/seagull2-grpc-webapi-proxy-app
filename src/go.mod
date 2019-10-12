module github.com/wooln/seagull2-grpc-webapi-proxy-app

go 1.12

require (
	Foo_Contracts v0.0.0-00010101000000-000000000000
	github.com/grpc-ecosystem/grpc-gateway v1.11.3 // indirect
	github.com/kardianos/service v1.0.0 // indirect
	github.com/wooln/seagull2-grpc-webapi-proxy v0.0.0
	google.golang.org/grpc v1.24.0 // indirect
)

replace github.com/wooln/seagull2-grpc-webapi-proxy => ../seagull2-grpc-webapi-proxy/src

replace Foo_Contracts => ../../DataGovernanceLab/Go/Foo_Contracts //TODO 路径引用
