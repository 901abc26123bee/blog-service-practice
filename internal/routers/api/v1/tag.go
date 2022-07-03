package v1

import "github.com/gin-gonic/gin"

type Tag struct {}

func NewTag() Tag {
	return Tag{}
}

// @Summary Get MultiTags
// @Produce  json
// @Param name query string false "tag name" maxlength(100)
// @Param state query int false "state" Enums(0, 1) default(1)
// @Param page query int false "page"
// @Param page_size query int false "page size"
// @Success 200 {object} model.TagSwagger "Success"
// @Failure 400 {object} errcode.Error "Bad Request"
// @Failure 500 {object} errcode.Error "Internal Server Error"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {}

// @Summary Add Tag
// @Produce  json
// @Param name body string true "tag name" minlength(3) maxlength(100)
// @Param state body int false "state" Enums(0, 1) default(1)
// @Param created_by body string true "created by" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "Success"
// @Failure 400 {object} errcode.Error "Bad Request"
// @Failure 500 {object} errcode.Error "Internal Server Error"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {}

// @Summary Update Tag
// @Produce  json
// @Param id path int true "Tag ID"
// @Param name body string false "tag name" minlength(3) maxlength(100)
// @Param state body int false "state" Enums(0, 1) default(1)
// @Param modified_by body string true "modified by" minlength(3) maxlength(100)
// @Success 200 {array} model.Tag "Success"
// @Failure 400 {object} errcode.Error "Bad Request"
// @Failure 500 {object} errcode.Error "Internal Server Error"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {}

// @Summary Delete tag
// @Produce  json
// @Param id path int true "Tag ID"
// @Success 200 {string} string "Success"
// @Failure 400 {object} errcode.Error "Bad Request"
// @Failure 500 {object} errcode.Error "Internal Server Error"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {}