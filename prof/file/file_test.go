package file

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
)

var random []byte = []byte("hhgakjgangnaganaaaanbakjgagmanmablajgbbahgajgangkla")
var lineSep []byte = []byte("\n")

func createRandoms(b *testing.B) []byte {
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
		content := strings.Join(lines, "\n")
		fmt.Println(content)
	}
}

func BenchmarkReadByte(b *testing.B) {
	randoms := createRandoms(b)
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		cp := make([]byte, len(randoms))
		copy(cp, randoms)
		reader := bytes.NewReader(cp)
		b.StartTimer()

		var s []int
		var off int64
		var loopCount int
		byteNUM := 1024
		for {
			readByte := make([]byte, byteNUM)
			n, err := reader.ReadAt(readByte, off)
			if err != nil {
				if err == io.EOF {
					break
				} else {
					b.Error("failed file")
				}
			}
			off += int64(n)

			c := bytes.Count(readByte, lineSep)
			var lastIdx int
			for i := 0; i < c; i++ {
				idx := bytes.Index(readByte, lineSep)
				if idx < 0 {
					break
				}
				s = append(s, (loopCount*byteNUM)+idx)

				readByte = readByte[lastIdx+idx:]
				lastIdx = idx
			}

			loopCount++
		}
		offset := int64(s[len(s)-21])
		bc := make([]byte, s[len(s)-1]-s[len(s)-21]+byteNUM)
		_, err := reader.ReadAt(bc, offset)
		if err != nil {
			b.Error("error")
		}
		content := string(bc)
		fmt.Println(content)
	}
}
