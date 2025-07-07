package filesize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int64
		wantErr bool
	}{
		{
			name:    "IEC format parsing. GiB",
			input:   "1.5 GiB",
			want:    1610612736,
			wantErr: false,
		},
		{
			name:    "IEC format parsing. MiB",
			input:   "100 MiB",
			want:    104857600,
			wantErr: false,
		},
		{
			name:    "IEC format parsing. KiB",
			input:   "1 KiB",
			want:    1024,
			wantErr: false,
		},
		{
			name:    "Standart format parsing. GiB",
			input:   "1.5 GB",
			want:    1610612736,
			wantErr: false,
		},
		{
			name:    "Standart format parsing. MiB",
			input:   "100 MB",
			want:    104857600,
			wantErr: false,
		},
		{
			name:    "Case-sensitive_1",
			input:   "1.5 gib",
			want:    1610612736,
			wantErr: false,
		},
		{
			name:    "Case-sensitive_2",
			input:   "1.5 Gb",
			want:    1610612736,
			wantErr: false,
		},
		{
			name:    "Without mesure units",
			input:   "2048",
			want:    2048,
			wantErr: false,
		},
		{
			name:    "IEC format parsing. MiB",
			input:   "100 MiB",
			want:    104857600,
			wantErr: false,
		},
		{
			name:    "Large size",
			input:   "100 PB",
			want:    112589990684262400,
			wantErr: false,
		},
		{
			name:    "Spaces in source string",
			input:   "  1.5 GiB  ",
			want:    1610612736,
			wantErr: false,
		},
		{
			name:    "Error handling. Empty string",
			input:   "",
			want:    0,
			wantErr: true,
		},
		{
			name:    "Error handling. Empty invalid input",
			input:   "%^^&",
			want:    0,
			wantErr: true,
		},
		{
			name:    "Error handling. Unknown unit",
			input:   "1.5 XB",
			want:    0,
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Parse(tc.input)

			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, result)
			}
		})
	}
}

func TestReversibility(t *testing.T) {
	t.Run("Reversibility of operations", func(t *testing.T) {
		result, err := Parse(Format(1500))

		if err != nil {
			assert.Error(t, err)
		}
		assert.InEpsilon(t, 1500, result, 0.03)
	})
}
