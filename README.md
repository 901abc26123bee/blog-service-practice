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

### setup
1. `docker-compose up build`
2. run `setup_mysql.sql` in `go_blog-service/storage/` to setup table and auth data
3. visit `http://127.0.0.1:8080/swagger/index.html`

## Note
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
[Go database/sql tutorial](http://go-database-sql.org/index.html)

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
open ~/.zshrc
```

### Tags API
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

### Article API
```
http://127.0.0.1:8080/api/v1/articles
// place token in header
// form-data
	cover_image_url:https://imgur.dcard.tw/k5IcoJB.gif
	created_by:tina
	tag_id:1
	title:test article
	desc:test create article
	content:learning golang, gin and gorm

http://127.0.0.1:8080/api/v1/articles/1
// place token in header
// form data
cover_image_url:https://imgur.dcard.tw/k5IcoJB.gif
tag_id:1
title:test article 5
desc:test create article 5
content:learning golang, gin and gorm 5
id:1
modified_by:alice

curl -X GET http://127.0.0.1:8080/api/v1/articles -F tag_id=1 -F state=1  -H 'token: xxx...'
curl -X GET http://127.0.0.1:8080/api/v1/articles/1 -H 'token: xxx...'
curl -X DELETE http://127.0.0.1:8080/api/v1/articles/{1} -H 'token: xxx...'

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


### upload file
- 將應用服務和文件服務給拆分開來，因為從安全角度來講，文件資源不應當與應用資源擺放在一起，否則會有風險
- 需要設置文件服務去提供靜態資源的訪問，才能實現讓外部請求項目 HTTP Server 時同時提供靜態資源的訪問
	`r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))`

```
$ curl -X POST http://127.0.0.1:8080/upload/file -F file=@{file_path} -F type=1
$ curl -X POST http://127.0.0.1:8080/upload/file -F file=@./w644.jpeg -F type=1
{"file_access_url":"http://127.0.0.1:8080/static/3b136bf60441cadd0369301b5eb1b2bf.jpeg"}
```
visit `http://127.0.0.1:8080/static/3b136bf60441cadd0369301b5eb1b2bf.jpeg`


### token
Base64
- Base64編碼可用於在HTTP環境下傳遞較長的標識資訊。
- 在其他應用程式中，也常常需要把二進位資料編碼為適合放在URL（包括隱藏表單域）中的形式。此時，採用Base64編碼不僅比較簡短，同時也具有不可讀性，即所編碼的資料不會被人用肉眼所直接看到。

Base64UrlEncode
- Base64UrlEncode 是Base64 算法的變種，為什麼要變呢，原因是在業界中我們經常可以看到JWT 令牌會被放入Header 或Query Param 中（也就是URL）。
- 而在URL 中，一些個別字符是有特殊意義的，例如：“+”、“/”、“=” 等等，因此在Base64UrlEncode 算法中，會對其進行替換，例如：“+” 替換為“-”、“/” 替換成“_”、“=” 會被進行忽略處理，以此來保證JWT 令牌的在URL 中的可用性和準確性。

```shell
go get -u github.com/dgrijalva/jwt-go@v3.2.0
```


```shell
# get valid token
curl -X POST \
  'http://127.0.0.1:8080/auth' \
  -d 'app_key=tdlemon' \
  -d 'app_secret=qwe123hahaha'

{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBfa2V5IjoiYjhkODg2NDM3NDcyYjkyMDQxZWMwMjY0NWE1MjA0ZDgiLCJhcHBfc2VjcmV0IjoiZjM1OTI2Y2Y0OTIzYjhjOWQyNjQxNjM2ZDZkODNiYTQiLCJleHAiOjE2NTk5ODA2MzEsImlzcyI6ImJsb2ctc2VydmljZSJ9.qRpOh6nSaG4cWAPQt5QFhr1I-PD7OxsZorGB4h1ddw0"}
```

```shell
# visit without token
curl -X GET http://127.0.0.1:8080/api/v1/tags

# using wrong token
curl -X GET http://127.0.0.1:8080/api/v1/tags -H 'token: eyJhbGciOiJIUzI1NiIsInRxxx'

# validation timeout
curl -X GET http://127.0.0.1:8080/api/v1/tags -H 'token: eyJhbGciOiJIUzI1NiIsInRxxx'

# success
curl -X GET http://127.0.0.1:8080/api/v1/tags -H 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBfa2V5IjoiYjhkODg2NDM3NDcyYjkyMDQxZWMwMjY0NWE1MjA0ZDgiLCJhcHBfc2VjcmV0IjoiZjM1OTI2Y2Y0OTIzYjhjOWQyNjQxNjM2ZDZkODNiYTQiLCJleHAiOjE2NTk5ODA2MzEsImlzcyI6ImJsb2ctc2VydmljZSJ9.qRpOh6nSaG4cWAPQt5QFhr1I-PD7OxsZorGB4h1ddw0'

{"list":[{"id":2,"created_by":"tinawong","modified_by":"","created_on":1658564501,"modified_on":1658564501,"deleted_on":0,"is_del":0,"name":"Java","state":1},{"id":4,"created_by":"tinawong","modified_by":"","created_on":1658584135,"modified_on":1658584135,"deleted_on":0,"is_del":0,"name":"JavaScript","state":1}],"pager":{"page":1,"page_size":10,"total_rows":2}}
```

```shell
curl --help
-H, --header LINE   Pass custom header LINE to server (H)
 	==> place token in header while using postman
-F, --form CONTENT  Specify HTTP multipart POST data (H)
	==> use form-date in request body while POST
```

### log middleware
```shell
go get -u gopkg.in/gomail.v2
```
### network visit ratelimit
```shell
go get -u github.com/juju/ratelimit@v1.0.1
```

### Go failing
[Go failing - expected 'package', found 'EOF'](https://stackoverflow.com/questions/31110191/go-failing-expected-package-found-eof)

```shell
go run main.go
```

### docker
Dockerfile
```shell
docker build -t 'blog-service' .
docker run 'blog-service'
```

docker-compose
```shell
docker-compose up --build
docker-compose down
```

visit `http://127.0.0.1:8000/swagger/index.html`


## Project Reference
[《Go 语言编程之旅：一起用 Go 做项目》](https://github.com/go-programming-tour-book)