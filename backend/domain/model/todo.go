package model

//Todoアプリのドメインモデルを定義
type Todo struct {
	ID        int    `json:"id"`
	Task      string `json:"task"`
	LimitDate string `json:"limitDate"`
	Status    bool   `json:"status"`
}

type Todos []Todos
