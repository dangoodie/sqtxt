package buffer

import (
	"bufio"
	"bytes"
	"fmt"

	structs "github.com/dangoodie/sqtxt/internal/structs"
)

// Buffer represents the text buffer
type Buffer struct {
	data [][]byte
}

// NewBuffer creates a new buffer
func NewBuffer(data []byte) *Buffer {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	b := &Buffer{}
	for scanner.Scan() {
		line := scanner.Bytes()
		b.data = append(b.data, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading buffer:", err)
		panic(err)
	}

	return b
}

// Read returns the buffer data
func (b *Buffer) Read() string {
	var buf bytes.Buffer
	for _, line := range b.data {
		buf.Write(line)
		buf.WriteByte('\n')
	}
	return buf.String()
}

// Insert inserts data at the specified index
func (b *Buffer) Insert(p structs.Position, data []byte) {
	row := p.Row
	col := p.Col

	if len(b.data) == 0 {
		b.data = append(b.data, []byte{})
	}

	if row >= len(b.data) {
		b.data = append(b.data, []byte{})
	}

	if col >= len(b.data[row]) {
		b.data[row] = append(b.data[row], data...)
	} else {
		b.data[row] = append(b.data[row][:col], append(data, b.data[row][col:]...)...)
	}

	// fmt.Println("Buffer data:", b.data)
}

// Delete deletes data at the specified index
func (b *Buffer) Delete(p structs.Position) {
	row := p.Row
	col := p.Col

	b.data[row] = append(b.data[row][:col], b.data[row][col+1:]...)
}

// Size returns the size of the buffer
func (b *Buffer) Size() int {
	return len(b.data)
}

// LineLength returns the length of a line
func (b *Buffer) LineLength(row int) int {
	return len(b.data[row])
}

// NumLines returns the height of the buffer
func (b *Buffer) NumLines() int {
	return len(b.data)
}
