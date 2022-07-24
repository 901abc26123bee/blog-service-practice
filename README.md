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

### translations
https://github.com/go-playground/validator/tree/master/translations

### zsh
```
vim ~/.zshrc
```

###
```
curl -X POST http://127.0.0.1:8080/api/v1/tags -F 'name=GoLang' -F created_by=tinawong
curl -X POST http://127.0.0.1:8080/api/v1/tags -F 'name=Java' -F created_by=tinawong
curl -X POST http://127.0.0.1:8080/api/v1/tags -F 'name=JavaScript' -F created_by=tinawong

curl -X GET 'http://127.0.0.1:8080/api/v1/tags?page=1&page_size=2'
curl -X GET 'http://127.0.0.1:8080/api/v1/tags?page=2&page_size=2'

curl -X PUT http://127.0.0.1:8080/api/v1/tags/{1} -F state=0 -F modified_by=atom -F 'name=Python'
curl -X PUT http://127.0.0.1:8080/api/v1/tags/{1} -F state=0 -F modified_by=roger -F 'name=Python'
curl -X DELETE  http://127.0.0.1:8080/api/v1/tags/{1}
```

### GORM
- When using struct type in GORM to update, GORM will not make changes to fields with a value of zero.
- guess that the more fundamental reason is that when identifying the value of this field in this structure, it is difficult to determine whether it is really a zero value, or whether the external incoming just happens to be a zero value of this type. GORM does not do too much special identification in this area.

db.Update()
- 使用 struct 的方式只會更新 non-zero value，也就是說如果值是 false 或 0 這類的都不會被保存
	使用 map 的方式才會更新 zero-value，因此個人會建議使用 map 的方式來避免可能的 bug
- 會建議使用 map 的方式（而非 struct）來進行更新避免可能的 bug

db.Save()
- 使用 db.Save() 時，沒給的值會全部被清空（包括 createdAt、除了 updatedAt），所以如果要用 save，一定要把舊有的資料先拉出來，對要改變的內容進行修改後，在呼叫 save：
- 使用 db.Save() 時所有的欄位都會更新，因此若沒給的話會導致該值變回 zero-value，記得一定要先把資料拉出來後，修改要改的資料後，在呼叫 Save。

[Delete](https://gorm.io/docs/delete.html)
Soft Delete

```go
// 1. gorm.DeletedAt
type User struct {
  ID      int
  Deleted gorm.DeletedAt
  Name    string
}
// 2. Delete Flag
//    - Unix Second
// 		- Use 1 / 0 AS Delete Flag
```