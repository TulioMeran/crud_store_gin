package controller

import (
	"crud_store_gin/dto/request"
	"crud_store_gin/dto/response"
	"crud_store_gin/helper"
	"crud_store_gin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TagsController struct {
	tagsService service.TagsService
}

func NewTagsController(service service.TagsService) *TagsController {
	return &TagsController{
		tagsService: service,
	}
}

// FindAll			godoc
// @Summary			FindAll tags
// @Description		Return all the tags from the data.
// @Produce			application/json
// @Tags			tags
// @Success			200 {object} response.Response{}
// @Router			/tags [get]
func (controller *TagsController) FindAll(c *gin.Context) {

	tagsResponse := controller.tagsService.FindAll()

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   tagsResponse,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}

// FindById			godoc
// @Summary			FindById tags
// @Description		Return a single tag
// @Param			tagId path int true "return tags by id"
// @Produce			application/json
// @Tags			tags
// @Success			200 {object} response.Response{}
// @Router			/tags/{tagId} [get]
func (controller *TagsController) FindById(c *gin.Context) {

	tagId := c.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	tagResponse := controller.tagsService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   tagResponse,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}

// CreateTags		godoc
// @Summary			Create tags
// @Description		Save tags data in DB.
// @Param			tags body request.CreateTagsRequest true "Create tags"
// @Produce			application/json
// @Tags			tags
// @Success			200 {object} response.Response{}
// @Router			/tags [post]
func (controller *TagsController) Create(c *gin.Context) {
	createTagsRequest := request.CreateTagsRequest{}
	err := c.ShouldBindJSON(&createTagsRequest)
	helper.ErrorPanic(err)

	controller.tagsService.Create(createTagsRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}

// UpdateTags		godoc
// @Summary			Update tags
// @Description		Update tags data in DB.
// @Param			tags body request.UpdateTagsRequest true "Update tags"
// @Param			tagId path int true "update tags by id"
// @Produce			application/json
// @Tags			tags
// @Success			200 {object} response.Response{}
// @Router			/tags/{tagId} [patch]
func (controller *TagsController) Update(c *gin.Context) {
	updateTagsRequest := request.UpdateTagsRequest{}
	err := c.ShouldBindJSON(&updateTagsRequest)
	helper.ErrorPanic(err)

	tagId := c.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	updateTagsRequest.Id = id
	controller.tagsService.Update(updateTagsRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)

}

// DeleteTags		godoc
// @Summary			Delete tags
// @Description		Delete tags data in DB.
// @Param			tagId path int true "update tags by id"
// @Produce			application/json
// @Tags			tags
// @Success			200 {object} response.Response{}
// @Router			/tags/{tagId} [delete]
func (controller *TagsController) Delete(c *gin.Context) {
	tagId := c.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	controller.tagsService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)

}
