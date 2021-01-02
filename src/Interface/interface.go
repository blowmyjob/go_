package Interface

import "fmt"

type DataWriter interface {
	WriteData(data interface{}) error
}

type File struct{}

func (file *File) WriteData(data interface{}) error {
	fmt.Println("file")
	return nil
}
