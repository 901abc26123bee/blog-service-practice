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

// @Summary Get an article
// @Produce json
// @Param id path int true "Article ID"
// @Success 200 {object} model.Article "Success"
// @Failure 400 {object} errcode.Error "Bad Request"
// @Failure 500 {object} errcode.Error "Internal Server Error"
// @Router /api/v1/articles/{id} [get]
func (a Article) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}

// @Summary Get Article
// @Produce json
// @Param name query string false "article name"
// @Param tag_id query int false "Tag ID"
// @Param state query int false "state"
// @Param page query int false "page"
// @Param page_size query int false "page size"
// @Success 200 {object} model.ArticleSwagger "Success"
// @Failure 400 {object} errcode.Error "Bad Request"
// @Failure 500 {object} errcode.Error "Internal Server Error"
// @Router /api/v1/articles [get]
func (a Article) List(c *gin.Context) {}

// @Summary Create Article
// @Produce json
// @Param tag_id body string true "Article ID"
// @Param title body string true "article title"
// @Param desc body string false "article description"
// @Param cover_image_url body string true "cover image url"
// @Param content body string true "content"
// @Param created_by body int true "created by"
// @Param state body int false "state"
// @Success 200 {object} model.Article "Success"
// @Failure 400 {object} errcode.Error "Bad Request"
// @Failure 500 {object} errcode.Error "Internal Server Error"
// @Router /api/v1/articles [post]
func (a Article) Create(c *gin.Context) {}

// @Summary Update Article
// @Produce json
// @Param tag_id body string false "Tag ID"
// @Param title body string false "article title"
// @Param desc body string false "article description"
// @Param cover_image_url body string false "cover image url"
// @Param content body string false "article content"
// @Param modified_by body string true "modified by"
// @Success 200 {object} model.Article "Success"
// @Failure 400 {object} errcode.Error "Bad Request"
// @Failure 500 {object} errcode.Error "Internal Server Error"
// @Router /api/v1/articles/{id} [put]
func (a Article) Update(c *gin.Context) {}

// @Summary Delete Article
// @Produce  json
// @Param id path int true "Article ID"
// @Success 200 {string} string "Success"
// @Failure 400 {object} errcode.Error "Bad Request"
// @Failure 500 {object} errcode.Error "Internal Server Error"
// @Router /api/v1/articles/{id} [delete]
func (a Article) Delete(c *gin.Context) {}