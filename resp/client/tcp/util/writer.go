package util

import "io"

type Writer struct {
	r *io.Writer
}

func (w *Writer) Write(buf *[]byte) error {
	_, err := (*(w.r)).Write(*buf)
	return err
}

func NewWriter(r io.Writer) *Writer {
	return &Writer{r: &r}
}
