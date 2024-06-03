package editor

// Buffer represents the text buffer
type Buffer struct {
    data []byte
}

// NewBuffer creates a new buffer
func NewBuffer(data []byte) *Buffer {
    return &Buffer{data: data}
}
// Add methods to manipulate the buffer

