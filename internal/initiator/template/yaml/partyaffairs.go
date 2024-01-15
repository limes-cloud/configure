package yaml

var PartyAffairs = `
env: ${Env.Keyword}
server:
  http:
    addr: ${PartyAffairsServer.Host}:${PartyAffairsServer.HttpPort}
    timeout: ${PartyAffairsServer.Timeout}
    marshal:
      emitUnpopulated: true
      useProtoNames: true
  grpc:
    addr: ${PartyAffairsServer.Host}:${PartyAffairsServer.GrpcPort}
    timeout: ${PartyAffairsServer.Timeout}
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
      maxLifetime: 2h #最大生存时间
      maxOpenConn: 20 #最大连接数量
      maxIdleConn: 10 #最大空闲数量
      logLevel: 4 #日志等级
      slowThreshold: 2s #慢sql阈值


`
