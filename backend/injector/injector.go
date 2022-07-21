package injector

import (
	"github.com/cosmtrek/air/domain/repository"
	"github.com/cosmtrek/air/handler"
	"github.com/cosmtrek/air/infrastructure"
	"github.com/cosmtrek/air/usecase"
)

func InjectDB() infrastructure.SqlHandler{
	sqlhandler := infrastructure.NewSqlHandler()
	return *sqlhandler
}

/*
TodoRepository(interface)に実装であるSqlHandler(struct)を渡し生成する。
*/
func InjectTodoRepository() repository.TodoRepository {
	sqlHandler := InjectDB()
	return infrastructure.NewTodoRepository(sqlHandler)
}

func InjectTodoUsecase() usecase.TodoUsecase {
	TodoRepo := InjectTodoRepository()
	return usecase.NewTodoUsecase(TodoRepo)
}

func InjectTodoHandler() handler.TodoHandler {
	return handler.NewTodoHandler(InjectTodoUsecase())
}

