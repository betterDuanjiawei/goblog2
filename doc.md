# 学习笔记
* goland 导入项目后import里的包报红
[参考](https://blog.csdn.net/weixin_37719934/article/details/108399699)
需要检查修改 goland设置里的 GOPATH和 GOMODULE
* 由于要使用 air 自动重载工具,把自动保存关掉,但是不能把 unsafe的第一个选项也取消勾选了,要不然会无法自动加载新增和修改的文件和文件夹的
[参考](http://www.lanrendemo.com/view/MTUxNQ.html)
关闭自动保存并标记未保存的文件*号
1、Preferences -> Appearance & Behavior -> System Settings，取消use safe等4个。(第一个不要取消勾选了)
2、Preferences -> Editor -> General -> Editor Tabs 勾选 Mark modified(*)
[设置 GoLand 保存时自动格式化](http://www.lanrendemo.com/view/MTUxNA.html)
* [设置 GOPATH](https://studygolang.com/articles/17598)
```
错误写法:
go env -w GOPATH="/Users/v_duanjiawei/go/src"
warning: go env -w GOPATH=... does not override conflicting OS environment variable (不能覆盖 os系统的环境变量)

Bash
export $GOPATH=$HOME/go
source ~/.bash_profile

Zsh
export $GOPATH=$HOME/go
source ~/.zshrc

set -x -U GOPATH $HOME/go
-x 用来指定你要导出的变量  -U 设置成全局的环境

go env GOPATH 查看路径
```
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
* go 默认 http.ServMux; 
* httpRouter 速度最快的路由器,被gin采用, 但是它不支持路由命名 高性能,路由功能相对来说简单的项目中,比如 api和微服务
* gorilla/mux; 全栈的 web 开发中,功能强大
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

## gorilla/mux 精准匹配 ; net/http 长度优先匹配
* 精准匹配 指路由只会匹配准确指定的规则，这个比较好理解，也是较常见的匹配方式。 动态内容
* 长度优先匹配 一般用在静态路由上（不支持动态元素如正则和 URL 路径参数），优先匹配字符数较多的规则。 静态内容
```
router.HandleFunc("/", defaultHandler)
router.HandleFunc("/about", aboutHandler)
使用 长度优先匹配 规则的 http.ServeMux 会把除了 /about 这个匹配的以外的所有 URI 都使用 defaultHandler 来处理。
而使用 精准匹配 的 gorilla/mux 会把以上两个规则精准匹配到两个链接，/ 为首页，/about 为关于，除此之外都是 404 未找到。
```
* 使用
```
    router.NotFoundHandler = http.HandlerFunc(notFountHandler)
	http.HandleFunc("/no-exits", notFountHandler)
```

## http.HandleFunc()  http.HandlerFunc()区别

## gorilla/mux
* 指定Methods()方法来区分请求
```
router.HandleFunc("/articles", articlesIndexHandler).Methods("GET").Name("articles.index")
router.HandleFunc("/articles", articlesStoreHandler).Methods("POST").Name("articles.store")
```
* 请求路径参数和正则匹配
```
router.HandleFunc("/articles/{id:[0-9]+}", articlesShowHandler).Methods("GET").Name("articles.show")

{id:[0-9]+} 限定了一个或者多个数字
使用{name} 花括号来设置路径参数
在有正则表达式的情况下,使用:区分,第一部分是名称,第二部分是正则表达式
```
* 获取请求路径参数
```
vars := mux.Vars(r)
mux 提供的方法 mux.Vars(r)会将URL 路径参数解析为键值对应的Map;
使用 vars["id"]的形式来读取
```
* 路由命名
```
router.HandleFunc("/", homeHandler).Methods("GET").Name("home")
Name()方法来给路由命名,传参是路由的名称;通过这个名称来获取 URL
homeURL, _ := router.Get("home").URL()
articlesURL, _ := router.Get("articles.show").URL("id", 2)
```
* 加载中间件 使用 gorilla/mux 的 mux.Use()方法加载中间件
```
router.Use(forceHTMLMiddleware)

func forceHTMLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 设置标头
		w.Header().Set("Content-Type", "text/html;charset=utf-8")
		// 继续处理请求
		next.ServeHTTP(w, r)
	})
}
```
* gorilla StrictSlash(value bool)
```
router.NewRouter().StrictSlash(true) 
URL 校正301跳转,第二次会是没有/的 url,处理 get请求可以,但是处理 post请求,301跳转之后会变成 get请求
```
* gorilla执行顺序,先匹配路由,再执行中间件


## GOMODULE
* go.mod go.sum 分别相当于 php的 composer.json 和 composer.lock
* go.mod
```
module 当前项目也算一个module
go 1.15  指定了最低的go 版本要求
require () 项目所需依赖 // indirect 未使用的
```
* go.sum  保存着依赖包的版本和哈希值, 不仅会保存直接依赖包的哈希值,间接依赖的包的哈希值也会被保存
```
两种 hash值,前者为GO MODULES打包整个模块文件zip后再进行 hash值,而后者针对go.mod的hash
go.sum是保证下载源码100%正确的重要依据.
github.com/gorilla/mux v1.7.4 h1:VuZ8uybHlWmqV03+zRzdwKL4tUnIp1MAQtp1mIFE1bc=
github.com/gorilla/mux v1.7.4/go.mod h1:DVbg23sWSpFRCP0SfiEN6jmj59UnW/n46BH5rLB71So=
```
* go mod tidy 整理go module 依赖,会把未使用的modules移除掉
* 源码包的存放位置: $GOPATH/pkg/mod
* go clean -modcache 清空本地下载的 go module 缓存
* 下载依赖,当执行go run 或 go build时候,go 会基于自动go.mod文件自动拉取依赖. go mod download 下载项目所需依赖
* go module 命令
```
go mod init     生成 go.mod 文件
go mod download 下载 go.mod中指明的所有依赖
go mod tidy     整理现有依赖,清除未使用的
go mod graph    查看现有的依赖结构
go mod edit     编辑 go.mod 文件
go mod vendor   导出项目的所以依赖到 vendor目录(生成的目录在当前项目目录下)
go mod verify   校验一个模块是否被篡改过
go mod why      查看为什么要依赖某个模块 // go mod why github.com/gorilla/mux  
```
* GO111MODULE 因为在 go 1.11版本添加,故命名为GO111MODULE
```
设置选项:GO111MODULE="on"
auto: 1.11-1.15的默认值, 表示项目中有 go.mod 文件的话启用 go module
on: 开启,未来会是默认值
off: 关闭,不推荐
```

## GOPROXY 
* GOPROXY 此变量用于设置 go模块代理,其作用是拉取源码时候能够脱离传统的VCS方式,直接通过镜像站点来快速拉取.
* GOPROXY默认值:https://proxy.golang.org,direct;设置国内的代理模块:go env -w GOPROXY=https://goproxy.cn,direct
* 加 direct表示:告诉 go get 抓取源码包时先尝试https://goproxy.cn,如果遇到404等错误,再尝试从源地址抓取
* 设置值为 off,禁止在后续的操作使用任何go模块代理

## GOSUMDB
* go checksum database的缩写,用于在拉取模块版本时候,保证拉取到的模块代码包未见过篡改,若发现不一致将会立即停止
* GOSUMDB="sum.golang.org" GOSUMDB可以被Go Module Proxy代理, goproxy.cn 同样支持代理 sum.golang.org
* 可以设置off,会禁止在后续操作中校验模块哈希

## GONOPROXY GONOSUMDB GOPRIVATE
* 三个环境变量用在了私有模块
* GONOPROXY 设置不走 go proxy 的 url规则
* GONOSUMDB 设置不检查哈希的 URL规则
* GOPRIVATE 设置私有模块的URL规则,会同时设置以上两个变量
一般私有仓库,直接使用GOPRIVATE即可
* go env -w GOPRIVATE="*.example.com" 使用通配符,example.com 的子域名都不走 go proxy 和go checksum database,但是这里不包括 example.com 本身

## strings.TrimSuffix(s, suffix string) string 函数,可以用来移除后缀

## http.Request 的方法和字段
* r *http.Request
* r.ParseForm() 从请求中解析出请求参数,必须是执行完这段解析代码,后面的 r.Form 和 r.PostForm 才可以读取到数据
* r.Form 存储了 post get put 参数,在使用之前需要调用 ParseForm 方法 比 r.PostForm 多了 URL参数里的数据  ?test=data
* r.PostForm 存储了 post put 参数,在使用之前需要调用 ParseForm 方法
* 如果不想获取所有的请求内容,而是逐个获取的话,无需使用 r.ParseForm()可以直接使用 r.FormValue() 和 r.PostFormValue() 方法获取数据


## 统计字符串长度
* len("中国") 6 len()可以用来统计字符串 切片 通道的长度,utf8一个汉字占用3个字节
* utf8.RuneCountInString("中国") 2

## html/template
* html 代码字符串的写法
```
html := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<title>创建文章 —— 我的技术博客</title>
		<style type="text/css">.error {color: red;}</style>
	</head>
	<body>
		<form action="{{ .URL }}" method="post">
			<p><input type="text" name="title" value="{{ .Title }}"></p>
			{{ with .Errors.title }}
				<p class="error">{{ . }}</p>
			{{ end }}
			<p><textarea name="body" cols="30" rows="10">{{ .Body }}}</textarea></p>
			{{ with .Errors.body }}
				<p class="error">{{ . }}</p>
			{{ end }}
			<p><button type="submit">提交</button></p>
		</form>
	</body>
	</html>
`
		storeURL, _ := router.Get("articles.store").URL()

		data := ArticlesFormData{
			Title:  title,
			Body:   body,
			URL:    storeURL,
			Errors: errors,
		}

		tmpl, err := template.New("create-form").Parse(html)
		if err != nil {
			panic(err)
		}
		tmpl.Execute(w, data)
```
* html文件的写法 template.ParseFiles("resources/views/articles/create.gohtml")
```
storeURL, _ := router.Get("articles.store").URL()

		data := ArticlesFormData{
			Title:  title,
			Body:   body,
			URL:    storeURL,
			Errors: errors,
		}

		tmpl, err := template.ParseFiles("resources/views/articles/create.gohtml")
		if err != nil {
			panic(err)
		}
		tmpl.Execute(w, data)
```
* 模板文件的后缀.gohtml 常见的其他后缀:.tmpl; .tpl; .gohtml;
* 模板语法
```
{{ }} 双层大括号是默认的模板界定符号.用于在 html模板文件中界定模板语法
{{ . }}语法 .表示当前的对象,当我们传入一个结构体对象时候,我们可以使用.来访问结构体对应的字段. 当我们传入的变量是 map时候,也可以在模板文件中通过.根据 key来取值
with 关键字
{{with pipeline}}}T1{{end}} 如果 pipeline 为空则不产生输出,否则将.设为 pipeline 的值并执行 T1,不修改外面的.
{{ with pipeline }}T1{{ else }}T0 {{ end }} 如果 pipeline 为空则不改变.,并执行 T0,否则.设置为 pipeline 的值并执行 T1
with 区块外, {{ . }} 代表传入模板的数据,而在 with区块内,则代表pipeline里的数据
如:{{ with .Error.title }}这个区块中, {{ . }}代表 .Error.title

pipeline 产生数据的操作, go的模板语法支持使用管道符号|连接多个命令

注释: {{/*这是一个注释*/}} 执行的时候会忽略,可以多行,注释不能嵌套,而且必须紧贴临界符
变量: 可以在模板中声明变量$variable := {{ . }} $variable是变量的名称,在后续的代码中可以使用该变量了
移除空格: {{- .Name -}} {{- 移除模板内容左侧的所有空白符号, -}} 移除模板内容右侧的所有空白符号, -要紧挨{{ 和 }},和模板变量之间还有空格
条件判断:
{{ if pipeline }}T1{{ end }}
{{ if pipeline }}T1 {{ else }} T0 {{ end }}
{{ if pipeline }} T1 {{ else if pipeline }} T0 {{ end }}
range遍历:
{{ range pipeline }} T1 {{ end }} 如果 pipeline的长度为0,不会有任何输出
{{ range pipeline }} T1 {{ else }} T0 {{ end }} 如果 pipeline的长度为0, 则会执行 T0
修改默认标识符:
防止和 vue angularJs 冲突,修改 go模板引擎默认的标识符号:
template.New("test").Delims("{[", "]}").ParseFile("filename.gohtml")
```

## go 操作数据库方式
* database/sql 用硬编码 sql语句来执行
* ORM GORM 对象关系映射来的方式抽象的操作数据库

## database/sql
* database/sql 只提供了一套操作数据库的接口和规范, 可以用多种数据库驱动

## mysql驱动
* go get github.com/go-sql-driver/mysql
* _ "github.com/go-sql-driver/mysql" 匿名导入的方式来记载驱动


## init()
[详解 Go 语言中的 init () 函数](https://learnku.com/go/t/47178)
* init函数通常用于:
    1. 变量初始化
    2. 注册器 sql.Register()
    3. 检查/修复状态
    4. 运行计算
* 包的初始化过程:
    1. 初始化导入的包(递归导入)
    2. 计算并为块中的变量分配初始值
    3. 在包中执行初始化函数
```
package main
import "fmt"
var _ int64=s()
func init(){
  fmt.Println("init function --->")
}
func s() int64{
  fmt.Println("function s() --->")
  return 1
}
func main(){
  fmt.Println("main --->")
}

function s() —>
init function —>
main —>
```
* 即使程序包被多次导入,也只需要初始化一次
* init 不需要传入参数,也没有函数返回. 无法引用
* init() 用于程序执行前包的初始化的函数
如果只需要一个包的 init 函数,不需要这个包的其他方法,就可以用匿名导入的方式
```
mysql/driver.go
func init() {
    sql.Register("mysql", &MySQLDriver{})
}

特点:
1 init函数是用于程序执行前做包的初始化的函数，比如初始化包里的变量等
2 每个包可以拥有多个init函数,执行顺序是从上往下执行。  To ensure reproducible initialization behavior, build systems are encouraged to present multiple files belonging to the same package in lexical file name order to a compiler. 
3 包的每个源文件也可以拥有多个init函数
4 同一个包中多个init函数的执行顺序go语言没有明确的定义  （应该是顺序执行）
5 不同包的init函数按照包导入的依赖关系决定该初始化函数的执行顺序
6 init函数不能被其他函数调用，而是在main函数执行之前，自动被调用
```

