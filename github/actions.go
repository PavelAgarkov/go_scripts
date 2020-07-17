package github

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const ApiUrl = "https://api.github.com"

var issueResponce *Issue

var (
	action         = flag.String("action", "read", "Указывает какую операцию запускать")
	ownerName      = flag.String("owner", "", "Имя владельца репозитория")
	repositoryName = flag.String("repository", "", "Имя репозитория")
	issue          = flag.String("iss_number", "", "Номер вопроса")
)

func readIssue() {
	str := []string{"repos", *ownerName, *repositoryName, "issues", *issue}
	url := ApiUrl + strings.Join(str, "/")

	response, err := StartGetRequestWithJsonResponce(url)
	if err != nil {
		issueResponce = nil
	}
	issueResponce = response
}

func openIssue() {
	body := &RequestOpenIssue{
		Title: "Found",
		Body:  "I'm having a problem with this.",
		//Assignees: []string{"octocat"},
		//Milestone: 1,
		Labels: []string{"bug"},
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	url := []string{ApiUrl, "repos", *ownerName, *repositoryName, "issues"}
	adress := strings.Join(url, "/")

	request, err := http.NewRequest("POST", adress, buf)
	client := &http.Client{}

	request.SetBasicAuth("PavelAgarkov", "Sportteam21")

	responce, err := client.Do(request)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Body: %s\n", responce)

}

func editIssue() {

}

func closeIssue() {

}

func ReadFlagsFromTerminal() {
	flag.Parse()

	switch *action {
	case "read":
		readIssue()
		break
	case "open":
		openIssue()
		break
	case "edit":
		editIssue()
		break
	case "close":
		closeIssue()
	}

	if issueResponce != nil {
		data, err := json.MarshalIndent(issueResponce, "", " ")
		if err != nil {
			log.Fatalf("Сбой маршалинга JSON: %s", err)
		}
		fmt.Printf("%s\n", data)
	}
}
