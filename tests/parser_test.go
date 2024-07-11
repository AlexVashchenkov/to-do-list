package tests

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"to-do-list/internal/models"
)

func Test_ParseString(t *testing.T) {
	t.Parallel()

	var (
		controller        *models.Controller
		ErrInvalidCommand = errors.New("invalid command")
	)

	tests := []struct {
		str     string
		nice    bool
		wantErr error
	}{
		{
			str:     "add",
			nice:    true,
			wantErr: nil,
		},
		{
			str:     "add",
			nice:    true,
			wantErr: nil,
		},
		{
			str: "add",

			nice:    true,
			wantErr: nil,
		},
		{
			str:     "add",
			nice:    true,
			wantErr: nil,
		},
		{
			str:     "connect",
			nice:    false,
			wantErr: ErrInvalidCommand,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.str, func(t *testing.T) {
			t.Parallel()

			nice, err := controller.Parse(tt.str)
			assert.Equal(t, tt.nice, nice)
			assert.Equal(t, err, tt.wantErr)
		})
	}
}
