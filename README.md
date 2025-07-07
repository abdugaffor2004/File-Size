# File-Size

A lightweight Go package for converting between bytes and human-readable file sizes with support for both binary and decimal units.

## Features

- Convert bytes to human-readable format (B, KB, MB, GB, TB, PB, EB)
- Support for both binary (1024) and decimal (1000) bases
- IEC format with binary suffixes (KiB, MiB, GiB)
- Standard format with decimal suffixes (KB, MB, GB)
- Parse human-readable sizes back to bytes
- Configurable precision and separators
- Clean, predictable output

## Installation

```bash
go get github.com/abdugaffor2004/filesize
```

## Usage

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/abdugaffor2004/filesize"
)

func main() {
    // Default format (IEC binary units)
    fmt.Println(filesize.Format(1024))        // Output: 1 KiB
    fmt.Println(filesize.Format(1048576))     // Output: 1 MiB
    fmt.Println(filesize.Format(1073741824))  // Output: 1 GiB
}
```

### Advanced Usage

```go
// Custom formatting options
options := filesize.Options{
    Base:      1000,           // Decimal base
    Precision: 2,              // Two decimal places
    Separator: "",             // No separator
    Format:    "standart",     // Standard units (KB, MB)
}

result := filesize.FormatWithOptions(1024, options)
fmt.Println(result) // Output: 1.02KB
```

### Parsing

```go
// Parse human-readable sizes back to bytes
bytes, err := filesize.Parse("1.5 KiB")
if err != nil {
    log.Fatal(err)
}
fmt.Println(bytes) // Output: 1536

bytes, err = filesize.Parse("2.5 MB")
fmt.Println(bytes) // Output: 2621440
```

## API

### Functions

- `Format(bytes int64) string` - Convert bytes to human-readable format with default options
- `FormatWithOptions(bytes int64, opts Options) string` - Convert with custom options
- `Parse(s string) (int, error)` - Parse human-readable string back to bytes

### Options

```go
type Options struct {
    Base      uint    // Base for conversion: 1024 (binary) or 1000 (decimal)
    Precision uint    // Number of decimal places to show
    Separator string // Character(s) used to separate number and unit
    Format    string // Output format: "IEC" or "standart"
}
```

### Constants

- `BaseBinary = 1024` - Binary base for IEC units
- `BaseDecimal = 1000` - Decimal base for standard units
- `FormatIEC = "IEC"` - IEC format (KiB, MiB, GiB)
- `FormatStandart = "standart"` - Standard format (KB, MB, GB)

## Examples

| Input (bytes) | Format | Output |
|---------------|--------|--------|
| `1024` | Default | `"1 KiB"` |
| `1024` | Standard | `"1 KB"` |
| `1048576` | IEC | `"1 MiB"` |
| `1000000` | Standard | `"1 MB"` |
| `0` | Any | `"0 B"` |

| Input (string) | Parsed (bytes) |
|----------------|----------------|
| `"1.5 KiB"` | `1536` |
| `"2.5 MB"` | `2621440` |
| `"1 GB"` | `1073741824` |
| `"1 GiB"` | `1073741824` |
