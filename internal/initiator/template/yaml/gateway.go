package yaml

var Gateway = `
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
        - '*'
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
  - path: /manager/*
    timeout: ${ManagerServer.Timeout}
    protocol: HTTP
    responseFormat: true
    backends:
      - target: ${ManagerServer.Host}:${ManagerServer.HttpPort}
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
  - path: /resource/*
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

`
