package file

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var random []byte = []byte("hhgakjgangnaganaaaanbakjgagmanmablajgbbahgajgangkla")
var lineSep []byte = []byte("\n")

func createRandoms(b *testing.B) []byte {
	b.Helper()

	randoms := make([]byte, 0, 10000000000)
	for i := 0; i < 1000; i++ {
		randoms = append(randoms, random...)
		randoms = append(randoms, lineSep...)
	}
	return randoms
}

func BenchmarkScanner(b *testing.B) {
	randoms := createRandoms(b)
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		cp := make([]byte, len(randoms))
		copy(cp, randoms)
		reader := bytes.NewReader(cp)
		b.StartTimer()

		scanner := bufio.NewScanner(reader)
		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		lines = lines[len(lines)-20:]
		_ = strings.Join(lines, "\n")
	}
}
