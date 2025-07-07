package filesize

import (
	"fmt"
	"math"
	"strconv"
)

var units = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

// BaseBinary base used for binary-based units like KiB, MiB
var BaseBinary uint = 1024

// BaseDecimal base used for decimal-based units like KB, MB
var BaseDecimal uint = 1000

// FormatStandart windows-like standard format using decimal-based units (e.g., KB, MB)
var FormatStandart = "standart"

// FormatIEC IEC format binary-based format (e.g., KiB, MiB) and decimal-based (e.g., KB, MB)
var FormatIEC = "IEC"

// DefaultPrecision default number of decimal places used in formatting
var DefaultPrecision uint = 1

// DefaultSeparator default separator between value and unit
var DefaultSeparator = " "

// Options struct allows customizing the behavior of FormatWithOptions function.
type Options struct {
	//Base base for conversion: 1024 or 1000
	Base uint

	// Precision number of decimal places to show
	Precision uint

	// Seperator character(s) used seperate number and unit
	Separator string

	// Format output in standart or IEC format
	Format string
}

// Format converts a given number of bytes into a human-readable string using default options.
func Format(bytes int64) string {
	return FormatWithOptions(bytes, Options{Base: 1024, Precision: 1, Format: FormatIEC})
}

// FormatWithOptions converts a given number of bytes into a human-readable string using custom options.
// Supports different bases (1024/1000), precision levels, separators, and output formats.
func FormatWithOptions(bytes int64, opts Options) string {
	if bytes == 0 {
		return "0" + DefaultSeparator + "B"
	}

	rawBytes := float64(bytes)
	pow := math.Floor(math.Log2(math.Abs(rawBytes)) / 10)

	if opts.Separator == "" {
		opts.Separator = DefaultSeparator
	}

	if opts.Format == "" {
		opts.Format = FormatIEC
	}

	if opts.Base == 0 {
		opts.Base = BaseBinary
	}

	if opts.Base == BaseDecimal {
		pow = math.Floor(math.Log10(math.Abs(rawBytes)) / 3)
	}

	converted := rawBytes / math.Pow(float64(opts.Base), pow)
	precised := toFixed(converted, opts.Precision)
	unit := determineMesureUnit(pow, opts.Format, opts.Base)

	return precised + opts.Separator + unit
}

func determineMesureUnit(pow float64, format string, base uint) string {
	stdSuffix := units[int(pow)]

	if format == FormatIEC {
		return iecSuffix(stdSuffix, base)
	}

	return stdSuffix
}

func iecSuffix(sf string, base uint) string {
	if base == BaseDecimal || sf == "B" {
		return sf
	}

	return sf[:1] + "i" + sf[1:2]
}

func toFixed(n float64, precision uint) string {
	if n == math.Trunc(n) || precision == 0 {
		return fmt.Sprintf("%.0f", n)
	}
	format := "%." + strconv.Itoa(int(precision)) + "f"

	return fmt.Sprintf(format, n)
}
