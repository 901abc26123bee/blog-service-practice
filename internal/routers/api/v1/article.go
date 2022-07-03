package v1

import (
	"blog-service/pkg/app"
	"blog-service/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}
/*
$ curl -v http://127.0.0.1:8080/api/v1/articles/1
> GET /api/v1/articles/1 HTTP/1.1
...
< HTTP/1.1 500 Internal Server Error
...
{"code":10000000,"msg":"Server Internal Error"}%
*/
func (a Article) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}
func (a Article) List(c *gin.Context) {}
func (a Article) Create(c *gin.Context) {}
func (a Article) Update(c *gin.Context) {}
func (a Article) Delete(c *gin.Context) {}