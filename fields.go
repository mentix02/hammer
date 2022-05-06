package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/bxcodec/faker/v3"
)

type FieldType string

type Fields map[string]FieldType

const (
	typeUrl      FieldType = "url"
	typeText     FieldType = "text"
	typeDate     FieldType = "date"
	typeTime     FieldType = "time"
	typeEmail    FieldType = "email"
	typeNumber   FieldType = "number"
	typePassword FieldType = "password"
)

func readFieldsFile(fieldsFileName string) *Fields {
	jsonFile, err := os.Open(fieldsFileName)

	if err != nil {
		log.Fatalln(err)
	}

	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(jsonFile)

	var fields Fields
	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &fields)

	if err != nil {
		log.Fatalln(err)
	}

	return &fields
}

func fakeNumber() string {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(11) + 1

	number := make([]rune, n)

	for i := 0; i < n; i++ {
		if i == 0 {
			// left side of numbers cannot be 0
			number[0] = (r.Int31n(9) + 1) + '0'
		} else {
			number[i] = r.Int31n(10) + '0'
		}
	}

	return string(number)
}

func generateDataFromFieldType(fieldType FieldType) string {
	switch fieldType {
	case typeEmail:
		return faker.Email()
	case typeText:
		return faker.Sentence()
	case typePassword:
		return faker.Password()
	case typeNumber:
		return fakeNumber()
	case typeUrl:
		return faker.URL()
	case typeDate:
		return faker.Date()
	case typeTime:
		fakeTimeSlice := strings.Split(faker.TimeString(), ":")[0:2]
		return strings.Join(fakeTimeSlice, ":")
	}
	return ""
}
