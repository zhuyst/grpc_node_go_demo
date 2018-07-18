'use strict';

module.exports = appInfo => {
  const config = exports = {};

  // use for cookie sign key, should change to your own and keep security
  config.keys = appInfo.name + '_1531908274808_8695';

  // egg-grpc配置
  config.grpc = {
    endpoint: 'localhost:50051', // 服务端地址
  };

  return config;
};
