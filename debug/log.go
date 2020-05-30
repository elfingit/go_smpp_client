package debug

import (
	"io"
	"os"
)

type Log struct{}

func (log *Log) Write(p []byte) (n int, err error) {
	return io.WriteString(os.Stdout, string(p))
}
