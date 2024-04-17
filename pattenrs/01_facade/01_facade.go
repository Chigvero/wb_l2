package main

import (
	"fmt"
	"io"
	_ "io"
	"log"
	"os"
	_ "os"
)

/*
	Фасад.
	Вид: Структурный.
	Суть паттерна - предоставление простого или урезанного интерфейса для работы
	со сложной подсистемой (фреймворка например).
	+: изоляция клиента от сложной подсистемы, тем самым ее облегченное использование
	-: риск создания перегруженного объекта
*/

type FileFacade struct {
}

func (f *FileFacade) ReadFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	fileContent := make([]byte, fileInfo.Size())
	_, err = file.Read(fileContent)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return fileContent, err
}

func (f *FileFacade) WriteFile(filePath string, data []byte) error {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (f *FileFacade) CopyFile(fromFile string, toFile string) error {
	srcFile, err := os.Open(fromFile)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	data, err := f.ReadFile(fromFile)
	if err != nil {
		return err
	}

	dstFile, err := os.OpenFile(toFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = dstFile.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (f *FileFacade) RemoveFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	file := FileFacade{}

	read, err := file.ReadFile("/home/andy/forLinux")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(read))

	err = file.CopyFile("/home/andy/forLinux", "/home/andy/newwFile.txt")
	if err != nil {
		log.Fatal(err)
	}

	err = file.WriteFile("/home/andy/newwFile.txt", []byte("/home/andy/newwFile.txt"))
	if err != nil {
		log.Fatal(err)
	}

	err = file.RemoveFile("/home/andy/newwFile.txt")
}
