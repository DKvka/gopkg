package split

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type part struct {
	offset, size int64
}

func (p *part) Offset() int64 {
	return p.offset
}

func (p *part) Size() int64 {
	return p.size
}

// GetParts divides the given file into Parts which can be used for concurrently reading a file
// a part has the methods part.Size() and part.Offset()
func GetParts(path string, partCount int) ([]part, error) {
	const maxLineLen = 100

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	size := stat.Size()
	partSize := size / int64(partCount)

	buffer := make([]byte, maxLineLen)
	parts := make([]part, 0, partCount)
	offset := int64(0)

	for i := 0; i < partCount; i++ {
		if i == partCount-1 {
			if offset < size {
				parts = append(parts, part{offset, size - offset})
			}
			break
		}

		seekOffset := max(offset+partSize-maxLineLen, 0)
		_, err = file.Seek(seekOffset, io.SeekStart)
		if err != nil {
			return nil, err
		}

		n, _ := io.ReadFull(file, buffer)
		chunk := buffer[:n]
		newline := bytes.LastIndexByte(chunk, '\n')

		if newline < 0 {
			return nil, fmt.Errorf("newline not found at offset %d", offset+partSize-maxLineLen)
		}

		remaining := len(chunk) - newline - 1
		nextOffset := seekOffset + int64(len(chunk)) - int64(remaining)
		parts = append(parts, part{offset, nextOffset - offset})
		offset = nextOffset
	}

	return parts, nil
}
