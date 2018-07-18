'use strict';

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
