package filesize

import (
	"fmt"
	"math"
	"strconv"
)

var units = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

var BaseBinary = 1024
var BaseDecimal = 1000
var FormatStandart = "standart"
var FormatIEC = "IEC"
var DefaultPrecision = 1
var DefaultSeparator = " "

type Options struct {
	Base      int
	Precision int
	Separator string
	Format    string
}

func Format(bytes int64) string {
	return FormatWithOptions(bytes, Options{Base: 1024, Precision: 1, Format: FormatIEC})
}

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

func determineMesureUnit(pow float64, format string, base int) string {
	stdSuffix := units[int(pow)]

	if format == FormatIEC {
		return iecSuffix(stdSuffix, base)
	}

	return stdSuffix
}

func iecSuffix(sf string, base int) string {
	if base == BaseDecimal || sf == "B" {
		return sf
	}

	return sf[:1] + "i" + sf[1:2]
}

func toFixed(n float64, precision int) string {
	if n == math.Trunc(n) || precision == 0 {
		return fmt.Sprintf("%.0f", n)
	}
	format := "%." + strconv.Itoa(precision) + "f"

	return fmt.Sprintf(format, n)
}
