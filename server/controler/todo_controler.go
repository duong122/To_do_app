package controler

import (
	"net/http"
	"server/data/request"
	"server/data/response"
	"server/helper"
	"server/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type TodoControler struct {
	TodoService service.TodoService
}

func NewTodoController(todoService service.TodoService) *TodoControler {
	return &TodoControler{TodoService: todoService}
}

func (controller *TodoControler) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	todoCreateRequest := request.TodoCreateRequest{}
	helper.ReadRequestBody(requests, &todoCreateRequest)

	controller.TodoService.Create(requests.Context(), todoCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *TodoControler) Update(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	todoUpadateRequest := request.TodoUpdateRequest{}
	helper.ReadRequestBody(requests, &todoUpadateRequest)

	todoId := params.ByName("todoId")
	id, err := strconv.Atoi(todoId)
	helper.PanicIfError(err)
	todoUpadateRequest.Id = id

	controller.TodoService.Update(requests.Context(), todoUpadateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *TodoControler) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	todoId := params.ByName("todoId")
	id, err := strconv.Atoi(todoId)
	helper.PanicIfError(err)

	controller.TodoService.Delete(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *TodoControler) FindAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	result := controller.TodoService.FindAll(requests.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   result,
	}
	helper.WriteResponseBody(writer, webResponse)
}

func (controller *TodoControler) FindById(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	todoId := params.ByName("todoId")
	id, err := strconv.Atoi(todoId)
	helper.PanicIfError(err)

	result := controller.TodoService.FindById(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}
