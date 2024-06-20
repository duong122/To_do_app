package controler

import (
	"fmt"
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
	todoUpdateRequest := request.TodoUpdateRequest{}
	helper.ReadRequestBody(requests, &todoUpdateRequest)

	todoId := params.ByName("todoId")
	fmt.Print("todoId: ", todoId)
	id, err := strconv.Atoi(todoId)
	helper.PanicIfError(err)
	todoUpdateRequest.Id = id

	controller.TodoService.Update(requests.Context(), todoUpdateRequest)
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
