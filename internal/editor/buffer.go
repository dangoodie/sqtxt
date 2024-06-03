package editor

// Buffer represents the text buffer
type Buffer struct {
    data []byte
}

// NewBuffer creates a new buffer
func NewBuffer(data []byte) Buffer {
    return Buffer{
        data: data,
    }
}

// Read returns the buffer data
func (b *Buffer) Read() []byte {
    return b.data
}

// Write writes data to the buffer
func (b *Buffer) Write(data []byte) {
    b.data = data
}

// Insert inserts data at the specified index
func (b *Buffer) Insert(i int, data []byte) {
    b.data = append(b.data[:i], append(data, b.data[i:]...)...)
}

// Delete deletes data at the specified index
func (b *Buffer) Delete(i int) {
    b.data = append(b.data[:i], b.data[i+1:]...)
}
