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
* r.Method 获取当期请求的方式:GET POST

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

## air 自动重载
* 安装
```
go env -w GOPROXY=https://goproxy.cn
GO111MODULE=on go get -u github.com/cosmtrek/air
```
* tmp 文件目录用来存放编译后的文件,还有 build 错误日志:build-errors.log; 需要用.gitignore 忽略上传代码库

## go get
[参考](http://c.biancheng.net/view/123.html)
* go get 一键获取代码 编译和安装
-d 只下载不安装
-u 强制使用网络去更新包和它的依赖包
-v 显示执行的命令

## Content-Type 标头
* Content-Type 响应标头用于告诉客户端内容的类型,客户端再根据这个信息将内容正确的呈现给用户
```
text/html   HTML 文档
text/plain  文本内容
text/css    CSS 样式文件
text/javascript JS 脚本文件
application/json    JSON格式的数据
application/xml     XML 格式的数据
image/png       PNG图片
```
* Content-Type: text/plain; charset=utf-8
* w.Header().Set() 设置 Content-Type的时候, 后面的参数是多个用;组合的,
```
w.Header().Set("Content-type", "text/html") // 错误,虽然显示的是html内容,但是编码有问题.
w.Header().Set("Content-type", "text/html; charset=utf-8")
```

## web 响应状态码
* web 响应和请求结构类似,由响应行 响应头部 响应体组成
    1. 响应行: 协议 响应状态码和状态描述: HTTP/1.1 200 OK
    2. 响应头部: 包含各种头部字段信息,如 cookie, Content-Type等头部信息
    3. 响应体: 携带客户端想要的响应数据,格式和编码由 Content-Type决定
* 响应状态码的固定值和意义
    1. 100~199: 表示服务端成功接收客户端的请求,要求客户端继续提交下一次请求才能处理完整个处理过程
    2. 200~299: 表示服务端成功接收并已经完成整个处理过程,200 成功
    3. 300~399: 为完成请求,客户端需要进一步细化请求.302:客户端请求的资源已经移动到一个新地址,使用302表示将资源重定向;304客户端请求的资源未发生改变的时候,使用304,告诉客户端从本地缓存中获取
    4. 400~499: 客户端的请求有错误:404 请求的资源在 web 服务器中找不到;403 服务器拒绝客户的访问,一般是权限不够;499 服务端处理时间过长,客户端不耐烦了,关闭了
    5.500~599: 服务器内部错误 500

## http.HandleFunc ServeMux DefaultServeMux
* 
```
http.ListenAndServe(":3000", nil) // nil其实默认是 defaultServeMux

func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)
}

// The handler is typically nil, in which case the DefaultServeMux is used.
//
// ListenAndServe always returns a non-nil error.
func ListenAndServe(addr string, handler Handler) error {
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}

router := http.NewServeMux()
func NewServeMux() *ServeMux { return new(ServeMux) }
```
* http.ServMux 缺点:
    1. 不支持 url路径参数 /articles/2
    2. 不支持请求方法过滤  GET POST
    3. 不支持路由命名 路由命名是一种允许我们快速修改 URL的方式

## strings.Split() strings.SplitN()
* 有 N没 N,区别在于是否返回指定个数的切割字字符串. 没有 n默认传-1,代表全部返回
```
        //以str为分隔符，将s切分成多个子串，结果中**不包含**str本身。如果str为空则将s切分成Unicode字符列表。
        //如果s中没有str子串，则将整个s作为[]string的第一个元素返回。
        //参数n表示最多切分出几个子串，超出的部分将不再切分，最后一个n包含了所有剩下的不切分。
        //如果n为0，则返回nil；如果n小于0，则不限制切分个数，全部切分

func Split(s, sep string) []string { return genSplit(s, sep, 0, -1) }
func SplitN(s, sep string, n int) []string { return genSplit(s, sep, 0, n) }
```


