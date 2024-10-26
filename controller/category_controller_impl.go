package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"rest-api-native/helper"
	"rest-api-native/model/web"
	"rest-api-native/service"
	"strconv"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (c *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryResponse := c.CategoryService.Create(request.Context(), categoryCreateRequest)
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (c *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	id, err := strconv.Atoi(params.ByName("categoryId"))
	helper.PanicIfError(err)
	categoryUpdateRequest.Id = id

	categoryResponse := c.CategoryService.Update(request.Context(), categoryUpdateRequest)
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (c *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("categoryId"))
	helper.PanicIfError(err)

	c.CategoryService.Delete(request.Context(), id)
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   "Category deleted",
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (c *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("categoryId"))
	helper.PanicIfError(err)

	categoryResponse := c.CategoryService.FindById(request.Context(), id)
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (c *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponses := c.CategoryService.FindAll(request.Context())
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponses,
	}

	helper.WriteResponseBody(writer, webResponse)
}
