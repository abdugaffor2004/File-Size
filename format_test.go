package filesize

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	tests := []struct {
		name  string
		input int64
		want  string
	}{
		{
			name:  "Zero value",
			input: 0,
			want:  "0 B",
		},
		{
			name:  "Byte value. More than zero",
			input: 512,
			want:  "512 B",
		},
		{
			name:  "Round KB value in IEC format. ",
			input: 1024,
			want:  "1 KiB",
		},
		{
			name:  "non-round KB value in IEC format. ",
			input: 1536,
			want:  "1.5 KiB",
		},
		{
			name:  "non-round MB value in IEC format. ",
			input: 1048576,
			want:  "1 MiB",
		},
		{
			name:  "Negative KB value in IEC format. ",
			input: -1024,
			want:  "-1 KiB",
		},
		{
			name:  "Maximum number",
			input: math.MaxInt64,
			want:  "8 EiB",
		},
		{
			name:  "Minimum number",
			input: math.MinInt64,
			want:  "-8 EiB",
		},
		{
			name:  "On the edge",
			input: 1023,
			want:  "1023 B",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Format(tc.input)
			assert.Equal(t, tc.want, result)
		})
	}
}

func TestFormatWithOptions(t *testing.T) {
	tests := []struct {
		name  string
		input int64
		opts  Options
		want  string
	}{
		{
			name:  "Binary windows-style for (0 < 1000 < 10^6) bytes",
			input: 1024,
			opts:  Options{Base: BaseBinary, Format: FormatStandard},
			want:  "1KB",
		},
		{
			name:  "Binary windows-style for (10^3 < 10^6 < 10^9) bytes",
			input: 1048576,
			opts:  Options{Base: BaseBinary, Format: FormatStandard},
			want:  "1MB",
		},
		{
			name:  "Binary system IEC format",
			input: 1024,
			opts:  Options{Base: BaseBinary, Format: FormatIEC},
			want:  "1KiB",
		},
		{
			name:  "Decimal system",
			input: 1000,
			opts:  Options{Base: BaseDecimal, Format: FormatStandard},
			want:  "1KB",
		},
		{
			name:  "Decimal system IEC format",
			input: 1000,
			opts:  Options{Base: BaseDecimal, Format: FormatIEC},
			want:  "1KB",
		},
		{
			name:  "Precision up to 0 with rounding",
			input: 1536,
			opts:  Options{Precision: 0},
			want:  "2KiB",
		},
		{
			name:  "Precision up to 2",
			input: 1536,
			opts:  Options{Precision: 2},
			want:  "1.50KiB",
		},
		{
			name:  "Precision up to 3",
			input: 1234567,
			opts:  Options{Precision: 3},
			want:  "1.177MiB",
		},
		{
			name:  "Space separator",
			input: 1024,
			opts:  Options{Separator: " "},
			want:  "1 KiB",
		},
		// {
		// 	name: "Without separator",
		// 	input: 1024,
		// 	opts: Options{Separator: ""},
		// 	want: "1KiB",
		// },
		{
			name:  "Underscore separator",
			input: 1024,
			opts:  Options{Separator: "_"},
			want:  "1_KiB",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := FormatWithOptions(tc.input, tc.opts)
			assert.Equal(t, tc.want, result)
		})
	}
}
