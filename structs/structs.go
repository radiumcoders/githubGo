package structs

import "time"

type Todo struct {
	Title     string
	Completed bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTodo(title string) *[]Todo {
	return &[]Todo{
		{
			Title:     title,
			Completed: false,
			CreatedAt: time.Now().In(time.Local),
			UpdatedAt: time.Now().In(time.Local),
		},
	}
}
