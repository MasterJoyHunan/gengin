## gengin 一个生成 gin 框架的脚手架

gengin 基于 go-zero 开发，是一个 go-zero 的插件，只需定义 .api 文件，一行命令就可以生成整个项目的基础文件

### 前言

正所谓站在巨人的肩膀上，看的比较远。

本项目基于 go-zero 开发的插件，只需定义 api 文件，即可生成 gin 框架工程目录，借鉴了 go-zero 的思想，
配置大于约定。指在提高开发人员的工作效率，减少沟通成本

使用本项目生成的代码，只会依赖 gin 框架，其他什么依赖都没有，非常简洁优雅

### 为什么需要它

你当然可以使用 goctl 生成一个 http 单体服务，但是你可能更习惯使用 gin 框架的老司机，
在新项目你在纠结如何选型的时，到底是纠结使用 gin 搭建简单的服务，还是使用 go-zero 框架利用其完善的组件的时候，
希望本项目对你有所帮助

以下用户可以考虑使用该项目

* 习惯使用 gin 框架
* 希望每个项目都是使用一样的规范，一样的依赖，从而减少学习成本
* 项目是一个简单的单体服务
* 项目里面使用自己习惯的第三方库，可能不需要用到 go-zero 的组件如
  * mysql
  * redis
  * logger
  * 限流
  * 熔断
  * .......
* 希望定义了一个 .api 文件，基础代码自动生成
* 希望用到 go-zero 丰富的生态

### 使用方法

#### 安装

前提是安装了go-zero，如果没有安装，则进行安装

go 1.16 以下使用
```sh
go get -u github.com/zeromicro/go-zero/tools/goctl
```
go 1.16 及以上使用
```sh
go install github.com/zeromicro/go-zero/tools/goctl@v1.3.8
```

再安装本项目，作为 go-zero 的插件

go 1.16 以下使用
```sh
go get -u github.com/MasterJoyHunan/gengin
```
go 1.16 及以上使用
```sh
go install github.com/MasterJoyHunan/gengin@v1.4.1
```

#### 初始化一个 GO 项目

```sh
mkdir you-application
cd you-application
go mod init you-app-pkg-name
```

#### 在项目下定义 you-app.api 文件

[api语法指南](https://go-zero.dev/cn/docs/design/grammar)

you-app.api 文件内容示例

```api
syntax = "v1"

info(
	title: "some app"
)

type bookRequest {
    Name string `json:"name"` // 姓名
    Age int `json:"age"`      // 年龄
}

type bookResponse {
    Code int `json:"code"` // 业务码
    Msg string `json:"msg"` // 业务消息
}

@server(
    jwt: Auth
    group: book
    middleware: SomeMiddleware,CorsMiddleware
    prefix: /v1
)

service someapp {
    @doc "获取所有书本信息"
    @handler getBookList
    get /book (bookRequest) returns (bookResponse)

    @doc "获取书本信息"
    @handler getBook
    get /book/:id (bookRequest) returns (bookResponse)

    @doc "添加书本信息"
    @handler addBook
    post /book (bookRequest) returns (bookResponse)

    @doc "获取书本信息"
    @handler editBook
    put /book/:id (bookRequest) returns (bookResponse)
}
```

#### 在项目下生成 gin 项目

[go-zero 插件使用教程](https://go-zero.dev/cn/docs/goctl/plugin)

```sh
goctl api plugin -p gengin -api xxx.api -dir .
```

#### 生成的目录结构如下

```
├─config       # 配置文件对应的 struct
├─etc          # yaml 配置文件
├─handler      # 控制器层
├─logic        # 服务层
├─middleware   # 中间件层
├─routes       # 路由定义
├─types        # 请求与相应的 struct
you-app.go     # main 文件
```

### 重新执行程序那些文件会重新生成

```
├─config       # 如果文件已存在，不会重新生成
├─etc          # 如果文件已存在，不会重新生成
├─handler      # 如果文件已存在，不会重新生成
├─logic        # 如果文件已存在，不会重新生成
├─middleware   # 如果文件已存在，不会重新生成
├─routes       # gengin生成的文件会重新生成，请不要修改。手动新增的文件不会
├─types        # gengin生成的文件会重新生成，请不要修改。手动新增的文件不会
you-app.go     # 如果文件已存在，不会重新生成
```

### 注意事项

* request 和 response 不支持数组格式
* get 和 post 请求都是使用 from 接受参数.
* 在 api 文件中定义 tag 为 path 的 tag 会转换为 uri，方便 gin 框架处理
* 在 api 文件中定义 @server 下的 jwt 会自动转换为一个 middleware，需要手动选择自己的熟悉 jwt 框架自行处理
* 在 api 文件中定义 @server 下的 group 使用驼峰命名，会自动分割文件夹，如：group=helloWord
  * logic => logic/hello/word/some_logic.go
  * handler => handler/hello/word/some_handle.go
* 在 api 文件中定义 @server 下的 group 如果为空，则会放入对应的根目录，如： 
  * logic => logic/some_logic.go
  * handler => handler/some_handle.go

### 其他

如果需要对 .api 文件生成 swagger 文档，请参考[https://github.com/zeromicro/goctl-swagger](https://github.com/zeromicro/goctl-swagger)

如果觉得该项目对你有所帮助，请不要吝啬你的小手，帮忙点个 stars

如果对本项目有更好的建议或意见，欢迎提交 pr / issues，或者联系本人 tanwuyang88@gmail.com

再次感谢 [go-zero](https://github.com/zeromicro/go-zero)

### 协议

[MIT](https://github.com/MasterJoyHunan/gengin/blob/master/LICENSE)