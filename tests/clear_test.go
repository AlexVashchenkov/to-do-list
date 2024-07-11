package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"to-do-list/internal/models"
)

func Test_ClearToDos(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		storage    models.Storage
		expStorage models.Storage
	}{
		{
			name: "Очистка списка задач",
			storage: models.NewStorage(
				[]models.ToDo{
					models.NewToDo("1", "", true),
					models.NewToDo("2", "", true),
					models.NewToDo("3", "", false),
				}),
			expStorage: models.NewStorage([]models.ToDo{}),
		},
		{
			name: "Пустой список задач",
			storage: models.NewStorage(
				[]models.ToDo{}),
			expStorage: models.NewStorage([]models.ToDo{}),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.storage.ClearAllToDos()
			assert.Equal(t, tt.storage, tt.expStorage)
		})
	}
}

func Test_ClearCompletedToDos(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		storage    models.Storage
		expStorage models.Storage
	}{
		{
			name: "Очистка только выполненных задач",
			storage: models.NewStorage(
				[]models.ToDo{
					models.NewToDo("1", "", true),
					models.NewToDo("2", "", true),
					models.NewToDo("3", "", false),
				}),
			expStorage: models.NewStorage([]models.ToDo{
				models.NewToDo("1", "", true),
				models.NewToDo("2", "", true),
			}),
		},
		{
			name: "Пустой список задач",
			storage: models.NewStorage(
				[]models.ToDo{}),
			expStorage: models.NewStorage([]models.ToDo{}),
		},
		{
			name: "Пустой список задач",
			storage: models.NewStorage(
				[]models.ToDo{}),
			expStorage: models.NewStorage([]models.ToDo{}),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.storage.ClearCompletedToDos()
			assert.Equal(t, tt.storage, tt.expStorage)
		})
	}
}
