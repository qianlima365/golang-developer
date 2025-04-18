### 前提
go版本为1.13及以上
### 官方文档
如果你想更深层次的了解GO MODULE的意义及开发者们的顾虑,可以直接访问官方文档(EN)
https://github.com/golang/go/wiki/Modules

### go module介绍
go module是go官方自带的go依赖管理库,在1.13版本正式推荐使用

go module可以将某个项目(文件夹)下的所有依赖整理成一个 go.mod 文件,里面写入了依赖的版本等

使用go module之后我们可不用将代码放置在src下了

使用 go module 管理依赖后会在项目根目录下生成两个文件 go.mod 和 go.sum。

go.mod 中会记录当前项目的所依赖，文件格式如下所示：


```
module github.com/gosoon/audit-webhook

go 1.12

require (
github.com/elastic/go-elasticsearch v0.0.0
github.com/gorilla/mux v1.7.2
github.com/gosoon/glog v0.0.0-20180521124921-a5fbfb162a81
)
```


go.sum记录每个依赖库的版本和哈希值，文件格式如下所示：

```
github.com/elastic/go-elasticsearch v0.0.0 h1:Pd5fqOuBxKxv83b0+xOAJDAkziWYwFinWnBO0y+TZaA=
github.com/elastic/go-elasticsearch v0.0.0/go.mod h1:TkBSJBuTyFdBnrNqoPc54FN0vKf5c04IdM4zuStJ7xg=
github.com/gorilla/mux v1.7.2 h1:zoNxOV7WjqXptQOVngLmcSQgXmgk4NMz1HibBchjl/I=
github.com/gorilla/mux v1.7.2/go.mod h1:1lud6UwP+6orDFRuTfBEV8e9/aOM/c4fVVCaMa2zaAs=
github.com/gosoon/glog v0.0.0-20180521124921-a5fbfb162a81 h1:JP0LU0ajeawW2xySrbhDqtSUfVWohZ505Q4LXo+hCmg=
github.com/gosoon/glog v0.0.0-20180521124921-a5fbfb162a81/go.mod h1:1e0N9vBl2wPF6qYa+JCRNIZnhxSkXkOJfD2iFw3eOfg=
```


### 开启go module
(1) go 版本 >= v1.11

(2) 设置GO111MODULE环境变量

要使用go module 首先要设置GO111MODULE=on，GO111MODULE 有三个值，off、on、auto，off 和 on 即关闭和开启，auto 则会根据当前目录下是否有 go.mod 文件来判断是否使用 modules 功能。无论使用哪种模式，module 功能默认不在 GOPATH 目录下查找依赖文件，所以使用 modules 功能时请设置好代理。

在使用 go module 时，将 GO111MODULE 全局环境变量设置为 off，在需要使用的时候再开启，避免在已有项目中意外引入 go module。

windows:
```
set GO111MODULE=on
```
mac:
```
export GO111MODULE=on
```

然后输入
```
go env
```
查看 GO111MODULE 选项为 on 代表修改成功

### GO PROXY
go module 的目的是依赖管理,所以使用 go module 时你可以舍弃 go get 命令(但是不是禁止使用, 如果要指定包的版本或更新包可使用go get,平时没有必要使用)
因go的网络问题, 所以推荐使用 goproxy.cn设置

```
// 阿里云镜像
GOPROXY=https://mirrors.aliyun.com/goproxy/
// 中国golang镜像
GOPROXY=https://goproxy.io
// 七牛云为中国的gopher提供了一个免费合法的代理goproxy.cn，其已经开源。只需一条简单命令就可以使用该代理：
go env -w GOPROXY=https://goproxy.cn,direct
```


初始化为你的项目第一次使用 GO MODULE(项目中还没有go.mod文件)
进入你的项目文件夹

```
cd xxx/xxx/test/
go mod init test(test为项目名)
```

我们会发现在项目根目录会出现一个 go.mod 文件
注意,此时的 go.mod 文件只标识了项目名和go的版本,这是正常的,因为只是初始化了

#### 检测依赖
```
go mod tidy
```

tidy会检测该文件夹目录下所有引入的依赖,写入 go.mod 文件
写入后你会发现 go.mod 文件有所变动
例如：
```
module test

go 1.13

require (
github.com/gin-contrib/sessions v0.0.1
github.com/gin-contrib/sse v0.1.0 // indirect
github.com/gin-gonic/gin v1.4.0
github.com/go-redis/redis v6.15.6+incompatible
github.com/go-sql-driver/mysql v1.4.1
github.com/golang/protobuf v1.3.2 // indirect
github.com/jinzhu/gorm v1.9.11
github.com/json-iterator/go v1.1.7 // indirect
github.com/kr/pretty v0.1.0 // indirect
github.com/mattn/go-isatty v0.0.10 // indirect
github.com/sirupsen/logrus v1.2.0
github.com/ugorji/go v1.1.7 // indirect
golang.org/x/sys v0.0.0-20191025021431-6c3a3bfe00ae // indirect
gopkg.in/yaml.v2 v2.2.4
)
```


此时依赖还是没有下载的
下载依赖我们需要将依赖下载至本地,而不是使用 go get
```
go mod download
```

如果你没有设置 GOPROXY 为国内镜像,这步百分百会卡到死
此时会将依赖全部下载至 GOPATH 下,会在根目录下生成 go.sum 文件, 该文件是依赖的详细依赖, 但是我们开头说了,我们的项目是没有放到 GOPATH 下的,那么我们下载至 GOPATH 下是无用的,照样找不到这些包

#### 导入依赖
```
go mod vendor
```

执行此命令,会将刚才下载至 GOPATH 下的依赖转移至该项目根目录下的 vendor(自动新建) 文件夹下, 此时我们就可以使用这些依赖了
goland设置开启 Go Module可能是因为 GO MODULE 功能还需完善,GOLAND默认是关闭该功能的,我们需要手动打开(不排除之后更新会不会改成默认开启)
打开设置Go->打开折叠->选择Go Modules(vgo)->选中两个复选框，在Proxy选择框中填入镜像源
依赖更新这里的更新不是指版本的更新,而是指引入新依赖
依赖更新请从检测依赖部分一直执行即可,即
```
go mod tidy
go mod download
go mod vendor
```


新增依赖有同学会问,不使用 go get ,我怎么在项目中加新包呢?
直接项目中 import 这个包,之后更新依赖即可
在协作中使用 GO MODULE要注意的是, 在项目管理中,如使用git,请将 vendor 文件夹放入白名单,不然项目中带上包体积会很大
git设置白名单方式为在git托管的项目根目录新建 .gitignore 文件设置忽略即可.但是 go.mod 和 go.sum 不要忽略
另一人clone项目后在本地进行依赖更新(同上方依赖更新)即可
### GOMODULE常用命令
```
go mod init # 初始化go.mod
go mod tidy # 更新依赖文件
go mod download # 下载依赖文件
go mod vendor # 将依赖转移至本地的vendor文件
go mod edit # 手动修改依赖文件
go mod graph # 打印依赖图
go mod verify # 校验依赖
```