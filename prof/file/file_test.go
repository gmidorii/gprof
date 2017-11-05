package file

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"testing"
)

var random []byte = []byte("hhgakjgangnaganaaaanbakjgagmanmablajgbbahgajgangkla")
var lineSep []byte = []byte("\n")

func makeReader(b *testing.B) io.Reader {
	b.Helper()

	randoms := make([]byte, 0, 10000000000)
	for i := 0; i < 1000; i++ {
		randoms = append(randoms, random...)
		randoms = append(randoms, lineSep...)
	}
	return bytes.NewReader(randoms)
}

func BenchmarkFileCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		reader := makeReader(b)
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
