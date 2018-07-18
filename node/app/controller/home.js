'use strict';

const Controller = require('egg').Controller;

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

module.exports = HomeController;
