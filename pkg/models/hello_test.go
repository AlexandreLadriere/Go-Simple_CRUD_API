package models

import (
	"testing"
)

func TestGetHellos(t *testing.T) {
	result := GetHellos() // should return [{Hello en} {Bonjour fr}]
	if len(result) != 2 && result[0].Lang != "en" && result[0].Msg != "Hello" && result[1].Lang != "fr" && result[1].Msg != "Bonjour" {
		t.Errorf("GetHellos() failed, expected %v, got %v", "[{Hello en} {Bonjour fr}]", result)
	}
}

func TestGetHello(t *testing.T) {
	emptyResult := GetHello("it") // should return ""
	if emptyResult != "" {
		t.Errorf("GetHello(\"it\") failed, expected %v, got %v", "", emptyResult)
	}
	result := GetHello("fr") // should return "Bonjour"
	if result != "Bonjour" {
		t.Errorf("GetHello(\"fr\") failed, expected %v, got %v", "Bonjour", result)
	}
}

func TestCreateNewHello(t *testing.T) {
	result := CreateNewHello(Hello {Msg: "Hallo", Lang: "de",}) // should add {Msg: "Hallo", Lang: "de",}
	result1 := GetHellos() // result1[2].Lang should be "de" and result1[2].Msg should be "Hallo"
	if result1[2].Lang != "de" && result1[2].Msg != "Hallo" && len(result1) != 3 {
		t.Errorf("CreateNewHello(Hello {Msg: \"Hallo\", Lang: \"de\",}) failed, expected %v, got %v", "[{Hello en} {Bonjour fr} {Hallo de}]", result1)
	}
	if result.Msg != "Hallo" && result.Lang != "de" {
		t.Errorf("CreateNewHello(Hello {Msg: \"Hallo\", Lang: \"de\",}) failed, expected %v, got %v", "{Hallo de}", result)
	}
}

func TestDeleteHello(t *testing.T) {
	emptyResult := DeleteHello("it") // should return {}
	if emptyResult.Msg != "" && emptyResult.Lang != "" {
		t.Errorf("DeleteHello(\"it\") failed, expected %v, got %v", "{}", emptyResult)
	}
	result := DeleteHello("en") // should return {Hello en}
	if result.Lang != "en" && result.Msg != "Hello" {
		t.Errorf("DeleteHello(\"en\") failed, expected %v, got %v", "{Hello en}", result)
	}
	result1 := GetHellos() // should return [{Bonjour fr}]
	if len(result1) != 1 && result1[0].Msg != "Bonjour" && result1[0].Lang != "fr" {
		t.Errorf("DeleteHello(\"en\") failed, expected %v, got %v", "[{Bonjour fr}]", result1)
	}
}
