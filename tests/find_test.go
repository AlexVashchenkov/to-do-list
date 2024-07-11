package tests

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"to-do-list/internal/models"
)

func Test_FindToDo(t *testing.T) {
	t.Parallel()

	var (
		ErrNoToDosFound      = errors.New("no to-dos found")
		ErrInvalidToDoNumber = errors.New("invalid to-do number")
	)

	tests := []struct {
		name     string
		storage  models.Storage
		number   int
		wantToDo *models.ToDo
		wantErr  error
	}{
		{
			name:     "Поиск задачи в списке",
			storage:  models.NewStorage([]models.ToDo{{Name: "1", Description: ""}}),
			number:   1,
			wantErr:  nil,
			wantToDo: &models.ToDo{Name: "1", Description: ""},
		},
		{
			name:     "Пустой список задач",
			storage:  models.NewStorage([]models.ToDo{}),
			number:   1,
			wantErr:  ErrNoToDosFound,
			wantToDo: nil,
		},
		{
			name:     "Несуществующий номер задачи",
			storage:  models.NewStorage([]models.ToDo{{Name: "1", Description: ""}}),
			number:   2,
			wantErr:  ErrInvalidToDoNumber,
			wantToDo: nil,
		},
		{
			name:     "Неверный номер задачи",
			storage:  models.NewStorage([]models.ToDo{{Name: "1", Description: ""}}),
			number:   0,
			wantErr:  ErrInvalidToDoNumber,
			wantToDo: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			toDo, err := tt.storage.FindByNumber(tt.number)
			assert.Equal(t, err, tt.wantErr)
			assert.Equal(t, toDo, tt.wantToDo)
		})
	}
}
