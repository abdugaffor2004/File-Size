package main

import (
	"fmt"
	"math"
	"strconv"
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
	fmt.Println(FormatWithOptions(math.MinInt64, Options{Base: BaseBinary, Format: FormatIEC, Precision: 0}))

}

func FormatWithOptions(bytes int64, opts Options) string {
	fbytes := float64(bytes)
	pow := math.Floor(math.Log2(math.Abs(fbytes)) / 10)

	if opts.Precision == 0 {
		opts.Precision = DefaultPrecision
	}

	if opts.Separator == "" {
		opts.Separator = DefaultSeparator
	}

	if opts.Base == BaseDecimal {
		pow = math.Floor(math.Log10(math.Abs(fbytes)) / 3)
	}

	converted := fbytes / math.Pow(float64(opts.Base), pow)

	return strconv.FormatFloat(converted, 'b', opts.Precision, 64) + opts.Separator + determineMesureUnit(pow, opts.Format, opts.Base)
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

	if base == BaseDecimal {
		return sf
	}

	return sf[:1] + "i" + sf[1:2]
}


