package godebug

import (
	"regexp"
	"strconv"
	"strings"
)

func getFileDetails(stack string) (string, int) { // file name, line that called godebug.Break(), and a "preview line" to begin from
	re, _ := regexp.Compile(`(?m)^(.*?)(:)(\d+)`)
	res := re.FindAllStringSubmatch(stack, -1)
	fn := strings.Trim(res[1][1], " \r\n\t")
	breakLine, _ := strconv.Atoi(res[1][3])

	return fn, breakLine
}
