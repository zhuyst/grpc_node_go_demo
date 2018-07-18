# 定义.proto

新建文件 **hello.proto** 定义一个 **HelloService** ，并且有一个 **SayHello** 方法。

```protobuf
syntax = "proto3";

package demo;
 b 
service HelloService {
    rpc SayHello (HelloRequest) returns (HelloResponse) {}
}

message HelloRequest {
    string code = 1;
    string message = 2;
}

message HelloResponse {
    string code = 1;
    string message = 2;
}
```

# Node.js

官网教程：[Node Quick Start](https://grpc.io/docs/quickstart/node.html)

官方例子：[grpc/examples/node/](https://github.com/grpc/grpc/tree/master/examples/node)

Node.js使用：
* 开发框架：[egg.js](https://github.com/eggjs/egg)
* grpc插件：[egg-grpc](https://github.com/eggjs/egg-grpc)

将 **hello.proto** 放入 **app/proto/** 下

## Node.js客户端

在 **config/plugin.js** 启动egg-grpc插件
```js
// egg-grpc插件
exports.grpc = {
  enable: true,
  package: 'egg-grpc',
};
```

在 **config/config.default.js** 配置grpc
```js
  // egg-grpc配置
  config.grpc = {
    endpoint: 'localhost:50051', // 服务端地址
  };
```

在 **app/controller/home.js** 调用服务端
```js
class HomeController extends Controller {
  async index() {
    const ctx = this.ctx;

    // 获得HelloService实例
    const helloService = ctx.grpc.demo.helloService;

    // 向服务端发送请求
    const result = await helloService.sayHello({
      code: '0',
      message: '来自Node客户端的OK',
    });

    // 打印服务端响应内容
    console.log(result);
    ctx.body = result;
  }
}
```

## Node.js服务端

一般服务端是在项目启动时进行加载，所以在 **app.js** 定义项目启动时执行的方法
```js
const PROTO_FILE_PATH = __dirname + '/app/proto/hello.proto'; // proto文件位置
const PORT = ':50051'; // RPC服务端端口

const grpc = require('grpc');
const protoLoader = require('@grpc/proto-loader');

module.exports = app => {
  app.beforeStart(async () => {
    // 新建一个grpc服务器
    const server = new grpc.Server();

    // 异步加载服务
    await protoLoader.load(PROTO_FILE_PATH).then(packageDefinition => {

      // 获取proto
      const helloProto = grpc.loadPackageDefinition(packageDefinition);

      // 获取package
      const grpc_demo = helloProto.demo;

      // 定义HelloService的SayHello实现
      const sayHello = (call, callback) => {
        // 打印客户端请求内容
        console.log(call.request);

        // 响应客户端
        callback(null, {
          code: '0',
          message: '来自Node服务端的OK',
        });
      };

      // 将sayHello方法作为HelloService的实现放入grpc服务器中
      server.addService(grpc_demo.HelloService.service, { sayHello });
    });

    // 启动服务监听
    server.bind(`0.0.0.0${PORT}`, grpc.ServerCredentials.createInsecure());
    server.start();
  });
};
```

# Go

官网教程：[Go Quick Start](https://grpc.io/docs/quickstart/go.html)

官方例子：[grpc-go/examples/helloworld/](https://github.com/grpc/grpc-go/tree/master/examples/helloworld)

Go按照教程直接导入对应的包即可，然后
1. 将 **hello.proto** 放入  **proto/** 下
2. 然后使用 `protoc -I proto/ proto/hello.proto --go_out=plugins=grpc:proto` 生成 **hello.pb.go**

## Go客户端
```go
const (
	address     = "localhost:50051" // 服务端地址
)

func main() {
	// 启动grpc客户端，连接grpc服务端
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 使用连接，创建HelloService实例
	helloService := pb.NewHelloServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 调用SayHello，向服务端发送消息
	response, err := helloService.SayHello(ctx, &pb.HelloRequest{
		Code: "0",
		Message: "来自GO客户端的OK",
	})
	if err != nil {
		log.Fatalf("could not sayHello: %v", err)
	}

	// 打印服务端回应内容
	log.Printf("SayHello: Code: %s,Message: %s", response.Code, response.Message)
}
```

## Go服务端

```go
const (
	port = ":50051" // RPC服务端端口
)

type HelloService struct{}

// 实现hello的HelloServiceServer中的SayHello方法
// 表示实现了HelloServiceServer接口
func (s *HelloService) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {

	// 打印客户端请求内容
	log.Printf("SayHello: Code: %s,Message: %s", request.Code, request.Message)

	// 响应客户端
	return &pb.HelloResponse{
		Code: "0",
		Message: "来自GO服务端的OK",
	}, nil
}

func main() {
	// 启动服务监听
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 新建一个grpc服务器
	s := grpc.NewServer()

	// 将实现HelloService注册到grpc服务器中
	pb.RegisterHelloServiceServer(s, &HelloService{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
```
