package main

import (
	"flag"
	"net/url"
	"os"
	"sync"
)

var fields *Fields
var endpoint string

func main() {

	flag.StringVar(&endpoint, "url", "", "url of form")
	nFlag := flag.Int("n", 1, "number of submissions to post")
	fileNameFlag := flag.String("file", "", "json file with field names & types")

	flag.Parse()

	if *fileNameFlag == "" || *nFlag < 1 || endpoint == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	wg := new(sync.WaitGroup)
	fields = readFieldsFile(*fileNameFlag)

	submissionData := make(url.Values, len(*fields))

	for field, fieldType := range *fields {
		submissionData.Set(field, generateDataFromFieldType(fieldType))
	}

	for i := 0; i < *nFlag; i++ {
		wg.Add(1)
		go submissionRequestWorker(i, wg)
	}

	wg.Wait()

}
