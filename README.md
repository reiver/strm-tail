# strm-tail

A driver for the **go-strm** Go programming language library, that provides the **TAIL** Ender command.

**TAIL** returns rows from the end of the data table stream.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/strm-tail

[![GoDoc](https://godoc.org/github.com/reiver/strm-tail?status.svg)](https://godoc.org/github.com/reiver/strm-tail)

## Example
```
package main

import (
	. "github.com/reiver/strm-csv"
	. "github.com/reiver/strm-tail"
	. "github.com/reiver/strm-stdout"
)

func main() {
	Begin(CSV, "table.csv").
		Strm(TAIL, 12).
	End(STDOUT, "tsv")
}
```

(Note that in that example dot imports were used.)

## See Also

For more information about **go-strm** and for a list of other drivers, see:
https://github.com/reiver/go-strm
