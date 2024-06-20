package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "fmt"
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
	result, errExec := tx.ExecContext(ctx, Sql, todo.Name, todo.Id)
	row_effected, errRows := result.RowsAffected()
	fmt.Printf("result: %d\n", row_effected)
	fmt.Print("errRows: \n", errRows)
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

	Sql := "select id, name from todo where id=$1"
	result, errQuery := tx.QueryContext(ctx, Sql, todoId)
	helper.PanicIfError(errQuery)
	defer result.Close()

	todo := model.Todo{}

	if result.Next() {
		err := result.Scan(&todo.Id, &todo.Name) // lỗi xảy ra ở dòng này, tôi đã ghi vào  biến todoId thay vì ghi vào todo.Id
		helper.PanicIfError(err)
		fmt.Print("todoId: ", todo.Id) // hàm findbyId nhận được id là 0 -> dẫn đến hàm todo.Id nhân đuọc luôn mặc định là 0
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
