package service

import (
	"context"
	"server/data/request"
	"server/data/response"
)

type TodoService interface {
	Create(ctx context.Context, request request.TodoCreateRequest)
	Update(ctx context.Context, request request.TodoUpdateRequest)
	Delete(ctx context.Context, todoId int)
	FindById(ctx context.Context, bookId int) response.TodoResponse
	FindAll(ctx context.Context) []response.TodoResponse
}
