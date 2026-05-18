package logic

import (
	"errors"
	"mime/multipart"
)

func readHeader(file *multipart.FileHeader, size int) ([]byte, error) {
	f, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf := make([]byte, size)
	_, err = f.Read(buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func IsPDF(file *multipart.FileHeader) error {
	buf, err := readHeader(file, 4)
	if err != nil {
		return err
	}

	if buf[0] != 0x25 || // %
		buf[1] != 0x50 || // P
		buf[2] != 0x44 || // D
		buf[3] != 0x46 { // F
		return errors.New("invalid pdf file")
	}

	return nil
}

func IsZip(file *multipart.FileHeader) error {
	buf, err := readHeader(file, 4)
	if err != nil {
		return err
	}

	if buf[0] != 0x50 || buf[1] != 0x4B {
		return errors.New("invalid zip file")
	}

	return nil
}