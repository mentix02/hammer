# Documentation

Hammer is a blazingly fast CLI tool to send concurrent POST requests to a web service.
It is super simple to use & configure. Simply create a JSON file with keys as POST data
names and values as either of `text`, `url`, `email`, `number`, `time`, `date`, or `password`
and hammer will automatically mock fake data and send it as form data to the url specefied
as a command line argument.

Example invocation (file: contact.json) - 
```json
{
    "name": "text",
    "email": "email",
    "phone": "number",
    "message": "text"
}
```

```bash
$ hammer -file=contact.json -url=https://httpbin.org/post -n 10
```

The most important feature of hammer comes with the `-n` flag - the number of requests to
send concurrently. It uses lightweight goroutines to execute all the requests in parallel 
(hopefully) and does so super quickly.

## Installation

You can download the platform specific binaries from the release page or use `go get` to
install it from source (you must have Go installed)

```sh
$ go get github.com/mentix02/hammer
```

## Usage

Hammer requires a JSON file to generate fake form data. Create a JSON file with your keys
as the field names and the values as one of the following types - 

```
url
date
text
time
email
number
password
```

Save your file and pass it to the `-file` flag. Next pass the endpoint your web service
expects a POST request to the `-url` flag. Finally, if you wish to send a bunch of requests
(as you probably would since you're using this and not curl), pass a number to the `-n` flag.

### Warning

Remember - all the requests will be attempted to be sent concurrently and Hammer is so fast
that your service may not be able to handle all requests. There **may** be timeouts.

This is **NOT** a DDOS too and should not be used as such. Any damages to any web services 
caused by the usage of this tool will bear no responsibility on Hammer's end. Use -n at your
own caution.
