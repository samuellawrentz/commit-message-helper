package jira

import (
	"commit-helper/data"
	"commit-helper/models"
	"commit-helper/storage"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
	"github.com/c-bata/go-prompt"
)

func promptUser(message string) string {
    return prompt.Input(message, func(d prompt.Document) []prompt.Suggest {
        return []prompt.Suggest{}
    })
}

func GetUserDetails() models.UserDetails {
    var userDetails models.UserDetails
    userDetails.Username = promptUser("Enter your JIRA username: ")
    userDetails.Token = promptUser("Enter your JIRA token: ")
    userDetails.Domain = promptUser("Enter your Company JIRA Domain (Eg: company.atlassian.net): ")
    userDetails.JqlURL = "https://" + userDetails.Domain + "/rest/api/3/search/?jql=" +
          url.QueryEscape("updated >= -20d AND project = CNTO AND assignee in (currentUser()) order by updated DESC") + "&maxResults=15"

    return userDetails
}

func FetchIssues(userDetails models.UserDetails) []models.Issue {
    fmt.Println("Fetching Latest JIRA issues...")
    client := &http.Client{}
    req, _ := http.NewRequest("GET", userDetails.JqlURL, nil)
    basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(userDetails.Username+":"+userDetails.Token))
    req.Header.Add("Authorization", basicAuth)
    req.Close = true

    response, err := client.Do(req)

    if err != nil {
        fmt.Println("Error: ", err)
        return []models.Issue{}
    }

    defer response.Body.Close()
    body, _ := io.ReadAll(response.Body)
    var responseObj models.Response
    err = json.Unmarshal(body, &responseObj)
    if err != nil {
        fmt.Println("Error: ", err)
        return []models.Issue{}
    }
    jiraIssues := responseObj.Issues

    // Write fetched issues to JSON file
    data.GetData().SetTicketData(models.Storage{Tickets: jiraIssues, LastUpdated: time.Now().String()})
    storage.WriteToStorage()
    return jiraIssues
}
