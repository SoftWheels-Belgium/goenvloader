package goenvloader

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"regexp"

	"github.com/mitchellh/mapstructure"
)

func removeCarriageReturn(text string) string {
	charriage, err := regexp.Compile(`[\r\n\t]`)
	if err != nil {
		panic(err)
	}
	spaceSuccession, err := regexp.Compile(`  +`)
	if err != nil {
		panic(err)
	}
	reformLines, err := regexp.Compile(`(\w+=)`)
	if err != nil {
		panic(err)
	}
	firstCarriage, err := regexp.Compile(`^[\n\r]`)
	if err != nil {
		panic(err)
	}
	comments, err := regexp.Compile(`#.*`)
	if err != nil {
		panic(err)
	}
	refractText := comments.ReplaceAllString(text, "")
	refractText = charriage.ReplaceAllString(refractText, " ")
	refractText = spaceSuccession.ReplaceAllString(refractText, " ")
	refractText = reformLines.ReplaceAllString(refractText, "\n$1")
	refractText = firstCarriage.ReplaceAllString(refractText, "")
	return string(refractText)
}

// Load read an environment file
// filename (string) : the path of the environment file
func Load(filename string, config interface{}) {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := removeCarriageReturn(string(file))
	input := map[string]interface{}{}
	for _, line := range strings.Split(lines, "\n") {
		lineSplitted := strings.SplitN(line, "=", 2)
		key := strings.Trim(lineSplitted[0], " ")
		value := strings.Trim(strings.Trim(lineSplitted[1], " "), "\"")
		realValue, err := strconv.Atoi(value)
		if err == nil {
			input[key] = realValue
		} else {
			input[key] = value
		}
	}
	err = mapstructure.Decode(input, config) 
	if err != nil {
		panic(err)
	}
	v := reflect.ValueOf(config).Elem()
	if v.NumField() != len(input) {
		panic(errors.New(fmt.Sprintf("Expected %d fields but found %d\n", v.NumField(), len(input))))
	}
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).IsZero() {
			panic(errors.New("Missing value for " + v.Type().Field(i).Name + " field"))
		}
	}
}

func LoadToMap(filename string) map[string]string {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := removeCarriageReturn(string(file))
	env := make(map[string]string)
	for _, line := range strings.Split(lines, "\n") {
		lineSplitted := strings.SplitN(line, "=", 2)
		key := strings.Trim(lineSplitted[0], " ")
		value := strings.Trim(strings.Trim(lineSplitted[1], " "), "\"")
		env[key] = value
	}
	return env	
}
