package yaml

var Resource = `
env: ${Env.Keyword}
server:
  http:
    addr: ${ResourceServer.Host}:${ResourceServer.HttpPort}
    timeout: ${ResourceServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    addr: ${ResourceServer.Host}:${ResourceServer.GrpcPort}
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

`
