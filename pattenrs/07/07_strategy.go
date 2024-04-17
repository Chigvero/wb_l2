package main

import (
	"errors"
	"fmt"
	"log"
)

/*
	Стратегия - поведенческий паттерн, который определяет семейство схожих алгоритмов
	и помещает каждый из них в свой собственный класс, после чего алгоритмы можно из-
	менять прямо во время исполнения программы.

	+:
		1)	Горячая замена алгоритмов на лету.
		2)	Изолирует код и алгоритмы от остальных классов.
		3)	Уход от наследования к делегированию.
		4)	Реализует принцип открытости / закрытости.

	-:
		1)	Усложняет программу из-за добавления допольнительных классов.
		2)	Клиент должен знать отличия стратегий чтобы выбрать наиболее подходящую.

*/

type ConvertationStrategy interface {
	Convert(inputFile string, outputFile string) error
}

type JPEGConvertStrategy struct{}

func (j *JPEGConvertStrategy) Convert(inputFile string, outputFile string) error {
	if inputFile == "" {
		return errors.New("Error input file")
	}
	fmt.Println("File converted to JPEG format")
	return nil
}

type PNGConvertStrategy struct{}

func (p *PNGConvertStrategy) Convert(inputFile string, outputFile string) error {
	if inputFile == "" {
		return errors.New("Error input file")
	}
	fmt.Println("File converted to PNG format")
	return nil
}

type BMPConvertStrategy struct{}

func (b *BMPConvertStrategy) Convert(inputFile string, outputFile string) error {
	if inputFile == "" {
		return errors.New("Error input file")
	}
	fmt.Println("File converted to BMP format")
	return nil
}

type FileConverter struct {
	strategy ConvertationStrategy
}

func NewFileConverter(strategy ConvertationStrategy) *FileConverter {
	return &FileConverter{strategy: strategy}
}

func (fc *FileConverter) SetStrategy(strategy ConvertationStrategy) {
	fc.strategy = strategy
}

func (fc *FileConverter) ConvertFile(inputFile string, outputFile string) error {
	return fc.strategy.Convert(inputFile, outputFile)
}

func main() {
	converter := NewFileConverter(nil)

	jpegStrategy := &JPEGConvertStrategy{}
	converter.SetStrategy(jpegStrategy)

	err := converter.ConvertFile("input.png", "out.jpeg")
	if err != nil {
		log.Fatal(err)
	}

	pngStrategy := &PNGConvertStrategy{}
	converter.SetStrategy(pngStrategy)

	err = converter.ConvertFile("input.jpeg", "out.png")
	if err != nil {
		log.Fatal(err)
	}

	bmpStrategy := &BMPConvertStrategy{}
	converter.SetStrategy(bmpStrategy)

	err = converter.ConvertFile("input.png", "out.bmp")
	if err != nil {
		log.Fatal(err)
	}

}
