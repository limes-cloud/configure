/*
 Navicat Premium Data Transfer

 Source Server         : dev
 Source Server Type    : MySQL
 Source Server Version : 80200
 Source Host           : localhost:3306
 Source Schema         : configure_hello

 Target Server Type    : MySQL
 Target Server Version : 80200
 File Encoding         : 65001

 Date: 24/03/2024 14:23:53
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for business
-- ----------------------------
DROP TABLE IF EXISTS `business`;
CREATE TABLE `business` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '修改时间',
  `server_id` int unsigned NOT NULL COMMENT '服务id',
  `keyword` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '变量标识',
  `type` varchar(32) NOT NULL COMMENT '变量类型',
  `description` varchar(128) NOT NULL COMMENT '变量描述',
  PRIMARY KEY (`id`),
  KEY `idx_business_created_at` (`created_at`),
  KEY `idx_business_updated_at` (`updated_at`),
  KEY `fk_business_server` (`server_id`),
  CONSTRAINT `fk_business_server` FOREIGN KEY (`server_id`) REFERENCES `server` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='业务变量';

-- ----------------------------
-- Records of business
-- ----------------------------
BEGIN;
INSERT INTO `business` VALUES (1, 1711261397, 1711261397, 2, 'AuthSkipRoles', 'object', '跳过权限检测的角色列表');
INSERT INTO `business` VALUES (2, 1711261397, 1711261397, 2, 'DefaultUserPassword', 'string', '默认账号密码');
INSERT INTO `business` VALUES (3, 1711261397, 1711261397, 2, 'JwtWhitelist', 'object', 'jwt校验白名单');
INSERT INTO `business` VALUES (4, 1711261397, 1711261397, 2, 'LoginPrivatePath', 'string', 'rsa解密私钥路径');
INSERT INTO `business` VALUES (5, 1711261397, 1711261397, 2, 'Setting', 'object', '系统设置');
INSERT INTO `business` VALUES (6, 1711261397, 1711261397, 3, 'AcceptTypes', 'object', '允许上传的文件后缀');
INSERT INTO `business` VALUES (7, 1711261397, 1711261397, 3, 'MaxChunkCount', 'int', '最大切片数量');
INSERT INTO `business` VALUES (8, 1711261397, 1711261397, 3, 'MaxChunkSize', 'int', '单个切片最大大小（M）');
INSERT INTO `business` VALUES (9, 1711261397, 1711261397, 3, 'MaxSingularSize', 'int', '单个文件最大大小（M），超过后被切片');
COMMIT;

-- ----------------------------
-- Table structure for business_value
-- ----------------------------
DROP TABLE IF EXISTS `business_value`;
CREATE TABLE `business_value` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '修改时间',
  `env_id` int unsigned NOT NULL COMMENT '环境id',
  `business_id` int unsigned NOT NULL COMMENT '业务变量id',
  `value` text NOT NULL COMMENT '业务变量id',
  PRIMARY KEY (`id`),
  KEY `idx_business_value_created_at` (`created_at`),
  KEY `idx_business_value_updated_at` (`updated_at`),
  KEY `fk_business_values` (`business_id`),
  KEY `fk_business_value_env` (`env_id`),
  CONSTRAINT `fk_business_value_env` FOREIGN KEY (`env_id`) REFERENCES `env` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_business_values` FOREIGN KEY (`business_id`) REFERENCES `business` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='业务变量值';

-- ----------------------------
-- Records of business_value
-- ----------------------------
BEGIN;
INSERT INTO `business_value` VALUES (1, 1711261397, 1711261397, 1, 1, '[\"superAdmin\"]');
INSERT INTO `business_value` VALUES (2, 1711261397, 1711261397, 2, 1, '[\"superAdmin\"]');
INSERT INTO `business_value` VALUES (3, 1711261397, 1711261397, 3, 1, '[\"superAdmin\"]');
INSERT INTO `business_value` VALUES (4, 1711261397, 1711261397, 1, 2, '12345678');
INSERT INTO `business_value` VALUES (5, 1711261397, 1711261397, 2, 2, '12345678');
INSERT INTO `business_value` VALUES (6, 1711261397, 1711261397, 3, 2, '12345678');
INSERT INTO `business_value` VALUES (7, 1711261397, 1711261397, 1, 3, '{\"GET:/manager/v1/setting\":true,\"POST:/manager/v1/login/captcha\":true,\"POST:/manager/v1/login\":true,\"POST:/manager/v1/token/refresh\":true}');
INSERT INTO `business_value` VALUES (8, 1711261397, 1711261397, 2, 3, '{\"GET:/manager/v1/setting\":true,\"POST:/manager/v1/login/captcha\":true,\"POST:/manager/v1/login\":true,\"POST:/manager/v1/token/refresh\":true}');
INSERT INTO `business_value` VALUES (9, 1711261397, 1711261397, 3, 3, '{\"GET:/manager/v1/setting\":true,\"POST:/manager/v1/login/captcha\":true,\"POST:/manager/v1/login\":true,\"POST:/manager/v1/token/refresh\":true}');
INSERT INTO `business_value` VALUES (10, 1711261397, 1711261397, 1, 4, 'static/cert/login.pem');
INSERT INTO `business_value` VALUES (11, 1711261397, 1711261397, 2, 4, 'static/cert/login.pem');
INSERT INTO `business_value` VALUES (12, 1711261397, 1711261397, 3, 4, 'static/cert/login.pem');
INSERT INTO `business_value` VALUES (13, 1711261397, 1711261397, 1, 5, '{\"name\":\"go-platform\",\"title\":\"go-platform 快速开发框架\",\"desc\":\"开放协作，拥抱未来，插件化编程实现\",\"copyright\":\"Copyright © 2023 qlime.cn. All rights reserved.\",\"logo\":\"http://p3-armor.byteimg.com/tos-cn-i-49unhts6dw/dfdba5317c0c20ce20e64fac803d52bc.svg~tplv-49unhts6dw-image.image\"}');
INSERT INTO `business_value` VALUES (14, 1711261397, 1711261397, 2, 5, '{\"name\":\"go-platform\",\"title\":\"go-platform 快速开发框架\",\"desc\":\"开放协作，拥抱未来，插件化编程实现\",\"copyright\":\"Copyright © 2023 qlime.cn. All rights reserved.\",\"logo\":\"http://p3-armor.byteimg.com/tos-cn-i-49unhts6dw/dfdba5317c0c20ce20e64fac803d52bc.svg~tplv-49unhts6dw-image.image\"}');
INSERT INTO `business_value` VALUES (15, 1711261397, 1711261397, 3, 5, '{\"name\":\"go-platform\",\"title\":\"go-platform 快速开发框架\",\"desc\":\"开放协作，拥抱未来，插件化编程实现\",\"copyright\":\"Copyright © 2023 qlime.cn. All rights reserved.\",\"logo\":\"http://p3-armor.byteimg.com/tos-cn-i-49unhts6dw/dfdba5317c0c20ce20e64fac803d52bc.svg~tplv-49unhts6dw-image.image\"}');
INSERT INTO `business_value` VALUES (16, 1711261397, 1711261397, 1, 6, '[\"jpg\",\"png\",\"txt\",\"ppt\",\"pptx\",\"mp4\"]');
INSERT INTO `business_value` VALUES (17, 1711261397, 1711261397, 2, 6, '[\"jpg\",\"png\",\"txt\",\"ppt\",\"pptx\",\"mp4\"]');
INSERT INTO `business_value` VALUES (18, 1711261397, 1711261397, 3, 6, '[\"jpg\",\"png\",\"txt\",\"ppt\",\"pptx\",\"mp4\"]');
INSERT INTO `business_value` VALUES (19, 1711261397, 1711261397, 1, 7, '200');
INSERT INTO `business_value` VALUES (20, 1711261397, 1711261397, 2, 7, '200');
INSERT INTO `business_value` VALUES (21, 1711261397, 1711261397, 3, 7, '200');
INSERT INTO `business_value` VALUES (22, 1711261397, 1711261397, 1, 8, '200');
INSERT INTO `business_value` VALUES (23, 1711261397, 1711261397, 2, 8, '200');
INSERT INTO `business_value` VALUES (24, 1711261397, 1711261397, 3, 8, '200');
INSERT INTO `business_value` VALUES (25, 1711261397, 1711261397, 1, 9, '400');
INSERT INTO `business_value` VALUES (26, 1711261397, 1711261397, 2, 9, '400');
INSERT INTO `business_value` VALUES (27, 1711261397, 1711261397, 3, 9, '400');
COMMIT;

-- ----------------------------
-- Table structure for configure
-- ----------------------------
DROP TABLE IF EXISTS `configure`;
CREATE TABLE `configure` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '修改时间',
  `server_id` int unsigned NOT NULL COMMENT '服务id',
  `env_id` int unsigned NOT NULL COMMENT '环境id',
  `content` text NOT NULL COMMENT '配置内容',
  `version` varchar(32) NOT NULL COMMENT '配置版本',
  `format` varchar(32) NOT NULL COMMENT '配置格式',
  `description` varchar(128) DEFAULT NULL COMMENT '配置描述',
  PRIMARY KEY (`id`),
  KEY `idx_configure_created_at` (`created_at`),
  KEY `idx_configure_updated_at` (`updated_at`),
  KEY `fk_configure_server` (`server_id`),
  KEY `fk_configure_env` (`env_id`),
  CONSTRAINT `fk_configure_env` FOREIGN KEY (`env_id`) REFERENCES `env` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_configure_server` FOREIGN KEY (`server_id`) REFERENCES `server` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='配置内容';

-- ----------------------------
-- Records of configure
-- ----------------------------
BEGIN;
INSERT INTO `configure` VALUES (1, 1711261397, 1711261397, 1, 1, '\ndebug: true\naddr: 0.0.0.0:7080\nname: gateway\nversion: v1\nmiddlewares:\n  - name: bbr\n  - name: cors\n    options:\n      allowCredentials: true\n      allowOrigins:\n        - \'*\'\n      allowMethods:\n        - GET\n        - POST\n        - PUT\n        - DELETE\n        - OPTIONS\n      AllowHeaders:\n        - Content-Type\n        - Content-Length\n        - Authorization\n      ExposeHeaders:\n        - Content-Length\n        - Access-Control-Allow-Headers\n  - name: tracing\n  - name: logging\n  - name: transcoder\nendpoints:\n  - path: /manager/v1/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7001\n  - path: /manager/client/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7001\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/user-center/client/token/parse\n          method: POST\n  - path: /configure/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:6081\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/manager/v1/auth\n          method: POST\n  - path: /resource/v1/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7003\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/manager/v1/auth\n          method: POST\n          whiteList:\n            - path: /resource/v1/static/*\n              method: GET\n  - path: /resource/client/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7003\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/user-center/client/token/parse\n          method: POST\n  - path: /user-center/v1/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7004\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/manager/v1/auth\n          method: POST\n  - path: /user-center/client/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7004\n  - path: /party-affairs/v1/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7100\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/manager/v1/auth\n          method: POST\n  - path: /party-affairs/client/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7100\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/user-center/client/token/parse\n          method: POST\n', '1276f4fafc63c4735e0e13d2ecb43451', 'yaml', '自动初始化');
INSERT INTO `configure` VALUES (2, 1711261397, 1711261397, 2, 1, '\nenv: TEST\nserver:\n  http:\n    host: 127.0.0.1\n    port: 7001\n    timeout: 10s\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: 127.0.0.1\n    port: 8001\n    timeout: 10s\nlog:\n  level: 0\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: mysql #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: root\n      password: root\n      host: 127.0.0.1\n      port: 3306\n      dbName: manager\n      option: ?charset=utf8mb4&parseTime=True&loc=Local\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 4 #日志等级\n      slowThreshold: 2s #慢sql阈值\nredis:\n  cache:\n    enable: true\n    host: 127.0.0.1:6379\n    username: \n    password: \ncaptcha:\n  login:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 240\n    skew: 0.7\n    refresh: true\n    dotCount: 80\n  changePassword:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: captcha\nloader:\n  login: static/cert/login.pem\nemail:\n  template:\n    captcha:\n      subject: 验证码发送通知\n      path: static/template/email/default.html\n      from: 青岑云科技\n      type: text/html\n  user: 860808187@qq.com\n  name: 青岑云\n  host: smtp.qq.com\n  port: 25\n  password: xxx\njwt:\n  redis: cache\n  secret: limes-cloud-test\n  expire: 2h\n  renewal: 120s\n  whitelist: {\"GET:/manager/client/v1/dictionary/value\":true,\"GET:/manager/v1/setting\":true,\"POST:/manager/v1/user/login\":true,\"POST:/manager/v1/user/login/captcha\":true,\"POST:/manager/v1/user/token/refresh\":true,\"POST:/resource/client/v1/upload\":true,\"POST:/resource/client/v1/upload/prepare\":true}\nauthentication:\n  db: system\n  redis: cache\n  roleKey: role_keyword\n  skipRole: [\"superAdmin\"]\nbusiness:\n  defaultUserPassword: 12345678\n  setting: {\"name\":\"go-platform\",\"title\":\"go-platform 快速开发框架\",\"desc\":\"开放协作，拥抱未来，插件化编程实现\",\"copyright\":\"Copyright © 2023 qlime.cn. All rights reserved.\",\"logo\":\"http://p3-armor.byteimg.com/tos-cn-i-49unhts6dw/dfdba5317c0c20ce20e64fac803d52bc.svg~tplv-49unhts6dw-image.image\"}\n', '3e40138da612990fece6905ab31f4027', 'yaml', '自动初始化');
INSERT INTO `configure` VALUES (3, 1711261397, 1711261397, 3, 1, '\nenv: TEST\nserver:\n  http:\n    host: 127.0.0.1\n    port: 7003\n    timeout: 10s\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: 127.0.0.1\n    port: 8003\n    timeout: 10s\nlog:\n  level: 0\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: mysql #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: root\n      password: root\n      host: 127.0.0.1\n      port: 3306\n      dbName: resource\n      option: ?charset=utf8mb4&parseTime=True&loc=Local\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 3 #日志等级\n      slowThreshold: 2s #慢sql阈值\nfile:\n  storage: local\n  serverPath: /resource/v1/static\n  localDir: static\n  maxSingularSize: 400\n  maxChunkSize: 200\n  maxChunkCount: 200\n  acceptTypes: [\"jpg\",\"png\",\"txt\",\"ppt\",\"pptx\",\"mp4\"]\n\n', '326986017e67d2c8aa1f0727a69927e9', 'yaml', '自动初始化');
INSERT INTO `configure` VALUES (4, 1711261397, 1711261397, 4, 1, '\nenv: TEST\nserver:\n  http:\n    host: 127.0.0.1\n    port: 7004\n    timeout: 10s\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: 127.0.0.1\n    port: 8004\n    timeout: 10s\nclient:\n  - server: Resource\n    type: direct\n    backends:\n      - target: 127.0.0.1:8003\nlog:\n  level: 0\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: mysql #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: root\n      password: root\n      host: 127.0.0.1\n      port: 3306\n      dbName: user_center\n      option: ?charset=utf8mb4&parseTime=True&loc=Local\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 4 #日志等级\n      slowThreshold: 2s #慢sql阈值\nredis:\n  cache:\n    enable: true\n    host: 127.0.0.1:6379\n    username: \n    password: \ncaptcha:\n  loginImage:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 200\n    skew: 0.7\n    dotCount: 80\n    refresh: true\n  bindImage:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 200\n    skew: 0.7\n    dotCount: 80\n    refresh: true\n  registerImage:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 200\n    skew: 0.7\n    dotCount: 80\n    refresh: true\n  loginEmail:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: login\n  bindEmail:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: bind\n  registerEmail:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: register\nloader:\n  password: static/cert/password.pem\nemail:\n  template:\n    login:\n      subject: 登录验证码发送通知\n      path: static/template/email/default.html\n      type: text/html\n    bind:\n      subject: 绑定验证码发送通知\n      path: static/template/email/default.html\n      type: text/html\n    register:\n      subject: 注册验证码发送通知\n      path: static/template/email/default.html\n      type: text/html\n  user: 860808187@qq.com\n  name: 青岑云\n  host: smtp.qq.com\n  port: 25\n  password: xxx\njwt:\n  redis: cache\n  secret: limes-cloud-client-test\n  expire: 2h\n  renewal: 360s\n  whitelist: {\"*:/user-center/client/v1/*\":true,\"*:/user-center/v1/*\":true,\"POST:/user-center/client/token/refresh\":true}\nbusiness:\n  service:\n    resource: 127.0.0.1:8003\n', '27c78261ae5fdd453127ae7b2760ec19', 'yaml', '自动初始化');
INSERT INTO `configure` VALUES (5, 1711261397, 1711261397, 5, 1, '\nenv: TEST\nserver:\n  http:\n    host: 127.0.0.1\n    port: 7100\n    timeout: 10s\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: 127.0.0.1\n    port: 8100\n    timeout: 10s\nclient:\n  - server: Resource\n    type: direct\n    backends:\n      - target: 127.0.0.1:8003\n  - server: UserCenter\n    type: direct\n    backends:\n      - target: 127.0.0.1:8004\nlog:\n  level: 0\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: mysql #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: root\n      password: root\n      host: 127.0.0.1\n      port: 3306\n      dbName: party_affairs\n      option: ?charset=utf8mb4&parseTime=True&loc=Local\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 4 #日志等级\n      slowThreshold: 2s #慢sql阈值\nbusiness:\n  auth:\n    yiBan:\n      appId: e4750b34230b48e1\n      appSecret: b0891a7f6018e5a76b085e3cb9548edd\n', 'dfff643162f3c96185e12ca22bf2033b', 'yaml', '自动初始化');
INSERT INTO `configure` VALUES (6, 1711261397, 1711261397, 1, 2, '\ndebug: true\naddr: 0.0.0.0:7080\nname: gateway\nversion: v1\nmiddlewares:\n  - name: bbr\n  - name: cors\n    options:\n      allowCredentials: true\n      allowOrigins:\n        - \'*\'\n      allowMethods:\n        - GET\n        - POST\n        - PUT\n        - DELETE\n        - OPTIONS\n      AllowHeaders:\n        - Content-Type\n        - Content-Length\n        - Authorization\n      ExposeHeaders:\n        - Content-Length\n        - Access-Control-Allow-Headers\n  - name: tracing\n  - name: logging\n  - name: transcoder\nendpoints:\n  - path: /manager/v1/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7001\n  - path: /manager/client/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7001\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/user-center/client/token/parse\n          method: POST\n  - path: /configure/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:6081\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/manager/v1/auth\n          method: POST\n  - path: /resource/v1/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7003\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/manager/v1/auth\n          method: POST\n          whiteList:\n            - path: /resource/v1/static/*\n              method: GET\n  - path: /resource/client/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7003\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/user-center/client/token/parse\n          method: POST\n  - path: /user-center/v1/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7004\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/manager/v1/auth\n          method: POST\n  - path: /user-center/client/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7004\n  - path: /party-affairs/v1/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7100\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/manager/v1/auth\n          method: POST\n  - path: /party-affairs/client/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7100\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/user-center/client/token/parse\n          method: POST\n', '1276f4fafc63c4735e0e13d2ecb43451', 'yaml', '自动初始化');
INSERT INTO `configure` VALUES (7, 1711261397, 1711261397, 2, 2, '\nenv: PRE\nserver:\n  http:\n    host: 127.0.0.1\n    port: 7001\n    timeout: 10s\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: 127.0.0.1\n    port: 8001\n    timeout: 10s\nlog:\n  level: 0\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: mysql #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: root\n      password: root\n      host: 127.0.0.1\n      port: 3306\n      dbName: manager\n      option: ?charset=utf8mb4&parseTime=True&loc=Local\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 4 #日志等级\n      slowThreshold: 2s #慢sql阈值\nredis:\n  cache:\n    enable: true\n    host: 127.0.0.1:6379\n    username: \n    password: \ncaptcha:\n  login:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 240\n    skew: 0.7\n    refresh: true\n    dotCount: 80\n  changePassword:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: captcha\nloader:\n  login: static/cert/login.pem\nemail:\n  template:\n    captcha:\n      subject: 验证码发送通知\n      path: static/template/email/default.html\n      from: 青岑云科技\n      type: text/html\n  user: 860808187@qq.com\n  name: 青岑云\n  host: smtp.qq.com\n  port: 25\n  password: xxx\njwt:\n  redis: cache\n  secret: limes-cloud-pre\n  expire: 2h\n  renewal: 120s\n  whitelist: {\"GET:/manager/client/v1/dictionary/value\":true,\"GET:/manager/v1/setting\":true,\"POST:/manager/v1/user/login\":true,\"POST:/manager/v1/user/login/captcha\":true,\"POST:/manager/v1/user/token/refresh\":true,\"POST:/resource/client/v1/upload\":true,\"POST:/resource/client/v1/upload/prepare\":true}\nauthentication:\n  db: system\n  redis: cache\n  roleKey: role_keyword\n  skipRole: [\"superAdmin\"]\nbusiness:\n  defaultUserPassword: 12345678\n  setting: {\"name\":\"go-platform\",\"title\":\"go-platform 快速开发框架\",\"desc\":\"开放协作，拥抱未来，插件化编程实现\",\"copyright\":\"Copyright © 2023 qlime.cn. All rights reserved.\",\"logo\":\"http://p3-armor.byteimg.com/tos-cn-i-49unhts6dw/dfdba5317c0c20ce20e64fac803d52bc.svg~tplv-49unhts6dw-image.image\"}\n', '33229ad0f2711424f9ccd19b807ac680', 'yaml', '自动初始化');
INSERT INTO `configure` VALUES (8, 1711261397, 1711261397, 3, 2, '\nenv: PRE\nserver:\n  http:\n    host: 127.0.0.1\n    port: 7003\n    timeout: 10s\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: 127.0.0.1\n    port: 8003\n    timeout: 10s\nlog:\n  level: 0\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: mysql #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: root\n      password: root\n      host: 127.0.0.1\n      port: 3306\n      dbName: resource\n      option: ?charset=utf8mb4&parseTime=True&loc=Local\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 3 #日志等级\n      slowThreshold: 2s #慢sql阈值\nfile:\n  storage: local\n  serverPath: /resource/v1/static\n  localDir: static\n  maxSingularSize: 400\n  maxChunkSize: 200\n  maxChunkCount: 200\n  acceptTypes: [\"jpg\",\"png\",\"txt\",\"ppt\",\"pptx\",\"mp4\"]\n\n', '6b8bb3f6ac2f8bb1bd3c1dec5cc7a11e', 'yaml', '自动初始化');
INSERT INTO `configure` VALUES (9, 1711261397, 1711261397, 4, 2, '\nenv: PRE\nserver:\n  http:\n    host: 127.0.0.1\n    port: 7004\n    timeout: 10s\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: 127.0.0.1\n    port: 8004\n    timeout: 10s\nclient:\n  - server: Resource\n    type: direct\n    backends:\n      - target: 127.0.0.1:8003\nlog:\n  level: 0\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: mysql #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: root\n      password: root\n      host: 127.0.0.1\n      port: 3306\n      dbName: user_center\n      option: ?charset=utf8mb4&parseTime=True&loc=Local\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 4 #日志等级\n      slowThreshold: 2s #慢sql阈值\nredis:\n  cache:\n    enable: true\n    host: 127.0.0.1:6379\n    username: \n    password: \ncaptcha:\n  loginImage:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 200\n    skew: 0.7\n    dotCount: 80\n    refresh: true\n  bindImage:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 200\n    skew: 0.7\n    dotCount: 80\n    refresh: true\n  registerImage:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 200\n    skew: 0.7\n    dotCount: 80\n    refresh: true\n  loginEmail:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: login\n  bindEmail:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: bind\n  registerEmail:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: register\nloader:\n  password: static/cert/password.pem\nemail:\n  template:\n    login:\n      subject: 登录验证码发送通知\n      path: static/template/email/default.html\n      type: text/html\n    bind:\n      subject: 绑定验证码发送通知\n      path: static/template/email/default.html\n      type: text/html\n    register:\n      subject: 注册验证码发送通知\n      path: static/template/email/default.html\n      type: text/html\n  user: 860808187@qq.com\n  name: 青岑云\n  host: smtp.qq.com\n  port: 25\n  password: xxx\njwt:\n  redis: cache\n  secret: limes-cloud-client-pre\n  expire: 2h\n  renewal: 360s\n  whitelist: {\"*:/user-center/client/v1/*\":true,\"*:/user-center/v1/*\":true,\"POST:/user-center/client/token/refresh\":true}\nbusiness:\n  service:\n    resource: 127.0.0.1:8003\n', '080c4cdf44dbf8451577b4edf1a4f702', 'yaml', '自动初始化');
INSERT INTO `configure` VALUES (10, 1711261397, 1711261397, 5, 2, '\nenv: PRE\nserver:\n  http:\n    host: 127.0.0.1\n    port: 7100\n    timeout: 10s\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: 127.0.0.1\n    port: 8100\n    timeout: 10s\nclient:\n  - server: Resource\n    type: direct\n    backends:\n      - target: 127.0.0.1:8003\n  - server: UserCenter\n    type: direct\n    backends:\n      - target: 127.0.0.1:8004\nlog:\n  level: 0\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: mysql #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: root\n      password: root\n      host: 127.0.0.1\n      port: 3306\n      dbName: party_affairs\n      option: ?charset=utf8mb4&parseTime=True&loc=Local\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 4 #日志等级\n      slowThreshold: 2s #慢sql阈值\nbusiness:\n  auth:\n    yiBan:\n      appId: e4750b34230b48e1\n      appSecret: b0891a7f6018e5a76b085e3cb9548edd\n', '6311dada0a48b1d307068a92da31a3f1', 'yaml', '自动初始化');
INSERT INTO `configure` VALUES (11, 1711261397, 1711261397, 1, 3, '\ndebug: true\naddr: 0.0.0.0:7080\nname: gateway\nversion: v1\nmiddlewares:\n  - name: bbr\n  - name: cors\n    options:\n      allowCredentials: true\n      allowOrigins:\n        - \'*\'\n      allowMethods:\n        - GET\n        - POST\n        - PUT\n        - DELETE\n        - OPTIONS\n      AllowHeaders:\n        - Content-Type\n        - Content-Length\n        - Authorization\n      ExposeHeaders:\n        - Content-Length\n        - Access-Control-Allow-Headers\n  - name: tracing\n  - name: logging\n  - name: transcoder\nendpoints:\n  - path: /manager/v1/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7001\n  - path: /manager/client/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7001\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/user-center/client/token/parse\n          method: POST\n  - path: /configure/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:6081\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/manager/v1/auth\n          method: POST\n  - path: /resource/v1/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7003\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/manager/v1/auth\n          method: POST\n          whiteList:\n            - path: /resource/v1/static/*\n              method: GET\n  - path: /resource/client/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7003\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/user-center/client/token/parse\n          method: POST\n  - path: /user-center/v1/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7004\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/manager/v1/auth\n          method: POST\n  - path: /user-center/client/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7004\n  - path: /party-affairs/v1/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7100\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/manager/v1/auth\n          method: POST\n  - path: /party-affairs/client/*\n    timeout: 10s\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: 127.0.0.1:7100\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:7080/user-center/client/token/parse\n          method: POST\n', '1276f4fafc63c4735e0e13d2ecb43451', 'yaml', '自动初始化');
INSERT INTO `configure` VALUES (12, 1711261397, 1711261397, 2, 3, '\nenv: PROD\nserver:\n  http:\n    host: 127.0.0.1\n    port: 7001\n    timeout: 10s\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: 127.0.0.1\n    port: 8001\n    timeout: 10s\nlog:\n  level: 0\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: mysql #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: manager\n      password: maGh8TzkjfyLbkYA\n      host: 127.0.0.1\n      port: 3306\n      dbName: manager\n      option: ?charset=utf8mb4&parseTime=True&loc=Local\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 4 #日志等级\n      slowThreshold: 2s #慢sql阈值\nredis:\n  cache:\n    enable: true\n    host: 127.0.0.1:6379\n    username: \n    password: \ncaptcha:\n  login:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 240\n    skew: 0.7\n    refresh: true\n    dotCount: 80\n  changePassword:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: captcha\nloader:\n  login: static/cert/login.pem\nemail:\n  template:\n    captcha:\n      subject: 验证码发送通知\n      path: static/template/email/default.html\n      from: 青岑云科技\n      type: text/html\n  user: 860808187@qq.com\n  name: 青岑云\n  host: smtp.qq.com\n  port: 25\n  password: xxx\njwt:\n  redis: cache\n  secret: limes-cloud-prod\n  expire: 2h\n  renewal: 120s\n  whitelist: {\"GET:/manager/client/v1/dictionary/value\":true,\"GET:/manager/v1/setting\":true,\"POST:/manager/v1/user/login\":true,\"POST:/manager/v1/user/login/captcha\":true,\"POST:/manager/v1/user/token/refresh\":true,\"POST:/resource/client/v1/upload\":true,\"POST:/resource/client/v1/upload/prepare\":true}\nauthentication:\n  db: system\n  redis: cache\n  roleKey: role_keyword\n  skipRole: [\"superAdmin\"]\nbusiness:\n  defaultUserPassword: 12345678\n  setting: {\"name\":\"go-platform\",\"title\":\"go-platform 快速开发框架\",\"desc\":\"开放协作，拥抱未来，插件化编程实现\",\"copyright\":\"Copyright © 2023 qlime.cn. All rights reserved.\",\"logo\":\"http://p3-armor.byteimg.com/tos-cn-i-49unhts6dw/dfdba5317c0c20ce20e64fac803d52bc.svg~tplv-49unhts6dw-image.image\"}\n', '76292d1df248620c1739b515730481a4', 'yaml', '自动初始化');
INSERT INTO `configure` VALUES (13, 1711261397, 1711261397, 3, 3, '\nenv: PROD\nserver:\n  http:\n    host: 127.0.0.1\n    port: 7003\n    timeout: 10s\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: 127.0.0.1\n    port: 8003\n    timeout: 10s\nlog:\n  level: 0\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: mysql #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: resource\n      password: Ti7MaKJJznywNBJb\n      host: 127.0.0.1\n      port: 3306\n      dbName: resource\n      option: ?charset=utf8mb4&parseTime=True&loc=Local\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 3 #日志等级\n      slowThreshold: 2s #慢sql阈值\nfile:\n  storage: local\n  serverPath: /resource/v1/static\n  localDir: static\n  maxSingularSize: 400\n  maxChunkSize: 200\n  maxChunkCount: 200\n  acceptTypes: [\"jpg\",\"png\",\"txt\",\"ppt\",\"pptx\",\"mp4\"]\n\n', '2cae49bbdd935a89dd1c26a22e358ff3', 'yaml', '自动初始化');
INSERT INTO `configure` VALUES (14, 1711261397, 1711261397, 4, 3, '\nenv: PROD\nserver:\n  http:\n    host: 127.0.0.1\n    port: 7004\n    timeout: 10s\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: 127.0.0.1\n    port: 8004\n    timeout: 10s\nclient:\n  - server: Resource\n    type: direct\n    backends:\n      - target: 127.0.0.1:8003\nlog:\n  level: 0\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: mysql #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: user_center\n      password: Ti7MaKJJznywNBJb\n      host: 127.0.0.1\n      port: 3306\n      dbName: user_center\n      option: ?charset=utf8mb4&parseTime=True&loc=Local\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 4 #日志等级\n      slowThreshold: 2s #慢sql阈值\nredis:\n  cache:\n    enable: true\n    host: 127.0.0.1:6379\n    username: \n    password: \ncaptcha:\n  loginImage:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 200\n    skew: 0.7\n    dotCount: 80\n    refresh: true\n  bindImage:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 200\n    skew: 0.7\n    dotCount: 80\n    refresh: true\n  registerImage:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 200\n    skew: 0.7\n    dotCount: 80\n    refresh: true\n  loginEmail:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: login\n  bindEmail:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: bind\n  registerEmail:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: register\nloader:\n  password: static/cert/password.pem\nemail:\n  template:\n    login:\n      subject: 登录验证码发送通知\n      path: static/template/email/default.html\n      type: text/html\n    bind:\n      subject: 绑定验证码发送通知\n      path: static/template/email/default.html\n      type: text/html\n    register:\n      subject: 注册验证码发送通知\n      path: static/template/email/default.html\n      type: text/html\n  user: 860808187@qq.com\n  name: 青岑云\n  host: smtp.qq.com\n  port: 25\n  password: xxx\njwt:\n  redis: cache\n  secret: limes-cloud-client-prod\n  expire: 2h\n  renewal: 360s\n  whitelist: {\"*:/user-center/client/v1/*\":true,\"*:/user-center/v1/*\":true,\"POST:/user-center/client/token/refresh\":true}\nbusiness:\n  service:\n    resource: 127.0.0.1:8003\n', '88d1870916daa43f25d80b47592212f4', 'yaml', '自动初始化');
INSERT INTO `configure` VALUES (15, 1711261397, 1711261397, 5, 3, '\nenv: PROD\nserver:\n  http:\n    host: 127.0.0.1\n    port: 7100\n    timeout: 10s\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: 127.0.0.1\n    port: 8100\n    timeout: 10s\nclient:\n  - server: Resource\n    type: direct\n    backends:\n      - target: 127.0.0.1:8003\n  - server: UserCenter\n    type: direct\n    backends:\n      - target: 127.0.0.1:8004\nlog:\n  level: 0\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: mysql #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: party_affairs\n      password: Ti7MaKJJznywNBJb\n      host: 127.0.0.1\n      port: 3306\n      dbName: party_affairs\n      option: ?charset=utf8mb4&parseTime=True&loc=Local\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 4 #日志等级\n      slowThreshold: 2s #慢sql阈值\nbusiness:\n  auth:\n    yiBan:\n      appId: e4750b34230b48e1\n      appSecret: b0891a7f6018e5a76b085e3cb9548edd\n', '8cd386b9b9f1f4be35be3695edf65380', 'yaml', '自动初始化');
COMMIT;

-- ----------------------------
-- Table structure for env
-- ----------------------------
DROP TABLE IF EXISTS `env`;
CREATE TABLE `env` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '修改时间',
  `keyword` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '环境标识',
  `name` varchar(64) NOT NULL COMMENT '环境名称',
  `description` varchar(128) NOT NULL COMMENT '环境描述',
  `token` varchar(128) NOT NULL COMMENT '环境token',
  `status` tinyint(1) DEFAULT '0' COMMENT '环境状态',
  PRIMARY KEY (`id`),
  KEY `idx_env_created_at` (`created_at`),
  KEY `idx_env_updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='环境信息';

-- ----------------------------
-- Records of env
-- ----------------------------
BEGIN;
INSERT INTO `env` VALUES (1, 1711261397, 1711261397, 'TEST', '测试环境', '用于本地测试', '1025D32F6CA7A2A320FE091B22C5DF3C', 1);
INSERT INTO `env` VALUES (2, 1711261397, 1711261397, 'PRE', '预发布环境', '用于上线前测试', 'F862AA25AF814B8981AA6B49E390F5C3', 1);
INSERT INTO `env` VALUES (3, 1711261397, 1711261397, 'PROD', '生产环境', '用于线上真实环境', '5B655B7D4A51BF79C974C9F27C27D992', 1);
COMMIT;

-- ----------------------------
-- Table structure for resource
-- ----------------------------
DROP TABLE IF EXISTS `resource`;
CREATE TABLE `resource` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '修改时间',
  `keyword` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '变量标识',
  `description` varchar(128) NOT NULL COMMENT '变量描述',
  `fields` varchar(256) NOT NULL COMMENT '变量字段集合',
  `tag` varchar(32) NOT NULL COMMENT '变量标签',
  `private` tinyint(1) DEFAULT '0' COMMENT '是否私有',
  PRIMARY KEY (`id`),
  KEY `idx_resource_created_at` (`created_at`),
  KEY `idx_resource_updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='资源变量';

-- ----------------------------
-- Records of resource
-- ----------------------------
BEGIN;
INSERT INTO `resource` VALUES (1, 1711261397, 1711261397, 'Env', '环境标识信息', 'Keyword', 'env', 0);
INSERT INTO `resource` VALUES (2, 1711261397, 1711261397, 'AdminJwt', '后台管理服务jwt配置信息', 'Secret,Expire,Renewal,Whitelist', 'jwt', 0);
INSERT INTO `resource` VALUES (3, 1711261397, 1711261397, 'ClientJwt', '客户端服务jwt配置信息', 'Secret,Expire,Renewal,Whitelist', 'jwt', 0);
INSERT INTO `resource` VALUES (4, 1711261397, 1711261397, 'Redis', 'redis中间件配置信息', 'Host,Username,Password,Port', 'redis', 0);
INSERT INTO `resource` VALUES (5, 1711261397, 1711261397, 'Email', '邮箱服务配置信息', 'Username,Company,Password,Host,Port', 'email', 0);
INSERT INTO `resource` VALUES (6, 1711261397, 1711261397, 'GatewayServer', '通用网关服务配置信息', 'HttpPort,Host,Timeout', 'server', 0);
INSERT INTO `resource` VALUES (7, 1711261397, 1711261397, 'ManagerServer', '管理中心服务配置信息', 'Host,HttpPort,GrpcPort,Timeout', 'server', 0);
INSERT INTO `resource` VALUES (8, 1711261397, 1711261397, 'ManagerDatabase', '管理中心数据库配置', 'Username,Password,Type,Port,Database,Option,Host', 'mysql', 1);
INSERT INTO `resource` VALUES (9, 1711261397, 1711261397, 'ResourceServer', '资源中心服务配置信息', 'Host,HttpPort,GrpcPort,Timeout', 'server', 0);
INSERT INTO `resource` VALUES (10, 1711261397, 1711261397, 'ResourceDatabase', '资源中心数据库配置信息', 'Username,Password,Type,Port,Database,Option,Host', 'mysql', 1);
INSERT INTO `resource` VALUES (11, 1711261397, 1711261397, 'ConfigureServer', '配置中心服务配置信息', 'Host,HttpPort,GrpcPort,Timeout', 'server', 0);
INSERT INTO `resource` VALUES (12, 1711261397, 1711261397, 'UserCenterServer', '用户中心服务配置信息', 'Host,HttpPort,GrpcPort,Timeout', 'server', 0);
INSERT INTO `resource` VALUES (13, 1711261397, 1711261397, 'UserCenterDatabase', '用户中心数据库配置信息', 'Username,Password,Type,Port,Database,Option,Host', 'mysql', 1);
INSERT INTO `resource` VALUES (14, 1711261397, 1711261397, 'PartyAffairsServer', '信号灯服务配置信息', 'Host,HttpPort,GrpcPort,Timeout', 'server', 0);
INSERT INTO `resource` VALUES (15, 1711261397, 1711261397, 'PartyAffairsDatabase', '信号灯数据库配置信息', 'Username,Password,Type,Port,Database,Option,Host', 'mysql', 1);
COMMIT;

-- ----------------------------
-- Table structure for resource_server
-- ----------------------------
DROP TABLE IF EXISTS `resource_server`;
CREATE TABLE `resource_server` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `server_id` int unsigned NOT NULL COMMENT '服务id',
  `resource_id` int unsigned NOT NULL COMMENT '资源id',
  PRIMARY KEY (`id`),
  UNIQUE KEY `sr` (`server_id`,`resource_id`),
  KEY `idx_resource_server_created_at` (`created_at`),
  KEY `fk_resource_resource_server` (`resource_id`),
  CONSTRAINT `fk_resource_resource_server` FOREIGN KEY (`resource_id`) REFERENCES `resource` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_resource_server_server` FOREIGN KEY (`server_id`) REFERENCES `server` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='资源服务信息';

-- ----------------------------
-- Records of resource_server
-- ----------------------------
BEGIN;
INSERT INTO `resource_server` VALUES (1, 1711261397, 2, 8);
INSERT INTO `resource_server` VALUES (2, 1711261397, 3, 10);
INSERT INTO `resource_server` VALUES (3, 1711261397, 4, 13);
INSERT INTO `resource_server` VALUES (4, 1711261397, 5, 15);
COMMIT;

-- ----------------------------
-- Table structure for resource_value
-- ----------------------------
DROP TABLE IF EXISTS `resource_value`;
CREATE TABLE `resource_value` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '修改时间',
  `env_id` int unsigned NOT NULL COMMENT '环境id',
  `resource_id` int unsigned NOT NULL COMMENT '资源id',
  `value` text NOT NULL COMMENT '资源id',
  PRIMARY KEY (`id`),
  UNIQUE KEY `er` (`env_id`,`resource_id`),
  KEY `idx_resource_value_created_at` (`created_at`),
  KEY `idx_resource_value_updated_at` (`updated_at`),
  KEY `fk_resource_resource_value` (`resource_id`),
  CONSTRAINT `fk_resource_resource_value` FOREIGN KEY (`resource_id`) REFERENCES `resource` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_resource_value_env` FOREIGN KEY (`env_id`) REFERENCES `env` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=46 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='资源变量值';

-- ----------------------------
-- Records of resource_value
-- ----------------------------
BEGIN;
INSERT INTO `resource_value` VALUES (1, 1711261397, 1711261397, 1, 1, '{\"Keyword\":\"TEST\"}');
INSERT INTO `resource_value` VALUES (2, 1711261397, 1711261397, 2, 1, '{\"Keyword\":\"PRE\"}');
INSERT INTO `resource_value` VALUES (3, 1711261397, 1711261397, 3, 1, '{\"Keyword\":\"PROD\"}');
INSERT INTO `resource_value` VALUES (4, 1711261397, 1711261397, 1, 2, '{\"Renewal\":\"120s\",\"Whitelist\":{\"POST:/manager/v1/user/login\":true,\"POST:/manager/v1/user/token/refresh\":true,\"GET:/manager/client/v1/dictionary/value\":true,\"POST:/resource/client/v1/upload/prepare\":true,\"POST:/resource/client/v1/upload\":true,\"GET:/manager/v1/setting\":true,\"POST:/manager/v1/user/login/captcha\":true},\"Secret\":\"limes-cloud-test\",\"Expire\":\"2h\"}');
INSERT INTO `resource_value` VALUES (5, 1711261397, 1711261397, 2, 2, '{\"Secret\":\"limes-cloud-pre\",\"Expire\":\"2h\",\"Renewal\":\"120s\",\"Whitelist\":{\"GET:/manager/v1/setting\":true,\"POST:/manager/v1/user/login/captcha\":true,\"POST:/manager/v1/user/login\":true,\"POST:/manager/v1/user/token/refresh\":true,\"GET:/manager/client/v1/dictionary/value\":true,\"POST:/resource/client/v1/upload/prepare\":true,\"POST:/resource/client/v1/upload\":true}}');
INSERT INTO `resource_value` VALUES (6, 1711261397, 1711261397, 3, 2, '{\"Secret\":\"limes-cloud-prod\",\"Expire\":\"2h\",\"Renewal\":\"120s\",\"Whitelist\":{\"POST:/manager/v1/user/login\":true,\"POST:/manager/v1/user/token/refresh\":true,\"GET:/manager/client/v1/dictionary/value\":true,\"POST:/resource/client/v1/upload/prepare\":true,\"POST:/resource/client/v1/upload\":true,\"GET:/manager/v1/setting\":true,\"POST:/manager/v1/user/login/captcha\":true}}');
INSERT INTO `resource_value` VALUES (7, 1711261397, 1711261397, 1, 3, '{\"Secret\":\"limes-cloud-client-test\",\"Expire\":\"2h\",\"Renewal\":\"360s\",\"Whitelist\":{\"*:/user-center/v1/*\":true,\"*:/user-center/client/v1/*\":true,\"POST:/user-center/client/token/refresh\":true}}');
INSERT INTO `resource_value` VALUES (8, 1711261397, 1711261397, 2, 3, '{\"Secret\":\"limes-cloud-client-pre\",\"Expire\":\"2h\",\"Renewal\":\"360s\",\"Whitelist\":{\"*:/user-center/client/v1/*\":true,\"POST:/user-center/client/token/refresh\":true,\"*:/user-center/v1/*\":true}}');
INSERT INTO `resource_value` VALUES (9, 1711261397, 1711261397, 3, 3, '{\"Expire\":\"2h\",\"Renewal\":\"360s\",\"Whitelist\":{\"*:/user-center/v1/*\":true,\"*:/user-center/client/v1/*\":true,\"POST:/user-center/client/token/refresh\":true},\"Secret\":\"limes-cloud-client-prod\"}');
INSERT INTO `resource_value` VALUES (10, 1711261397, 1711261397, 1, 4, '{\"Host\":\"127.0.0.1\",\"Username\":\"\",\"Password\":\"\",\"Port\":6379}');
INSERT INTO `resource_value` VALUES (11, 1711261397, 1711261397, 2, 4, '{\"Host\":\"127.0.0.1\",\"Username\":\"\",\"Password\":\"\",\"Port\":6379}');
INSERT INTO `resource_value` VALUES (12, 1711261397, 1711261397, 3, 4, '{\"Host\":\"127.0.0.1\",\"Username\":\"\",\"Password\":\"\",\"Port\":6379}');
INSERT INTO `resource_value` VALUES (13, 1711261397, 1711261397, 1, 5, '{\"Username\":\"860808187@qq.com\",\"Company\":\"青岑云\",\"Password\":\"xxx\",\"Host\":\"smtp.qq.com\",\"Port\":25}');
INSERT INTO `resource_value` VALUES (14, 1711261397, 1711261397, 2, 5, '{\"Port\":25,\"Username\":\"860808187@qq.com\",\"Company\":\"青岑云\",\"Password\":\"xxx\",\"Host\":\"smtp.qq.com\"}');
INSERT INTO `resource_value` VALUES (15, 1711261397, 1711261397, 3, 5, '{\"Username\":\"860808187@qq.com\",\"Company\":\"青岑云\",\"Password\":\"xxx\",\"Host\":\"smtp.qq.com\",\"Port\":25}');
INSERT INTO `resource_value` VALUES (16, 1711261397, 1711261397, 1, 6, '{\"Host\":\"127.0.0.1\",\"HttpPort\":7080,\"Timeout\":\"10s\"}');
INSERT INTO `resource_value` VALUES (17, 1711261397, 1711261397, 2, 6, '{\"Host\":\"127.0.0.1\",\"HttpPort\":7080,\"Timeout\":\"10s\"}');
INSERT INTO `resource_value` VALUES (18, 1711261397, 1711261397, 3, 6, '{\"Host\":\"127.0.0.1\",\"HttpPort\":7080,\"Timeout\":\"10s\"}');
INSERT INTO `resource_value` VALUES (19, 1711261397, 1711261397, 1, 7, '{\"Host\":\"127.0.0.1\",\"HttpPort\":7001,\"GrpcPort\":8001,\"Timeout\":\"10s\"}');
INSERT INTO `resource_value` VALUES (20, 1711261397, 1711261397, 2, 7, '{\"Host\":\"127.0.0.1\",\"HttpPort\":7001,\"GrpcPort\":8001,\"Timeout\":\"10s\"}');
INSERT INTO `resource_value` VALUES (21, 1711261397, 1711261397, 3, 7, '{\"Host\":\"127.0.0.1\",\"HttpPort\":7001,\"GrpcPort\":8001,\"Timeout\":\"10s\"}');
INSERT INTO `resource_value` VALUES (22, 1711261397, 1711261397, 1, 8, '{\"Port\":\"3306\",\"Type\":\"mysql\",\"Database\":\"manager\",\"Option\":\"?charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\",\"Username\":\"root\",\"Password\":\"root\",\"Host\":\"127.0.0.1\"}');
INSERT INTO `resource_value` VALUES (23, 1711261397, 1711261397, 2, 8, '{\"Password\":\"root\",\"Host\":\"127.0.0.1\",\"Port\":\"3306\",\"Type\":\"mysql\",\"Database\":\"manager\",\"Option\":\"?charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\",\"Username\":\"root\"}');
INSERT INTO `resource_value` VALUES (24, 1711261397, 1711261397, 3, 8, '{\"Username\":\"manager\",\"Password\":\"maGh8TzkjfyLbkYA\",\"Host\":\"127.0.0.1\",\"Port\":\"3306\",\"Type\":\"mysql\",\"Database\":\"manager\",\"Option\":\"?charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\"}');
INSERT INTO `resource_value` VALUES (25, 1711261397, 1711261397, 1, 9, '{\"Host\":\"127.0.0.1\",\"HttpPort\":7003,\"GrpcPort\":8003,\"Timeout\":\"10s\"}');
INSERT INTO `resource_value` VALUES (26, 1711261397, 1711261397, 2, 9, '{\"HttpPort\":7003,\"GrpcPort\":8003,\"Timeout\":\"10s\",\"Host\":\"127.0.0.1\"}');
INSERT INTO `resource_value` VALUES (27, 1711261397, 1711261397, 3, 9, '{\"GrpcPort\":8003,\"Timeout\":\"10s\",\"Host\":\"127.0.0.1\",\"HttpPort\":7003}');
INSERT INTO `resource_value` VALUES (28, 1711261397, 1711261397, 1, 10, '{\"Password\":\"root\",\"Host\":\"127.0.0.1\",\"Port\":\"3306\",\"Type\":\"mysql\",\"Database\":\"resource\",\"Option\":\"?charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\",\"Username\":\"root\"}');
INSERT INTO `resource_value` VALUES (29, 1711261397, 1711261397, 2, 10, '{\"Type\":\"mysql\",\"Database\":\"resource\",\"Option\":\"?charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\",\"Username\":\"root\",\"Password\":\"root\",\"Host\":\"127.0.0.1\",\"Port\":\"3306\"}');
INSERT INTO `resource_value` VALUES (30, 1711261397, 1711261397, 3, 10, '{\"Database\":\"resource\",\"Option\":\"?charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\",\"Username\":\"resource\",\"Password\":\"Ti7MaKJJznywNBJb\",\"Host\":\"127.0.0.1\",\"Port\":\"3306\",\"Type\":\"mysql\"}');
INSERT INTO `resource_value` VALUES (31, 1711261397, 1711261397, 1, 11, '{\"Host\":\"127.0.0.1\",\"HttpPort\":6081,\"GrpcPort\":6082,\"Timeout\":\"10s\"}');
INSERT INTO `resource_value` VALUES (32, 1711261397, 1711261397, 2, 11, '{\"Host\":\"127.0.0.1\",\"HttpPort\":6081,\"GrpcPort\":6082,\"Timeout\":\"10s\"}');
INSERT INTO `resource_value` VALUES (33, 1711261397, 1711261397, 3, 11, '{\"GrpcPort\":6082,\"Timeout\":\"10s\",\"Host\":\"127.0.0.1\",\"HttpPort\":6081}');
INSERT INTO `resource_value` VALUES (34, 1711261397, 1711261397, 1, 12, '{\"Timeout\":\"10s\",\"Host\":\"127.0.0.1\",\"HttpPort\":7004,\"GrpcPort\":8004}');
INSERT INTO `resource_value` VALUES (35, 1711261397, 1711261397, 2, 12, '{\"Host\":\"127.0.0.1\",\"HttpPort\":7004,\"GrpcPort\":8004,\"Timeout\":\"10s\"}');
INSERT INTO `resource_value` VALUES (36, 1711261397, 1711261397, 3, 12, '{\"Timeout\":\"10s\",\"Host\":\"127.0.0.1\",\"HttpPort\":7004,\"GrpcPort\":8004}');
INSERT INTO `resource_value` VALUES (37, 1711261397, 1711261397, 1, 13, '{\"Password\":\"root\",\"Host\":\"127.0.0.1\",\"Port\":\"3306\",\"Type\":\"mysql\",\"Database\":\"user_center\",\"Option\":\"?charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\",\"Username\":\"root\"}');
INSERT INTO `resource_value` VALUES (38, 1711261397, 1711261397, 2, 13, '{\"Password\":\"root\",\"Host\":\"127.0.0.1\",\"Port\":\"3306\",\"Type\":\"mysql\",\"Database\":\"user_center\",\"Option\":\"?charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\",\"Username\":\"root\"}');
INSERT INTO `resource_value` VALUES (39, 1711261397, 1711261397, 3, 13, '{\"Option\":\"?charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\",\"Username\":\"user_center\",\"Password\":\"Ti7MaKJJznywNBJb\",\"Host\":\"127.0.0.1\",\"Port\":\"3306\",\"Type\":\"mysql\",\"Database\":\"user_center\"}');
INSERT INTO `resource_value` VALUES (40, 1711261397, 1711261397, 1, 14, '{\"Host\":\"127.0.0.1\",\"HttpPort\":7100,\"GrpcPort\":8100,\"Timeout\":\"10s\"}');
INSERT INTO `resource_value` VALUES (41, 1711261397, 1711261397, 2, 14, '{\"Host\":\"127.0.0.1\",\"HttpPort\":7100,\"GrpcPort\":8100,\"Timeout\":\"10s\"}');
INSERT INTO `resource_value` VALUES (42, 1711261397, 1711261397, 3, 14, '{\"Host\":\"127.0.0.1\",\"HttpPort\":7100,\"GrpcPort\":8100,\"Timeout\":\"10s\"}');
INSERT INTO `resource_value` VALUES (43, 1711261397, 1711261397, 1, 15, '{\"Database\":\"party_affairs\",\"Option\":\"?charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\",\"Username\":\"root\",\"Password\":\"root\",\"Host\":\"127.0.0.1\",\"Port\":\"3306\",\"Type\":\"mysql\"}');
INSERT INTO `resource_value` VALUES (44, 1711261397, 1711261397, 2, 15, '{\"Database\":\"party_affairs\",\"Option\":\"?charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\",\"Username\":\"root\",\"Password\":\"root\",\"Host\":\"127.0.0.1\",\"Port\":\"3306\",\"Type\":\"mysql\"}');
INSERT INTO `resource_value` VALUES (45, 1711261397, 1711261397, 3, 15, '{\"Username\":\"party_affairs\",\"Password\":\"Ti7MaKJJznywNBJb\",\"Host\":\"127.0.0.1\",\"Port\":\"3306\",\"Type\":\"mysql\",\"Database\":\"party_affairs\",\"Option\":\"?charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\"}');
COMMIT;

-- ----------------------------
-- Table structure for server
-- ----------------------------
DROP TABLE IF EXISTS `server`;
CREATE TABLE `server` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '修改时间',
  `keyword` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '服务标识',
  `name` varchar(64) NOT NULL COMMENT '服务名称',
  `description` varchar(128) NOT NULL COMMENT '服务描述',
  PRIMARY KEY (`id`),
  KEY `idx_server_created_at` (`created_at`),
  KEY `idx_server_updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='服务信息';

-- ----------------------------
-- Records of server
-- ----------------------------
BEGIN;
INSERT INTO `server` VALUES (1, 1711261397, 1711261397, 'Gateway', '通用网关', '主要负责前端到后端的转发');
INSERT INTO `server` VALUES (2, 1711261397, 1711261397, 'Manager', '管理中心', '主要负责系统的基础管理');
INSERT INTO `server` VALUES (3, 1711261397, 1711261397, 'Resource', '资源中心', '主要负责静态资源的管理');
INSERT INTO `server` VALUES (4, 1711261397, 1711261397, 'UserCenter', '用户中心', '主要负责业务用户的管理');
INSERT INTO `server` VALUES (5, 1711261397, 1711261397, 'PartyAffairs', '信号灯系统', '指尖上的党务系统');
COMMIT;

-- ----------------------------
-- Table structure for template
-- ----------------------------
DROP TABLE IF EXISTS `template`;
CREATE TABLE `template` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '修改时间',
  `server_id` int unsigned NOT NULL COMMENT '模板id',
  `content` text NOT NULL COMMENT '模板内容',
  `version` varchar(32) NOT NULL COMMENT '模板版本',
  `is_use` tinyint(1) DEFAULT '0' COMMENT '是否使用',
  `format` varchar(32) NOT NULL COMMENT '模板格式',
  `description` varchar(128) NOT NULL COMMENT '模板描述',
  `compare` text NOT NULL COMMENT '变更详情',
  PRIMARY KEY (`id`),
  UNIQUE KEY `sv` (`server_id`,`version`),
  KEY `idx_template_created_at` (`created_at`),
  KEY `idx_template_updated_at` (`updated_at`),
  CONSTRAINT `fk_template_server` FOREIGN KEY (`server_id`) REFERENCES `server` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='配置模板';

-- ----------------------------
-- Records of template
-- ----------------------------
BEGIN;
INSERT INTO `template` VALUES (1, 1711261397, 1711261397, 1, '\ndebug: true\naddr: 0.0.0.0:${GatewayServer.HttpPort}\nname: gateway\nversion: v1\nmiddlewares:\n  - name: bbr\n  - name: cors\n    options:\n      allowCredentials: true\n      allowOrigins:\n        - \'*\'\n      allowMethods:\n        - GET\n        - POST\n        - PUT\n        - DELETE\n        - OPTIONS\n      AllowHeaders:\n        - Content-Type\n        - Content-Length\n        - Authorization\n      ExposeHeaders:\n        - Content-Length\n        - Access-Control-Allow-Headers\n  - name: tracing\n  - name: logging\n  - name: transcoder\nendpoints:\n  - path: /manager/v1/*\n    timeout: ${ManagerServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  - path: /manager/client/*\n    timeout: ${ManagerServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse\n          method: POST\n  - path: /configure/*\n    timeout: ${ConfigureServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/manager/v1/auth\n          method: POST\n  - path: /resource/v1/*\n    timeout: ${ResourceServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/manager/v1/auth\n          method: POST\n          whiteList:\n            - path: /resource/v1/static/*\n              method: GET\n  - path: /resource/client/*\n    timeout: ${ResourceServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse\n          method: POST\n  - path: /user-center/v1/*\n    timeout: ${UserCenterServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/manager/v1/auth\n          method: POST\n  - path: /user-center/client/*\n    timeout: ${UserCenterServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  - path: /party-affairs/v1/*\n    timeout: ${PartyAffairsServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/manager/v1/auth\n          method: POST\n  - path: /party-affairs/client/*\n    timeout: ${PartyAffairsServer.Timeout}\n    protocol: HTTP\n    responseFormat: true\n    backends:\n      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n    middlewares:\n      - name: auth\n        options:\n          url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse\n          method: POST\n', '96884AEB89D8', 1, 'yaml', '初始化模板', '');
INSERT INTO `template` VALUES (2, 1711261397, 1711261397, 2, '\nenv: ${Env.Keyword}\nserver:\n  http:\n    host: ${ManagerServer.Host}\n    port: ${ManagerServer.HttpPort}\n    timeout: ${ManagerServer.Timeout}\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: ${ManagerServer.Host}\n    port: ${ManagerServer.GrpcPort}\n    timeout: ${ManagerServer.Timeout}\nlog:\n  level: 0\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: ${ManagerDatabase.Type} #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: ${ManagerDatabase.Username}\n      password: ${ManagerDatabase.Password}\n      host: ${ManagerDatabase.Host}\n      port: ${ManagerDatabase.Port}\n      dbName: ${ManagerDatabase.Database}\n      option: ${ManagerDatabase.Option}\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 4 #日志等级\n      slowThreshold: 2s #慢sql阈值\nredis:\n  cache:\n    enable: true\n    host: ${Redis.Host}:${Redis.Port}\n    username: ${Redis.Username}\n    password: ${Redis.Password}\ncaptcha:\n  login:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 240\n    skew: 0.7\n    refresh: true\n    dotCount: 80\n  changePassword:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: captcha\nloader:\n  login: ${LoginPrivatePath}\nemail:\n  template:\n    captcha:\n      subject: 验证码发送通知\n      path: static/template/email/default.html\n      from: 青岑云科技\n      type: text/html\n  user: ${Email.Username}\n  name: ${Email.Company}\n  host: ${Email.Host}\n  port: ${Email.Port}\n  password: ${Email.Password}\njwt:\n  redis: cache\n  secret: ${AdminJwt.Secret}\n  expire: ${AdminJwt.Expire}\n  renewal: ${AdminJwt.Renewal}\n  whitelist: ${AdminJwt.Whitelist}\nauthentication:\n  db: system\n  redis: cache\n  roleKey: role_keyword\n  skipRole: ${AuthSkipRoles}\nbusiness:\n  defaultUserPassword: ${DefaultUserPassword}\n  setting: ${Setting}\n', '15768C4C6F57', 1, 'yaml', '初始化模板', '');
INSERT INTO `template` VALUES (3, 1711261397, 1711261397, 3, '\nenv: ${Env.Keyword}\nserver:\n  http:\n    host: ${ResourceServer.Host}\n    port: ${ResourceServer.HttpPort}\n    timeout: ${ResourceServer.Timeout}\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: ${ResourceServer.Host}\n    port: ${ResourceServer.GrpcPort}\n    timeout: ${ResourceServer.Timeout}\nlog:\n  level: 0\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: ${ResourceDatabase.Type} #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: ${ResourceDatabase.Username}\n      password: ${ResourceDatabase.Password}\n      host: ${ResourceDatabase.Host}\n      port: ${ResourceDatabase.Port}\n      dbName: ${ResourceDatabase.Database}\n      option: ${ResourceDatabase.Option}\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 3 #日志等级\n      slowThreshold: 2s #慢sql阈值\nfile:\n  storage: local\n  serverPath: /resource/v1/static\n  localDir: static\n  maxSingularSize: ${MaxSingularSize}\n  maxChunkSize: ${MaxChunkSize}\n  maxChunkCount: ${MaxChunkCount}\n  acceptTypes: ${AcceptTypes}\n\n', '56945B81FA4D', 1, 'yaml', '初始化模板', '');
INSERT INTO `template` VALUES (4, 1711261397, 1711261397, 4, '\nenv: ${Env.Keyword}\nserver:\n  http:\n    host: ${UserCenterServer.Host}\n    port: ${UserCenterServer.HttpPort}\n    timeout: ${UserCenterServer.Timeout}\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: ${UserCenterServer.Host}\n    port: ${UserCenterServer.GrpcPort}\n    timeout: ${UserCenterServer.Timeout}\nclient:\n  - server: Resource\n    type: direct\n    backends:\n      - target: ${ResourceServer.Host}:${ResourceServer.GrpcPort}\nlog:\n  level: 0\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: ${UserCenterDatabase.Type} #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: ${UserCenterDatabase.Username}\n      password: ${UserCenterDatabase.Password}\n      host: ${UserCenterDatabase.Host}\n      port: ${UserCenterDatabase.Port}\n      dbName: ${UserCenterDatabase.Database}\n      option: ${UserCenterDatabase.Option}\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 4 #日志等级\n      slowThreshold: 2s #慢sql阈值\nredis:\n  cache:\n    enable: true\n    host: ${Redis.Host}:${Redis.Port}\n    username: ${Redis.Username}\n    password: ${Redis.Password}\ncaptcha:\n  loginImage:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 200\n    skew: 0.7\n    dotCount: 80\n    refresh: true\n  bindImage:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 200\n    skew: 0.7\n    dotCount: 80\n    refresh: true\n  registerImage:\n    type: image\n    length: 6\n    expire: 180s\n    redis: cache\n    height: 80\n    width: 200\n    skew: 0.7\n    dotCount: 80\n    refresh: true\n  loginEmail:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: login\n  bindEmail:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: bind\n  registerEmail:\n    type: email\n    length: 6\n    expire: 180s\n    redis: cache\n    template: register\nloader:\n  password: static/cert/password.pem\nemail:\n  template:\n    login:\n      subject: 登录验证码发送通知\n      path: static/template/email/default.html\n      type: text/html\n    bind:\n      subject: 绑定验证码发送通知\n      path: static/template/email/default.html\n      type: text/html\n    register:\n      subject: 注册验证码发送通知\n      path: static/template/email/default.html\n      type: text/html\n  user: ${Email.Username}\n  name: ${Email.Company}\n  host: ${Email.Host}\n  port: ${Email.Port}\n  password: ${Email.Password}\njwt:\n  redis: cache\n  secret: ${ClientJwt.Secret}\n  expire: ${ClientJwt.Expire}\n  renewal: ${ClientJwt.Renewal}\n  whitelist: ${ClientJwt.Whitelist}\nbusiness:\n  service:\n    resource: ${ResourceServer.Host}:${ResourceServer.GrpcPort}\n', 'DE24124137B1', 1, 'yaml', '初始化模板', '');
INSERT INTO `template` VALUES (5, 1711261397, 1711261397, 5, '\nenv: ${Env.Keyword}\nserver:\n  http:\n    host: ${PartyAffairsServer.Host}\n    port: ${PartyAffairsServer.HttpPort}\n    timeout: ${PartyAffairsServer.Timeout}\n    marshal:\n      emitUnpopulated: true\n      useProtoNames: true\n  grpc:\n    host: ${PartyAffairsServer.Host}\n    port: ${PartyAffairsServer.GrpcPort}\n    timeout: ${PartyAffairsServer.Timeout}\nclient:\n  - server: Resource\n    type: direct\n    backends:\n      - target: ${ResourceServer.Host}:${ResourceServer.GrpcPort}\n  - server: UserCenter\n    type: direct\n    backends:\n      - target: ${UserCenterServer.Host}:${UserCenterServer.GrpcPort}\nlog:\n  level: 0\n  output:\n    - stdout\n    - file\n  file:\n    name: ./tmp/runtime/output.log\n    maxSize: 1\n    maxBackup: 5\n    maxAge: 1\n    compress: false\ndatabase:\n  system:\n    enable: true #是否启用数据库\n    drive: ${PartyAffairsDatabase.Type} #数据库类型\n    autoCreate: true #是否自动创建数据库\n    connect:\n      username: ${PartyAffairsDatabase.Username}\n      password: ${PartyAffairsDatabase.Password}\n      host: ${PartyAffairsDatabase.Host}\n      port: ${PartyAffairsDatabase.Port}\n      dbName: ${PartyAffairsDatabase.Database}\n      option: ${PartyAffairsDatabase.Option}\n    config:\n      transformError:\n        enable: true\n      initializer:\n        enable: true\n        path: deploy/data.sql\n      maxLifetime: 2h #最大生存时间\n      maxOpenConn: 20 #最大连接数量\n      maxIdleConn: 10 #最大空闲数量\n      logLevel: 4 #日志等级\n      slowThreshold: 2s #慢sql阈值\nbusiness:\n  auth:\n    yiBan:\n      appId: e4750b34230b48e1\n      appSecret: b0891a7f6018e5a76b085e3cb9548edd\n', 'D825BF2B6742', 1, 'yaml', '初始化模板', '');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
