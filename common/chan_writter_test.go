package common

import (
	"testing"
)

func TestChanWriter_Write_HappyPath(t *testing.T) {
	ch := make(chan []byte, 1)

	writer := NewChanWriter(ch)
	n, err := writer.Write([]byte("Hello, world!"))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if n != 13 {
		t.Errorf("Expected to write 13 bytes, but wrote %d", n)
	}

	data := <-ch
	if string(data) != "Hello, world!" {
		t.Errorf("Expected to write 'Hello, world!', but wrote '%s'", data)
	}
}

func TestChanWriter_Write_EmptyData(t *testing.T) {
	ch := make(chan []byte, 1)

	writer := NewChanWriter(ch)
	n, err := writer.Write([]byte(""))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if n != 0 {
		t.Errorf("Expected to write 0 bytes, but wrote %d", n)
	}

	data := <-ch
	if string(data) != "" {
		t.Errorf("Expected to write '', but wrote '%s'", data)
	}
}
