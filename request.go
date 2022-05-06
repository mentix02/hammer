package main

import (
	"log"
	"net/http"
	"net/url"
	"sync"
)

func submissionRequestWorker(idx int, wg *sync.WaitGroup) {

	defer wg.Done()

	submissionData := make(url.Values, len(*fields))

	for field, fieldType := range *fields {
		submissionData.Set(field, generateDataFromFieldType(fieldType))
	}

	resp, err := http.PostForm(endpoint, submissionData)

	if err != nil {
		log.Println(err)
	} else if int(resp.StatusCode/100) != 2 && int(resp.StatusCode%100) != 3 {
		log.Printf("req %d: %d\n", idx, resp.StatusCode)
	}

}
