## Learning Outputs

Integration Testing is defined as a type of testing where software modules are integrated logically and tested as a group. A typical software project consists of multiple software modules, coded by different programmers. The purpose of this level of testing is to expose defects in the interaction between these software modules when they are integrated

## Implementation Description

this repository contains source code and integration tests for consumer api.
when started, it will start restful api on port `8080`.
You can read comments in source codes to understand what each function does.
Apart from that, here is a simple api spec about what each endpoint does:
| endpoint | method | description |
|---|---|---|
| /home | GET | basic endpoint to return plain text. Acts like healtcheck.
| /api/{endpoint} | GET | parameter based api endpoint for getting random object from mock api. When requested, it will fetch json from mock api and return one random object from list. endpoint can be either `repos`, `branches` or `commits`
| /api/internal/ | GET | internal api endpoint. for testing purposes, you can ignore.
| /api/internal/calculateResponseTime/{endpoint} | GET | after providing parameter, it will make a request to /api/{endpoint} api and calculate total request duration. this endpoint will not directly use mock api. The main purpose is to calculate response time of api requests.


## Setup

you will need golang >1.19 version. after that, to start api, you need to run following command:
```bash
go run main.go
```
It will download dependencies for the first time and listen on port 8080. Now api is ready.

Now you can run integration tests. For that, run following:
```bash
cd test
go test -v
```

