package dice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDice_Parse(t *testing.T) {

	tests := []struct {
		in      string
		expect  *Dice
		wantErr assert.ErrorAssertionFunc
	}{
		{
			in: "d%",
			expect: &Dice{
				Rolls:    1,
				DiceSize: 100,
			},
		},
		{
			in: "1d6",
			expect: &Dice{
				Rolls:    1,
				DiceSize: 6,
			},
		},
		{
			in: "2d123",
			expect: &Dice{
				Rolls:    2,
				DiceSize: 123,
			},
		},
		{
			in:      "2",
			wantErr: assert.Error,
		},
		{
			in:      "d",
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if tt.wantErr == nil {
				tt.wantErr = assert.NoError
			}
			res := &Dice{}
			err := res.Parse(tt.in)
			if !tt.wantErr(t, err) {
				return
			}

			if err == nil {
				assert.Equal(t, tt.expect, res)
			}

		})
	}
}
