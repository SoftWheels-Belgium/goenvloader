// goenvloader is a package to load environment variables from a file
// It allows to load environment variables from a file into a struct or a map
// The file must be in the format of a .env file
// example: key=value
// The file can contain comments, they must start with a #
// These lines are ignored
// example: # this is a comment
// The file can contain empty lines, they are ignored
package goenvloader

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"

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

// Load loads the environment variables from a file into a struct
//
// Input :
//   - filename : the name of the file to load
//   - config : a pointer to the struct to load the environment variables into
//
// Errors :
//   - if the file cannot be read
//   - if the structure does not match the file
//
// Effect :
//   - the structure is filled with the environment variables
//
// Example :
//
//	type Config struct {
//		Username string `env:"USERNAME"`
//		Password string `env:"PASSWORD"`
//	}
//
// Then the file .env contains :
//
//	USERNAME=foo
//	PASSWORD=bar
//
// DEPRECATED : use LoadToMap instead
func Load(filename string, config interface{}) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		return err
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
		return err
	}
	v := reflect.ValueOf(config).Elem()
	if v.NumField() != len(input) {
		return fmt.Errorf("expected %d fields but found %d", v.NumField(), len(input))
	}
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).IsZero() {
			return errors.New("Missing value for " + v.Type().Field(i).Name + " field")
		}
	}
	return nil
}

func loadToMap2(filename string) (map[string]string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	elements := regexp.MustCompile(`[\w|\d]*=(.*)`)
	env := make(map[string]string)
	for _, match := range elements.FindAll(file, -1) {
		lineSplitted := bytes.Split(match, []byte("="))
		key := string(bytes.Trim(lineSplitted[0], " "))
		value := string(bytes.Trim(bytes.Trim(lineSplitted[1], " "), "\""))
		env[key] = value
	}
	return env, nil
}

func loadToMap3(filename string) (map[string]string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	m := make(map[string]string)
	for _, line := range strings.Split(string(file), "\n") {
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		keyValue := strings.Split(line, "=")
		if len(keyValue) != 2 {
			continue
		}
		key := strings.Trim(keyValue[0], " ")
		value := strings.Trim(strings.Trim(keyValue[1], " "), "\"")
		m[key] = value
	}
	return m, nil
}

// LoadToMap loads the environment variables from a file into a map
//
// Input :
//   - filename : the name of the file to load
//
// Errors :
//   - if the file cannot be read
func LoadToMap(filename string) (map[string]string, error) {
	return loadToMap3(filename)
}
