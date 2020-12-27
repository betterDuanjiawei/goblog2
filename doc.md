# 学习笔记

## 2.舞台布置

### 2.2 Go Docs
* 启动本地 doc服务
```
go doc -http=:6060  -http参数是表示以 http服务的形式指定地址喝端口为:6060,:前默认为 localhost
```
* 如果提示 go doc命令不存在,需要下载安装 go doc
```
go get golang.org/x/tools/cmd/godoc
```
* 如需加速,请参考[Wiki：Go 技巧分享：Go 国内加速镜像 ](https://learnku.com/go/wikis/38122)


## 3. 开始编码

## r http.Request 用户请求信息
* r.URL.Query() 获取用户参数
* r.URL.Path    获取当前请求路径
* r.Header.Get("User-Agent")    获取用户客户端信息

## w http.ResponseWriter 返回给用户的响应
* w.WriterHeader(http.StatusInternalServerError) 返回状态码:500
* w.Header().Set("name", "my name is djw") 设置返回表头,注意这里是函数的链式调用

* go 标准包
```
标准库包名	功能简介
bufio	带缓冲的 I/O 操作
bytes	实现字节操作
container	封装堆、列表和环形列表等容器
crypto	加密算法
database	数据库驱动和接口
debug	各种调试文件格式访问及调试功能
encoding	常见算法如 JSON、XML、Base64 等
flag	命令行解析
fmt	格式化操作
go	Go 语言的词法、语法树、类型等。可通过这个包进行代码信息提取和修改
html	HTML 转义及模板系统
image	常见图形格式的访问及生成
io	实现 I/O 原始访问接口及访问封装
math	数学库
net	网络库，支持 Socket、HTTP、邮件、RPC、SMTP 等
os	操作系统平台不依赖平台操作封装
path	兼容各操作系统的路径操作实用函数
plugin	Go 1.7 加入的插件系统。支持将代码编译为插件，按需加载
reflect	语言反射支持。可以动态获得代码中的类型信息，获取和修改变量的值
regexp	正则表达式封装
runtime	运行时接口
sort	排序接口
strings	字符串转换、解析及实用函数
time	时间接口
text	文本模板及 Token 词法器
```

## http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request))
* http.HandleFunc 里传参的 / 意味着 任意路径。并不是根目录的意思
* 第二个参数是函数类型,函数签名必须满足 w r 的类型

## go 的字符连接符
* go + ; php .; js + ; shell 没有,直接连接即可



