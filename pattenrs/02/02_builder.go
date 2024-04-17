package main

import (
	"fmt"
	"os"
)

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

type Config struct {
	DatabaseURL string
	CasheSize   int
	LogLevel    string
	Host        string
	Port        int
	UserName    string
	Password    string
}

type ConfigBuilder struct {
	config *Config
}

func NewConfigBuilder() *ConfigBuilder {
	return &ConfigBuilder{
		config: &Config{},
	}
}

func (b *ConfigBuilder) SetDatabaseURL(url string) *ConfigBuilder {
	b.config.DatabaseURL = url
	return b
}

func (b *ConfigBuilder) SetCasheSize(size int) *ConfigBuilder {
	b.config.CasheSize = size
	return b
}

func (b *ConfigBuilder) SetLogLevel(level string) *ConfigBuilder {
	b.config.LogLevel = level
	return b
}

func (b *ConfigBuilder) SetUserName(user string) *ConfigBuilder {
	b.config.UserName = user
	return b
}

func (b *ConfigBuilder) SetPassword(password string) *ConfigBuilder {
	b.config.Password = password
	return b
}

func (b *ConfigBuilder) Build() *Config {
	return b.config
}

func main() {
	builder := NewConfigBuilder()
	builder.SetDatabaseURL("localhost:8080/mydb")
	builder.SetUserName("Builder")
	builder.SetPassword("I_am_building_your_object")
	config := builder.Build()

	file, err := os.Create("config.txt")
	if err != nil {
		fmt.Println("Error creating file: ", err)
		panic(err)
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "DatabaseURL=%s\nCasheSize=%d\nLogLevel=%s\nHost=%s\nPort=%d\nUserName=%s\nPassword=%s\n", config.DatabaseURL, config.CasheSize, config.LogLevel, config.Host, config.Port, config.UserName, config.Password)
	if err != nil {
		fmt.Println("Error writting to file: ", err)
		panic(err)
	}

	fmt.Println("Config file created successfully!")

}
