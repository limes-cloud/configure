package yaml

var Manager = `
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
`
