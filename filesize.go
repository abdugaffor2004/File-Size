package filesize

var (
	stdUnits = [7]string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	iecUnits = [7]string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB"}
)

const (
	// BaseBinary base used for binary-based units like KiB, MiB
	BaseBinary uint = 1024

	// BaseDecimal base used for decimal-based units like KB, MB
	BaseDecimal uint = 1000
)

const (
	// FormatStandart windows-like standard format using decimal-based units (e.g., KB, MB)
	FormatStandard = "standard"

	// FormatIEC IEC format binary-based format (e.g., KiB, MiB) and decimal-based (e.g., KB, MB)
	FormatIEC = "IEC"
)
