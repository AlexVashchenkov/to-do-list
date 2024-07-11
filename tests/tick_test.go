package tests

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"to-do-list/internal/models"
)

func Test_TickToDo(t *testing.T) {
	t.Parallel()

	var (
		ErrNoToDosFound      = errors.New("no to-dos found")
		ErrInvalidToDoNumber = errors.New("invalid to-do number")
		ErrTaskAlreadyDone   = errors.New("todo is already completed")
	)

	tests := []struct {
		name    string
		storage models.Storage
		number  int
		wantErr error
	}{
		{
			name:    "Отметка задачи, как выполненной",
			storage: models.NewStorage([]models.ToDo{models.NewToDo("1", "", true)}),
			number:  1,
			wantErr: nil,
		},
		{
			name:    "Задача уже выполнена",
			storage: models.NewStorage([]models.ToDo{models.NewToDo("1", "", false)}),
			number:  1,
			wantErr: ErrTaskAlreadyDone,
		},
		{
			name:    "Пустой список задач",
			storage: models.NewStorage([]models.ToDo{}),
			number:  1,
			wantErr: ErrNoToDosFound,
		},
		{
			name:    "Несуществующий номер задачи",
			storage: models.NewStorage([]models.ToDo{{Name: "1", Description: ""}}),
			number:  2,
			wantErr: ErrInvalidToDoNumber,
		},
		{
			name:    "Неверный номер задачи",
			storage: models.NewStorage([]models.ToDo{{Name: "1", Description: ""}}),
			number:  0,
			wantErr: ErrInvalidToDoNumber,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.storage.Tick(tt.number)
			assert.Equal(t, err, tt.wantErr)
		})
	}
}
