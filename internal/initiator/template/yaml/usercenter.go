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


`
