package filesize

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

func Parse(s string) (int, error) {
	var rawBytes int
	prepared := strings.ToUpper(strings.TrimSpace(s))

	if prepared == "" {
		return 0, errors.New("empty input")
	}

	if isBytes(prepared) {
		return strconv.Atoi(prepared)
	}

	for i, sfx := range units[1:] {
		if stdCutted, ok := strings.CutSuffix(prepared, sfx); ok {
			rawBytes = calcRawBytes(stdCutted, i+1)
			break
		}

		if iecCutted, ok := strings.CutSuffix(prepared, strings.ToUpper(iecSuffix(sfx, BaseBinary))); ok {
			rawBytes = calcRawBytes(iecCutted, i+1)
			break
		}
	}

	if cutted, ok := strings.CutSuffix(prepared, "B"); rawBytes == 0 && ok {
		rawBytes = calcRawBytes(cutted, 0)
	}

	if rawBytes == 0 {
		return 0, errors.New("invalid input")
	}

	return rawBytes, nil
}

func calcRawBytes(bytes string, i int) (rawBytes int) {
	trimmed := strings.TrimSpace(bytes)
	convedBytes, _ := strconv.ParseFloat(trimmed, 32)
	pow := float64(i)
	rawBytes = int(convedBytes * math.Pow(float64(BaseBinary), pow))

	return
}

func isBytes(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}
