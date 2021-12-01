package filereader

import (
	"bufio"
	"os"
)

type FileReader struct {
	file *os.File
	buffer *bufio.Scanner
	OutCh chan string
	DoneCh chan struct{}
}

func NewFileReader(path string, outCh chan string) (FileReader, error) {
	f, err := os.Open(path)
	if err != nil {
		return FileReader{}, err
	}
	fileReader := FileReader{
		file: f,
		buffer: bufio.NewScanner(f),
		OutCh: outCh,
		DoneCh: make(chan struct{}),
	}
	return fileReader, nil
}

func (f FileReader) ReadAllToCh() {
	for f.buffer.Scan() {
		f.OutCh <- f.buffer.Text()
	}
	f.DoneCh <- struct{}{}
}

func (f FileReader) Close() {
	f.file.Close()
}