package yaml

var UserCenter = `
env: ${Env.Keyword}
server:
  http:
    addr: ${UserCenterServer.Host}:${UserCenterServer.HttpPort}
    timeout: ${UserCenterServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    addr: ${UserCenterServer.Host}:${UserCenterServer.GrpcPort}
    timeout: ${UserCenterServer.Timeout}
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
  prefix: /user-center/admin/
  secret: ${ClientJwt.Secret}
  expire: ${ClientJwt.Expire}
  renewal: ${ClientJwt.Renewal}
  whitelist: ${ClientJwt.Whitelist}
business:
  service:
    resource: ${ResourceServer.Host}:${ResourceServer.GrpcPort}
`
