## structure
```
blog-service
├── configs
├── docs
├── global
├── internal
│   ├── dao
│   ├── middleware
│   ├── model
│   ├── routers
│   └── service
├── pkg
├── storage
├── scripts
└── third_party
```

### Problem Solving:
#### Swag Command
```shell
$ swag -v
zsh: command not found: swag
```
##### Check if swag is downloaded
[swaggo/swag](https://github.com/swaggo/swag)
Download swag by using:
```shell
$ go get -u github.com/swaggo/swag/cmd/swag
# 1.16 or newer
$ go install github.com/swaggo/swag/cmd/swag@latest
```
you should see `swag` under `GOPATH/bin` after downloading.
```shell
$ go env
...
set GOPATH="/Users/xxx/go"
set GOROOT="/usr/local/Cellar/go/1.18.1/libexec"
...
$ cd /Users/xxx/go/bin
$ ls
swag
```

##### Solution 1:
```shell
$ open ~/.zshrc
...
# add
export PATH=$GOPATH/bin
```
##### Solution 2:
Use absolute path in command
```shell
# $GOPATH/bin/swag -v
$ /Users/xxx/go/bin/swag -v
swag version v1.8.3
$ /Users/xxx/go/bin/swag init
2022/07/03 23:48:18 Generate swagger docs....
...
2022/07/03 23:48:19 create docs.go at  docs/docs.go
2022/07/03 23:48:19 create swagger.json at  docs/swagger.json
2022/07/03 23:48:19 create swagger.yaml at  docs/swagger.yaml
```

#### Swagger page
[undefined: "github.com/swaggo/swag".Spec #1126](https://github.com/swaggo/swag/issues/1126)
http://127.0.0.1:8080/swagger/index.html

```shell
go install github.com/swaggo/swag/cmd/swag@v1.7.8
```

- Remember to import docs, so that the swagger index.html page can read the document.
- Run `swag init` every time after you modied your swagger annotations to produce the new docs.
	```go
	import _ "blog-service/docs"
	```


#### mysql connection
Remember to import mysql-driver when you are using `gorm`.
```go
import (
	...
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
```



