package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/go-playground/validator"
	"gopkg.in/yaml.v2"
)

func main() {
	cfgPath := flag.String("config", "default.yaml", "Path to config file")
	flag.Parse()

	content, err := ioutil.ReadFile(*cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	cfg := struct {
		Line   string   `yaml:"line" validate:"required"`
		Array  []string `yaml:"array"`
		Struct struct {
			Field1 struct {
				Dsdsds string `yaml:"dsdsdsd"`
			} `yaml:"field1"`
			Field2 int
		} `yaml:"struct"`
	}{}
	err = yaml.Unmarshal(content, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	validate := validator.New()
	if err = validate.Struct(cfg); err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg.Line)
}
