## ouho 后台服务目录结构

### cache

1. redis 本机内存及redis封装

- 全局引用的变量
- 对象池
- redis封装在该处


### config

2. config 配置类信息(yaml)

- config存放各类配置文件，及对配置文件的解析
- config解析并初始化 logger、redis、mysql、mongo等


### constant

3. constant 常量包

- 全局常量定义在此处
- 错误码及错误信息定义在此处
- 一些静态数据在此处

### routes

4. routes 路由及路由分组管理

- 业务接口的路由入口管理注册在这里
![image](http://note.youdao.com/yws/res/7125/24D95E13D3B5495DA86CC4CAD11C9ECE)


### service

5. service 业务逻辑层

- 业务逻辑相关，业务逻辑的检验 业务逻辑的实现等都在这里

### models

6. models 数据库层(mysql、mongo)

- 数据库的存取结构定义在此
- 所有对数据库的直接操作，实现在此
- 对数据库的封装定义在此

***一个表(一个集合)对应一个go文件，文件名以表名(集合名)命名， 根据需要适当增加目录。***

### protos

7. protos 通信协议定义在此

- 第三方协议（比如 protobuf）
- 自定义协议（比如 tcp自定义）
- 请求返回协议，也就是通用的请求返回结构

### library

8. library 通用的基本类库

- logger
- redis
- mysql
- mongo
- middleware
- ......
- tools (这里放自己本地封装的工具类)

### plugins

9. plugins 第三方及自定义服务分装类

- 自定义服务（如 tcp、udp）
- 第三方服务分装（诸如：grpc、 websocket，阿里OSS，阿里短信，网易IM）
- 或是自实现的文件服务器

***按服务适当增加目录***