package tests

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"to-do-list/internal/models"
)

func Test_DeleteToDo(t *testing.T) {
	t.Parallel()

	var (
		ErrNoToDosFound      = errors.New("no to-dos found")
		ErrInvalidToDoNumber = errors.New("invalid to-do number")
	)

	tests := []struct {
		name    string
		storage models.Storage
		number  int
		wantErr error
	}{
		{
			name:    "Удаление задачи из списка задач",
			storage: models.NewStorage([]models.ToDo{{Name: "1", Description: ""}}),
			number:  1,
			wantErr: nil,
		},
		{
			name:    "Удаление задачи из пустого списка",
			storage: models.NewStorage([]models.ToDo{}),
			number:  0,
			wantErr: ErrNoToDosFound,
		},
		{
			name:    "Несуществующий номер задачи",
			storage: models.NewStorage([]models.ToDo{{Name: "1", Description: ""}}),
			number:  0,
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

			err := tt.storage.DeleteByNumber(tt.number)
			assert.Equal(t, err, tt.wantErr)
		})
	}
}
