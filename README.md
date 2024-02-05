## Easy Note

> Forked from https://github.com/cloudwego/kitex-examples/tree/main/bizdemo/easy_note

用于个人学习。

### 项目启动

1. 克隆项目到本地/远程主机：

```bash
git clone https://github.com/Breeze-P/ez-note.git
cd ez-note
```

2. 安装`docker`、`docker-compose`
3. 使用`docker-compose`部署项目：

```bash
docker-compose up
```



### TODO

* 项目部署后，需要手动重启user、note服务，原因是mysql还没有初始化，user、note无法登陆`note`账户。



### 重写记录

* 初始化`idl`目录，应用`kitex`命令生成代码
* 重写`cmd`目录
* 修改`docker-compose.yal`实现联合部署



### 学习所得

微服务相较于单机服务区别：

* http服务器的handler由调用本机的service改成rpc调用其他server的服务



Get ETCD 服务注册与发现

* 根据servicename获取服务

Get Jaeger 微服务监控



errno要统一处理

resp要统一处理

* rpc的resp：pack
  * status_code
  * status_msg
* http的resp



脚本生成的：

* router（only in hertz
* handler/rpc-service
* dto
* req/resp



hanlder做服务启动后的错误处理，返回错误信息给前端



rpc的server结构：

* handler/rpc-service-imp（interface是生成的；创建server的时候注册上去
* service
* dal



rpc的Client结构：

* service（调用server的handler；供其自身的service调用



需要trace的：

* rpc的server和client
* gormopentraing
* jaeger监视器



需要etcd注册&发现的：

* rpc的server和client



两种环境变量：

* 服务主体对中间件的调用`pky/constants`
* 中间件自身启动的环境变量`.env`
