# demo-chatroom

go+iris+jwt+mysql+gorm+viper，iris 项目实战简易聊天室，登录、注册、私聊、群聊。

## 项目启动

后端进行了升级改造，接口的返回格式有较大改动，前端尚未进行适配，暂时无法正常运行，参考效果请看demo网址！！！

项目目前只写了 mysql 相关的适配，其他数据库暂时没有支持。

项目只需要配置端口和数据库地址、用户名、密码以及数据库名，并提前建好数据库即可启动，不需要建表，便于快速启动

```cmd
git clone https://github.com/JabinGP/demo-chatroom.git
cd demo-chatroom
// 复制config.toml.example 为 config.toml 并填写数据库信息，或者可选修改端口号
go run main.go
```

默认为 8888 端口，启动后访问`http://localhost:8888`即可，或者访问演示地址`http://mike.jabingp.cn:8888`

### 前端

用了 react，但是没用 ui 框架，很多的小细节上表现并不好，凑合着看吧，重点放在后端。

聊天框设置了窗口自动滚动到底端，但是 api 是 react 提供的，发现在许多浏览器上并不兼容，使用 chrome 浏览器可以解决这个问题。

注册后手动返回选择登陆，消息框里面的红色名称为公共发言，灰色名称为私聊发言、可以在红色的框里面指定接收者的名称，如果不指定的话，默认是公共发言，指定后只有对应的用户能看到信息。

蓝色框内显示自己的用户名，点击即直接退出登录。

### 后端

api 格式基于 restful 设计，登录功能使用 jwt 完成，许多接口需要登录状态，请求的时候需要携带 JWT，具体请看[golang iris 的 jwt 实践](https://segmentfault.com/a/1190000021187448)，另外便于测试，JWT 签发有效时间只设置了 20 分钟，过期需要重新登录。

| 功能                | 请求方式 | 地址                                |
| ------------------- | -------- | ----------------------------------- |
| 获取登录 token      | POST     | http://localhost:8888/v1/login      |
| 查找用户            | GET      | http://localhost:8888/v1/user       |
| 注册                | POST     | http://localhost:8888/v1/user       |
| 用户自己修改信息    | PUT      | http://localhost:8888/v1/user       |
| 用户发送信息        | POST     | http://localhost:8888/v1/message    |
| 用户获取信息        | GET      | http://localhost:8888/v1/message    |
| 用户获取 token 信息 | GET      | http://localhost:8888/v1/token/info |

详细请求参数可以在[demo-chatroom 的 postman-api 文档](https://documenter.getpostman.com/view/7019042/SWECVaGy?version=latest)里查看。

或者查看源码，请求参数在`model/reqo`里面查看，响应参数可以在`model/reso`里查看

## 前言

聊天功能 AJAX 不是最好的选择，WebSocket 比较好，但是被要求使用了 AJAX 所以没有选择后者。

项目的前端比较简陋，因为只是作为 demo 使用。

英语不是很好，代码注释用英语只是因为懒得切换输入法。

第一次用 go 开发 web 项目，也是第一次用 react 写前端，由于前端没怎么注重项目结构（xjbx），就不放源码了，把项目编译后放在了 assets 文件夹下，可读性很差，但是可以和后端一起启动，不需要单独启动前端，比较方便查看效果。如果还有时间会考虑用原生写一个极简版的供大家参考原理。

第一次用 ORM 操作数据库，感觉好难用，我还是宁愿手写 sql，好多想要的效果翻半天文档都找不到解决方案，后期有机会考虑用 sqlx 重构。

## 项目来源

最近对 Go 比较有兴趣，又接到任务编写一个简易聊天室，发现目前 iris 的项目实践比较少，只有一些 HelloWorld 级别的示例，于是决定用 Go 来做，然后开源出来供大互相参考借鉴，当然项目结构如何设计完全基于我有限的开发经验，对于不合理的地方，请给出你宝贵的意见。

## 项目要点

这个项目有如下要求

1. 基于 AJAX
2. 前端页面需要无刷新（自动更新数据）
3. 登录功能
4. 注册功能（要求用户有用户名、密码、性别、年龄、兴趣爱好）

## 实现思路

### 登陆功能

登陆功能这次选用`JWT`来实现，`JWT`和`Session`各自的优劣就不再赘述。

### 基于 AJAX，且无刷新

基于 AJAX 是所有前后端分离项目的必备，因此这个功能不过多讨论，这里重点在于无刷新，难点在哪？

#### 用户操作需求

用户的操作逻辑是，在聊天室里面发送数据，然后数据就被发出去，聊天界面要显示出自己发送的数据，以及要实时更新别人发出来的数据。

#### 前端的操作逻辑

前端和后端之间是通过 AJAX 来交流的，前端发送数据和后端发送数据可以表现为

- 发送数据：前端将需要发送的数据以 JSON 格式携带在请求里，请求对于结构
- 获取数据：前端请求后台获取消息的接口，获取最新消息

这里有什么问题？问就在前端永远只能主动发起请求，而后端永远只能接受请求。这意味着最新的消息永远无法实时地从后端主动发送给前端，最新的消息只能先存放在后端，然后等待前端发起请求，后端才能返回数据。

由于后端是没有能力主动推送消息给前端的，因此用户获取最新数据的解决方法是前端设置一个定时器`每隔一段比较短的时间就请求一次后台接口（轮询）`，这样就能不断更新数据。

#### 后端的操作逻辑

前端已经确定使用 AJAX 定时轮询后台接口来获取最新数据，为了数据实时性，轮询间隔会`小于1s`，这样又会带来另个问题，后端在如此频繁的请求下，一定不能每次都将所有数据都传输出去，一是数据大小导致的网路传输效率、流量成本问题，二是数据大小导致的前端判断新数据的效率问题，那么后端每次必须都返回前端还没有接收过的数据，而问题在于--后端怎么知道前端已经接收了哪些信息？

这个就要利用到消息的`自增主键`，只需要前端每次请求的时候都携带上前端已经接收的`最后的消息的主键`，由于主键是不重复且自增的，我们可以很轻松的找出比该主键大的数据，也就是前端还没接收到的数据。

## 项目技术栈

- 语言

  - Golang
  - HTML
  - CSS
  - JavaScript

- 框架

  - Iris 后端框架
  - React 前端框架
  - Gorm 数据库 ORM 框架
  - Viper 多类型配置文件读取支持

- 数据存储

  - Mysql 经典数据库

- 技术
  - JWT 签发登陆令牌
  - AJAX 异步请求后端数据

## 数据库设计结构

> 由于使用了 Gorm 数据库 ORM 框架，以下表都是自动生成的，自带了`xxxxxx_at`字段

基于如上的需求，设计了`users`和`messages`两个表

### users

关键字段

- id
- username
- passwd
- gender
- age
- interest

数据库表结构

| Field      | Type               | Null | Key | Default | Extra          |
| :--------- | :----------------- | :--- | :-- | :------ | :------------- |
| id         | int\(10\) unsigned | NO   | PRI | NULL    | auto_increment |
| created_at | timestamp          | YES  |     | NULL    |                |
| updated_at | timestamp          | YES  |     | NULL    |                |
| deleted_at | timestamp          | YES  | MUL | NULL    |                |
| username   | varchar\(255\)     | YES  |     | NULL    |                |
| passwd     | varchar\(255\)     | YES  |     | NULL    |                |
| gender     | bigint\(20\)       | YES  |     | NULL    |                |
| age        | bigint\(20\)       | YES  |     | NULL    |                |
| interest   | varchar\(255\)     | YES  |     | NULL    |                |

### messages

关键字段

- id
- sender_id -> 对应消息发送者
- receiver_id -> 对应消息接受者
- content
- send_time

数据库表结构

| Field       | Type               | Null | Key | Default | Extra          |
| :---------- | :----------------- | :--- | :-- | :------ | :------------- |
| id          | int\(10\) unsigned | NO   | PRI | NULL    | auto_increment |
| created_at  | timestamp          | YES  |     | NULL    |                |
| updated_at  | timestamp          | YES  |     | NULL    |                |
| deleted_at  | timestamp          | YES  | MUL | NULL    |                |
| sender_id   | int\(10\) unsigned | YES  |     | NULL    |                |
| receiver_id | int\(10\) unsigned | YES  |     | NULL    |                |
| content     | varchar\(255\)     | YES  |     | NULL    |                |
| send_time   | timestamp          | YES  |     | NULL    |                |

## 项目结构

> 以下结构出于个人经验，有不当之处请给出宝贵意见

- route 路由层，负责将"\xxx\xxx"请求映射到对应的函数
- middleware 中间件层，可以在执行函数前后进行拦截并处理，如登陆拦截
- controller 控制层，存放与"\xxx\xxx"请求对应的函数，根据请求，调用业务层，并将数据进行格式封装返回
- service 业务层，调用持久层完成业务逻辑
- dao 持久层，可以理解为 sql 执行到函数执行的封装，由于使用了 ORM，本项目没有 dao 层目录
- database 提供数据库连接
- model 定义一系列结构体
  - pojo 业务逻辑实体，如 User，Message
  - reqo 请求数据实体，对应 controller 中的每个方法
  - reso 响应数据实体，对应 controller 中的每个方法
- config 读取配置，并提供单实例的配置文件实体供外访问
- tool 工具层
- assets 静态资源目录，存放静态资源（前端文件）

### pojo、reqo、reso 都是什么

- pojo

  很好理解，就是数据库对应的实体，但不要求与数据库字段一一对应

- reqo(request object)、reso(response object)

  不同接口请求的时候，可以携带的参数以及响应的数据也不同，所以为每一个接口设计一个对应的请求实体和响应实体

### controller、service、dao 到底有什么区别

> 以下为个人理解

- Controller

  主要职责是，接受请求的请求参数，转换为 reqo，进行简单的请求参数验证（我个人的定义与数据库无关的验证，如非空、非零），调用 Service 层的函数获取 pojo 结果，并将 pojo 结果转换封装为 reso 返回。

- Service

  主要职责是，对 Dao 层的接口进一步封装，提供通用的接口给 Controller 调用，返回数据可以是 pojo，在 Service 内需要进行数据的验证，如（新增用户，校验用户名是否重复）。

- Dao

  这里基本上一个方法直接对应一条 sql 语句，不做任何的验证，认为接收到的数据是可靠的（已经经过了 Controller 和 Service 两层的参数验证了），返回数据可以是 pojo。
