package main

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// Visitor - интерфейс посетителя.
type Visitor interface {
	VisitJSON(data []byte) error
	VisitXML(data []byte) error
	VisitCSV(data [][]string) error
}

type Person struct {
	Name string `xml:"name"`
	Age  int    `xml:"age"`
}

// DataFormat - абстрактный класс, представляющий формат данных.
type DataFormat interface {
	Accept(visitor Visitor) error
}

// JSONData - класс, представляющий данные в формате JSON.
type JSONData struct {
	Data []byte
}

// Accept - метод для принятия посетителя.
func (j *JSONData) Accept(visitor Visitor) error {
	return visitor.VisitJSON(j.Data)
}

// XMLData - класс, представляющий данные в формате XML.
type XMLData struct {
	Data []byte
}

// Accept - метод для принятия посетителя.
func (x *XMLData) Accept(visitor Visitor) error {
	return visitor.VisitXML(x.Data)
}

// CSVData - класс, представляющий данные в формате CSV.
type CSVData struct {
	Data [][]string
}

// Accept - метод для принятия посетителя.
func (c *CSVData) Accept(visitor Visitor) error {
	return visitor.VisitCSV(c.Data)
}

// ExtractVisitor - класс-посетитель, который извлекает определенные поля из каждого элемента данных.
type ExtractVisitor struct {
	Fields []string
}

// VisitJSON - метод для посещения JSON-данных.
func (e *ExtractVisitor) VisitJSON(data []byte) error {
	fmt.Println("VisitJson: ")
	var obj map[string]interface{}
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return err
	}

	for _, field := range e.Fields {
		val, ok := obj[field]
		if ok {
			fmt.Println(field, ":", val)
		}
	}
	fmt.Printf("END JSON\n\n")
	return nil
}

// VisitXML - метод для посещения XML-данных.
func (e *ExtractVisitor) VisitXML(data []byte) error {
	fmt.Println("VisitXML: ")
	xmlData := `
	<person>
		<name>Alice</name>
		<age>25</age>
	</person>`
	var p Person

	//var obj map[string]interface{}
	err := xml.Unmarshal([]byte(xmlData), &p)
	if err != nil {
		fmt.Println("ERROR!")
		return err
	}

	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
	fmt.Printf("END XML\n\n")
	return nil
}

func (e *ExtractVisitor) VisitCSV(data [][]string) error {
	fmt.Println("VisitCSV: ")
	header := data[0]

	for _, row := range data[1:] {
		for _, field := range e.Fields {
			index := -1
			for i, h := range header {
				if h == field {
					index = i
					break
				}
			}
			if index >= 0 {
				fmt.Println(field, ":", row[index])
			}
		}
	}
	fmt.Printf("END CSV\n\n")
	return nil
}

func visitFormats(format DataFormat, extract *ExtractVisitor) error {
	err := format.Accept(extract)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// Создание данных для посещения.
	jsonData := &JSONData{Data: []byte(`{"name": "Alice", "age": 25}`)}
	xmlData := &XMLData{Data: []byte(`<person><name>Alice</name><age>25</age></person>`)}
	csvData := &CSVData{Data: [][]string{{"name", "age"}, {"Alice", "25"}, {"Bob", "30"}}}

	// Создание посетителя.
	extractVisitor := &ExtractVisitor{Fields: []string{"name", "age"}}
	obj := make(map[int]DataFormat)
	obj[0] = jsonData
	obj[1] = xmlData
	obj[2] = csvData

	for _, value := range obj {
		if err := visitFormats(value, extractVisitor); err != nil {
			fmt.Println(err)
		}
	}
}
