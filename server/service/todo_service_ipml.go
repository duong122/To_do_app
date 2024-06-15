package service

import (
	"context"
	"server/data/request"
	"server/data/response"
	"server/helper"
	"server/model"
	"server/repository"
)

type TodoServiceImpl struct {
	TodoRepository repository.TodoRepository
}

func NewTodoServiceImpl(todoRepository repository.TodoRepository) TodoService {
	return &TodoServiceImpl{TodoRepository: todoRepository}
}

func (b *TodoServiceImpl) Create(ctx context.Context, request request.TodoCreateRequest) {
	todo := model.Todo{
		Name: request.Name,
	}
	b.TodoRepository.Save(ctx, todo)
}

func (b *TodoServiceImpl) Delete(ctx context.Context, todoId int) {
	todo, err := b.TodoRepository.FindById(ctx, todoId)
	helper.PanicIfError(err)
	b.TodoRepository.Delete(ctx, todo.Id)
}

func (b *TodoServiceImpl) FindAll(ctx context.Context) []response.TodoResponse {
	todos := b.TodoRepository.FindAll(ctx)

	var todoResp []response.TodoResponse

	for _, value := range todos {
		todo := response.TodoResponse{Id: value.Id, Name: value.Name}
		todoResp = append(todoResp, todo)
	}
	return todoResp
}

func (b *TodoServiceImpl) FindById(ctx context.Context, todoId int) response.TodoResponse {
	todo, err := b.TodoRepository.FindById(ctx, todoId)
	helper.PanicIfError(err)
	return response.TodoResponse(todo)
}

func (b *TodoServiceImpl) Update(ctx context.Context, request request.TodoUpdateRequest) {
	todo, err := b.TodoRepository.FindById(ctx, request.Id)
	helper.PanicIfError(err)

	todo.Name = request.Name
	b.TodoRepository.Update(ctx, todo)

}
