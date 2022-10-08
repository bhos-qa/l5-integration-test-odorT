package controllers

import (
	"awesomeProject/models"
	"awesomeProject/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)


// GetRandom is a function that returns a random message based on endpoint request parameter
// @endpoint can be either "repos", "branches" or "commits"
func GetRandom(rw http.ResponseWriter, r *http.Request) {
	var repos = models.RepoModel{}
	var branches = models.BranchModel{}
	var commits = models.CommitModel{}
	var response = models.ResponseModel{}

	var endpoint = mux.Vars(r)["endpoint"]
	var randomNumber int

	data, err := http.Get(utils.BuildURL(r))
	if err != nil {
		log.Fatalln(err)
	}

	readData, err := io.ReadAll(data.Body)
	if err != nil {
		log.Println(err)
	}

	rand.Seed(time.Now().UnixMicro())

	switch endpoint {
	case "commits":
		err = json.Unmarshal(readData, &commits)
		randomNumber = rand.Intn(len(commits)) // this is for getting random repo from repos list.
		response.Message = fmt.Sprintf("Random commit message: %s with id: %s", commits[randomNumber].Message, commits[randomNumber].ID)
	case "repos":
		err = json.Unmarshal(readData, &repos)
		randomNumber = rand.Intn(len(repos))
		response.Message = fmt.Sprintf("Random repo name: %s with id: %s", repos[randomNumber].Name, repos[randomNumber].ID)
	case "branches":
		err = json.Unmarshal(readData, &branches)
		randomNumber = rand.Intn(len(branches))
		response.Message = fmt.Sprintf("Random branch name: %s with id: %s", branches[randomNumber].Name, branches[randomNumber].ID)
	}
	if err != nil {
		log.Println(err)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	err = json.NewEncoder(rw).Encode(response)
	if err != nil {
		log.Println(err)
	}
}

// GetInternal will return static text message in json format about internal endpoint
func GetInternal(rw http.ResponseWriter, r *http.Request) {
	var response = models.ResponseModel{}

	response.Message = "Internal Route"

	rw.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(rw).Encode(response)
	if err != nil {
		log.Println(err)
	}
}

// GetResponseTime will calculate the request duration and return the response time in seconds in json format
// It will request the endpoint of GetRandom function
func GetResponseTime(rw http.ResponseWriter, r *http.Request) {
	var endpoint = mux.Vars(r)["endpoint"]
	var response = models.ResponseModel{}

	start := time.Now()

	data, err := http.Get(utils.BuildInternalURL(r))
	if err != nil {
		log.Fatalln(err)
	}

	readData, err := io.ReadAll(data.Body)
	if err != nil {
		log.Println(err)
	}

	elapsed := time.Since(start).Seconds()
	response.Message = fmt.Sprintf("Time for processing request to %s endpoint: %fs", endpoint, elapsed)

	log.Print("Response" + string(readData))

	rw.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(rw).Encode(response)
	if err != nil {
		log.Println(err)
	}
}
