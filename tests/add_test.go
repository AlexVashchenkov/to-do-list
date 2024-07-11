package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"to-do-list/internal/models"
)

func Test_AddToDo(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		storage models.Storage
		toDo    models.ToDo
	}{
		{
			name:    "Добавление задачи в список задач",
			storage: models.NewStorage([]models.ToDo{{Name: "1", Description: ""}}),
			toDo:    models.NewToDo("some task", "", true),
		},
		{
			name:    "Добавление задачи в пустой список",
			storage: models.NewStorage([]models.ToDo{}),
			toDo:    models.NewToDo("some task", "Description", true),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			length := len(tt.storage.ListOfToDos())
			tt.storage.Add(tt.toDo)
			assert.Equal(t, len(tt.storage.ListOfToDos()), length+1)
			assert.Contains(t, tt.storage.ListOfToDos(), tt.toDo)
		})
	}
}
