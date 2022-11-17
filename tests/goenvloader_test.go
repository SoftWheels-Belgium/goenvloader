package test

import (
	"testing"

	"github.com/JustinCassart/goenvloader"
)

type goodStruct struct {
	Name string
	Age	 int
}

type badStruct struct {
	Name string
	Age  string
}

type tooBigStruct struct {
	Name string
	Age  int
	Pwd  string
}

type tooSmallStruct struct {
	Name string
}

type subStruct1 struct {
	Name string
}

type subStruct2 struct {
	Age int
}

type composedStruct struct {
	subStruct1
	subStruct2
}

const ENVPATH = "test.env"

func TestGoodStruct(t *testing.T) {
	s := goodStruct{}
	goenvloader.Load(ENVPATH, &s)
}

func assertPanic(t *testing.T, s interface{}, errorMessage string) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf(errorMessage)
		}
	}()
	goenvloader.Load(ENVPATH, &s)

}

func TestBadType(t *testing.T) {
	assertPanic(t, &badStruct{}, "No error found for bad type field")
}

func TestTooBigStruct(t *testing.T) {
	assertPanic(t, &tooBigStruct{}, "No error found for too big struct")
}

func TestTooSmallStruct(t *testing.T) {
	assertPanic(t, &tooSmallStruct{}, "No error found for too small struct")
}

func TestComposedStruct(t *testing.T) {
	assertPanic(t, &composedStruct{}, "No error found for composed struct")
}
