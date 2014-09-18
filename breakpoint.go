package godebug

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

type breakpoint struct {
	id      string // Hash of filename + line number
	enabled bool
}

func generateBreakpointId(file string, line int) string {
	h := md5.New()
	r := fmt.Sprintf("%s - %d", file, line)
	io.WriteString(h, r)
	return hex.EncodeToString(h.Sum(nil))
}
