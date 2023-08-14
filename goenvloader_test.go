package goenvloader

import (
	"testing"
)

type goodStruct struct {
	Name string
	Age  int
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

type realQueries struct {
	Resources  string
	Tyre       string
	WebParams  string
	TaskParams string
	Scheduler  string
}

const ENVPATH = "test.env"

func TestGoodStruct(t *testing.T) {
	s := goodStruct{}
	Load(ENVPATH, &s)
	if s.Age != 1 {
		t.Errorf("Age error : expected 1 but found %d\n", s.Age)
	}
	if s.Name != "GOENVLOADER" {
		t.Errorf("Name error : expected 'GOENVLOADER' but found '%s'\n", s.Name)
	}
}

func assertPanic(t *testing.T, s interface{}, errorMessage string) {
	if err := Load(ENVPATH, &s); err == nil {
		t.Errorf(errorMessage)
	}
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
	Load("./longtest.env", &s)
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
	Load("./withoutquotes.env", &s)
	expectedQuery := "select * from test where test.id > 0"
	if s.Query != expectedQuery {
		t.Errorf("Withoutquotes error : expected '%s' but found '%s'\n", expectedQuery, s.Query)
	}
	if s.Value != 3 {
		t.Errorf("Withoutquotes error : expected 3 but found %d\n", s.Value)
	}
}

func TestRealQueries(t *testing.T) {
	s := realQueries{}
	Load("./realqueries.sql", &s)
}

func compareMap(t *testing.T, firstMap, secondMap map[string]string) {
	for k, v := range firstMap {
		v2 := secondMap[k]
		if v != v2 {
			t.Errorf("Bad value for '%s' : expected '%s' but found '%s'", k, v, v2)
		}
	}
}

func TestLoadToMap(t *testing.T) {
	env, err := LoadToMap("./testmap.env")
	if err != nil {
		t.Error(err)
	}
	expectedEnv := map[string]string{
		"key1": "hello bonjour",
		"key2": "test",
		"key3": "long long long message",
	}
	if len(env) != len(expectedEnv) {
		t.Errorf("Expected %d elements but found %d", len(expectedEnv), len(env))
	}
	compareMap(t, env, expectedEnv)
	compareMap(t, expectedEnv, env)
}

func TestPasswordWithHashTag(t *testing.T) {
	env, err := loadToMap2("./passwordwithhashtag.env")
	if err != nil {
		t.Error(err)
	}
	expectedEnv := map[string]string{
		"PASS": "test#123",
		"USER": "user",
	}
	if len(env) != len(expectedEnv) {
		t.Errorf("Expected %d elements but found %d", len(expectedEnv), len(env))
	}
	compareMap(t, env, expectedEnv)
}

func TestPasswordWithHashTag2(t *testing.T) {
	env, err := loadToMap3("./passwordwithhashtag.env")
	if err != nil {
		t.Error(err)
	}
	expectedEnv := map[string]string{
		"PASS": "test#123",
		"USER": "user",
	}
	if len(env) != len(expectedEnv) {
		t.Errorf("Expected %d elements but found %d", len(expectedEnv), len(env))
	}
	compareMap(t, env, expectedEnv)
}

func BenchmarkLoadToMap2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loadToMap2("./passwordwithhashtag.env")
	}
}

func BenchmarkLoadToMap3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loadToMap3("./passwordwithhashtag.env")
	}
}
