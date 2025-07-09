package filesize

import (
	"fmt"
	"math"
)

const (
	logDecimalPow = 3
	logBinaryPow  = 10
)

// Options struct allows customizing the behavior of FormatWithOptions function.
type Options struct {
	// Base base for conversion: 1024 or 1000
	Base uint

	// Precision number of decimal places to show
	Precision uint

	// Separator character(s) used separate number and unit
	Separator string

	// Format output in standard or IEC format
	Format string
}

// Format converts a given number of bytes into a human-readable string using default options.
func Format(bytes int64) string {
	return FormatWithOptions(bytes, Options{Base: BaseBinary, Precision: 1, Format: FormatIEC, Separator: " "})
}

// FormatWithOptions converts a given number of bytes into a human-readable string using custom options.
// Supports different bases (1024/1000), precision levels, separators, and output formats.
func FormatWithOptions(bytes int64, opts Options) string {
	if opts.Format == "" {
		opts.Format = FormatIEC
	}

	if opts.Base == 0 {
		opts.Base = BaseBinary
	}

	if bytes == 0 {
		return "0" + opts.Separator + "B"
	}

	rawBytes := float64(bytes)
	absBytes := math.Abs(rawBytes)
	var pow float64

	if absBytes >= float64(opts.Base) {
		if opts.Base == BaseDecimal {
			pow = math.Floor(math.Log10(absBytes) / logDecimalPow)
		} else {
			pow = math.Floor(math.Log2(absBytes) / logBinaryPow)
		}
	}

	converted := rawBytes / math.Pow(float64(opts.Base), pow)
	formatted := formatNumber(converted, opts.Precision)
	unit := determineUnit(pow, opts.Format)

	return formatted + opts.Separator + unit
}

func determineUnit(pow float64, format string) string {
	powIdx := int(pow)

	if format == FormatIEC {
		return iecUnits[powIdx]
	}

	return stdUnits[powIdx]
}

func formatNumber(n float64, precision uint) string {
	if n == math.Trunc(n) || precision == 0 {
		return fmt.Sprintf("%.0f", n)
	}

	format := fmt.Sprintf("%%.%df", precision)

	return fmt.Sprintf(format, n)
}
