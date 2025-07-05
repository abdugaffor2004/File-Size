package main

import (
	"fmt"
	"math"
)

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

func main() {
	// fmt.Println(FormatWithOptions(1536, Options{Base: BaseBinary, Format: FormatIEC, Precision: 2}))
	fmt.Println(Format(-1024))
}

func Format(bytes int64) string {
	return FormatWithOptions(bytes, Options{Base: 1024, Precision: 1, Format: FormatIEC})
}

func FormatWithOptions(bytes int64, opts Options) string {
	fbytes := float64(bytes)
	pow := math.Floor(math.Log2(math.Abs(fbytes)) / 10)
	unit := determineMesureUnit(pow, opts.Format, opts.Base)

	if opts.Base == BaseDecimal {
		pow = math.Floor(math.Log10(math.Abs(fbytes)) / 3)
	}

	converted := fbytes / math.Pow(float64(opts.Base), pow)
	formatted := fmt.Sprint(toFixed(converted, opts.Precision))

	if opts.Separator == "" {
		opts.Separator = DefaultSeparator
	}

	if bytes == 0 {
		formatted = "0"
	}

	return formatted + opts.Separator + unit
}

func determineMesureUnit(pow float64, format string, base int) string {
	var stdSuffix string

	switch pow {
	case 1:
		stdSuffix = "KB"
	case 2:
		stdSuffix = "MB"
	case 3:
		stdSuffix = "GB"
	case 4:
		stdSuffix = "TB"
	case 5:
		stdSuffix = "PB"
	case 6:
		stdSuffix = "EB"

	default:
		stdSuffix = "B"
	}

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

func toFixed(n float64, precision int) float64 {
	scale := math.Pow(10, float64(precision))

	return math.Round(n*scale) / scale
}
