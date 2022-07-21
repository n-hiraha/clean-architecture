package repository

import "github.com/cosmtrek/air/domain/model"

//ここではClean Architectureの依存性の順番を守るためinterface（抽象）のみを定義
//実際の実装はinfrastructure層
type TodoRepository interface {
	FindAll() (todos []*model.Todo, err error)
	Find(word string) (todos []*model.Todo, err error)
	Create(todo *model.Todo) (*model.Todo, error)
	Update(todo *model.Todo) (*model.Todo, error)
}
