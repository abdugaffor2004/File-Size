package filesize

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

// Parse converts a human-readable string (e.g., "1.5 KiB") into the corresponding number of bytes.
// Supports both IEC (KiB, MiB) and SI (KB, MB) formats with appropriate bases.
// Returns an error if input is invalid or cannot be parsed.
func Parse(s string) (int64, error) {
	prepared := strings.ToUpper(strings.TrimSpace(s))
	if prepared == "" {
		return 0, errors.New("empty input")
	}

	if bytes, ok := parseNumber(prepared); ok {
		return bytes, nil
	}

	if bytes, ok := parseWithUnits(prepared, stdUnits); ok {
		return bytes, nil
	}

	if bytes, ok := parseWithUnits(prepared, iecUnits); ok {
		return bytes, nil
	}

	return 0, errors.New("invalid input")
}

func parseWithUnits(s string, units [7]string) (int64, bool) {
	for i, unit := range units {
		upperedUnit := strings.ToUpper(unit)
		if stdCutted, ok := strings.CutSuffix(s, upperedUnit); ok {
			bytes, err := convertToBytes(stdCutted, i)
			if err != nil {
				continue
			}

			return bytes, true
		}
	}

	return 0, false
}

func convertToBytes(numStr string, pow int) (int64, error) {
	trimmed := strings.TrimSpace(numStr)
	if trimmed == "" {
		return 0, errors.New("empty number")
	}

	convedBytes, err := strconv.ParseFloat(trimmed, 64)
	if err != nil {
		return 0, err
	}

	rawBytes := convedBytes * math.Pow(float64(BaseBinary), float64(pow))

	return int64(rawBytes), nil
}

func parseNumber(s string) (int64, bool) {
	bytes, err := strconv.ParseInt(s, 10, 64)

	return bytes, err == nil
}
