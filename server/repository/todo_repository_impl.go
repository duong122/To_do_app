package repository

import (
	"context"
	"database/sql"
	"errors"
	"server/helper"
	"server/model"
)

type TodoRepositoryIpml struct {
	Db *sql.DB
}

func NewTodoRepository(Db *sql.DB) TodoRepository {
	return &TodoRepositoryIpml{Db: Db}
}

func (b *TodoRepositoryIpml) Update(ctx context.Context, todo model.Todo) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Sql := "update todo set name=$1 where id=$2"
	_, errExec := tx.ExecContext(ctx, Sql, todo.Name, todo.Id)
	helper.PanicIfError(errExec)
}

func (b *TodoRepositoryIpml) Delete(ctx context.Context, todoId int) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Sql := "delete from todo where id = $1"
	_, errExec := tx.ExecContext(ctx, Sql, todoId)
	helper.PanicIfError(errExec)
}

func (b *TodoRepositoryIpml) FindAll(ctx context.Context) []model.Todo {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Sql := "select id, name from todo"
	result, errQuery := tx.QueryContext(ctx, Sql)
	helper.PanicIfError(errQuery)
	defer result.Close()

	var todos []model.Todo
	for result.Next() {
		todo := model.Todo{}
		err = result.Scan(&todo.Id, &todo.Name)
		helper.PanicIfError(err)

		todos = append(todos, todo)
	}

	return todos
}

func (b *TodoRepositoryIpml) FindById(ctx context.Context, todoId int) (model.Todo, error) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Sql := "select id,name from todo where id=$1"
	result, errQuery := tx.QueryContext(ctx, Sql, todoId)
	helper.PanicIfError(errQuery)
	defer result.Close()

	todo := model.Todo{}

	if result.Next() {
		err := result.Scan(&todoId, todo.Name)
		helper.PanicIfError(err)
		return todo, nil
	} else {
		return todo, errors.New("todo id not found")
	}
}

func (b *TodoRepositoryIpml) Save(ctx context.Context, todo model.Todo) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Sql := "insert into todo(name) values($1)"
	_, errExec := tx.ExecContext(ctx, Sql, todo.Name)
	helper.PanicIfError(errExec)
}
