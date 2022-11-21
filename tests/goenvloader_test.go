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

type longStruct struct {
	Name  string
	Other string
}

type withoutQuotes struct {
	Query string
	Value int
}

const ENVPATH = "test.env"

func TestGoodStruct(t *testing.T) {
	s := goodStruct{}
	goenvloader.Load(ENVPATH, &s)
	if s.Age != 1 {
		t.Errorf("Age error : expected 1 but found %d\n", s.Age)
	}
	if s.Name != "GOENVLOADER" {
		t.Errorf("Name error : expected 'GOENVLOADER' but found '%s'\n", s.Name)
	}
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

func TestLongStruct(t *testing.T) {
	s := longStruct{}
	goenvloader.Load("./longtest.env", &s)
	expectedName := "select * from test where test.id > 0"
	if s.Name != expectedName {
		t.Errorf("Longenv error : expected '%s' but found '%s'\n", expectedName, s.Name)
	}
	if s.Other != "Hello" {
		t.Errorf("Longenv error : expected 'Hello' but found %s\n", s.Other)
	}
}

func TestWithoutQuotes(t *testing.T) {
	s := withoutQuotes{}
	goenvloader.Load("./withoutquotes.env", &s)
	expectedQuery := "select * from test where test.id > 0"
	if s.Query != expectedQuery {
		t.Errorf("Withoutquotes error : expected '%s' but found '%s'\n", expectedQuery, s.Query)
	}
	if s.Value != 3 {
		t.Errorf("Withoutquotes error : expected 3 but found %d\n", s.Value)
	}
}
