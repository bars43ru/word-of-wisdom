package pow

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
)

type Reader struct {
	in io.Reader
}

func wrapReader(in io.Reader) *Reader {
	return &Reader{in: in}
}

func (r *Reader) readNext() ([]byte, error) {
	sizeData := int64(0)
	err := binary.Read(r.in, binary.LittleEndian, &sizeData)
	if err != nil {
		return nil, fmt.Errorf("read size package data: %w", err)
	}
	buffer := make([]byte, sizeData)
	_, err = io.ReadFull(r.in, buffer)
	if err != nil {
		return nil, fmt.Errorf("read package data: %w", err)
	}
	return buffer, nil
}

func (r *Reader) read(o interface{}) error {
	data, err := r.readNext()
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, &o); err != nil {
		return fmt.Errorf("json unmarshal responce: %w", err)
	}
	return nil
}

type Writer struct {
	out io.Writer
}

func wrapWriter(out io.Writer) *Writer {
	return &Writer{out: out}
}

func (w *Writer) writeNext(b []byte) error {
	sizeData := int64(len(b))
	err := binary.Write(w.out, binary.LittleEndian, sizeData)
	if err != nil {
		return fmt.Errorf("write size package data: %w", err)
	}
	_, err = w.out.Write(b)
	if err != nil {
		return fmt.Errorf("write package data: %w", err)
	}
	return nil
}

func (w *Writer) write(o interface{}) error {
	data, err := json.Marshal(o)
	if err != nil {
		return fmt.Errorf("json marshal request: %w", err)
	}
	return w.writeNext(data)
}
