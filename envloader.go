package goenvloader

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/mitchellh/mapstructure"
)

func main() {}

// Load read an environment file
// filename (string) : the path of the environment file
func Load(filename string, config interface{}) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	lines := bufio.NewScanner(file)
	input := map[string]interface{}{}
	for lines.Scan() {
		line := lines.Text()
		lineSplitted := strings.Split(line, "=")
		key := lineSplitted[0]
		value := lineSplitted[1]
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
