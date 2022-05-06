package main

import (
	"net/url"
	"regexp"
	"testing"
)

var numRegex = regexp.MustCompile(`\d+`)

// TestReadFieldsFile tests that the fields file is read
// correctly and that the fields are returned as a map
// with the correct types. Else, an error is returned.
func TestReadFieldsFile(t *testing.T) {

	expectedFields := make(map[string]FieldType, 4)

	expectedFields["name"] = typeText
	expectedFields["email"] = typeEmail
	expectedFields["phone"] = typeNumber
	expectedFields["message"] = typeText

	fields := readFieldsFile("examples/forms/contact.json")

	if len(*fields) != len(expectedFields) {
		t.Errorf("Expected %d fields, got %d", len(expectedFields), len(*fields))
	}

	for fieldName, fieldType := range *fields {
		if fieldType != expectedFields[fieldName] {
			t.Errorf("Expected %s type to be %s, got %s", fieldName, expectedFields[fieldName], fieldType)
		}
	}

}

func TestFakeNumber(t *testing.T) {
	if !numRegex.MatchString(fakeNumber()) {
		t.Errorf("fakeNumber() did not return a number")
	}
}

func TestGenerateDataFromFieldType(t *testing.T) {

	var value string

	var typeToRegex = map[FieldType]*regexp.Regexp{
		typeNumber: numRegex,
		typeText:   regexp.MustCompile(`[A-Za-z]+`),
		typeTime:   regexp.MustCompile(`^\d{1,2}:\d{2}$`),
		typeDate:   regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`),
	}

	for fieldType, re := range typeToRegex {
		value = generateDataFromFieldType(fieldType)
		if !re.MatchString(value) {
			t.Errorf("Expected %s to match %s, got %s", fieldType, re.String(), value)
		}
	}

	// custom test for url since it's
	// not easily regexable (if at all)
	value = generateDataFromFieldType(typeUrl)
	_, err := url.ParseRequestURI(value)
	if err != nil {
		t.Errorf("Expected %s to be a valid url, got %s", value, err)
	}

}
