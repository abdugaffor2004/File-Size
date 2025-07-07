package filesize

var stdUnits = [7]string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
var iecUnits = [6]string{"KiB", "MiB", "GiB", "TiB", "PiB", "EiB"}

const (
	BaseBinary  uint = 1024 // BaseBinary base used for binary-based units like KiB, MiB
	BaseDecimal uint = 1000 // BaseDecimal base used for decimal-based units like KB, MB
)

const (
	FormatStandart = "standart" // FormatStandart windows-like standard format using decimal-based units (e.g., KB, MB)
	FormatIEC      = "IEC"      // FormatIEC IEC format binary-based format (e.g., KiB, MiB) and decimal-based (e.g., KB, MB)
)
