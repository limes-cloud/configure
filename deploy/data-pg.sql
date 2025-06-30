


-- ----------------------------
-- Table structure for business
-- ----------------------------
DROP TABLE IF EXISTS "public"."business";
CREATE TABLE "public"."business" (
                                     "id" int8 NOT NULL,
                                     "created_at" int8,
                                     "updated_at" int8,
                                     "server_id" int8 NOT NULL,
                                     "keyword" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
                                     "type" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
                                     "description" varchar(128) COLLATE "pg_catalog"."default" NOT NULL
)
;
COMMENT ON COLUMN "public"."business"."id" IS '主键ID';
COMMENT ON COLUMN "public"."business"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."business"."updated_at" IS '修改时间';
COMMENT ON COLUMN "public"."business"."server_id" IS '服务id';
COMMENT ON COLUMN "public"."business"."keyword" IS '变量标识';
COMMENT ON COLUMN "public"."business"."type" IS '变量类型';
COMMENT ON COLUMN "public"."business"."description" IS '变量描述';
COMMENT ON TABLE "public"."business" IS '业务变量';

-- ----------------------------
-- Records of business
-- ----------------------------
INSERT INTO "public"."business" VALUES (1, 1712995716, 1712995716, 2, 'AuthSkipRoles', 'object', '跳过权限检测的角色列表');
INSERT INTO "public"."business" VALUES (2, 1712995716, 1712995716, 2, 'DefaultUserPassword', 'string', '默认账号密码');
INSERT INTO "public"."business" VALUES (4, 1712995716, 1712995716, 2, 'LoginPrivatePath', 'string', 'rsa解密私钥路径');
INSERT INTO "public"."business" VALUES (5, 1712995716, 1712995716, 2, 'Setting', 'object', '系统设置');
INSERT INTO "public"."business" VALUES (6, 1712995716, 1719465188, 3, 'DefaultAcceptTypes', 'object', '允许上传的文件后缀');
INSERT INTO "public"."business" VALUES (8, 1712995716, 1719465209, 3, 'ChunkSize', 'int', '单个切片最大大小（M）');
INSERT INTO "public"."business" VALUES (9, 1712995716, 1719465230, 3, 'DefaultMaxSize', 'int', '文件最大大小（M）');
INSERT INTO "public"."business" VALUES (15, 1719462652, 1719462652, 2, 'ChangePasswordType', 'string', '修改密码方式');
INSERT INTO "public"."business" VALUES (16, 1723029091, 1723029091, 9, 'botSetting', 'object', 'bot设置');
INSERT INTO "public"."business" VALUES (17, 1723029177, 1723029177, 9, 'botId', 'string', '机器人id');
INSERT INTO "public"."business" VALUES (18, 1723029204, 1723029204, 9, 'botKey', 'string', '机器人密钥');

-- ----------------------------
-- Table structure for business_value
-- ----------------------------
DROP TABLE IF EXISTS "public"."business_value";
CREATE TABLE "public"."business_value" (
                                           "id" int8 NOT NULL,
                                           "created_at" int8,
                                           "updated_at" int8,
                                           "env_id" int8 NOT NULL,
                                           "business_id" int8 NOT NULL,
                                           "value" text COLLATE "pg_catalog"."default" NOT NULL
)
;
COMMENT ON COLUMN "public"."business_value"."id" IS '主键ID';
COMMENT ON COLUMN "public"."business_value"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."business_value"."updated_at" IS '修改时间';
COMMENT ON COLUMN "public"."business_value"."env_id" IS '环境id';
COMMENT ON COLUMN "public"."business_value"."business_id" IS '业务变量id';
COMMENT ON COLUMN "public"."business_value"."value" IS '业务变量id';
COMMENT ON TABLE "public"."business_value" IS '业务变量值';

-- ----------------------------
-- Records of business_value
-- ----------------------------
INSERT INTO "public"."business_value" VALUES (52, 1719371012, 1719465259, 1, 6, '["jpg","png","txt","ppt","pptx","mp4","pdf"]');
INSERT INTO "public"."business_value" VALUES (53, 1719371012, 1719465259, 2, 6, '["jpg","png","txt","ppt","pptx","mp4"]');
INSERT INTO "public"."business_value" VALUES (54, 1719371012, 1719465259, 3, 6, '["jpg","png","txt","ppt","pptx","mp4"]');
INSERT INTO "public"."business_value" VALUES (61, 1719382347, 1719465239, 1, 9, '10');
INSERT INTO "public"."business_value" VALUES (62, 1719382347, 1719465239, 2, 9, '10');
INSERT INTO "public"."business_value" VALUES (63, 1719382347, 1719465239, 3, 9, '10');
INSERT INTO "public"."business_value" VALUES (64, 1719382352, 1719382352, 1, 8, '1');
INSERT INTO "public"."business_value" VALUES (65, 1719382352, 1719382352, 2, 8, '1');
INSERT INTO "public"."business_value" VALUES (66, 1719382352, 1719382352, 3, 8, '1');
INSERT INTO "public"."business_value" VALUES (73, 1719462662, 1719462662, 1, 15, 'password');
INSERT INTO "public"."business_value" VALUES (74, 1719462662, 1719462662, 2, 15, 'password');
INSERT INTO "public"."business_value" VALUES (75, 1719462662, 1719462662, 3, 15, 'email');
INSERT INTO "public"."business_value" VALUES (76, 1719463071, 1719463071, 1, 4, 'static/cert/login.pem');
INSERT INTO "public"."business_value" VALUES (77, 1719463071, 1719463071, 2, 4, 'static/cert/login.pem');
INSERT INTO "public"."business_value" VALUES (78, 1719463071, 1719463071, 3, 4, 'static/cert/login.pem');
INSERT INTO "public"."business_value" VALUES (79, 1719463092, 1719463092, 1, 1, '["superAdmin"]');
INSERT INTO "public"."business_value" VALUES (80, 1719463092, 1719463092, 2, 1, '["superAdmin"]');
INSERT INTO "public"."business_value" VALUES (81, 1719463092, 1719463092, 3, 1, '["superAdmin"]');
INSERT INTO "public"."business_value" VALUES (82, 1719463137, 1719463137, 1, 2, '12345678');
INSERT INTO "public"."business_value" VALUES (83, 1719463137, 1719463137, 2, 2, '12345678');
INSERT INTO "public"."business_value" VALUES (84, 1719463137, 1719463137, 3, 2, '12345678');
INSERT INTO "public"."business_value" VALUES (85, 1719463217, 1719463217, 1, 5, '{"logo":"http://p3-armor.byteimg.com/tos-cn-i-49unhts6dw/dfdba5317c0c20ce20e64fac803d52bc.svg~tplv-49unhts6dw-image.image","watermark":"go-platform","title":"统一应用管理平台","desc":"开放协作，拥抱未来，统一应用管理平台","copyright":"Copyright © 2024 lime.qlime.cn. All rights reserved."}');
INSERT INTO "public"."business_value" VALUES (86, 1719463217, 1719463217, 2, 5, '{"title":"统一应用管理平台","desc":"开放协作，拥抱未来，统一应用管理平台","copyright":"Copyright © 2024 lime.qlime.cn. All rights reserved.","logo":"http://p3-armor.byteimg.com/tos-cn-i-49unhts6dw/dfdba5317c0c20ce20e64fac803d52bc.svg~tplv-49unhts6dw-image.image","watermark":"go-platform"}');
INSERT INTO "public"."business_value" VALUES (87, 1719463217, 1719463217, 3, 5, '{"title":"统一应用管理平台","desc":"开放协作，拥抱未来，统一应用管理平台","copyright":"Copyright © 2024 lime.qlime.cn. All rights reserved.","logo":"http://p3-armor.byteimg.com/tos-cn-i-49unhts6dw/dfdba5317c0c20ce20e64fac803d52bc.svg~tplv-49unhts6dw-image.image","watermark":"go-platform"}');
INSERT INTO "public"."business_value" VALUES (94, 1723029160, 1728496834, 1, 16, '{"logo":"http://p3-armor.byteimg.com/tos-cn-i-49unhts6dw/dfdba5317c0c20ce20e64fac803d52bc.svg~tplv-49unhts6dw-image.image","desc":"你好，同学。我是负责贫困资助的AI助手，请问你有什么想要咨询的问题吗？","guiding":[{"type":"question","text":"你好，同学"}],"name":"我的AI"}');
INSERT INTO "public"."business_value" VALUES (95, 1723029160, 1728496834, 2, 16, '{"logo":"http://p3-armor.byteimg.com/tos-cn-i-49unhts6dw/dfdba5317c0c20ce20e64fac803d52bc.svg~tplv-49unhts6dw-image.image","desc":"你好，同学。我是负责贫困资助的AI助手，请问你有什么想要咨询的问题吗？","guiding":[{"type":"question","text":"你好，同学"}],"name":"我的AI"}');
INSERT INTO "public"."business_value" VALUES (96, 1723029160, 1728496834, 3, 16, '{"name":"贫困补助AI助手","logo":"https://gimg3.baidu.com/topone/src=https%3A%2F%2Fbkimg.cdn.bcebos.com%2Fsmart%2F8326cffc1e178a82b90138118155648da97739124271-bkimg-process%2Cv_1%2Crw_1%2Crh_1%2Cmaxl_800%2Cpad_1%3Fx-bce-process%3Dimage%2Fresize%2Cm_pad%2Cw_348%2Ch_348%2Ccolor_ffffff\u0026refer=http%3A%2F%2Fwww.baidu.com\u0026app=2011\u0026size=w931\u0026n=0\u0026g=0n\u0026er=404\u0026q=75\u0026fmt=auto\u0026maxorilen2heic=2000000?sec=1728666000\u0026t=fe09e144ea94ccedec88400db060be55","desc":"你好，同学。我是负责贫困资助的AI助手，请问你有什么想要咨询的问题吗？","guiding":[{"type":"question","text":"你好，同学"}]}');
INSERT INTO "public"."business_value" VALUES (97, 1723029192, 1723029192, 1, 17, '8f0ad162-7cc2-4c5a-8dc8-60434d4a27d9');
INSERT INTO "public"."business_value" VALUES (98, 1723029192, 1723029192, 2, 17, '8f0ad162-7cc2-4c5a-8dc8-60434d4a27d9');
INSERT INTO "public"."business_value" VALUES (99, 1723029192, 1723029192, 3, 17, '8f0ad162-7cc2-4c5a-8dc8-60434d4a27d9');
INSERT INTO "public"."business_value" VALUES (100, 1723029215, 1723029215, 1, 18, 'bce-v3/ALTAK-vXBk2DvNrk0ashSpsZ2jc/ee04a3d76a9dd2325a46e074c5a4c7bc1ebf493d');
INSERT INTO "public"."business_value" VALUES (101, 1723029215, 1723029215, 2, 18, 'bce-v3/ALTAK-vXBk2DvNrk0ashSpsZ2jc/ee04a3d76a9dd2325a46e074c5a4c7bc1ebf493d');
INSERT INTO "public"."business_value" VALUES (102, 1723029215, 1723029215, 3, 18, 'bce-v3/ALTAK-vXBk2DvNrk0ashSpsZ2jc/ee04a3d76a9dd2325a46e074c5a4c7bc1ebf493d');

-- ----------------------------
-- Table structure for configure
-- ----------------------------
DROP TABLE IF EXISTS "public"."configure";
CREATE TABLE "public"."configure" (
                                      "id" int8 NOT NULL,
                                      "created_at" int8,
                                      "updated_at" int8,
                                      "server_id" int8 NOT NULL,
                                      "env_id" int8 NOT NULL,
                                      "content" text COLLATE "pg_catalog"."default" NOT NULL,
                                      "version" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
                                      "format" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
                                      "description" varchar(128) COLLATE "pg_catalog"."default"
)
;
COMMENT ON COLUMN "public"."configure"."id" IS '主键ID';
COMMENT ON COLUMN "public"."configure"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."configure"."updated_at" IS '修改时间';
COMMENT ON COLUMN "public"."configure"."server_id" IS '服务id';
COMMENT ON COLUMN "public"."configure"."env_id" IS '环境id';
COMMENT ON COLUMN "public"."configure"."content" IS '配置内容';
COMMENT ON COLUMN "public"."configure"."version" IS '配置版本';
COMMENT ON COLUMN "public"."configure"."format" IS '配置格式';
COMMENT ON COLUMN "public"."configure"."description" IS '配置描述';
COMMENT ON TABLE "public"."configure" IS '配置内容';

-- ----------------------------
-- Records of configure
-- ----------------------------
INSERT INTO "public"."configure" VALUES (1, 1712995716, 1750766734, 1, 1, 'addr: 0.0.0.0:7080
name: gateway
version: v1
middlewares:
  - name: bbr
  - name: cors
    options:
      allowCredentials: true
      allowOrigins:
        - ''*''
      allowMethods:
        - GET
        - POST
        - PUT
        - DELETE
        - OPTIONS
      AllowHeaders:
        - Content-Type
        - Content-Length
        - Authorization
      ExposeHeaders:
        - Content-Length
        - Access-Control-Allow-Headers
  - name: tracing
  - name: logging
  - name: transcoder
endpoints:
  - path: /manager/api/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7010
  - path: /manager/client/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7010
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/application/client/v1/auth
          method: POST
  - path: /configure/api/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:6081
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/manager/api/v1/auth
          method: POST
  - path: /resource/api/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7020
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/manager/api/v1/auth
          method: POST
          whiteList:
            - path: /resource/api/v1/static/*
              method: GET
            - path: /resource/api/v1/download/*
              method: GET
  - path: /resource/client/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7020
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/application/client/v1/auth
          method: POST
  - path: /application/api/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7030
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/manager/api/v1/auth
          method: POST
  - path: /application/client/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7030
  - path: /cron/api/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7040
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/api/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7100
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/client/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7100
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/application/client/v1/auth
          method: POST
  - path: /poverty/api/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7120
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/manager/api/v1/auth
          method: POST
  - path: /poverty/client/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7120
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/application/client/v1/auth
          method: POST', 'c11287e778a39783b8f7289777d202cc', 'yaml', '测试环境调整');
INSERT INTO "public"."configure" VALUES (2, 1712995716, 1750762888, 2, 1, 'test: 11
env: TEST
server:
  http:
    host: 127.0.0.1
    port: 7010
    timeout: 10s
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: 127.0.0.1
    port: 8010
    timeout: 10s
log:
  level: 0
  caller: true
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: mysql #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: root
      password: 123456
      host: 192.168.110.102
      port: 3306
      dbName: manager
      option: ?charset=utf8mb4&parseTime=True&loc=Local
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: 192.168.110.102:6379
    username:
    password: xj2023
captcha:
  login:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 240
    skew: 0.7
    refresh: true
    dotCount: 80
  changePassword:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: captcha
loader:
  login: static/cert/login.pem
email:
  template:
    captcha:
      subject: 验证码发送通知
      path: static/template/email/default.html
      from: 统一应用管理中心
      type: text/html
  user: 860808187@qq.com
  name: 青岑云
  host: smtp.qq.com
  port: 25
  password: fyudafdzqmhwbfbd
jwt:
  redis: cache
  secret: limes-cloud-test
  expire: 2h
  renewal: 600s
  whitelist: {"POST:/manager/api/v1/user/login":true,"POST:/manager/api/v1/user/logout":true,"POST:/manager/api/v1/user/token/refresh":true,"GET:/manager/api/v1/system/setting":true,"GET:/manager/api/v1/user/login/captcha":true}
authentication:
  db: system
  redis: cache
  roleKey: roleKeyword
  skipRole: ["superAdmin"]
client:
  - server: Resource
    type: direct
    backends:
      - target: 127.0.0.1:8020
business:
  changePasswordType: password
  defaultUserPassword: 12345678
  setting: {"logo":"http://p3-armor.byteimg.com/tos-cn-i-49unhts6dw/dfdba5317c0c20ce20e64fac803d52bc.svg~tplv-49unhts6dw-image.image","watermark":"go-platform","title":"统一应用管理平台","desc":"开放协作，拥抱未来，统一应用管理平台","copyright":"Copyright © 2024 lime.qlime.cn. All rights reserved."}
', 'bed099bd79102d04749c7c54f59241c8', 'yaml', '修改redis配置');
INSERT INTO "public"."configure" VALUES (3, 1712995716, 1719466309, 3, 1, '
env: TEST
server:
  http:
    host: 127.0.0.1
    port: 7020
    timeout: 10s
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: 127.0.0.1
    port: 8020
    timeout: 10s
log:
  level: 0
  caller: true
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: mysql #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: root
      password: root
      host: 127.0.0.1
      port: 3306
      dbName: resource
      option: ?charset=utf8mb4&parseTime=True&loc=Local
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 3 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: 127.0.0.1:6379
    username:
    password:
business:
  chunkSize: 1
  defaultMaxSize: 10
  defaultAcceptTypes: ["jpg","png","txt","ppt","pptx","mp4","pdf"]
  storage:
    type: local
    serverUrl: http://127.0.0.1:7080/resource/api/v1/static
    localDir: static
    temporaryExpire: 600s
    secret: limescloud
  export:
    serverUrl: http://127.0.0.1:7080/resource/api/v1/download
    localDir: static/export
    expire: 72h
', '0266de2d040fa6a857c691196eee79dc', 'yaml', '初始化模板');
INSERT INTO "public"."configure" VALUES (4, 1712995716, 1720269927, 5, 1, '
env: TEST
server:
  http:
    host: 127.0.0.1
    port: 7030
    timeout: 10s
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: 127.0.0.1
    port: 8030
    timeout: 10s
client:
  - server: Resource
    type: direct
    backends:
      - target: 127.0.0.1:8020
log:
  level: 0
  caller: true
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: mysql #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: root
      password: root
      host: 127.0.0.1
      port: 3306
      dbName: usercenter
      option: ?charset=utf8mb4&parseTime=True&loc=Local
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: 127.0.0.1:6379
    username:
    password:
captcha:
  loginImage:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 200
    skew: 0.7
    dotCount: 80
    refresh: true
  bindImage:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 200
    skew: 0.7
    dotCount: 80
    refresh: true
  registerImage:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 200
    skew: 0.7
    dotCount: 80
    refresh: true
  loginEmail:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: login
  bindEmail:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: bind
  registerEmail:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: register
loader:
  password: static/cert/password.pem
email:
  template:
    login:
      subject: 登录验证码发送通知
      path: static/template/email/default.html
      type: text/html
    bind:
      subject: 绑定验证码发送通知
      path: static/template/email/default.html
      type: text/html
    register:
      subject: 注册验证码发送通知
      path: static/template/email/default.html
      type: text/html
  user: 860808187@qq.com
  name: 青岑云
  host: smtp.qq.com
  port: 25
  password: fyudafdzqmhwbfbd
jwt:
  redis: cache
  secret: limes-cloud-client-test
  expire: 2h
  renewal: 600s
  whitelist: {"*:/usercenter/api/*":true,"GET:/usercenter/client/v1/app":true,"POST:/usercenter/client/v1/auth":true,"POST:/usercenter/client/v1/auth/captcha":true,"POST:/usercenter/client/v1/bind/*":true,"POST:/usercenter/client/v1/login/*":true,"POST:/usercenter/client/v1/logout":true,"POST:/usercenter/client/v1/register/*":true,"POST:/usercenter/client/v1/token/refresh":true}
business:
  defaultPasswordExpire: 10s
  defaultNickName: 用户-
', '7c27fdf3d7414e04b5e6c8d81365c6f5', 'yaml', '1');
INSERT INTO "public"."configure" VALUES (5, 1712995716, 1750772374, 6, 1, '
env: TEST
server:
  http:
    host: 127.0.0.1
    port: 7100
    timeout: 10s
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: 127.0.0.1
    port: 8100
    timeout: 10s
client:
  - server: Resource
    type: direct
    backends:
      - target: 127.0.0.1:8020
  - server: Application
    type: direct
    backends:
      - target: 127.0.0.1:8030
log:
  level: 0
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: mysql #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: root
      password: root
      host: 127.0.0.1
      port: 3306
      dbName: party_affairs
      option: ?charset=utf8mb4&parseTime=True&loc=Local
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
', '950eac12032a266c38b8737acfb3df9b', 'yaml', '启用');
INSERT INTO "public"."configure" VALUES (6, 1712995716, 1719557335, 1, 2, 'addr: 0.0.0.0:7080
name: gateway
version: v1
middlewares:
  - name: bbr
  - name: cors
    options:
      allowCredentials: true
      allowOrigins:
        - ''*''
      allowMethods:
        - GET
        - POST
        - PUT
        - DELETE
        - OPTIONS
      AllowHeaders:
        - Content-Type
        - Content-Length
        - Authorization
      ExposeHeaders:
        - Content-Length
        - Access-Control-Allow-Headers
  - name: tracing
  - name: logging
  - name: transcoder
endpoints:
  - path: /manager/api/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7010
  - path: /manager/client/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7010
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/usercenter/client/v1/auth
          method: POST
  - path: /configure/api/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:6081
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/manager/api/v1/auth
          method: POST
  - path: /resource/api/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7020
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/manager/api/v1/auth
          method: POST
          whiteList:
            - path: /resource/api/v1/static/*
              method: GET
            - path: /resource/api/v1/download/*
              method: GET
  - path: /resource/client/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7020
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/usercenter/api/v1/auth
          method: POST
  - path: /usercenter/api/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7030
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/manager/api/v1/auth
          method: POST
  - path: /usercenter/client/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7030
  - path: /cron/api/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7040
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/api/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7100
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/client/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7100
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/usercenter/api/v1/auth
          method: POST', 'e25350369a0436f3af879333dcb7cbe5', 'yaml', '初始化模板');
INSERT INTO "public"."configure" VALUES (7, 1712995716, 1719464076, 2, 2, 'test: 11
env: PRE
server:
  http:
    host: 127.0.0.1
    port: 7010
    timeout: 10s
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: 127.0.0.1
    port: 8010
    timeout: 10s
log:
  level: 0
  caller: true
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: mysql #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: root
      password: root
      host: 127.0.0.1
      port: 3306
      dbName: manager
      option: ?charset=utf8mb4&parseTime=True&loc=Local
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: 127.0.0.1:6379
    username:
    password:
captcha:
  login:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 240
    skew: 0.7
    refresh: true
    dotCount: 80
  changePassword:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: captcha
loader:
  login: static/cert/login.pem
email:
  template:
    captcha:
      subject: 验证码发送通知
      path: static/template/email/default.html
      from: 统一应用管理中心
      type: text/html
  user: 860808187@qq.com
  name: 青岑云
  host: smtp.qq.com
  port: 25
  password: xxx
jwt:
  redis: cache
  secret: limes-cloud-pre
  expire: 2h
  renewal: 600s
  whitelist: {"GET:/manager/api/v1/system/setting":true,"GET:/manager/api/v1/user/login/captcha":true,"POST:/manager/api/v1/user/login":true,"POST:/manager/api/v1/user/logout":true,"POST:/manager/api/v1/user/token/refresh":true}
authentication:
  db: system
  redis: cache
  roleKey: roleKeyword
  skipRole: ["superAdmin"]
client:
  - server: Resource
    type: direct
    backends:
      - target: 127.0.0.1:8010
business:
  changePasswordType: password
  defaultUserPassword: 12345678
  setting: {"title":"统一应用管理平台","desc":"开放协作，拥抱未来，统一应用管理平台","copyright":"Copyright © 2024 lime.qlime.cn. All rights reserved.","logo":"http://p3-armor.byteimg.com/tos-cn-i-49unhts6dw/dfdba5317c0c20ce20e64fac803d52bc.svg~tplv-49unhts6dw-image.image","watermark":"go-platform"}
', '1d2555ef65b788c69f3254a92719cc0b', 'yaml', '初始化模板');
INSERT INTO "public"."configure" VALUES (8, 1712995716, 1719467313, 3, 2, '
env: PRE
server:
  http:
    host: 127.0.0.1
    port: 7020
    timeout: 10s
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: 127.0.0.1
    port: 8020
    timeout: 10s
log:
  level: 0
  caller: true
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: mysql #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: root
      password: root
      host: 127.0.0.1
      port: 3306
      dbName: resource
      option: ?charset=utf8mb4&parseTime=True&loc=Local
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 3 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: 127.0.0.1:6379
    username:
    password:
business:
  chunkSize: 1
  defaultMaxSize: 10
  defaultAcceptTypes: ["jpg","png","txt","ppt","pptx","mp4"]
  storage:
    type: local
    serverUrl: http://127.0.0.1:7080/resource/api/v1/static
    localDir: static
    temporaryExpire: 600s
    secret: limescloud
  export:
    serverUrl: http://127.0.0.1:7080/resource/api/v1/download
    localDir: static/export
    expire: 72h
', '84cdc6dbc3e7e5d04af6f5f1da605013', 'yaml', '初始化模板');
INSERT INTO "public"."configure" VALUES (9, 1712995716, 1719474377, 5, 2, '
env: PRE
server:
  http:
    host: 127.0.0.1
    port: 7030
    timeout: 10s
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: 127.0.0.1
    port: 8030
    timeout: 10s
client:
  - server: Resource
    type: direct
    backends:
      - target: 127.0.0.1:8020
log:
  level: 0
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: mysql #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: root
      password: root
      host: 127.0.0.1
      port: 3306
      dbName: usercenter
      option: ?charset=utf8mb4&parseTime=True&loc=Local
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: 127.0.0.1:6379
    username:
    password:
captcha:
  loginImage:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 200
    skew: 0.7
    dotCount: 80
    refresh: true
  bindImage:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 200
    skew: 0.7
    dotCount: 80
    refresh: true
  registerImage:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 200
    skew: 0.7
    dotCount: 80
    refresh: true
  loginEmail:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: login
  bindEmail:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: bind
  registerEmail:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: register
loader:
  password: static/cert/password.pem
email:
  template:
    login:
      subject: 登录验证码发送通知
      path: static/template/email/default.html
      type: text/html
    bind:
      subject: 绑定验证码发送通知
      path: static/template/email/default.html
      type: text/html
    register:
      subject: 注册验证码发送通知
      path: static/template/email/default.html
      type: text/html
  user: 860808187@qq.com
  name: 青岑云
  host: smtp.qq.com
  port: 25
  password: xxx
jwt:
  redis: cache
  secret: limes-cloud-client-pre
  expire: 2h
  renewal: 600s
  whitelist: {"*:/usercenter/api/*":true,"*:/usercenter/client/*":true}
business:
  service:
    resource: 127.0.0.1:8020
', '0c3387a6828a72a708b98085b7b45d71', 'yaml', '初始化配置');
INSERT INTO "public"."configure" VALUES (10, 1712995716, 1712995716, 6, 2, '
env: PRE
server:
  http:
    host: 127.0.0.1
    port: 7100
    timeout: 10s
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: 127.0.0.1
    port: 8100
    timeout: 10s
client:
  - server: Resource
    type: direct
    backends:
      - target: 127.0.0.1:8003
  - server: UserCenter
    type: direct
    backends:
      - target: 127.0.0.1:8004
log:
  level: 0
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: mysql #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: root
      password: root
      host: 127.0.0.1
      port: 3306
      dbName: party_affairs
      option: ?charset=utf8mb4&parseTime=True&loc=Local
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
business:
  auth:
    yiBan:
      appId: e4750b34230b48e1
      appSecret: b0891a7f6018e5a76b085e3cb9548edd
', '6311dada0a48b1d307068a92da31a3f1', 'yaml', '自动初始化');
INSERT INTO "public"."configure" VALUES (11, 1712995716, 1729189107, 1, 3, 'addr: 0.0.0.0:7080
name: gateway
version: v1
middlewares:
  - name: bbr
  - name: cors
    options:
      allowCredentials: true
      allowOrigins:
        - ''*''
      allowMethods:
        - GET
        - POST
        - PUT
        - DELETE
        - OPTIONS
      AllowHeaders:
        - Content-Type
        - Content-Length
        - Authorization
      ExposeHeaders:
        - Content-Length
        - Access-Control-Allow-Headers
  - name: tracing
  - name: logging
  - name: transcoder
endpoints:
  - path: /manager/api/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7010
  - path: /manager/client/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7010
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/application/client/v1/auth
          method: POST
  - path: /configure/api/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:6081
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/manager/api/v1/auth
          method: POST
  - path: /resource/api/*
    timeout: 600s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7020
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/manager/api/v1/auth
          method: POST
          whiteList:
            - path: /resource/api/v1/static/*
              method: GET
            - path: /resource/api/v1/download/*
              method: GET
  - path: /resource/client/*
    timeout: 600s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7020
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/application/client/v1/auth
          method: POST
  - path: /application/api/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7030
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/manager/api/v1/auth
          method: POST
  - path: /application/client/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7030
  - path: /cron/api/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7040
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/api/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7100
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/client/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7100
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/application/client/v1/auth
          method: POST
  - path: /poverty/api/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7120
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/manager/api/v1/auth
          method: POST
  - path: /poverty/client/*
    timeout: 10s
    protocol: HTTP
    responseFormat: true
    backends:
      - target: 127.0.0.1:7120
    middlewares:
      - name: auth
        options:
          url: http://localhost:7080/application/client/v1/auth
          method: POST', '66164cc6f6d0a50cdfb1035b6672a706', 'yaml', '1');
INSERT INTO "public"."configure" VALUES (12, 1712995716, 1720364177, 2, 3, 'test: 11
env: PROD
server:
  http:
    host: 127.0.0.1
    port: 7010
    timeout: 10s
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: 127.0.0.1
    port: 8010
    timeout: 10s
log:
  level: 0
  caller: true
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: mysql #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: manager
      password: L8hjTy5GMZmdHJX3
      host: 127.0.0.1
      port: 3306
      dbName: manager
      option: ?charset=utf8mb4&parseTime=True&loc=Local
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: 127.0.0.1:6379
    username:
    password:
captcha:
  login:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 240
    skew: 0.7
    refresh: true
    dotCount: 80
  changePassword:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: captcha
loader:
  login: static/cert/login.pem
email:
  template:
    captcha:
      subject: 验证码发送通知
      path: static/template/email/default.html
      from: 统一应用管理中心
      type: text/html
  user: 860808187@qq.com
  name: 青岑云
  host: smtp.qq.com
  port: 25
  password: fyudafdzqmhwbfbd
jwt:
  redis: cache
  secret: limes-cloud-prod
  expire: 2h
  renewal: 600s
  whitelist: {"GET:/manager/api/v1/system/setting":true,"GET:/manager/api/v1/user/login/captcha":true,"POST:/manager/api/v1/user/login":true,"POST:/manager/api/v1/user/logout":true,"POST:/manager/api/v1/user/token/refresh":true}
authentication:
  db: system
  redis: cache
  roleKey: roleKeyword
  skipRole: ["superAdmin"]
client:
  - server: Resource
    type: direct
    backends:
      - target: 127.0.0.1:8020
business:
  changePasswordType: email
  defaultUserPassword: 12345678
  setting: {"title":"统一应用管理平台","desc":"开放协作，拥抱未来，统一应用管理平台","copyright":"Copyright © 2024 lime.qlime.cn. All rights reserved.","logo":"http://p3-armor.byteimg.com/tos-cn-i-49unhts6dw/dfdba5317c0c20ce20e64fac803d52bc.svg~tplv-49unhts6dw-image.image","watermark":"go-platform"}
', '42462633dd9a6685e39fb3572c9fd01e', 'yaml', '初始化模板');
INSERT INTO "public"."configure" VALUES (13, 1712995716, 1729189119, 3, 3, '
env: PROD
server:
  http:
    host: 127.0.0.1
    port: 7020
    timeout: 600s
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: 127.0.0.1
    port: 8020
    timeout: 600s
log:
  level: 0
  caller: true
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: mysql #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: resource
      password: ct7AYfaM8kc8LWHi
      host: 127.0.0.1
      port: 3306
      dbName: resource
      option: ?charset=utf8mb4&parseTime=True&loc=Local
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: 127.0.0.1:6379
    username:
    password:
business:
  chunkSize: 1
  defaultMaxSize: 10
  defaultAcceptTypes: ["jpg","png","txt","ppt","pptx","mp4"]
  storage:
    type: local
    serverUrl: https://prod-gw.qlime.cn/resource/api/v1/static
    localDir: static
    temporaryExpire: 600s
    secret: limescloud
  export:
    serverUrl: https://prod-gw.qlime.cn/resource/api/v1/download
    localDir: static/export
    expire: 72h
', 'de69d97e91f4e71bfdc3ad1e63f85976', 'yaml', '1');
INSERT INTO "public"."configure" VALUES (14, 1712995716, 1728067876, 5, 3, '
env: PROD
server:
  http:
    host: 127.0.0.1
    port: 7030
    timeout: 10s
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: 127.0.0.1
    port: 8030
    timeout: 10s
client:
  - server: Resource
    type: direct
    backends:
      - target: 127.0.0.1:8020
log:
  level: 0
  caller: true
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: mysql #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: usercenter
      password: eZEhkfYZRKRY5kYk
      host: 127.0.0.1
      port: 3306
      dbName: usercenter
      option: ?charset=utf8mb4&parseTime=True&loc=Local
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: 127.0.0.1:6379
    username:
    password:
captcha:
  loginImage:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 200
    skew: 0.7
    dotCount: 80
    refresh: true
  bindImage:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 200
    skew: 0.7
    dotCount: 80
    refresh: true
  registerImage:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 200
    skew: 0.7
    dotCount: 80
    refresh: true
  loginEmail:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: login
  bindEmail:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: bind
  registerEmail:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: register
loader:
  password: static/cert/password.pem
email:
  template:
    login:
      subject: 登录验证码发送通知
      path: static/template/email/default.html
      type: text/html
    bind:
      subject: 绑定验证码发送通知
      path: static/template/email/default.html
      type: text/html
    register:
      subject: 注册验证码发送通知
      path: static/template/email/default.html
      type: text/html
  user: 860808187@qq.com
  name: 青岑云
  host: smtp.qq.com
  port: 25
  password: fyudafdzqmhwbfbd
jwt:
  redis: cache
  secret: limes-cloud-client-prod
  expire: 2h
  renewal: 600s
  whitelist: {"*:/application/api/*":true,"GET:/application/client/v1/app":true,"POST:/application/client/v1/auth/captcha":true,"POST:/application/client/v1/bind/email":true,"POST:/application/client/v1/bind/password":true,"POST:/application/client/v1/login/email":true,"POST:/application/client/v1/login/oauth":true,"POST:/application/client/v1/login/password":true,"POST:/application/client/v1/register/email":true,"POST:/application/client/v1/register/password":true,"POST:/application/client/v1/token/refresh":true}
business:
  defaultPasswordExpire: 10s
  defaultNickName: 用户-
', '62d335b1dae6a68c8109449502fe82db', 'yaml', '新增应用中心配置');
INSERT INTO "public"."configure" VALUES (15, 1712995716, 1728071385, 6, 3, '
env: PROD
server:
  http:
    host: 127.0.0.1
    port: 7100
    timeout: 10s
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: 127.0.0.1
    port: 8100
    timeout: 10s
client:
  - server: Resource
    type: direct
    backends:
      - target: 127.0.0.1:8020
  - server: Application
    type: direct
    backends:
      - target: 127.0.0.1:8030
log:
  level: 0
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: mysql #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: party_affairs
      password: 2pH7MTcaRD3YZWWt
      host: 127.0.0.1
      port: 3306
      dbName: party_affairs
      option: ?charset=utf8mb4&parseTime=True&loc=Local
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
', 'a062a539daf9cc44818fd77687c65fd8', 'yaml', '同步应用中心配置');
INSERT INTO "public"."configure" VALUES (19, 1719635224, 1719719685, 4, 1, '
env: TEST
server:
  http:
    host: 127.0.0.1
    port: 7040
    timeout: 10s
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: 127.0.0.1
    port: 8040
    timeout: 10s
log:
  level: 0
  caller: true
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: true
redis:
  cache:
    enable: true
    host: 127.0.0.1:6379
    username:
    password:
database:
  system:
    enable: true #是否启用数据库
    drive: mysql #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: root
      password: root
      host: 127.0.0.1
      port: 3306
      dbName: cron
      option: ?charset=utf8mb4&parseTime=True&loc=Local
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
pool: #并非池配置
  size: 10000 #最大协程数量
  expiryDuration: 30s #过期时间
  preAlloc: true #是否预分配
  maxBlockingTasks: 1000 #最大的并发任务
  nonblocking: true #设置为true时maxBlockingTasks将失效，不限制并发任务
business:
  key: value
', '4a935e90c678304f74f519825ac5cf77', 'yaml', '初始化模板');
INSERT INTO "public"."configure" VALUES (20, 1720364197, 1720364784, 4, 3, '
env: PROD
server:
  http:
    host: 127.0.0.1
    port: 7040
    timeout: 10s
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: 127.0.0.1
    port: 8040
    timeout: 10s
log:
  level: 0
  caller: true
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: true
redis:
  cache:
    enable: true
    host: 127.0.0.1:6379
    username:
    password:
database:
  system:
    enable: true #是否启用数据库
    drive: mysql #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: cron
      password: 4b7fTCenTjLH5PeL
      host: 127.0.0.1
      port: 3306
      dbName: cron
      option: ?charset=utf8mb4&parseTime=True&loc=Local
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
pool: #并非池配置
  size: 10000 #最大协程数量
  expiryDuration: 30s #过期时间
  preAlloc: true #是否预分配
  maxBlockingTasks: 1000 #最大的并发任务
  nonblocking: true #设置为true时maxBlockingTasks将失效，不限制并发任务
business:
  key: value
', 'ef1f5d6a1f457576911903cf4f562c47', 'yaml', '初始化配置');
INSERT INTO "public"."configure" VALUES (21, 1723029369, 1728496848, 9, 3, '
env: PROD
server:
  http:
    host: 127.0.0.1
    port: 7120
    timeout: 10s
    marshal:
      emitUnpopulated: true
      useProtoNames: true
client:
  - server: Resource
    type: direct
    backends:
      - target: 127.0.0.1:8020
  - server: Application
    type: direct
    backends:
      - target: 127.0.0.1:8030
log:
  level: 0
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: mysql #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: poverty
      password: i22MzdthBnF5N2S3
      host: 127.0.0.1
      port: 3306
      dbName: poverty
      option: ?charset=utf8mb4&parseTime=True&loc=Local
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 100 #最大连接数量
      maxIdleConn: 50 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
pool: #并非池配置
  size: 1000 #最大协程数量
  expiryDuration: 30s #过期时间
  preAlloc: true #是否预分配
  maxBlockingTasks: 1000 #最大的并发任务
  nonblocking: true
business:
  aiBot:
    secret: bce-v3/ALTAK-vXBk2DvNrk0ashSpsZ2jc/ee04a3d76a9dd2325a46e074c5a4c7bc1ebf493d
    appid: 8f0ad162-7cc2-4c5a-8dc8-60434d4a27d9
    setting: {"name":"贫困补助AI助手","logo":"https://gimg3.baidu.com/topone/src=https%3A%2F%2Fbkimg.cdn.bcebos.com%2Fsmart%2F8326cffc1e178a82b90138118155648da97739124271-bkimg-process%2Cv_1%2Crw_1%2Crh_1%2Cmaxl_800%2Cpad_1%3Fx-bce-process%3Dimage%2Fresize%2Cm_pad%2Cw_348%2Ch_348%2Ccolor_ffffff\u0026refer=http%3A%2F%2Fwww.baidu.com\u0026app=2011\u0026size=w931\u0026n=0\u0026g=0n\u0026er=404\u0026q=75\u0026fmt=auto\u0026maxorilen2heic=2000000?sec=1728666000\u0026t=fe09e144ea94ccedec88400db060be55","desc":"你好，同学。我是负责贫困资助的AI助手，请问你有什么想要咨询的问题吗？","guiding":[{"type":"question","text":"你好，同学"}]}
', '1b15c3b73aa4e18b21a64b3e89dc46bf', 'yaml', '修改setting');

-- ----------------------------
-- Table structure for env
-- ----------------------------
DROP TABLE IF EXISTS "public"."env";
CREATE TABLE "public"."env" (
                                "id" int8 NOT NULL,
                                "created_at" int8,
                                "updated_at" int8,
                                "keyword" char(32) COLLATE "pg_catalog"."default" NOT NULL,
                                "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
                                "description" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
                                "token" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
                                "status" boolean
)
;
COMMENT ON COLUMN "public"."env"."id" IS '主键ID';
COMMENT ON COLUMN "public"."env"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."env"."updated_at" IS '修改时间';
COMMENT ON COLUMN "public"."env"."keyword" IS '环境标识';
COMMENT ON COLUMN "public"."env"."name" IS '环境名称';
COMMENT ON COLUMN "public"."env"."description" IS '环境描述';
COMMENT ON COLUMN "public"."env"."token" IS '环境token';
COMMENT ON COLUMN "public"."env"."status" IS '环境状态';
COMMENT ON TABLE "public"."env" IS '环境信息';

-- ----------------------------
-- Records of env
-- ----------------------------
INSERT INTO "public"."env" VALUES (1, 1712995716, 1719369848, 'TEST', '测试环境', '用于本地测试', '1025D32F6CA7A2A320FE091B22C5DF3C', true);
INSERT INTO "public"."env" VALUES (2, 1712995716, 1750762479, 'PRE', '预发布环境', '用于上线前测试', '862BBE3D5BE34A780305DA84A8DD5147', false);
INSERT INTO "public"."env" VALUES (3, 1712995716, 1750762481, 'PROD', '生产环境', '用于线上真实环境', '5B655B7D4A51BF79C974C9F27C27D992', false);

-- ----------------------------
-- Table structure for gorm_init
-- ----------------------------
DROP TABLE IF EXISTS "public"."gorm_init";
CREATE TABLE "public"."gorm_init" (
                                      "id" int8 NOT NULL,
                                      "init" boolean NOT NULL DEFAULT 0
)
;


-- ----------------------------
-- Records of gorm_init
-- ----------------------------
INSERT INTO "public"."gorm_init" VALUES (1, true);

-- ----------------------------
-- Table structure for resource
-- ----------------------------
DROP TABLE IF EXISTS "public"."resource";
CREATE TABLE "public"."resource" (
                                     "id" int8 NOT NULL,
                                     "created_at" int8,
                                     "updated_at" int8,
                                     "keyword" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
                                     "description" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
                                     "fields" varchar(256) COLLATE "pg_catalog"."default" NOT NULL,
                                     "tag" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
                                     "private" boolean
)
;
COMMENT ON COLUMN "public"."resource"."id" IS '主键ID';
COMMENT ON COLUMN "public"."resource"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."resource"."updated_at" IS '修改时间';
COMMENT ON COLUMN "public"."resource"."keyword" IS '变量标识';
COMMENT ON COLUMN "public"."resource"."description" IS '变量描述';
COMMENT ON COLUMN "public"."resource"."fields" IS '变量字段集合';
COMMENT ON COLUMN "public"."resource"."tag" IS '变量标签';
COMMENT ON COLUMN "public"."resource"."private" IS '是否私有';
COMMENT ON TABLE "public"."resource" IS '资源变量';

-- ----------------------------
-- Records of resource
-- ----------------------------
INSERT INTO "public"."resource" VALUES (1, 1712995716, 1712995716, 'Env', '环境标识信息', 'Keyword', 'env', false);
INSERT INTO "public"."resource" VALUES (2, 1712995716, 1712995716, 'AdminJwt', '后台管理服务jwt配置信息', 'Secret,Expire,Renewal,Whitelist', 'jwt', false);
INSERT INTO "public"."resource" VALUES (3, 1712995716, 1712995716, 'ClientJwt', '客户端服务jwt配置信息', 'Secret,Expire,Renewal,Whitelist', 'jwt', false);
INSERT INTO "public"."resource" VALUES (4, 1712995716, 1712995716, 'Redis', 'redis中间件配置信息', 'Host,Username,Password,Port', 'redis', false);
INSERT INTO "public"."resource" VALUES (5, 1712995716, 1712995716, 'Email', '邮箱服务配置信息', 'Username,Company,Password,Host,Port', 'email', false);
INSERT INTO "public"."resource" VALUES (6, 1712995716, 1712995716, 'GatewayServer', '通用网关服务配置信息', 'HttpPort,Host,Timeout', 'server', false);
INSERT INTO "public"."resource" VALUES (7, 1712995716, 1712995716, 'ManagerServer', '管理中心服务配置信息', 'Host,HttpPort,GrpcPort,Timeout', 'server', false);
INSERT INTO "public"."resource" VALUES (8, 1712995716, 1712995716, 'ManagerDatabase', '管理中心数据库配置', 'Username,Password,Type,Port,Database,Option,Host', 'mysql', false);
INSERT INTO "public"."resource" VALUES (9, 1712995716, 1712995716, 'ResourceServer', '资源中心服务配置信息', 'Host,HttpPort,GrpcPort,Timeout', 'server', false);
INSERT INTO "public"."resource" VALUES (10, 1712995716, 1712995716, 'ResourceDatabase', '资源中心数据库配置信息', 'Username,Password,Type,Port,Database,Option,Host', 'mysql', false);
INSERT INTO "public"."resource" VALUES (11, 1712995716, 1712995716, 'CronServer', '定时任务服务配置信息', 'Host,HttpPort,GrpcPort,Timeout', 'server', false);
INSERT INTO "public"."resource" VALUES (12, 1712995716, 1712995716, 'CronDatabase', '定时任务数据库配置信息', 'Username,Password,Type,Port,Database,Option,Host', 'mysql', false);
INSERT INTO "public"."resource" VALUES (13, 1712995716, 1712995716, 'ConfigureServer', '配置中心服务配置信息', 'Host,HttpPort,GrpcPort,Timeout', 'server', false);
INSERT INTO "public"."resource" VALUES (14, 1712995716, 1728064624, 'ApplicationServer', '应用中心服务配置信息', 'Host,HttpPort,GrpcPort,Timeout', 'server', false);
INSERT INTO "public"."resource" VALUES (15, 1712995716, 1728064632, 'ApplicationDatabase', '应用中心数据库配置信息', 'Username,Password,Type,Port,Database,Option,Host', 'mysql', true);
INSERT INTO "public"."resource" VALUES (16, 1712995716, 1712995716, 'PartyAffairsServer', '信号灯服务配置信息', 'Host,HttpPort,GrpcPort,Timeout', 'server', false);
INSERT INTO "public"."resource" VALUES (17, 1712995716, 1719371833, 'PartyAffairsDatabase', '信号灯数据库配置信息', 'Username,Password,Type,Port,Database,Option,Host', 'mysql', true);
INSERT INTO "public"."resource" VALUES (19, 1723028903, 1723028903, 'PovertyDatabase', '资助系统数据库信息', 'Username,Password,Type,Port,Database,Option,Host', 'mysql', true);
INSERT INTO "public"."resource" VALUES (20, 1723028953, 1723028953, 'PovertyServer', '资助系统服务信息', 'Host,Port,Timeout', 'server', true);

-- ----------------------------
-- Table structure for resource_server
-- ----------------------------
DROP TABLE IF EXISTS "public"."resource_server";
CREATE TABLE "public"."resource_server" (
                                            "id" int8 NOT NULL,
                                            "created_at" int8,
                                            "server_id" int8 NOT NULL,
                                            "resource_id" int8 NOT NULL
)
;
COMMENT ON COLUMN "public"."resource_server"."id" IS '主键ID';
COMMENT ON COLUMN "public"."resource_server"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."resource_server"."server_id" IS '服务id';
COMMENT ON COLUMN "public"."resource_server"."resource_id" IS '资源id';
COMMENT ON TABLE "public"."resource_server" IS '资源服务信息';

-- ----------------------------
-- Records of resource_server
-- ----------------------------
INSERT INTO "public"."resource_server" VALUES (1, 1712995716, 2, 8);
INSERT INTO "public"."resource_server" VALUES (2, 1712995716, 3, 10);
INSERT INTO "public"."resource_server" VALUES (3, 1712995716, 4, 12);
INSERT INTO "public"."resource_server" VALUES (6, NULL, 6, 17);
INSERT INTO "public"."resource_server" VALUES (7, NULL, 9, 19);
INSERT INTO "public"."resource_server" VALUES (8, NULL, 5, 15);

-- ----------------------------
-- Table structure for resource_value
-- ----------------------------
DROP TABLE IF EXISTS "public"."resource_value";
CREATE TABLE "public"."resource_value" (
                                           "id" int8 NOT NULL,
                                           "created_at" int8,
                                           "updated_at" int8,
                                           "env_id" int8 NOT NULL,
                                           "resource_id" int8 NOT NULL,
                                           "value" text COLLATE "pg_catalog"."default" NOT NULL
)
;
COMMENT ON COLUMN "public"."resource_value"."id" IS '主键ID';
COMMENT ON COLUMN "public"."resource_value"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."resource_value"."updated_at" IS '修改时间';
COMMENT ON COLUMN "public"."resource_value"."env_id" IS '环境id';
COMMENT ON COLUMN "public"."resource_value"."resource_id" IS '资源id';
COMMENT ON COLUMN "public"."resource_value"."value" IS '资源id';
COMMENT ON TABLE "public"."resource_value" IS '资源变量值';

-- ----------------------------
-- Records of resource_value
-- ----------------------------
INSERT INTO "public"."resource_value" VALUES (1, 1712995716, 1719371433, 1, 1, '{"Keyword":"TEST"}');
INSERT INTO "public"."resource_value" VALUES (2, 1712995716, 1719371433, 2, 1, '{"Keyword":"PRE"}');
INSERT INTO "public"."resource_value" VALUES (3, 1712995716, 1719371433, 3, 1, '{"Keyword":"PROD"}');
INSERT INTO "public"."resource_value" VALUES (4, 1712995716, 1719559394, 1, 2, '{"Renewal":"600s","Whitelist":{"GET:/manager/api/v1/system/setting":true,"GET:/manager/api/v1/user/login/captcha":true,"POST:/manager/api/v1/user/login":true,"POST:/manager/api/v1/user/logout":true,"POST:/manager/api/v1/user/token/refresh":true},"Secret":"limes-cloud-test","Expire":"2h"}');
INSERT INTO "public"."resource_value" VALUES (5, 1712995716, 1719559394, 2, 2, '{"Secret":"limes-cloud-pre","Expire":"2h","Renewal":"600s","Whitelist":{"GET:/manager/api/v1/user/login/captcha":true,"POST:/manager/api/v1/user/login":true,"POST:/manager/api/v1/user/logout":true,"POST:/manager/api/v1/user/token/refresh":true,"GET:/manager/api/v1/system/setting":true}}');
INSERT INTO "public"."resource_value" VALUES (6, 1712995716, 1719559394, 3, 2, '{"Whitelist":{"POST:/manager/api/v1/user/token/refresh":true,"GET:/manager/api/v1/system/setting":true,"GET:/manager/api/v1/user/login/captcha":true,"POST:/manager/api/v1/user/login":true,"POST:/manager/api/v1/user/logout":true},"Secret":"limes-cloud-prod","Expire":"2h","Renewal":"600s"}');
INSERT INTO "public"."resource_value" VALUES (7, 1712995716, 1728067848, 1, 3, '{"Secret":"limes-cloud-client-test","Expire":"2h","Renewal":"600s","Whitelist":{"POST:/application/client/v1/token/refresh":true,"POST:/application/client/v1/logout":true,"POST:/application/client/v1/register/*":true,"POST:/application/client/v1/bind/*":true,"GET:/application/client/v1/app":true,"POST:/application/client/v1/auth/captcha":true,"POST:/application/client/v1/auth":true,"*:/application/api/*":true,"POST:/application/client/v1/login/*":true}}');
INSERT INTO "public"."resource_value" VALUES (8, 1712995716, 1728067848, 2, 3, '{"Whitelist":{"POST:/application/client/v1/register/email":true,"POST:/application/client/v1/login/oauth":true,"POST:/application/client/v1/login/email":true,"POST:/application/client/v1/register/password":true,"POST:/application/client/v1/bind/password":true,"POST:/application/client/v1/bind/email":true,"POST:/application/client/v1/auth/captcha":true,"*:/application/api/*":true,"POST:/application/client/v1/token/refresh":true,"GET:/application/client/v1/app":true,"POST:/application/client/v1/login/password":true},"Secret":"limes-cloud-client-prod","Expire":"2h","Renewal":"600s"}');
INSERT INTO "public"."resource_value" VALUES (9, 1712995716, 1728067848, 3, 3, '{"Expire":"2h","Renewal":"600s","Whitelist":{"POST:/application/client/v1/bind/email":true,"POST:/application/client/v1/bind/password":true,"POST:/application/client/v1/login/oauth":true,"POST:/application/client/v1/login/password":true,"POST:/application/client/v1/token/refresh":true,"GET:/application/client/v1/app":true,"POST:/application/client/v1/login/email":true,"POST:/application/client/v1/register/password":true,"*:/application/api/*":true,"POST:/application/client/v1/auth/captcha":true,"POST:/application/client/v1/register/email":true},"Secret":"limes-cloud-client-prod"}');
INSERT INTO "public"."resource_value" VALUES (10, 1712995716, 1750762871, 1, 4, '{"Host":"192.168.110.102","Username":"","Password":"xj2023","Port":6379}');
INSERT INTO "public"."resource_value" VALUES (11, 1712995716, 1750762871, 2, 4, '{"Password":"","Port":6379,"Host":"127.0.0.1","Username":""}');
INSERT INTO "public"."resource_value" VALUES (12, 1712995716, 1750762871, 3, 4, '{"Username":"","Password":"","Port":6379,"Host":"127.0.0.1"}');
INSERT INTO "public"."resource_value" VALUES (13, 1712995716, 1714884855, 1, 5, '{"Username":"860808187@qq.com","Company":"青岑云","Password":"fyudafdzqmhwbfbd","Host":"smtp.qq.com","Port":25}');
INSERT INTO "public"."resource_value" VALUES (14, 1712995716, 1714884855, 2, 5, '{"Password":"xxx","Host":"smtp.qq.com","Port":25,"Username":"860808187@qq.com","Company":"青岑云"}');
INSERT INTO "public"."resource_value" VALUES (15, 1712995716, 1714884855, 3, 5, '{"Username":"860808187@qq.com","Company":"青岑云","Password":"fyudafdzqmhwbfbd","Host":"smtp.qq.com","Port":25}');
INSERT INTO "public"."resource_value" VALUES (16, 1712995716, 1712995716, 1, 6, '{"Host":"127.0.0.1","HttpPort":7080,"Timeout":"10s"}');
INSERT INTO "public"."resource_value" VALUES (17, 1712995716, 1712995716, 2, 6, '{"Host":"127.0.0.1","HttpPort":7080,"Timeout":"10s"}');
INSERT INTO "public"."resource_value" VALUES (18, 1712995716, 1712995716, 3, 6, '{"Host":"127.0.0.1","HttpPort":7080,"Timeout":"10s"}');
INSERT INTO "public"."resource_value" VALUES (19, 1712995716, 1719462382, 1, 7, '{"Host":"127.0.0.1","HttpPort":7010,"GrpcPort":8010,"Timeout":"10s"}');
INSERT INTO "public"."resource_value" VALUES (20, 1712995716, 1719462382, 2, 7, '{"Timeout":"10s","Host":"127.0.0.1","HttpPort":7010,"GrpcPort":8010}');
INSERT INTO "public"."resource_value" VALUES (21, 1712995716, 1719462382, 3, 7, '{"Host":"127.0.0.1","HttpPort":7010,"GrpcPort":8010,"Timeout":"10s"}');
INSERT INTO "public"."resource_value" VALUES (22, 1712995716, 1750762765, 1, 8, '{"Option":"?charset=utf8mb4\u0026parseTime=True\u0026loc=Local","Username":"root","Password":"123456","Host":"192.168.110.102","Port":"3306","Type":"mysql","Database":"manager"}');
INSERT INTO "public"."resource_value" VALUES (23, 1712995716, 1750762765, 2, 8, '{"Username":"root","Password":"root","Host":"192.168.110.102","Port":"3306","Type":"mysql","Database":"manager","Option":"?charset=utf8mb4\u0026parseTime=True\u0026loc=Local"}');
INSERT INTO "public"."resource_value" VALUES (24, 1712995716, 1750762765, 3, 8, '{"Password":"L8hjTy5GMZmdHJX3","Host":"192.168.110.102","Port":"3306","Type":"mysql","Database":"manager","Option":"?charset=utf8mb4\u0026parseTime=True\u0026loc=Local","Username":"manager"}');
INSERT INTO "public"."resource_value" VALUES (25, 1712995716, 1729189096, 1, 9, '{"Timeout":"10s","Host":"127.0.0.1","HttpPort":7020,"GrpcPort":8020}');
INSERT INTO "public"."resource_value" VALUES (26, 1712995716, 1729189096, 2, 9, '{"GrpcPort":8020,"Timeout":"10s","Host":"127.0.0.1","HttpPort":7020}');
INSERT INTO "public"."resource_value" VALUES (27, 1712995716, 1729189096, 3, 9, '{"Timeout":"600s","Host":"127.0.0.1","HttpPort":7020,"GrpcPort":8020}');
INSERT INTO "public"."resource_value" VALUES (28, 1712995716, 1720363971, 1, 10, '{"Host":"127.0.0.1","Port":"3306","Type":"mysql","Database":"resource","Option":"?charset=utf8mb4\u0026parseTime=True\u0026loc=Local","Username":"root","Password":"root"}');
INSERT INTO "public"."resource_value" VALUES (29, 1712995716, 1720363971, 2, 10, '{"Type":"mysql","Database":"resource","Option":"?charset=utf8mb4\u0026parseTime=True\u0026loc=Local","Username":"root","Password":"root","Host":"127.0.0.1","Port":"3306"}');
INSERT INTO "public"."resource_value" VALUES (30, 1712995716, 1720363971, 3, 10, '{"Password":"ct7AYfaM8kc8LWHi","Host":"127.0.0.1","Port":"3306","Type":"mysql","Database":"resource","Option":"?charset=utf8mb4\u0026parseTime=True\u0026loc=Local","Username":"resource"}');
INSERT INTO "public"."resource_value" VALUES (31, 1712995716, 1719462436, 1, 11, '{"Host":"127.0.0.1","HttpPort":7040,"GrpcPort":8040,"Timeout":"10s"}');
INSERT INTO "public"."resource_value" VALUES (32, 1712995716, 1719462436, 2, 11, '{"GrpcPort":8040,"Timeout":"10s","Host":"127.0.0.1","HttpPort":7040}');
INSERT INTO "public"."resource_value" VALUES (33, 1712995716, 1719462436, 3, 11, '{"Timeout":"10s","Host":"127.0.0.1","HttpPort":7040,"GrpcPort":8040}');
INSERT INTO "public"."resource_value" VALUES (34, 1712995716, 1720364767, 1, 12, '{"Password":"root","Host":"127.0.0.1","Port":"3306","Type":"mysql","Database":"cron","Option":"?charset=utf8mb4\u0026parseTime=True\u0026loc=Local","Username":"root"}');
INSERT INTO "public"."resource_value" VALUES (35, 1712995716, 1720364767, 2, 12, '{"Option":"?charset=utf8mb4\u0026parseTime=True\u0026loc=Local","Username":"root","Password":"root","Host":"127.0.0.1","Port":"3306","Type":"mysql","Database":"cron"}');
INSERT INTO "public"."resource_value" VALUES (36, 1712995716, 1720364767, 3, 12, '{"Password":"4b7fTCenTjLH5PeL","Host":"127.0.0.1","Port":"3306","Type":"mysql","Database":"cron","Option":"?charset=utf8mb4\u0026parseTime=True\u0026loc=Local","Username":"cron"}');
INSERT INTO "public"."resource_value" VALUES (37, 1712995716, 1712995716, 1, 13, '{"Host":"127.0.0.1","HttpPort":6081,"GrpcPort":6082,"Timeout":"10s"}');
INSERT INTO "public"."resource_value" VALUES (38, 1712995716, 1712995716, 2, 13, '{"Host":"127.0.0.1","HttpPort":6081,"GrpcPort":6082,"Timeout":"10s"}');
INSERT INTO "public"."resource_value" VALUES (39, 1712995716, 1712995716, 3, 13, '{"GrpcPort":6082,"Timeout":"10s","Host":"127.0.0.1","HttpPort":6081}');
INSERT INTO "public"."resource_value" VALUES (40, 1712995716, 1719462419, 1, 14, '{"Timeout":"10s","Host":"127.0.0.1","HttpPort":7030,"GrpcPort":8030}');
INSERT INTO "public"."resource_value" VALUES (41, 1712995716, 1719462419, 2, 14, '{"Host":"127.0.0.1","HttpPort":7030,"GrpcPort":8030,"Timeout":"10s"}');
INSERT INTO "public"."resource_value" VALUES (42, 1712995716, 1719462419, 3, 14, '{"Timeout":"10s","Host":"127.0.0.1","HttpPort":7030,"GrpcPort":8030}');
INSERT INTO "public"."resource_value" VALUES (43, 1712995716, 1720364038, 1, 15, '{"Port":"3306","Type":"mysql","Database":"usercenter","Option":"?charset=utf8mb4\u0026parseTime=True\u0026loc=Local","Username":"root","Password":"root","Host":"127.0.0.1"}');
INSERT INTO "public"."resource_value" VALUES (44, 1712995716, 1720364038, 2, 15, '{"Port":"3306","Type":"mysql","Database":"usercenter","Option":"?charset=utf8mb4\u0026parseTime=True\u0026loc=Local","Username":"root","Password":"root","Host":"127.0.0.1"}');
INSERT INTO "public"."resource_value" VALUES (45, 1712995716, 1720364038, 3, 15, '{"Database":"usercenter","Option":"?charset=utf8mb4\u0026parseTime=True\u0026loc=Local","Username":"usercenter","Password":"eZEhkfYZRKRY5kYk","Host":"127.0.0.1","Port":"3306","Type":"mysql"}');
INSERT INTO "public"."resource_value" VALUES (46, 1712995716, 1712995716, 1, 16, '{"GrpcPort":8100,"Timeout":"10s","Host":"127.0.0.1","HttpPort":7100}');
INSERT INTO "public"."resource_value" VALUES (47, 1712995716, 1712995716, 2, 16, '{"Host":"127.0.0.1","HttpPort":7100,"GrpcPort":8100,"Timeout":"10s"}');
INSERT INTO "public"."resource_value" VALUES (48, 1712995716, 1712995716, 3, 16, '{"HttpPort":7100,"GrpcPort":8100,"Timeout":"10s","Host":"127.0.0.1"}');
INSERT INTO "public"."resource_value" VALUES (49, 1712995716, 1720364104, 1, 17, '{"Type":"mysql","Database":"party_affairs","Option":"?charset=utf8mb4\u0026parseTime=True\u0026loc=Local","Username":"root","Password":"root","Host":"127.0.0.1","Port":"3306"}');
INSERT INTO "public"."resource_value" VALUES (50, 1712995716, 1720364104, 2, 17, '{"Type":"mysql","Database":"party_affairs","Option":"?charset=utf8mb4\u0026parseTime=True\u0026loc=Local","Username":"root","Password":"root","Host":"127.0.0.1","Port":"3306"}');
INSERT INTO "public"."resource_value" VALUES (51, 1712995716, 1720364104, 3, 17, '{"Type":"mysql","Database":"party_affairs","Option":"?charset=utf8mb4\u0026parseTime=True\u0026loc=Local","Username":"party_affairs","Password":"2pH7MTcaRD3YZWWt","Host":"127.0.0.1","Port":"3306"}');
INSERT INTO "public"."resource_value" VALUES (122, 1723029335, 1723029335, 1, 19, '{"Username":"root","Password":"root","Host":"127.0.0.1","Port":"3306","Type":"mysql","Database":"poverty","Option":"?charset=utf8mb4\u0026parseTime=True\u0026loc=Local"}');
INSERT INTO "public"."resource_value" VALUES (123, 1723029335, 1723029335, 2, 19, '{"Type":"mysql","Database":"poverty","Option":"?charset=utf8mb4\u0026parseTime=True\u0026loc=Local","Username":"root","Password":"root","Host":"127.0.0.1","Port":"3306"}');
INSERT INTO "public"."resource_value" VALUES (124, 1723029335, 1723029335, 3, 19, '{"Port":"3306","Type":"mysql","Database":"poverty","Option":"?charset=utf8mb4\u0026parseTime=True\u0026loc=Local","Username":"poverty","Password":"i22MzdthBnF5N2S3","Host":"127.0.0.1"}');
INSERT INTO "public"."resource_value" VALUES (125, 1723029352, 1723029352, 1, 20, '{"Timeout":"10s","Host":"127.0.0.1","Port":7120}');
INSERT INTO "public"."resource_value" VALUES (126, 1723029352, 1723029352, 2, 20, '{"Port":7120,"Timeout":"10s","Host":"127.0.0.1"}');
INSERT INTO "public"."resource_value" VALUES (127, 1723029352, 1723029352, 3, 20, '{"Timeout":"10s","Host":"127.0.0.1","Port":7120}');

-- ----------------------------
-- Table structure for server
-- ----------------------------
DROP TABLE IF EXISTS "public"."server";
CREATE TABLE "public"."server" (
                                   "id" int8 NOT NULL,
                                   "created_at" int8,
                                   "updated_at" int8,
                                   "keyword" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
                                   "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
                                   "description" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
                                   "status" boolean
)
;
COMMENT ON COLUMN "public"."server"."id" IS '主键ID';
COMMENT ON COLUMN "public"."server"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."server"."updated_at" IS '修改时间';
COMMENT ON COLUMN "public"."server"."keyword" IS '服务标识';
COMMENT ON COLUMN "public"."server"."name" IS '服务名称';
COMMENT ON COLUMN "public"."server"."description" IS '服务描述';
COMMENT ON TABLE "public"."server" IS '服务信息';

-- ----------------------------
-- Records of server
-- ----------------------------
INSERT INTO "public"."server" VALUES (1, 1712995716, 1719246376, 'Gateway', '通用网关', '主要负责前端到后端的转发', true);
INSERT INTO "public"."server" VALUES (2, 1712995716, 1750762065, 'Manager', '管理中心', '主要负责系统的基础管理', true);
INSERT INTO "public"."server" VALUES (3, 1712995716, 1719283436, 'Resource', '资源中心', '主要负责静态资源的管理', true);
INSERT INTO "public"."server" VALUES (4, 1712995716, 1719634247, 'Cron', '定时任务', '主要负责定时任务的管理', true);
INSERT INTO "public"."server" VALUES (5, 1712995716, 1728064298, 'Application', '应用中心', '主要负责业务应用的管理', true);
INSERT INTO "public"."server" VALUES (6, 1712995716, 1750772337, 'PartyAffairs', '信号灯系统', '指尖上的党务系统', true);
INSERT INTO "public"."server" VALUES (9, 1723027582, 1750772356, 'Poverty', '资助系统', '资助系统', true);

-- ----------------------------
-- Table structure for template
-- ----------------------------
DROP TABLE IF EXISTS "public"."template";
CREATE TABLE "public"."template" (
                                     "id" int8 NOT NULL,
                                     "created_at" int8,
                                     "updated_at" int8,
                                     "server_id" int8 NOT NULL,
                                     "content" text COLLATE "pg_catalog"."default" NOT NULL,
                                     "version" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
                                     "is_use" boolean,
                                     "format" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
                                     "description" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
                                     "compare" text COLLATE "pg_catalog"."default" NOT NULL
)
;
COMMENT ON COLUMN "public"."template"."id" IS '主键ID';
COMMENT ON COLUMN "public"."template"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."template"."updated_at" IS '修改时间';
COMMENT ON COLUMN "public"."template"."server_id" IS '模板id';
COMMENT ON COLUMN "public"."template"."content" IS '模板内容';
COMMENT ON COLUMN "public"."template"."version" IS '模板版本';
COMMENT ON COLUMN "public"."template"."is_use" IS '是否使用';
COMMENT ON COLUMN "public"."template"."format" IS '模板格式';
COMMENT ON COLUMN "public"."template"."description" IS '模板描述';
COMMENT ON COLUMN "public"."template"."compare" IS '变更详情';
COMMENT ON TABLE "public"."template" IS '配置模板';

-- ----------------------------
-- Records of template
-- ----------------------------
INSERT INTO "public"."template" VALUES (1, 1712995716, 1728064733, 1, '
debug: true
addr: 0.0.0.0:${GatewayServer.HttpPort}
name: gateway
version: v1
middlewares:
  - name: bbr
  - name: cors
    options:
      allowCredentials: true
      allowOrigins:
        - ''*''
      allowMethods:
        - GET
        - POST
        - PUT
        - DELETE
        - OPTIONS
      AllowHeaders:
        - Content-Type
        - Content-Length
        - Authorization
      ExposeHeaders:
        - Content-Length
        - Access-Control-Allow-Headers
  - name: tracing
  - name: logging
  - name: transcoder
endpoints:
  - path: /manager/v1/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
  - path: /manager/client/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse
          method: POST
  - path: /configure/*
    timeout: ${ConfigureServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/v1/auth
          method: POST
  - path: /resource/v1/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/v1/auth
          method: POST
          whiteList:
            - path: /resource/v1/static/*
              method: GET
  - path: /cron/v1/*
    timeout: ${CronServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${CronServer.Host}:${CronServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/v1/auth
          method: POST
  - path: /resource/client/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse
          method: POST
  - path: /user-center/v1/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/v1/auth
          method: POST
  - path: /user-center/client/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
  - path: /party-affairs/v1/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/v1/auth
          method: POST
  - path: /party-affairs/client/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse
          method: POST
', '54E50FB3522C', false, 'yaml', '初始化模板', '');
INSERT INTO "public"."template" VALUES (2, 1712995716, 1719473730, 2, '
env: ${Env.Keyword}
server:
  http:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.HttpPort}
    timeout: ${ManagerServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.GrpcPort}
    timeout: ${ManagerServer.Timeout}
log:
  level: 0
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: ${ManagerDatabase.Type} #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: ${ManagerDatabase.Username}
      password: ${ManagerDatabase.Password}
      host: ${ManagerDatabase.Host}
      port: ${ManagerDatabase.Port}
      dbName: ${ManagerDatabase.Database}
      option: ${ManagerDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
captcha:
  login:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 240
    skew: 0.7
    refresh: true
    dotCount: 80
  changePassword:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: captcha
loader:
  login: ${LoginPrivatePath}
email:
  template:
    captcha:
      subject: 验证码发送通知
      path: static/template/email/default.html
      from: 青岑云科技
      type: text/html
  user: ${Email.Username}
  name: ${Email.Company}
  host: ${Email.Host}
  port: ${Email.Port}
  password: ${Email.Password}
jwt:
  redis: cache
  secret: ${AdminJwt.Secret}
  expire: ${AdminJwt.Expire}
  renewal: ${AdminJwt.Renewal}
  whitelist: ${AdminJwt.Whitelist}
authentication:
  db: system
  redis: cache
  roleKey: role_keyword
  skipRole: ${AuthSkipRoles}
business:
  defaultUserPassword: ${DefaultUserPassword}
  setting: ${Setting}
', '15768C4C6F57', false, 'yaml', '初始化模板', '');
INSERT INTO "public"."template" VALUES (3, 1712995716, 1728072395, 3, '
env: ${Env.Keyword}
server:
  http:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.HttpPort}
    timeout: ${ResourceServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.GrpcPort}
    timeout: ${ResourceServer.Timeout}
log:
  level: 0
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: ${ResourceDatabase.Type} #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: ${ResourceDatabase.Username}
      password: ${ResourceDatabase.Password}
      host: ${ResourceDatabase.Host}
      port: ${ResourceDatabase.Port}
      dbName: ${ResourceDatabase.Database}
      option: ${ResourceDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 3 #日志等级
      slowThreshold: 2s #慢sql阈值
file:
  storage: local
  serverPath: /resource/v1/static
  localDir: static
  maxSingularSize: ${MaxSingularSize}
  maxChunkSize: ${MaxChunkSize}
  maxChunkCount: ${MaxChunkCount}
  acceptTypes: ${AcceptTypes}

', '56945B81FA4D', false, 'yaml', '初始化模板', '');
INSERT INTO "public"."template" VALUES (4, 1712995716, 1719719677, 4, '
env: ${Env.Keyword}
server:
  http:
    host: ${CronServer.Host}
    port: ${CronServer.HttpPort}
    timeout: ${CronServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${CronServer.Host}
    port: ${CronServer.GrpcPort}
    timeout: ${CronServer.Timeout}
log:
  level: 0
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: ${CronDatabase.Type} #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: ${CronDatabase.Username}
      password: ${CronDatabase.Password}
      host: ${CronDatabase.Host}
      port: ${CronDatabase.Port}
      dbName: ${CronDatabase.Database}
      option: ${CronDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 3 #日志等级
      slowThreshold: 2s #慢sql阈值
pool: #并非池配置
  size: 10000 #最大协程数量
  expiryDuration: 30s #过期时间
  preAlloc: true #是否预分配
  maxBlockingTasks: 1000 #最大的并发任务
  nonblocking: true #设置为true时maxBlockingTasks将失效，不限制并发任务
business:
  key: value
', 'F2893009A192', false, 'yaml', '初始化模板', '');
INSERT INTO "public"."template" VALUES (5, 1712995716, 1728064794, 5, '
env: ${Env.Keyword}
server:
  http:
    host: ${UserCenterServer.Host}
    port: ${UserCenterServer.HttpPort}
    timeout: ${UserCenterServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${UserCenterServer.Host}
    port: ${UserCenterServer.GrpcPort}
    timeout: ${UserCenterServer.Timeout}
client:
  - server: Resource
    type: direct
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.GrpcPort}
log:
  level: 0
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: ${UserCenterDatabase.Type} #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: ${UserCenterDatabase.Username}
      password: ${UserCenterDatabase.Password}
      host: ${UserCenterDatabase.Host}
      port: ${UserCenterDatabase.Port}
      dbName: ${UserCenterDatabase.Database}
      option: ${UserCenterDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
captcha:
  loginImage:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 200
    skew: 0.7
    dotCount: 80
    refresh: true
  bindImage:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 200
    skew: 0.7
    dotCount: 80
    refresh: true
  registerImage:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 200
    skew: 0.7
    dotCount: 80
    refresh: true
  loginEmail:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: login
  bindEmail:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: bind
  registerEmail:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: register
loader:
  password: static/cert/password.pem
email:
  template:
    login:
      subject: 登录验证码发送通知
      path: static/template/email/default.html
      type: text/html
    bind:
      subject: 绑定验证码发送通知
      path: static/template/email/default.html
      type: text/html
    register:
      subject: 注册验证码发送通知
      path: static/template/email/default.html
      type: text/html
  user: ${Email.Username}
  name: ${Email.Company}
  host: ${Email.Host}
  port: ${Email.Port}
  password: ${Email.Password}
jwt:
  redis: cache
  secret: ${ClientJwt.Secret}
  expire: ${ClientJwt.Expire}
  renewal: ${ClientJwt.Renewal}
  whitelist: ${ClientJwt.Whitelist}
business:
  service:
    resource: ${ResourceServer.Host}:${ResourceServer.GrpcPort}
', 'DE24124137B1', false, 'yaml', '初始化模板', '');
INSERT INTO "public"."template" VALUES (12, 1719463020, 1719473730, 2, 'test: 11
env: ${Env.Keyword}
server:
  http:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.HttpPort}
    timeout: ${ManagerServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.GrpcPort}
    timeout: ${ManagerServer.Timeout}
log:
  level: 0
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: ${ManagerDatabase.Type} #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: ${ManagerDatabase.Username}
      password: ${ManagerDatabase.Password}
      host: ${ManagerDatabase.Host}
      port: ${ManagerDatabase.Port}
      dbName: ${ManagerDatabase.Database}
      option: ${ManagerDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
captcha:
  login:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 240
    skew: 0.7
    refresh: true
    dotCount: 80
  changePassword:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: captcha
loader:
  login: ${LoginPrivatePath}
email:
  template:
    captcha:
      subject: 验证码发送通知
      path: static/template/email/default.html
      from: 统一应用管理中心
      type: text/html
  user: ${Email.Username}
  name: ${Email.Company}
  host: ${Email.Host}
  port: ${Email.Port}
  password: ${Email.Password}
jwt:
  redis: cache
  secret: ${AdminJwt.Secret}
  expire: ${AdminJwt.Expire}
  renewal: ${AdminJwt.Renewal}
  whitelist: ${AdminJwt.Whitelist}
authentication:
  db: system
  redis: cache
  roleKey: roleKeyword
  skipRole: ${AuthSkipRoles}
client:
  - server: Resource
    type: direct
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.GrpcPort}
business:
  changePasswordType: ${ChangePasswordType}
  defaultUserPassword: ${DefaultUserPassword}
  setting: ${Setting}
', 'B0041A124D61345FF898A0019133ABF2', false, 'yaml', '初始化模板', '[{"type":"add","key":"client","old":"","cur":"- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.GrpcPort}\n  server: Resource\n  type: direct\n"},{"type":"update","key":"email","old":"host: ${Email.Host}\nname: ${Email.Company}\npassword: ${Email.Password}\nport: ${Email.Port}\ntemplate:\n    captcha:\n        from: 青岑云科技\n        path: static/template/email/default.html\n        subject: 验证码发送通知\n        type: text/html\nuser: ${Email.Username}\n","cur":"host: ${Email.Host}\nname: ${Email.Company}\npassword: ${Email.Password}\nport: ${Email.Port}\ntemplate:\n    captcha:\n        from: 统一应用管理中心\n        path: static/template/email/default.html\n        subject: 验证码发送通知\n        type: text/html\nuser: ${Email.Username}\n"},{"type":"update","key":"authentication","old":"db: system\nredis: cache\nroleKey: role_keyword\nskipRole: ${AuthSkipRoles}\n","cur":"db: system\nredis: cache\nroleKey: roleKeyword\nskipRole: ${AuthSkipRoles}\n"},{"type":"update","key":"database","old":"system:\n    autoCreate: true\n    config:\n        initializer:\n            enable: false\n            path: deploy/data.sql\n        logLevel: 4\n        maxIdleConn: 10\n        maxLifetime: 2h\n        maxOpenConn: 20\n        slowThreshold: 2s\n        transformError:\n            enable: true\n    connect:\n        dbName: ${ManagerDatabase.Database}\n        host: ${ManagerDatabase.Host}\n        option: ${ManagerDatabase.Option}\n        password: ${ManagerDatabase.Password}\n        port: ${ManagerDatabase.Port}\n        username: ${ManagerDatabase.Username}\n    drive: ${ManagerDatabase.Type}\n    enable: true\n","cur":"system:\n    autoCreate: true\n    config:\n        initializer:\n            enable: true\n            path: deploy/data.sql\n        logLevel: 4\n        maxIdleConn: 10\n        maxLifetime: 2h\n        maxOpenConn: 20\n        slowThreshold: 2s\n        transformError:\n            enable: true\n    connect:\n        dbName: ${ManagerDatabase.Database}\n        host: ${ManagerDatabase.Host}\n        option: ${ManagerDatabase.Option}\n        password: ${ManagerDatabase.Password}\n        port: ${ManagerDatabase.Port}\n        username: ${ManagerDatabase.Username}\n    drive: ${ManagerDatabase.Type}\n    enable: true\n"},{"type":"update","key":"business","old":"defaultUserPassword: ${DefaultUserPassword}\nsetting: ${Setting}\n","cur":"changePasswordType: ${ChangePasswordType}\ndefaultUserPassword: ${DefaultUserPassword}\nsetting: ${Setting}\n"}]');
INSERT INTO "public"."template" VALUES (6, 1712995716, 1728071372, 6, '
env: ${Env.Keyword}
server:
  http:
    host: ${PartyAffairsServer.Host}
    port: ${PartyAffairsServer.HttpPort}
    timeout: ${PartyAffairsServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${PartyAffairsServer.Host}
    port: ${PartyAffairsServer.GrpcPort}
    timeout: ${PartyAffairsServer.Timeout}
client:
  - server: Resource
    type: direct
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.GrpcPort}
  - server: UserCenter
    type: direct
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.GrpcPort}
log:
  level: 0
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: ${PartyAffairsDatabase.Type} #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: ${PartyAffairsDatabase.Username}
      password: ${PartyAffairsDatabase.Password}
      host: ${PartyAffairsDatabase.Host}
      port: ${PartyAffairsDatabase.Port}
      dbName: ${PartyAffairsDatabase.Database}
      option: ${PartyAffairsDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
business:
  auth:
    yiBan:
      appId: e4750b34230b48e1
      appSecret: b0891a7f6018e5a76b085e3cb9548edd
', 'D825BF2B6742', false, 'yaml', '初始化模板', '');
INSERT INTO "public"."template" VALUES (7, 1713011130, 1719473730, 2, '
env: ${Env.Keyword}
server:
  http:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.HttpPort}
    timeout: ${ManagerServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.GrpcPort}
    timeout: ${ManagerServer.Timeout}
log:
  level: 0
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: ${ManagerDatabase.Type} #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: ${ManagerDatabase.Username}
      password: ${ManagerDatabase.Password}
      host: ${ManagerDatabase.Host}
      port: ${ManagerDatabase.Port}
      dbName: ${ManagerDatabase.Database}
      option: ${ManagerDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: false
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
captcha:
  login:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 240
    skew: 0.7
    refresh: true
    dotCount: 80
  changePassword:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: captcha
loader:
  login: ${LoginPrivatePath}
email:
  template:
    captcha:
      subject: 验证码发送通知
      path: static/template/email/default.html
      from: 青岑云科技
      type: text/html
  user: ${Email.Username}
  name: ${Email.Company}
  host: ${Email.Host}
  port: ${Email.Port}
  password: ${Email.Password}
jwt:
  redis: cache
  secret: ${AdminJwt.Secret}
  expire: ${AdminJwt.Expire}
  renewal: ${AdminJwt.Renewal}
  whitelist: ${AdminJwt.Whitelist}
authentication:
  db: system
  redis: cache
  roleKey: role_keyword
  skipRole: ${AuthSkipRoles}
business:
  defaultUserPassword: ${DefaultUserPassword}
  setting: ${Setting}
', '3ECF4F0622F7', false, 'yaml', '1', '');
INSERT INTO "public"."template" VALUES (8, 1714750273, 1728072395, 3, '
env: ${Env.Keyword}
server:
  http:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.HttpPort}
    timeout: ${ResourceServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.GrpcPort}
    timeout: ${ResourceServer.Timeout}
log:
  level: 0
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: ${ResourceDatabase.Type} #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: ${ResourceDatabase.Username}
      password: ${ResourceDatabase.Password}
      host: ${ResourceDatabase.Host}
      port: ${ResourceDatabase.Port}
      dbName: ${ResourceDatabase.Database}
      option: ${ResourceDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 3 #日志等级
      slowThreshold: 2s #慢sql阈值
business:
  storage:
    type: local
    serverPath: /resource/v1/static
    localDir: static
    maxSingularSize: ${MaxSingularSize}
    maxChunkSize: ${MaxChunkSize}
    maxChunkCount: ${MaxChunkCount}
    acceptTypes: ${AcceptTypes}
  export:
    localDir: static/export
    expire: 72h
', '3942130C8BEE', false, 'yaml', '初始化模板', '');
INSERT INTO "public"."template" VALUES (10, 1719379877, 1719473730, 2, 'test: 11
env: ${Env.Keyword}
server:
  http:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.HttpPort}
    timeout: ${ManagerServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.GrpcPort}
    timeout: ${ManagerServer.Timeout}
log:
  level: 0
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: ${ManagerDatabase.Type} #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: ${ManagerDatabase.Username}
      password: ${ManagerDatabase.Password}
      host: ${ManagerDatabase.Host}
      port: ${ManagerDatabase.Port}
      dbName: ${ManagerDatabase.Database}
      option: ${ManagerDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: false
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
captcha:
  login:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 240
    skew: 0.7
    refresh: true
    dotCount: 80
  changePassword:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: captcha
loader:
  login: ${LoginPrivatePath}
email:
  template:
    captcha:
      subject: 验证码发送通知
      path: static/template/email/default.html
      from: 青岑云科技
      type: text/html
  user: ${Email.Username}
  name: ${Email.Company}
  host: ${Email.Host}
  port: ${Email.Port}
  password: ${Email.Password}
jwt:
  redis: cache
  secret: ${AdminJwt.Secret}
  expire: ${AdminJwt.Expire}
  renewal: ${AdminJwt.Renewal}
  whitelist: ${AdminJwt.Whitelist}
authentication:
  db: system
  redis: cache
  roleKey: role_keyword
  skipRole: ${AuthSkipRoles}
business:
  defaultUserPassword: ${DefaultUserPassword}
  setting: ${Setting}
', 'A8A826623534726F8ED1F0AC6ACDFB8B', false, 'yaml', '1', '[{"type":"add","key":"test","old":"","cur":"11"}]');
INSERT INTO "public"."template" VALUES (11, 1719462171, 1728064733, 1, 'addr: 0.0.0.0:${GatewayServer.HttpPort}
name: gateway
version: v1
middlewares:
  - name: bbr
  - name: cors
    options:
      allowCredentials: true
      allowOrigins:
        - ''*''
      allowMethods:
        - GET
        - POST
        - PUT
        - DELETE
        - OPTIONS
      AllowHeaders:
        - Content-Type
        - Content-Length
        - Authorization
      ExposeHeaders:
        - Content-Length
        - Access-Control-Allow-Headers
  - name: tracing
  - name: logging
  - name: transcoder
endpoints:
  - path: /manager/api/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
  - path: /manager/client/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth
          method: POST
  - path: /configure/api/*
    timeout: ${ConfigureServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /resource/api/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
          whiteList:
            - path: /resource/api/v1/static/*
              method: GET
  - path: /resource/client/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse
          method: POST
  - path: /usercenter/api/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /usercenter/client/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
  - path: /cron/api/*
    timeout: ${CronServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${CronServer.Host}:${CronServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/api/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/client/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth
          method: POST', 'CE6DA7E554780B619C4C5E8EBD11E7FB', false, 'yaml', '初始化模板', '[{"type":"del","key":"debug","old":"true","cur":""},{"type":"update","key":"endpoints","old":"- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  path: /manager/v1/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse\n  path: /manager/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/v1/auth\n  path: /configure/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ConfigureServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/v1/auth\n        whiteList:\n            - method: GET\n              path: /resource/v1/static/*\n  path: /resource/v1/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/v1/auth\n  path: /cron/v1/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${CronServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse\n  path: /resource/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/v1/auth\n  path: /user-center/v1/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  path: /user-center/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/v1/auth\n  path: /party-affairs/v1/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse\n  path: /party-affairs/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n","cur":"- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  path: /manager/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n  path: /manager/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /configure/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ConfigureServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n        whiteList:\n            - method: GET\n              path: /resource/api/v1/static/*\n  path: /resource/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse\n  path: /resource/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /usercenter/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  path: /usercenter/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /cron/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${CronServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /partyaffairs/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\n  path: /partyaffairs/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n"}]');
INSERT INTO "public"."template" VALUES (13, 1719464064, 1719473730, 2, 'test: 11
env: ${Env.Keyword}
server:
  http:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.HttpPort}
    timeout: ${ManagerServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.GrpcPort}
    timeout: ${ManagerServer.Timeout}
log:
  level: 0
  caller: true
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: ${ManagerDatabase.Type} #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: ${ManagerDatabase.Username}
      password: ${ManagerDatabase.Password}
      host: ${ManagerDatabase.Host}
      port: ${ManagerDatabase.Port}
      dbName: ${ManagerDatabase.Database}
      option: ${ManagerDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
captcha:
  login:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 240
    skew: 0.7
    refresh: true
    dotCount: 80
  changePassword:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: captcha
loader:
  login: ${LoginPrivatePath}
email:
  template:
    captcha:
      subject: 验证码发送通知
      path: static/template/email/default.html
      from: 统一应用管理中心
      type: text/html
  user: ${Email.Username}
  name: ${Email.Company}
  host: ${Email.Host}
  port: ${Email.Port}
  password: ${Email.Password}
jwt:
  redis: cache
  secret: ${AdminJwt.Secret}
  expire: ${AdminJwt.Expire}
  renewal: ${AdminJwt.Renewal}
  whitelist: ${AdminJwt.Whitelist}
authentication:
  db: system
  redis: cache
  roleKey: roleKeyword
  skipRole: ${AuthSkipRoles}
client:
  - server: Resource
    type: direct
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.GrpcPort}
business:
  changePasswordType: ${ChangePasswordType}
  defaultUserPassword: ${DefaultUserPassword}
  setting: ${Setting}
', '794DBE4E2B65F3CA34F0CCF465169E87', false, 'yaml', '初始化模板', '[{"type":"update","key":"log","old":"file:\n    compress: false\n    maxAge: 1\n    maxBackup: 5\n    maxSize: 1\n    name: ./tmp/runtime/output.log\nlevel: 0\noutput:\n    - stdout\n    - file\n","cur":"caller: true\nfile:\n    compress: false\n    maxAge: 1\n    maxBackup: 5\n    maxSize: 1\n    name: ./tmp/runtime/output.log\nlevel: 0\noutput:\n    - stdout\n    - file\n"}]');
INSERT INTO "public"."template" VALUES (14, 1719465442, 1728072395, 3, '
env: ${Env.Keyword}
server:
  http:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.HttpPort}
    timeout: ${ResourceServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.GrpcPort}
    timeout: ${ResourceServer.Timeout}
log:
  level: 0
  caller: true
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: ${ResourceDatabase.Type} #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: ${ResourceDatabase.Username}
      password: ${ResourceDatabase.Password}
      host: ${ResourceDatabase.Host}
      port: ${ResourceDatabase.Port}
      dbName: ${ResourceDatabase.Database}
      option: ${ResourceDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 3 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
business:
  chunkSize: ${ChunkSize}
  defaultMaxSize: ${DefaultMaxSize}
  defaultAcceptTypes: ${DefaultAcceptTypes}
  storage:
    type: local
    serverUrl: ${GatewayServer.Host}:${ResourceServer.HttpPort}/resource/api/v1/static
    localDir: static
    temporaryExpire: 600s
    secret: limescloud
  export:
    serverUrl: ${GatewayServer.Host}:${ResourceServer.HttpPort}/resource/api/v1/download
    localDir: static/export
    expire: 72h
', '26E885D17255295754DFB7B3574A30E9', false, 'yaml', '初始化模板', '[{"type":"add","key":"redis","old":"","cur":"cache:\n    enable: true\n    host: ${Redis.Host}:${Redis.Port}\n    password: ${Redis.Password}\n    username: ${Redis.Username}\n"},{"type":"update","key":"business","old":"export:\n    expire: 72h\n    localDir: static/export\nstorage:\n    acceptTypes: ${AcceptTypes}\n    localDir: static\n    maxChunkCount: ${MaxChunkCount}\n    maxChunkSize: ${MaxChunkSize}\n    maxSingularSize: ${MaxSingularSize}\n    serverPath: /resource/v1/static\n    type: local\n","cur":"chunkSize: ${ChunkSize}\ndefaultAcceptTypes: ${DefaultAcceptTypes}\ndefaultMaxSize: ${DefaultMaxSize}\nexport:\n    expire: 72h\n    localDir: static/export\n    serverUrl: ${GatewayServer.Host}:${ResourceServer.HttpPort}/resource/api/v1/download\nstorage:\n    localDir: static\n    secret: limescloud\n    serverUrl: ${GatewayServer.Host}:${ResourceServer.HttpPort}/resource/api/v1/static\n    temporaryExpire: 600s\n    type: local\n"},{"type":"update","key":"log","old":"file:\n    compress: false\n    maxAge: 1\n    maxBackup: 5\n    maxSize: 1\n    name: ./tmp/runtime/output.log\nlevel: 0\noutput:\n    - stdout\n    - file\n","cur":"caller: true\nfile:\n    compress: false\n    maxAge: 1\n    maxBackup: 5\n    maxSize: 1\n    name: ./tmp/runtime/output.log\nlevel: 0\noutput:\n    - stdout\n    - file\n"}]');
INSERT INTO "public"."template" VALUES (15, 1719466300, 1728072395, 3, '
env: ${Env.Keyword}
server:
  http:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.HttpPort}
    timeout: ${ResourceServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.GrpcPort}
    timeout: ${ResourceServer.Timeout}
log:
  level: 0
  caller: true
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: ${ResourceDatabase.Type} #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: ${ResourceDatabase.Username}
      password: ${ResourceDatabase.Password}
      host: ${ResourceDatabase.Host}
      port: ${ResourceDatabase.Port}
      dbName: ${ResourceDatabase.Database}
      option: ${ResourceDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 3 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
business:
  chunkSize: ${ChunkSize}
  defaultMaxSize: ${DefaultMaxSize}
  defaultAcceptTypes: ${DefaultAcceptTypes}
  storage:
    type: local
    serverUrl: http://${GatewayServer.Host}:${GatewayServer.HttpPort}/resource/api/v1/static
    localDir: static
    temporaryExpire: 600s
    secret: limescloud
  export:
    serverUrl: http://${GatewayServer.Host}:${GatewayServer.HttpPort}/resource/api/v1/download
    localDir: static/export
    expire: 72h
', 'C2E8D51753B9C34DBC56F80137C492A7', false, 'yaml', '初始化模板', '[{"type":"update","key":"business","old":"chunkSize: ${ChunkSize}\ndefaultAcceptTypes: ${DefaultAcceptTypes}\ndefaultMaxSize: ${DefaultMaxSize}\nexport:\n    expire: 72h\n    localDir: static/export\n    serverUrl: ${GatewayServer.Host}:${ResourceServer.HttpPort}/resource/api/v1/download\nstorage:\n    localDir: static\n    secret: limescloud\n    serverUrl: ${GatewayServer.Host}:${ResourceServer.HttpPort}/resource/api/v1/static\n    temporaryExpire: 600s\n    type: local\n","cur":"chunkSize: ${ChunkSize}\ndefaultAcceptTypes: ${DefaultAcceptTypes}\ndefaultMaxSize: ${DefaultMaxSize}\nexport:\n    expire: 72h\n    localDir: static/export\n    serverUrl: http://${GatewayServer.Host}:${GatewayServer.HttpPort}/resource/api/v1/download\nstorage:\n    localDir: static\n    secret: limescloud\n    serverUrl: http://${GatewayServer.Host}:${GatewayServer.HttpPort}/resource/api/v1/static\n    temporaryExpire: 600s\n    type: local\n"}]');
INSERT INTO "public"."template" VALUES (16, 1719466963, 1728064733, 1, 'addr: 0.0.0.0:${GatewayServer.HttpPort}
name: gateway
version: v1
middlewares:
  - name: bbr
  - name: cors
    options:
      allowCredentials: true
      allowOrigins:
        - ''*''
      allowMethods:
        - GET
        - POST
        - PUT
        - DELETE
        - OPTIONS
      AllowHeaders:
        - Content-Type
        - Content-Length
        - Authorization
      ExposeHeaders:
        - Content-Length
        - Access-Control-Allow-Headers
  - name: tracing
  - name: logging
  - name: transcoder
endpoints:
  - path: /manager/api/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
  - path: /manager/client/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth
          method: POST
  - path: /configure/api/*
    timeout: ${ConfigureServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /resource/api/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
          whiteList:
            - path: /resource/api/v1/static/*
              method: GET
            - path: /resource/api/v1/download/*
              method: GET
  - path: /resource/client/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse
          method: POST
  - path: /usercenter/api/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /usercenter/client/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
  - path: /cron/api/*
    timeout: ${CronServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${CronServer.Host}:${CronServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/api/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/client/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth
          method: POST', 'D810212421E6DB836EA3CEA6C2E6DCFF', false, 'yaml', '初始化模板', '[{"type":"update","key":"endpoints","old":"- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  path: /manager/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n  path: /manager/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /configure/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ConfigureServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n        whiteList:\n            - method: GET\n              path: /resource/api/v1/static/*\n  path: /resource/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse\n  path: /resource/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /usercenter/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  path: /usercenter/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /cron/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${CronServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /partyaffairs/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\n  path: /partyaffairs/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n","cur":"- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  path: /manager/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n  path: /manager/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /configure/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ConfigureServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n        whiteList:\n            - method: GET\n              path: /resource/api/v1/static/*\n            - method: GET\n              path: /resource/api/v1/download/*\n  path: /resource/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse\n  path: /resource/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /usercenter/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  path: /usercenter/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /cron/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${CronServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /partyaffairs/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\n  path: /partyaffairs/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n"}]');
INSERT INTO "public"."template" VALUES (21, 1719559979, 1728064733, 1, 'addr: 0.0.0.0:${GatewayServer.HttpPort}
name: gateway
version: v1
middlewares:
  - name: bbr
  - name: cors
    options:
      allowCredentials: true
      allowOrigins:
        - ''*''
      allowMethods:
        - GET
        - POST
        - PUT
        - DELETE
        - OPTIONS
      AllowHeaders:
        - Content-Type
        - Content-Length
        - Authorization
      ExposeHeaders:
        - Content-Length
        - Access-Control-Allow-Headers
  - name: tracing
  - name: logging
  - name: transcoder
endpoints:
  - path: /manager/api/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
  - path: /manager/client/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth
          method: POST
  - path: /configure/api/*
    timeout: ${ConfigureServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /resource/api/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
          whiteList:
            - path: /resource/api/v1/static/*
              method: GET
            - path: /resource/api/v1/download/*
              method: GET
  - path: /resource/client/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth
          method: POST
  - path: /usercenter/api/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /usercenter/client/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
  - path: /cron/api/*
    timeout: ${CronServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${CronServer.Host}:${CronServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/api/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/client/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth
          method: POST', 'CD650FA575AAE93B796D3C395E324C48', false, 'yaml', '初始化模板', '[{"type":"update","key":"endpoints","old":"- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  path: /manager/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n  path: /manager/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /configure/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ConfigureServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n        whiteList:\n            - method: GET\n              path: /resource/api/v1/static/*\n            - method: GET\n              path: /resource/api/v1/download/*\n  path: /resource/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\n  path: /resource/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n        whiteList:\n            - method: POST\n              path: /usercenter/api/v1/auth\n  path: /usercenter/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  path: /usercenter/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /cron/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${CronServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /partyaffairs/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\n  path: /partyaffairs/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n","cur":"- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  path: /manager/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n  path: /manager/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /configure/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ConfigureServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n        whiteList:\n            - method: GET\n              path: /resource/api/v1/static/*\n            - method: GET\n              path: /resource/api/v1/download/*\n  path: /resource/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n  path: /resource/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /usercenter/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  path: /usercenter/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /cron/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${CronServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /partyaffairs/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\n  path: /partyaffairs/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n"}]');
INSERT INTO "public"."template" VALUES (17, 1719473730, 1719473730, 2, 'test: 11
env: ${Env.Keyword}
server:
  http:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.HttpPort}
    timeout: ${ManagerServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ManagerServer.Host}
    port: ${ManagerServer.GrpcPort}
    timeout: ${ManagerServer.Timeout}
log:
  level: 0
  caller: true
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: ${ManagerDatabase.Type} #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: ${ManagerDatabase.Username}
      password: ${ManagerDatabase.Password}
      host: ${ManagerDatabase.Host}
      port: ${ManagerDatabase.Port}
      dbName: ${ManagerDatabase.Database}
      option: ${ManagerDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
captcha:
  login:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 240
    skew: 0.7
    refresh: true
    dotCount: 80
  changePassword:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: captcha
loader:
  login: ${LoginPrivatePath}
email:
  template:
    captcha:
      subject: 验证码发送通知
      path: static/template/email/default.html
      from: 统一应用管理中心
      type: text/html
  user: ${Email.Username}
  name: ${Email.Company}
  host: ${Email.Host}
  port: ${Email.Port}
  password: ${Email.Password}
jwt:
  redis: cache
  secret: ${AdminJwt.Secret}
  expire: ${AdminJwt.Expire}
  renewal: ${AdminJwt.Renewal}
  whitelist: ${AdminJwt.Whitelist}
authentication:
  db: system
  redis: cache
  roleKey: roleKeyword
  skipRole: ${AuthSkipRoles}
client:
  - server: Resource
    type: direct
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.GrpcPort}
business:
  changePasswordType: ${ChangePasswordType}
  defaultUserPassword: ${DefaultUserPassword}
  setting: ${Setting}
', 'D69ED17B11BA6B9C08A906818FD83937', true, 'yaml', '初始化模板', '[{"type":"update","key":"client","old":"- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.GrpcPort}\n  server: Resource\n  type: direct\n","cur":"- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.GrpcPort}\n  server: Resource\n  type: direct\n"}]');
INSERT INTO "public"."template" VALUES (18, 1719557318, 1728064733, 1, 'addr: 0.0.0.0:${GatewayServer.HttpPort}
name: gateway
version: v1
middlewares:
  - name: bbr
  - name: cors
    options:
      allowCredentials: true
      allowOrigins:
        - ''*''
      allowMethods:
        - GET
        - POST
        - PUT
        - DELETE
        - OPTIONS
      AllowHeaders:
        - Content-Type
        - Content-Length
        - Authorization
      ExposeHeaders:
        - Content-Length
        - Access-Control-Allow-Headers
  - name: tracing
  - name: logging
  - name: transcoder
endpoints:
  - path: /manager/api/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
  - path: /manager/client/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth
          method: POST
  - path: /configure/api/*
    timeout: ${ConfigureServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /resource/api/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
          whiteList:
            - path: /resource/api/v1/static/*
              method: GET
            - path: /resource/api/v1/download/*
              method: GET
  - path: /resource/client/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth
          method: POST
  - path: /usercenter/api/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /usercenter/client/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
  - path: /cron/api/*
    timeout: ${CronServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${CronServer.Host}:${CronServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/api/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/client/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth
          method: POST', 'D0E935C5032DC43B988E70406A7B6111', false, 'yaml', '初始化模板', '[{"type":"update","key":"endpoints","old":"- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  path: /manager/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n  path: /manager/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /configure/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ConfigureServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n        whiteList:\n            - method: GET\n              path: /resource/api/v1/static/*\n            - method: GET\n              path: /resource/api/v1/download/*\n  path: /resource/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/user-center/client/token/parse\n  path: /resource/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /usercenter/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  path: /usercenter/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /cron/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${CronServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /partyaffairs/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\n  path: /partyaffairs/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n","cur":"- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  path: /manager/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n  path: /manager/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /configure/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ConfigureServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n        whiteList:\n            - method: GET\n              path: /resource/api/v1/static/*\n            - method: GET\n              path: /resource/api/v1/download/*\n  path: /resource/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\n  path: /resource/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /usercenter/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  path: /usercenter/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /cron/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${CronServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /partyaffairs/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\n  path: /partyaffairs/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n"}]');
INSERT INTO "public"."template" VALUES (19, 1719558297, 1728064794, 5, '
env: ${Env.Keyword}
server:
  http:
    host: ${UserCenterServer.Host}
    port: ${UserCenterServer.HttpPort}
    timeout: ${UserCenterServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${UserCenterServer.Host}
    port: ${UserCenterServer.GrpcPort}
    timeout: ${UserCenterServer.Timeout}
client:
  - server: Resource
    type: direct
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.GrpcPort}
log:
  level: 0
  caller: true
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: ${UserCenterDatabase.Type} #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: ${UserCenterDatabase.Username}
      password: ${UserCenterDatabase.Password}
      host: ${UserCenterDatabase.Host}
      port: ${UserCenterDatabase.Port}
      dbName: ${UserCenterDatabase.Database}
      option: ${UserCenterDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
captcha:
  loginImage:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 200
    skew: 0.7
    dotCount: 80
    refresh: true
  bindImage:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 200
    skew: 0.7
    dotCount: 80
    refresh: true
  registerImage:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 200
    skew: 0.7
    dotCount: 80
    refresh: true
  loginEmail:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: login
  bindEmail:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: bind
  registerEmail:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: register
loader:
  password: static/cert/password.pem
email:
  template:
    login:
      subject: 登录验证码发送通知
      path: static/template/email/default.html
      type: text/html
    bind:
      subject: 绑定验证码发送通知
      path: static/template/email/default.html
      type: text/html
    register:
      subject: 注册验证码发送通知
      path: static/template/email/default.html
      type: text/html
  user: ${Email.Username}
  name: ${Email.Company}
  host: ${Email.Host}
  port: ${Email.Port}
  password: ${Email.Password}
jwt:
  redis: cache
  secret: ${ClientJwt.Secret}
  expire: ${ClientJwt.Expire}
  renewal: ${ClientJwt.Renewal}
  whitelist: ${ClientJwt.Whitelist}
business:
  defaultPasswordExpire: 10s
  defaultNickName: 用户-
', '0D01EEE78E5662F3193361C9EEE34266', false, 'yaml', '初始化模板', '[{"type":"update","key":"business","old":"service:\n    resource: ${ResourceServer.Host}:${ResourceServer.GrpcPort}\n","cur":"defaultNickName: 用户-\ndefaultPasswordExpire: 10s\n"},{"type":"update","key":"log","old":"file:\n    compress: false\n    maxAge: 1\n    maxBackup: 5\n    maxSize: 1\n    name: ./tmp/runtime/output.log\nlevel: 0\noutput:\n    - stdout\n    - file\n","cur":"caller: true\nfile:\n    compress: false\n    maxAge: 1\n    maxBackup: 5\n    maxSize: 1\n    name: ./tmp/runtime/output.log\nlevel: 0\noutput:\n    - stdout\n    - file\n"}]');
INSERT INTO "public"."template" VALUES (20, 1719559466, 1728064733, 1, 'addr: 0.0.0.0:${GatewayServer.HttpPort}
name: gateway
version: v1
middlewares:
  - name: bbr
  - name: cors
    options:
      allowCredentials: true
      allowOrigins:
        - ''*''
      allowMethods:
        - GET
        - POST
        - PUT
        - DELETE
        - OPTIONS
      AllowHeaders:
        - Content-Type
        - Content-Length
        - Authorization
      ExposeHeaders:
        - Content-Length
        - Access-Control-Allow-Headers
  - name: tracing
  - name: logging
  - name: transcoder
endpoints:
  - path: /manager/api/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
  - path: /manager/client/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth
          method: POST
  - path: /configure/api/*
    timeout: ${ConfigureServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /resource/api/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
          whiteList:
            - path: /resource/api/v1/static/*
              method: GET
            - path: /resource/api/v1/download/*
              method: GET
  - path: /resource/client/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth
          method: POST
  - path: /usercenter/api/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
          whiteList:
            - path: /usercenter/api/v1/auth
              method: POST
  - path: /usercenter/client/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
  - path: /cron/api/*
    timeout: ${CronServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${CronServer.Host}:${CronServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/api/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/client/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth
          method: POST', 'DEBD584379CAA5B33A00199457BF87DC', false, 'yaml', '初始化模板', '[{"type":"update","key":"endpoints","old":"- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  path: /manager/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n  path: /manager/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /configure/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ConfigureServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n        whiteList:\n            - method: GET\n              path: /resource/api/v1/static/*\n            - method: GET\n              path: /resource/api/v1/download/*\n  path: /resource/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\n  path: /resource/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /usercenter/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  path: /usercenter/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /cron/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${CronServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /partyaffairs/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\n  path: /partyaffairs/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n","cur":"- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  path: /manager/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n  path: /manager/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /configure/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ConfigureServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n        whiteList:\n            - method: GET\n              path: /resource/api/v1/static/*\n            - method: GET\n              path: /resource/api/v1/download/*\n  path: /resource/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\n  path: /resource/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n        whiteList:\n            - method: POST\n              path: /usercenter/api/v1/auth\n  path: /usercenter/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  path: /usercenter/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /cron/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${CronServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /partyaffairs/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\n  path: /partyaffairs/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n"}]');
INSERT INTO "public"."template" VALUES (22, 1719634360, 1719719677, 4, '
env: ${Env.Keyword}
server:
  http:
    host: ${CronServer.Host}
    port: ${CronServer.HttpPort}
    timeout: ${CronServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${CronServer.Host}
    port: ${CronServer.GrpcPort}
    timeout: ${CronServer.Timeout}
log:
  level: 0
  caller: true
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: ${CronDatabase.Type} #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: ${CronDatabase.Username}
      password: ${CronDatabase.Password}
      host: ${CronDatabase.Host}
      port: ${CronDatabase.Port}
      dbName: ${CronDatabase.Database}
      option: ${CronDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
pool: #并非池配置
  size: 10000 #最大协程数量
  expiryDuration: 30s #过期时间
  preAlloc: true #是否预分配
  maxBlockingTasks: 1000 #最大的并发任务
  nonblocking: true #设置为true时maxBlockingTasks将失效，不限制并发任务
business:
  key: value
', 'C5A5395CF93921829673B76AEFC4AE1A', false, 'yaml', '初始化模板', '[{"type":"update","key":"log","old":"file:\n    compress: false\n    maxAge: 1\n    maxBackup: 5\n    maxSize: 1\n    name: ./tmp/runtime/output.log\nlevel: 0\noutput:\n    - stdout\n    - file\n","cur":"caller: true\nfile:\n    compress: false\n    maxAge: 1\n    maxBackup: 5\n    maxSize: 1\n    name: ./tmp/runtime/output.log\nlevel: 0\noutput:\n    - stdout\n    - file\n"},{"type":"update","key":"database","old":"system:\n    autoCreate: true\n    config:\n        initializer:\n            enable: true\n            path: deploy/data.sql\n        logLevel: 3\n        maxIdleConn: 10\n        maxLifetime: 2h\n        maxOpenConn: 20\n        slowThreshold: 2s\n        transformError:\n            enable: true\n    connect:\n        dbName: ${CronDatabase.Database}\n        host: ${CronDatabase.Host}\n        option: ${CronDatabase.Option}\n        password: ${CronDatabase.Password}\n        port: ${CronDatabase.Port}\n        username: ${CronDatabase.Username}\n    drive: ${CronDatabase.Type}\n    enable: true\n","cur":"system:\n    autoCreate: true\n    config:\n        initializer:\n            enable: true\n            path: deploy/data.sql\n        logLevel: 4\n        maxIdleConn: 10\n        maxLifetime: 2h\n        maxOpenConn: 20\n        slowThreshold: 2s\n        transformError:\n            enable: true\n    connect:\n        dbName: ${CronDatabase.Database}\n        host: ${CronDatabase.Host}\n        option: ${CronDatabase.Option}\n        password: ${CronDatabase.Password}\n        port: ${CronDatabase.Port}\n        username: ${CronDatabase.Username}\n    drive: ${CronDatabase.Type}\n    enable: true\n"}]');
INSERT INTO "public"."template" VALUES (23, 1719719677, 1719719677, 4, '
env: ${Env.Keyword}
server:
  http:
    host: ${CronServer.Host}
    port: ${CronServer.HttpPort}
    timeout: ${CronServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${CronServer.Host}
    port: ${CronServer.GrpcPort}
    timeout: ${CronServer.Timeout}
log:
  level: 0
  caller: true
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: true
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
database:
  system:
    enable: true #是否启用数据库
    drive: ${CronDatabase.Type} #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: ${CronDatabase.Username}
      password: ${CronDatabase.Password}
      host: ${CronDatabase.Host}
      port: ${CronDatabase.Port}
      dbName: ${CronDatabase.Database}
      option: ${CronDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
pool: #并非池配置
  size: 10000 #最大协程数量
  expiryDuration: 30s #过期时间
  preAlloc: true #是否预分配
  maxBlockingTasks: 1000 #最大的并发任务
  nonblocking: true #设置为true时maxBlockingTasks将失效，不限制并发任务
business:
  key: value
', '2494002DB6A751C727E337352C1084CA', true, 'yaml', '初始化模板', '[{"type":"add","key":"redis","old":"","cur":"cache:\n    enable: true\n    host: ${Redis.Host}:${Redis.Port}\n    password: ${Redis.Password}\n    username: ${Redis.Username}\n"},{"type":"update","key":"log","old":"caller: true\nfile:\n    compress: false\n    maxAge: 1\n    maxBackup: 5\n    maxSize: 1\n    name: ./tmp/runtime/output.log\nlevel: 0\noutput:\n    - stdout\n    - file\n","cur":"caller: true\nfile:\n    compress: true\n    maxAge: 1\n    maxBackup: 5\n    maxSize: 1\n    name: ./tmp/runtime/output.log\nlevel: 0\noutput:\n    - stdout\n    - file\n"}]');
INSERT INTO "public"."template" VALUES (24, 1720283172, 1728071372, 6, '
env: ${Env.Keyword}
server:
  http:
    host: ${PartyAffairsServer.Host}
    port: ${PartyAffairsServer.HttpPort}
    timeout: ${PartyAffairsServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${PartyAffairsServer.Host}
    port: ${PartyAffairsServer.GrpcPort}
    timeout: ${PartyAffairsServer.Timeout}
client:
  - server: Resource
    type: direct
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.GrpcPort}
  - server: UserCenter
    type: direct
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.GrpcPort}
log:
  level: 0
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: ${PartyAffairsDatabase.Type} #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: ${PartyAffairsDatabase.Username}
      password: ${PartyAffairsDatabase.Password}
      host: ${PartyAffairsDatabase.Host}
      port: ${PartyAffairsDatabase.Port}
      dbName: ${PartyAffairsDatabase.Database}
      option: ${PartyAffairsDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
', '5080946CA7F15A5B8AC9F29648DD41EA', false, 'yaml', '1', '[{"type":"del","key":"business","old":"auth:\n    yiBan:\n        appId: e4750b34230b48e1\n        appSecret: b0891a7f6018e5a76b085e3cb9548edd\n","cur":""}]');
INSERT INTO "public"."template" VALUES (25, 1720323538, 1728064733, 1, 'addr: 0.0.0.0:${GatewayServer.HttpPort}
name: gateway
version: v1
middlewares:
  - name: bbr
  - name: cors
    options:
      allowCredentials: true
      allowOrigins:
        - ''*''
      allowMethods:
        - GET
        - POST
        - PUT
        - DELETE
        - OPTIONS
      AllowHeaders:
        - Content-Type
        - Content-Length
        - Authorization
      ExposeHeaders:
        - Content-Length
        - Access-Control-Allow-Headers
  - name: tracing
  - name: logging
  - name: transcoder
endpoints:
  - path: /manager/api/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
  - path: /manager/client/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth
          method: POST
  - path: /configure/api/*
    timeout: ${ConfigureServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /resource/api/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
          whiteList:
            - path: /resource/api/v1/static/*
              method: GET
            - path: /resource/api/v1/download/*
              method: GET
  - path: /resource/client/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth
          method: POST
  - path: /usercenter/api/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
          whiteList:
            - path: /manager/api/v1/auth
              method: POST
  - path: /usercenter/client/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
  - path: /cron/api/*
    timeout: ${CronServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${CronServer.Host}:${CronServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/api/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/client/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth
          method: POST', '1802390697E147D536FEFB858CBB09BE', false, 'yaml', '新增usercenter auth', '[{"type":"update","key":"endpoints","old":"- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  path: /manager/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n  path: /manager/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /configure/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ConfigureServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n        whiteList:\n            - method: GET\n              path: /resource/api/v1/static/*\n            - method: GET\n              path: /resource/api/v1/download/*\n  path: /resource/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n  path: /resource/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /usercenter/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  path: /usercenter/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /cron/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${CronServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /partyaffairs/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\n  path: /partyaffairs/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n","cur":"- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  path: /manager/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n  path: /manager/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /configure/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ConfigureServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n        whiteList:\n            - method: GET\n              path: /resource/api/v1/static/*\n            - method: GET\n              path: /resource/api/v1/download/*\n  path: /resource/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n  path: /resource/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n        whiteList:\n            - method: POST\n              path: /manager/api/v1/auth\n  path: /usercenter/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  path: /usercenter/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /cron/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${CronServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /partyaffairs/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\n  path: /partyaffairs/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n"}]');
INSERT INTO "public"."template" VALUES (26, 1720323609, 1728064733, 1, 'addr: 0.0.0.0:${GatewayServer.HttpPort}
name: gateway
version: v1
middlewares:
  - name: bbr
  - name: cors
    options:
      allowCredentials: true
      allowOrigins:
        - ''*''
      allowMethods:
        - GET
        - POST
        - PUT
        - DELETE
        - OPTIONS
      AllowHeaders:
        - Content-Type
        - Content-Length
        - Authorization
      ExposeHeaders:
        - Content-Length
        - Access-Control-Allow-Headers
  - name: tracing
  - name: logging
  - name: transcoder
endpoints:
  - path: /manager/api/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
  - path: /manager/client/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth
          method: POST
  - path: /configure/api/*
    timeout: ${ConfigureServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /resource/api/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
          whiteList:
            - path: /resource/api/v1/static/*
              method: GET
            - path: /resource/api/v1/download/*
              method: GET
  - path: /resource/client/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth
          method: POST
  - path: /usercenter/api/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
          whiteList:
            - path: /usercenter/api/v1/auth
              method: POST
  - path: /usercenter/client/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
  - path: /cron/api/*
    timeout: ${CronServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${CronServer.Host}:${CronServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/api/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/client/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth
          method: POST', '6ED76BB69C66544B238C0F7CF70951FB', false, 'yaml', '1', '[{"type":"update","key":"endpoints","old":"- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  path: /manager/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n  path: /manager/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /configure/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ConfigureServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n        whiteList:\n            - method: GET\n              path: /resource/api/v1/static/*\n            - method: GET\n              path: /resource/api/v1/download/*\n  path: /resource/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n  path: /resource/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n        whiteList:\n            - method: POST\n              path: /manager/api/v1/auth\n  path: /usercenter/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  path: /usercenter/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /cron/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${CronServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /partyaffairs/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\n  path: /partyaffairs/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n","cur":"- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  path: /manager/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n  path: /manager/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /configure/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ConfigureServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n        whiteList:\n            - method: GET\n              path: /resource/api/v1/static/*\n            - method: GET\n              path: /resource/api/v1/download/*\n  path: /resource/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n  path: /resource/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n        whiteList:\n            - method: POST\n              path: /usercenter/api/v1/auth\n  path: /usercenter/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  path: /usercenter/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /cron/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${CronServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /partyaffairs/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\n  path: /partyaffairs/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n"}]');
INSERT INTO "public"."template" VALUES (27, 1720323742, 1728064733, 1, 'addr: 0.0.0.0:${GatewayServer.HttpPort}
name: gateway
version: v1
middlewares:
  - name: bbr
  - name: cors
    options:
      allowCredentials: true
      allowOrigins:
        - ''*''
      allowMethods:
        - GET
        - POST
        - PUT
        - DELETE
        - OPTIONS
      AllowHeaders:
        - Content-Type
        - Content-Length
        - Authorization
      ExposeHeaders:
        - Content-Length
        - Access-Control-Allow-Headers
  - name: tracing
  - name: logging
  - name: transcoder
endpoints:
  - path: /manager/api/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
  - path: /manager/client/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth
          method: POST
  - path: /configure/api/*
    timeout: ${ConfigureServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /resource/api/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
          whiteList:
            - path: /resource/api/v1/static/*
              method: GET
            - path: /resource/api/v1/download/*
              method: GET
  - path: /resource/client/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth
          method: POST
  - path: /usercenter/api/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /usercenter/client/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
  - path: /cron/api/*
    timeout: ${CronServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${CronServer.Host}:${CronServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/api/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/client/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth
          method: POST', '2B62DFD7A54233FC42AC86B7FFA9E044', false, 'yaml', '1', '[{"type":"update","key":"endpoints","old":"- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  path: /manager/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n  path: /manager/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /configure/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ConfigureServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n        whiteList:\n            - method: GET\n              path: /resource/api/v1/static/*\n            - method: GET\n              path: /resource/api/v1/download/*\n  path: /resource/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n  path: /resource/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n        whiteList:\n            - method: POST\n              path: /usercenter/api/v1/auth\n  path: /usercenter/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  path: /usercenter/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /cron/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${CronServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /partyaffairs/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/api/v1/auth\n  path: /partyaffairs/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n","cur":"- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  path: /manager/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n  path: /manager/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /configure/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ConfigureServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n        whiteList:\n            - method: GET\n              path: /resource/api/v1/static/*\n            - method: GET\n              path: /resource/api/v1/download/*\n  path: /resource/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n  path: /resource/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /usercenter/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  path: /usercenter/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /cron/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${CronServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /partyaffairs/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n  path: /partyaffairs/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n"}]');
INSERT INTO "public"."template" VALUES (28, 1720366812, 1728072395, 3, '
env: ${Env.Keyword}
server:
  http:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.HttpPort}
    timeout: ${ResourceServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.GrpcPort}
    timeout: ${ResourceServer.Timeout}
log:
  level: 0
  caller: true
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: ${ResourceDatabase.Type} #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: ${ResourceDatabase.Username}
      password: ${ResourceDatabase.Password}
      host: ${ResourceDatabase.Host}
      port: ${ResourceDatabase.Port}
      dbName: ${ResourceDatabase.Database}
      option: ${ResourceDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 3 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
business:
  chunkSize: ${ChunkSize}
  defaultMaxSize: ${DefaultMaxSize}
  defaultAcceptTypes: ${DefaultAcceptTypes}
  storage:
    type: local
    serverUrl: https://prod-gw.qlime.cn/resource/api/v1/static
    localDir: static
    temporaryExpire: 600s
    secret: limescloud
  export:
    serverUrl: https://prod-gw.qlime.cn/resource/api/v1/download
    localDir: static/export
    expire: 72h
', '7124FA23DF2383EFF43F397F75755BFB', false, 'yaml', '初始化配置', '[{"type":"update","key":"business","old":"chunkSize: ${ChunkSize}\ndefaultAcceptTypes: ${DefaultAcceptTypes}\ndefaultMaxSize: ${DefaultMaxSize}\nexport:\n    expire: 72h\n    localDir: static/export\n    serverUrl: http://${GatewayServer.Host}:${GatewayServer.HttpPort}/resource/api/v1/download\nstorage:\n    localDir: static\n    secret: limescloud\n    serverUrl: http://${GatewayServer.Host}:${GatewayServer.HttpPort}/resource/api/v1/static\n    temporaryExpire: 600s\n    type: local\n","cur":"chunkSize: ${ChunkSize}\ndefaultAcceptTypes: ${DefaultAcceptTypes}\ndefaultMaxSize: ${DefaultMaxSize}\nexport:\n    expire: 72h\n    localDir: static/export\n    serverUrl: https://prod-gw.qlime.cn/resource/api/v1/download\nstorage:\n    localDir: static\n    secret: limescloud\n    serverUrl: https://prod-gw.qlime.cn/resource/api/v1/static\n    temporaryExpire: 600s\n    type: local\n"}]');
INSERT INTO "public"."template" VALUES (29, 1723029256, 1728496716, 9, '
env: ${Env.Keyword}
server:
  http:
    host: ${PovertyServer.Host}
    port: ${PovertyServer.Port}
    timeout: ${PovertyServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
client:
  - server: Resource
    type: direct
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.GrpcPort}
  - server: UserCenter
    type: direct
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.GrpcPort}
log:
  level: 0
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: ${PovertyDatabase.Type} #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: ${PovertyDatabase.Username}
      password: ${PovertyDatabase.Password}
      host: ${PovertyDatabase.Host}
      port: ${PovertyDatabase.Port}
      dbName: ${PovertyDatabase.Database}
      option: ${PovertyDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 100 #最大连接数量
      maxIdleConn: 50 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
pool: #并非池配置
  size: 1000 #最大协程数量
  expiryDuration: 30s #过期时间
  preAlloc: true #是否预分配
  maxBlockingTasks: 1000 #最大的并发任务
  nonblocking: true
business:
  aiBot:
    secret: ${botKey}
    appid: ${botId}
    setting: ${botSetting}
', 'C590B3EEA4E514683EEBE01F293ACF39', false, 'yaml', '初始化提交', '');
INSERT INTO "public"."template" VALUES (30, 1723030749, 1728064733, 1, 'addr: 0.0.0.0:${GatewayServer.HttpPort}
name: gateway
version: v1
middlewares:
  - name: bbr
  - name: cors
    options:
      allowCredentials: true
      allowOrigins:
        - ''*''
      allowMethods:
        - GET
        - POST
        - PUT
        - DELETE
        - OPTIONS
      AllowHeaders:
        - Content-Type
        - Content-Length
        - Authorization
      ExposeHeaders:
        - Content-Length
        - Access-Control-Allow-Headers
  - name: tracing
  - name: logging
  - name: transcoder
endpoints:
  - path: /manager/api/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
  - path: /manager/client/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth
          method: POST
  - path: /configure/api/*
    timeout: ${ConfigureServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /resource/api/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
          whiteList:
            - path: /resource/api/v1/static/*
              method: GET
            - path: /resource/api/v1/download/*
              method: GET
  - path: /resource/client/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth
          method: POST
  - path: /usercenter/api/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /usercenter/client/*
    timeout: ${UserCenterServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
  - path: /cron/api/*
    timeout: ${CronServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${CronServer.Host}:${CronServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/api/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/client/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth
          method: POST
  - path: /poverty/api/*
    timeout: ${PovertyServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PovertyServer.Host}:${PovertyServer.Port}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /poverty/client/*
    timeout: ${PovertyServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PovertyServer.Host}:${PovertyServer.Port}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth
          method: POST', '634F4B998D8891899B0E918F95BC364B', false, 'yaml', '新增poverty', '[{"type":"update","key":"endpoints","old":"- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  path: /manager/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n  path: /manager/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /configure/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ConfigureServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n        whiteList:\n            - method: GET\n              path: /resource/api/v1/static/*\n            - method: GET\n              path: /resource/api/v1/download/*\n  path: /resource/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n  path: /resource/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /usercenter/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  path: /usercenter/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /cron/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${CronServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /partyaffairs/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n  path: /partyaffairs/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n","cur":"- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  path: /manager/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n  path: /manager/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /configure/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ConfigureServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n        whiteList:\n            - method: GET\n              path: /resource/api/v1/static/*\n            - method: GET\n              path: /resource/api/v1/download/*\n  path: /resource/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n  path: /resource/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /usercenter/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  path: /usercenter/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /cron/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${CronServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /partyaffairs/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n  path: /partyaffairs/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n- backends:\n    - target: ${PovertyServer.Host}:${PovertyServer.Port}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /poverty/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PovertyServer.Timeout}\n- backends:\n    - target: ${PovertyServer.Host}:${PovertyServer.Port}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n  path: /poverty/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PovertyServer.Timeout}\n"}]');
INSERT INTO "public"."template" VALUES (31, 1728064733, 1728064733, 1, 'addr: 0.0.0.0:${GatewayServer.HttpPort}
name: gateway
version: v1
middlewares:
  - name: bbr
  - name: cors
    options:
      allowCredentials: true
      allowOrigins:
        - ''*''
      allowMethods:
        - GET
        - POST
        - PUT
        - DELETE
        - OPTIONS
      AllowHeaders:
        - Content-Type
        - Content-Length
        - Authorization
      ExposeHeaders:
        - Content-Length
        - Access-Control-Allow-Headers
  - name: tracing
  - name: logging
  - name: transcoder
endpoints:
  - path: /manager/api/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
  - path: /manager/client/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/application/client/v1/auth
          method: POST
  - path: /configure/api/*
    timeout: ${ConfigureServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /resource/api/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
          whiteList:
            - path: /resource/api/v1/static/*
              method: GET
            - path: /resource/api/v1/download/*
              method: GET
  - path: /resource/client/*
    timeout: ${ResourceServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/application/client/v1/auth
          method: POST
  - path: /application/api/*
    timeout: ${ApplicationServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ApplicationServer.Host}:${ApplicationServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /application/client/*
    timeout: ${ApplicationServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ApplicationServer.Host}:${ApplicationServer.HttpPort}
  - path: /cron/api/*
    timeout: ${CronServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${CronServer.Host}:${CronServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/api/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /partyaffairs/client/*
    timeout: ${PartyAffairsServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/application/client/v1/auth
          method: POST
  - path: /poverty/api/*
    timeout: ${PovertyServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PovertyServer.Host}:${PovertyServer.Port}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth
          method: POST
  - path: /poverty/client/*
    timeout: ${PovertyServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${PovertyServer.Host}:${PovertyServer.Port}
    middlewares:
      - name: auth
        options:
          url: http://localhost:${GatewayServer.HttpPort}/application/client/v1/auth
          method: POST', '3020B1B4C417B67FE2480669A6087A34', true, 'yaml', '新增应用中心', '[{"type":"update","key":"endpoints","old":"- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  path: /manager/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n  path: /manager/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /configure/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ConfigureServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n        whiteList:\n            - method: GET\n              path: /resource/api/v1/static/*\n            - method: GET\n              path: /resource/api/v1/download/*\n  path: /resource/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n  path: /resource/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /usercenter/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}\n  path: /usercenter/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${UserCenterServer.Timeout}\n- backends:\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /cron/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${CronServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /partyaffairs/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n  path: /partyaffairs/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n- backends:\n    - target: ${PovertyServer.Host}:${PovertyServer.Port}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /poverty/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PovertyServer.Timeout}\n- backends:\n    - target: ${PovertyServer.Host}:${PovertyServer.Port}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/usercenter/client/v1/auth\n  path: /poverty/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PovertyServer.Timeout}\n","cur":"- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  path: /manager/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/application/client/v1/auth\n  path: /manager/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ManagerServer.Timeout}\n- backends:\n    - target: ${ConfigureServer.Host}:${ConfigureServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /configure/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ConfigureServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n        whiteList:\n            - method: GET\n              path: /resource/api/v1/static/*\n            - method: GET\n              path: /resource/api/v1/download/*\n  path: /resource/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/application/client/v1/auth\n  path: /resource/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ResourceServer.Timeout}\n- backends:\n    - target: ${ApplicationServer.Host}:${ApplicationServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /application/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ApplicationServer.Timeout}\n- backends:\n    - target: ${ApplicationServer.Host}:${ApplicationServer.HttpPort}\n  path: /application/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${ApplicationServer.Timeout}\n- backends:\n    - target: ${CronServer.Host}:${CronServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /cron/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${CronServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /partyaffairs/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n- backends:\n    - target: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/application/client/v1/auth\n  path: /partyaffairs/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PartyAffairsServer.Timeout}\n- backends:\n    - target: ${PovertyServer.Host}:${PovertyServer.Port}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/manager/api/v1/auth\n  path: /poverty/api/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PovertyServer.Timeout}\n- backends:\n    - target: ${PovertyServer.Host}:${PovertyServer.Port}\n  middlewares:\n    - name: auth\n      options:\n        method: POST\n        url: http://localhost:${GatewayServer.HttpPort}/application/client/v1/auth\n  path: /poverty/client/*\n  protocol: HTTP\n  responseFormat: true\n  timeout: ${PovertyServer.Timeout}\n"}]');
INSERT INTO "public"."template" VALUES (32, 1728064794, 1728064794, 5, '
env: ${Env.Keyword}
server:
  http:
    host: ${ApplicationServer.Host}
    port: ${ApplicationServer.HttpPort}
    timeout: ${ApplicationServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ApplicationServer.Host}
    port: ${ApplicationServer.GrpcPort}
    timeout: ${ApplicationServer.Timeout}
client:
  - server: Resource
    type: direct
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.GrpcPort}
log:
  level: 0
  caller: true
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: ${ApplicationDatabase.Type} #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: ${ApplicationDatabase.Username}
      password: ${ApplicationDatabase.Password}
      host: ${ApplicationDatabase.Host}
      port: ${ApplicationDatabase.Port}
      dbName: ${ApplicationDatabase.Database}
      option: ${ApplicationDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
captcha:
  loginImage:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 200
    skew: 0.7
    dotCount: 80
    refresh: true
  bindImage:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 200
    skew: 0.7
    dotCount: 80
    refresh: true
  registerImage:
    type: image
    length: 6
    expire: 180s
    redis: cache
    height: 80
    width: 200
    skew: 0.7
    dotCount: 80
    refresh: true
  loginEmail:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: login
  bindEmail:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: bind
  registerEmail:
    type: email
    length: 6
    expire: 180s
    redis: cache
    template: register
loader:
  password: static/cert/password.pem
email:
  template:
    login:
      subject: 登录验证码发送通知
      path: static/template/email/default.html
      type: text/html
    bind:
      subject: 绑定验证码发送通知
      path: static/template/email/default.html
      type: text/html
    register:
      subject: 注册验证码发送通知
      path: static/template/email/default.html
      type: text/html
  user: ${Email.Username}
  name: ${Email.Company}
  host: ${Email.Host}
  port: ${Email.Port}
  password: ${Email.Password}
jwt:
  redis: cache
  secret: ${ClientJwt.Secret}
  expire: ${ClientJwt.Expire}
  renewal: ${ClientJwt.Renewal}
  whitelist: ${ClientJwt.Whitelist}
business:
  defaultPasswordExpire: 10s
  defaultNickName: 用户-
', 'AB3C721A133F848EDC1A4A015271B274', true, 'yaml', '新增应用中心', '[{"type":"update","key":"server","old":"grpc:\n    host: ${UserCenterServer.Host}\n    port: ${UserCenterServer.GrpcPort}\n    timeout: ${UserCenterServer.Timeout}\nhttp:\n    host: ${UserCenterServer.Host}\n    marshal:\n        emitUnpopulated: true\n        useProtoNames: true\n    port: ${UserCenterServer.HttpPort}\n    timeout: ${UserCenterServer.Timeout}\n","cur":"grpc:\n    host: ${ApplicationServer.Host}\n    port: ${ApplicationServer.GrpcPort}\n    timeout: ${ApplicationServer.Timeout}\nhttp:\n    host: ${ApplicationServer.Host}\n    marshal:\n        emitUnpopulated: true\n        useProtoNames: true\n    port: ${ApplicationServer.HttpPort}\n    timeout: ${ApplicationServer.Timeout}\n"},{"type":"update","key":"database","old":"system:\n    autoCreate: true\n    config:\n        initializer:\n            enable: true\n            path: deploy/data.sql\n        logLevel: 4\n        maxIdleConn: 10\n        maxLifetime: 2h\n        maxOpenConn: 20\n        slowThreshold: 2s\n        transformError:\n            enable: true\n    connect:\n        dbName: ${UserCenterDatabase.Database}\n        host: ${UserCenterDatabase.Host}\n        option: ${UserCenterDatabase.Option}\n        password: ${UserCenterDatabase.Password}\n        port: ${UserCenterDatabase.Port}\n        username: ${UserCenterDatabase.Username}\n    drive: ${UserCenterDatabase.Type}\n    enable: true\n","cur":"system:\n    autoCreate: true\n    config:\n        initializer:\n            enable: true\n            path: deploy/data.sql\n        logLevel: 4\n        maxIdleConn: 10\n        maxLifetime: 2h\n        maxOpenConn: 20\n        slowThreshold: 2s\n        transformError:\n            enable: true\n    connect:\n        dbName: ${ApplicationDatabase.Database}\n        host: ${ApplicationDatabase.Host}\n        option: ${ApplicationDatabase.Option}\n        password: ${ApplicationDatabase.Password}\n        port: ${ApplicationDatabase.Port}\n        username: ${ApplicationDatabase.Username}\n    drive: ${ApplicationDatabase.Type}\n    enable: true\n"}]');
INSERT INTO "public"."template" VALUES (33, 1728071372, 1728071372, 6, '
env: ${Env.Keyword}
server:
  http:
    host: ${PartyAffairsServer.Host}
    port: ${PartyAffairsServer.HttpPort}
    timeout: ${PartyAffairsServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${PartyAffairsServer.Host}
    port: ${PartyAffairsServer.GrpcPort}
    timeout: ${PartyAffairsServer.Timeout}
client:
  - server: Resource
    type: direct
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.GrpcPort}
  - server: Application
    type: direct
    backends:
      - target: ${ApplicationServer.Host}:${ApplicationServer.GrpcPort}
log:
  level: 0
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: ${PartyAffairsDatabase.Type} #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: ${PartyAffairsDatabase.Username}
      password: ${PartyAffairsDatabase.Password}
      host: ${PartyAffairsDatabase.Host}
      port: ${PartyAffairsDatabase.Port}
      dbName: ${PartyAffairsDatabase.Database}
      option: ${PartyAffairsDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
', 'D0ACCE3CA222DC8765CD13AAFD841EDD', true, 'yaml', '修改应用中心', '[{"type":"update","key":"client","old":"- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.GrpcPort}\n  server: Resource\n  type: direct\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.GrpcPort}\n  server: UserCenter\n  type: direct\n","cur":"- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.GrpcPort}\n  server: Resource\n  type: direct\n- backends:\n    - target: ${ApplicationServer.Host}:${ApplicationServer.GrpcPort}\n  server: Application\n  type: direct\n"}]');
INSERT INTO "public"."template" VALUES (34, 1728072395, 1728072395, 3, '
env: ${Env.Keyword}
server:
  http:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.HttpPort}
    timeout: ${ResourceServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    host: ${ResourceServer.Host}
    port: ${ResourceServer.GrpcPort}
    timeout: ${ResourceServer.Timeout}
log:
  level: 0
  caller: true
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: ${ResourceDatabase.Type} #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: ${ResourceDatabase.Username}
      password: ${ResourceDatabase.Password}
      host: ${ResourceDatabase.Host}
      port: ${ResourceDatabase.Port}
      dbName: ${ResourceDatabase.Database}
      option: ${ResourceDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
redis:
  cache:
    enable: true
    host: ${Redis.Host}:${Redis.Port}
    username: ${Redis.Username}
    password: ${Redis.Password}
business:
  chunkSize: ${ChunkSize}
  defaultMaxSize: ${DefaultMaxSize}
  defaultAcceptTypes: ${DefaultAcceptTypes}
  storage:
    type: local
    serverUrl: https://prod-gw.qlime.cn/resource/api/v1/static
    localDir: static
    temporaryExpire: 600s
    secret: limescloud
  export:
    serverUrl: https://prod-gw.qlime.cn/resource/api/v1/download
    localDir: static/export
    expire: 72h
', 'A940188BE20724F72D54E3673D51C7FD', true, 'yaml', '新增数据库日志', '[{"type":"update","key":"database","old":"system:\n    autoCreate: true\n    config:\n        initializer:\n            enable: true\n            path: deploy/data.sql\n        logLevel: 3\n        maxIdleConn: 10\n        maxLifetime: 2h\n        maxOpenConn: 20\n        slowThreshold: 2s\n        transformError:\n            enable: true\n    connect:\n        dbName: ${ResourceDatabase.Database}\n        host: ${ResourceDatabase.Host}\n        option: ${ResourceDatabase.Option}\n        password: ${ResourceDatabase.Password}\n        port: ${ResourceDatabase.Port}\n        username: ${ResourceDatabase.Username}\n    drive: ${ResourceDatabase.Type}\n    enable: true\n","cur":"system:\n    autoCreate: true\n    config:\n        initializer:\n            enable: true\n            path: deploy/data.sql\n        logLevel: 4\n        maxIdleConn: 10\n        maxLifetime: 2h\n        maxOpenConn: 20\n        slowThreshold: 2s\n        transformError:\n            enable: true\n    connect:\n        dbName: ${ResourceDatabase.Database}\n        host: ${ResourceDatabase.Host}\n        option: ${ResourceDatabase.Option}\n        password: ${ResourceDatabase.Password}\n        port: ${ResourceDatabase.Port}\n        username: ${ResourceDatabase.Username}\n    drive: ${ResourceDatabase.Type}\n    enable: true\n"}]');
INSERT INTO "public"."template" VALUES (35, 1728496716, 1728496716, 9, '
env: ${Env.Keyword}
server:
  http:
    host: ${PovertyServer.Host}
    port: ${PovertyServer.Port}
    timeout: ${PovertyServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
client:
  - server: Resource
    type: direct
    backends:
      - target: ${ResourceServer.Host}:${ResourceServer.GrpcPort}
  - server: Application
    type: direct
    backends:
      - target: ${ApplicationServer.Host}:${ApplicationServer.GrpcPort}
log:
  level: 0
  output:
    - stdout
    - file
  file:
    name: ./tmp/runtime/output.log
    maxSize: 1
    maxBackup: 5
    maxAge: 1
    compress: false
database:
  system:
    enable: true #是否启用数据库
    drive: ${PovertyDatabase.Type} #数据库类型
    autoCreate: true #是否自动创建数据库
    connect:
      username: ${PovertyDatabase.Username}
      password: ${PovertyDatabase.Password}
      host: ${PovertyDatabase.Host}
      port: ${PovertyDatabase.Port}
      dbName: ${PovertyDatabase.Database}
      option: ${PovertyDatabase.Option}
    config:
      transformError:
        enable: true
      initializer:
        enable: true
        path: deploy/data.sql
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 100 #最大连接数量
      maxIdleConn: 50 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值
pool: #并非池配置
  size: 1000 #最大协程数量
  expiryDuration: 30s #过期时间
  preAlloc: true #是否预分配
  maxBlockingTasks: 1000 #最大的并发任务
  nonblocking: true
business:
  aiBot:
    secret: ${botKey}
    appid: ${botId}
    setting: ${botSetting}
', 'A52745FD56B8F35A884788911B482BA5', true, 'yaml', '修改应用中心', '[{"type":"update","key":"client","old":"- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.GrpcPort}\n  server: Resource\n  type: direct\n- backends:\n    - target: ${UserCenterServer.Host}:${UserCenterServer.GrpcPort}\n  server: UserCenter\n  type: direct\n","cur":"- backends:\n    - target: ${ResourceServer.Host}:${ResourceServer.GrpcPort}\n  server: Resource\n  type: direct\n- backends:\n    - target: ${ApplicationServer.Host}:${ApplicationServer.GrpcPort}\n  server: Application\n  type: direct\n"}]');

-- ----------------------------
-- Indexes structure for table business
-- ----------------------------
CREATE INDEX "fk_business_server" ON "public"."business" USING btree (
                                                                      "server_id" "pg_catalog"."int8_ops" ASC NULLS LAST
    );
CREATE INDEX "idx_business_created_at" ON "public"."business" USING btree (
                                                                           "created_at" "pg_catalog"."int8_ops" ASC NULLS LAST
    );
CREATE INDEX "idx_business_updated_at" ON "public"."business" USING btree (
                                                                           "updated_at" "pg_catalog"."int8_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Primary Key structure for table business
-- ----------------------------
ALTER TABLE "public"."business" ADD CONSTRAINT "business_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table business_value
-- ----------------------------
CREATE UNIQUE INDEX "env_id" ON "public"."business_value" USING btree (
                                                                       "env_id" "pg_catalog"."int8_ops" ASC NULLS LAST,
                                                                       "business_id" "pg_catalog"."int8_ops" ASC NULLS LAST
    );
CREATE INDEX "fk_business_value_env" ON "public"."business_value" USING btree (
                                                                               "env_id" "pg_catalog"."int8_ops" ASC NULLS LAST
    );
CREATE INDEX "fk_business_values" ON "public"."business_value" USING btree (
                                                                            "business_id" "pg_catalog"."int8_ops" ASC NULLS LAST
    );
CREATE INDEX "idx_business_value_created_at" ON "public"."business_value" USING btree (
                                                                                       "created_at" "pg_catalog"."int8_ops" ASC NULLS LAST
    );
CREATE INDEX "idx_business_value_updated_at" ON "public"."business_value" USING btree (
                                                                                       "updated_at" "pg_catalog"."int8_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Primary Key structure for table business_value
-- ----------------------------
ALTER TABLE "public"."business_value" ADD CONSTRAINT "business_value_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table configure
-- ----------------------------
CREATE INDEX "fk_configure_env" ON "public"."configure" USING btree (
                                                                     "env_id" "pg_catalog"."int8_ops" ASC NULLS LAST
    );
CREATE INDEX "fk_configure_server" ON "public"."configure" USING btree (
                                                                        "server_id" "pg_catalog"."int8_ops" ASC NULLS LAST
    );
CREATE INDEX "idx_configure_created_at" ON "public"."configure" USING btree (
                                                                             "created_at" "pg_catalog"."int8_ops" ASC NULLS LAST
    );
CREATE INDEX "idx_configure_updated_at" ON "public"."configure" USING btree (
                                                                             "updated_at" "pg_catalog"."int8_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Primary Key structure for table configure
-- ----------------------------
ALTER TABLE "public"."configure" ADD CONSTRAINT "configure_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table env
-- ----------------------------
CREATE INDEX "idx_env_created_at" ON "public"."env" USING btree (
                                                                 "created_at" "pg_catalog"."int8_ops" ASC NULLS LAST
    );
CREATE INDEX "idx_env_updated_at" ON "public"."env" USING btree (
                                                                 "updated_at" "pg_catalog"."int8_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Primary Key structure for table env
-- ----------------------------
ALTER TABLE "public"."env" ADD CONSTRAINT "env_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table gorm_init
-- ----------------------------
ALTER TABLE "public"."gorm_init" ADD CONSTRAINT "gorm_init_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table resource
-- ----------------------------
CREATE INDEX "idx_resource_created_at" ON "public"."resource" USING btree (
                                                                           "created_at" "pg_catalog"."int8_ops" ASC NULLS LAST
    );
CREATE INDEX "idx_resource_updated_at" ON "public"."resource" USING btree (
                                                                           "updated_at" "pg_catalog"."int8_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Primary Key structure for table resource
-- ----------------------------
ALTER TABLE "public"."resource" ADD CONSTRAINT "resource_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table resource_server
-- ----------------------------
CREATE INDEX "fk_resource_resource_server" ON "public"."resource_server" USING btree (
                                                                                      "resource_id" "pg_catalog"."int8_ops" ASC NULLS LAST
    );
CREATE INDEX "idx_resource_server_created_at" ON "public"."resource_server" USING btree (
                                                                                         "created_at" "pg_catalog"."int8_ops" ASC NULLS LAST
    );
CREATE UNIQUE INDEX "sr" ON "public"."resource_server" USING btree (
                                                                    "server_id" "pg_catalog"."int8_ops" ASC NULLS LAST,
                                                                    "resource_id" "pg_catalog"."int8_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Primary Key structure for table resource_server
-- ----------------------------
ALTER TABLE "public"."resource_server" ADD CONSTRAINT "resource_server_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table resource_value
-- ----------------------------

ALTER TABLE resource_value
    ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (START WITH 140);

-- 唯一约束（名称和列相同）
ALTER TABLE resource_value ADD CONSTRAINT er UNIQUE (env_id, resource_id);

CREATE INDEX "fk_resource_resource_value" ON "public"."resource_value" USING btree (
                                                                                    "resource_id" "pg_catalog"."int8_ops" ASC NULLS LAST
    );
CREATE INDEX "idx_resource_value_created_at" ON "public"."resource_value" USING btree (
                                                                                       "created_at" "pg_catalog"."int8_ops" ASC NULLS LAST
    );
CREATE INDEX "idx_resource_value_updated_at" ON "public"."resource_value" USING btree (
                                                                                       "updated_at" "pg_catalog"."int8_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Primary Key structure for table resource_value
-- ----------------------------
ALTER TABLE "public"."resource_value" ADD CONSTRAINT "resource_value_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table server
-- ----------------------------
CREATE INDEX "idx_server_created_at" ON "public"."server" USING btree (
                                                                       "created_at" "pg_catalog"."int8_ops" ASC NULLS LAST
    );
CREATE INDEX "idx_server_updated_at" ON "public"."server" USING btree (
                                                                       "updated_at" "pg_catalog"."int8_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Primary Key structure for table server
-- ----------------------------
ALTER TABLE "public"."server" ADD CONSTRAINT "server_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table template
-- ----------------------------
CREATE INDEX "idx_template_created_at" ON "public"."template" USING btree (
                                                                           "created_at" "pg_catalog"."int8_ops" ASC NULLS LAST
    );
CREATE INDEX "idx_template_updated_at" ON "public"."template" USING btree (
                                                                           "updated_at" "pg_catalog"."int8_ops" ASC NULLS LAST
    );
CREATE UNIQUE INDEX "sv" ON "public"."template" USING btree (
                                                             "server_id" "pg_catalog"."int8_ops" ASC NULLS LAST,
                                                             "version" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Primary Key structure for table template
-- ----------------------------
ALTER TABLE "public"."template" ADD CONSTRAINT "template_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Foreign Keys structure for table business
-- ----------------------------
ALTER TABLE "public"."business" ADD CONSTRAINT "fk_business_server" FOREIGN KEY ("server_id") REFERENCES "public"."server" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table business_value
-- ----------------------------
ALTER TABLE "public"."business_value" ADD CONSTRAINT "fk_business_value_env" FOREIGN KEY ("env_id") REFERENCES "public"."env" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "public"."business_value" ADD CONSTRAINT "fk_business_values" FOREIGN KEY ("business_id") REFERENCES "public"."business" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table configure
-- ----------------------------
ALTER TABLE "public"."configure" ADD CONSTRAINT "fk_configure_env" FOREIGN KEY ("env_id") REFERENCES "public"."env" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "public"."configure" ADD CONSTRAINT "fk_configure_server" FOREIGN KEY ("server_id") REFERENCES "public"."server" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table resource_server
-- ----------------------------
ALTER TABLE "public"."resource_server" ADD CONSTRAINT "fk_resource_resource_server" FOREIGN KEY ("resource_id") REFERENCES "public"."resource" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "public"."resource_server" ADD CONSTRAINT "fk_resource_server_server" FOREIGN KEY ("server_id") REFERENCES "public"."server" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table resource_value
-- ----------------------------
ALTER TABLE "public"."resource_value" ADD CONSTRAINT "fk_resource_resource_value" FOREIGN KEY ("resource_id") REFERENCES "public"."resource" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "public"."resource_value" ADD CONSTRAINT "fk_resource_value_env" FOREIGN KEY ("env_id") REFERENCES "public"."env" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table template
-- ----------------------------
ALTER TABLE "public"."template" ADD CONSTRAINT "fk_template_server" FOREIGN KEY ("server_id") REFERENCES "public"."server" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;


