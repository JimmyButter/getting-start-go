Nacos 安装
省流Tips：nacos官方文档的安装方式是docker-compose，我在下文用的是docker安装，并且已经有一台mysql 8.0服务器。
[https://nacos.io/zh-cn/docs/1.X/v2/quickstart/quick-start-docker/](https://nacos.io/zh-cn/docs/1.X/v2/quickstart/quick-start-docker/)
```java
docker run -p 8848:8848 --name nacos -d nacos/nacos-server:v2.1.0
```
建本地文件夹
```shell
cd ~/docker
mkdir nacos
mkdir -p nacos/{logs,conf}
```
把文件拷出来后删除docker
```shell
docker cp nacos:/home/nacos/logs/ ~/docker/nacos/
docker cp nacos:/home/nacos/conf/ ~/docker/nacos/
```
修改conf/application.properties
![image.png](https://cdn.nlark.com/yuque/0/2024/png/12869462/1712737746388-a63ede51-13ea-456c-86a6-1c5c90c01e7f.png#averageHue=%232d3037&clientId=u1a4a5956-263e-4&from=paste&height=407&id=u09bb6925&originHeight=814&originWidth=1846&originalType=binary&ratio=2&rotation=0&showTitle=false&size=252654&status=done&style=none&taskId=u647ceb11-e055-4d14-99af-e21225310ec&title=&width=923)
把对应的mysql库表建好
官方脚本：[https://raw.githubusercontent.com/alibaba/nacos/develop/distribution/conf/mysql-schema.sql](https://raw.githubusercontent.com/alibaba/nacos/develop/distribution/conf/mysql-schema.sql)
重新run Docker
```shell
docker run -d --name nacos -p 8848:8848  -p 9848:9848 -p 9849:9849 \
--privileged=true -e JVM_XMS=256m -e JVM_XMX=256m -e MODE=standalone \
-v ~/docker/nacos/logs:/home/nacos/logs \
-v ~/docker/nacos/conf:/home/nacos/conf \
--restart=always nacos/nacos-server:v2.1.0
```
运行成功，浏览器访问 http://127.0.0.1:8848/nacos 可以看到nacos控制台

通过这个请求可以查看实例列表
[http://127.0.0.1:8848/nacos/v1/ns/instance/list?serviceName=hertz.test.hello&namespaceId=d165c4e0-c76b-42cc-81bb-b173ffa5ad3d](http://127.0.0.1:8848/nacos/v1/ns/instance/list?serviceName=hertz.test.hello&namespaceId=d165c4e0-c76b-42cc-81bb-b173ffa5ad3d)
serviceName和namespaceId要替换成你实际的值
```json
{
  name: "DEFAULT_GROUP@@hertz.test.hello",
  groupName: "DEFAULT_GROUP",
  clusters: "",
  cacheMillis: 10000,
  hosts: [
  {
  instanceId: "127.0.0.1#8001#DEFAULT#DEFAULT_GROUP@@hertz.test.hello",
  ip: "127.0.0.1",
  port: 8001,
  weight: 10,
  healthy: true,
  enabled: true,
  ephemeral: true,
  clusterName: "DEFAULT",
  serviceName: "DEFAULT_GROUP@@hertz.test.hello",
  metadata: { },
  instanceHeartBeatInterval: 5000,
  instanceIdGenerator: "simple",
  instanceHeartBeatTimeOut: 15000,
  ipDeleteTimeout: 30000
  }
  ],
  lastRefTime: 1712651461836,
  checksum: "",
  allIPs: false,
  reachProtectionThreshold: false,
  valid: true
}
```
关于golang如何连接nacos，代码参考字节的hertz框架
[https://github.com/hertz-contrib/registry/blob/d07355f82f9f7ddddeb0231790a6c55210e5efd2/nacos/v2/examples/custom_config/client/main.go](https://github.com/hertz-contrib/registry/blob/d07355f82f9f7ddddeb0231790a6c55210e5efd2/nacos/v2/examples/custom_config/client/main.go)


