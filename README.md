# Summary
用于把grpc该代理成http webapi. 
本应用引用了类库seagull2-grpc-webapi-proxy

# Usage
使用时fork此仓库进行必要的重写.
1. 准备和引入[grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway/)生成的go文件
2. 准备配置文件，务必修改服务名，根据需要修改Mapping
```
{
	"Name": "服务名称",
	"DisplayName": "服务显示名",
	"Description": "服务描述",
	
    "WebAPIPort":"如:8081, 代称成webapi的端口, 不要缺少前面的冒号",
	"GrpcEndpointMapping": {
		"Greeter" : "localhost:8080",
		"GreeterNew" : "localhost:8080"
	},
	
	"Stderr": "C:\\builder_err.log",
	"Stdout": "C:\\builder_out.log"
}
```
3. 调用Proxy传入的[]RegisterAction类型的参数, EndpointKey为配置文件中的GrpcEndpointMapping对应的Key
```
func main()  {
	actions := []proxyLib.RegisterAction {
		proxyLib.RegisterAction{
			Action: gw.RegisterGreeterHandlerFromEndpoint, //grpc-gateway生成的
			EndpointKey: "Greeter",//对应配置文件中的GrpcEndpointMapping的Key
		},		
		proxyLib.RegisterAction{
			Action: gw.RegisterGreeterNewHandlerFromEndpoint,
			EndpointKey: "GreeterNew",
		},		
	}
	proxyLib.Proxy(actions)
}
```
4. 使用```go build```进行编译
5. ./xx.exe install 安装服务
6. ./xx.exe uninstall 卸载服务
