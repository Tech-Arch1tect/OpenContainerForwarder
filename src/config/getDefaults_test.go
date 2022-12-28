package config

import (
	"os"
	"testing"
)

// test getDefaultStr
func TestGetDefaultStr(t *testing.T) {
	// test default value
	key := "TEST"
	defaultValue := "default"
	expected := "default"
	result := getDefaultStr(key, defaultValue)
	if result != expected {
		t.Errorf("getDefaultStr(%v, %v) = %v; want %v", key, defaultValue, result, expected)
	}
	// test environment variable
	os.Setenv(key, "test")
	expected = "test"
	result = getDefaultStr(key, defaultValue)
	if result != expected {
		t.Errorf("getDefaultStr(%v, %v) = %v; want %v", key, defaultValue, result, expected)
	}
}

// test getDefaultInt
func TestGetDefaultInt(t *testing.T) {
	// test default value
	key := "TEST"
	defaultValue := 1
	expected := 1
	result := getDefaultInt(key, defaultValue)
	if result != expected {
		t.Errorf("getDefaultInt(%v, %v) = %v; want %v", key, defaultValue, result, expected)
	}
	// test environment variable
	os.Setenv(key, "2")
	expected = 2
	result = getDefaultInt(key, defaultValue)
	if result != expected {
		t.Errorf("getDefaultInt(%v, %v) = %v; want %v", key, defaultValue, result, expected)
	}
	// test invalid environment variable falls back to default value
	os.Setenv(key, "string")
	expected = 1
	result = getDefaultInt(key, defaultValue)
	if result != expected {
		t.Errorf("getDefaultInt(%v, %v) = %v; want %v", key, defaultValue, result, expected)
	}
}

// test getDefaultBool
func TestGetDefaultBool(t *testing.T) {
	// test default value
	key := "TEST"
	defaultValue := true
	expected := true
	result := getDefaultBool(key, defaultValue)
	if result != expected {
		t.Errorf("getDefaultBool(%v, %v) = %v; want %v", key, defaultValue, result, expected)
	}
	// test environment variable
	os.Setenv(key, "false")
	expected = false
	result = getDefaultBool(key, defaultValue)
	if result != expected {
		t.Errorf("getDefaultBool(%v, %v) = %v; want %v", key, defaultValue, result, expected)
	}
	// test invalid environment variable falls back to default value
	os.Setenv(key, "string")
	expected = true
	result = getDefaultBool(key, defaultValue)
	if result != expected {
		t.Errorf("getDefaultBool(%v, %v) = %v; want %v", key, defaultValue, result, expected)
	}
}
