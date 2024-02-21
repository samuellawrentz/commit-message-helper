package main

import (
	jira "commit-helper/adapters"
	"commit-helper/config"
	"commit-helper/models"
	"commit-helper/storage"
	"fmt"
	"os"
	"os/exec"

	"github.com/c-bata/go-prompt"
)

var configuration config.Config
    func updateTickets(tickets []models.Issue) {
        tickets = jira.FetchIssues(configuration.UserDetails)
    }


func main() {
    fmt.Println("Welcome to Commit Helper! 🚀")
    configuration = config.InitConfig()
    storage.InitStorage()

	storage := storage.LoadFromStorage()
    tickets := storage.Tickets


    go updateTickets(tickets)

	ticketIDs := prompt.Input("Ticket ID: ", getCompleter(tickets), prompt.OptionShowCompletionAtStart())
	commitType := prompt.Input("Commit type: ", commitTypeCompleter, prompt.OptionShowCompletionAtStart())
    commitMessage := prompt.Input("Enter commit message: ", func(d prompt.Document) []prompt.Suggest {
        return []prompt.Suggest{}
    })

    fullCommitMessage := (`"` + "[" + ticketIDs + "] " + commitType + ": " + commitMessage + `"`)

	cmd := exec.Command("/bin/sh", "-c", "git commit -m "+fullCommitMessage)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Got error: %s\n", err.Error())
	}
}

func commitTypeCompleter(d prompt.Document) []prompt.Suggest {
	return prompt.FilterContains(configuration.CommitTypes, d.GetWordBeforeCursor(), true)
}

func getCompleter(tickets []models.Issue) func(prompt.Document) []prompt.Suggest {
	return func(d prompt.Document) []prompt.Suggest {
		s := []prompt.Suggest{}
		for _, ticket := range tickets {
			s = append(s, prompt.Suggest{Text: ticket.Key, Description: ticket.Fields.Summary})
		}
		return prompt.FilterContains(s, d.GetWordBeforeCursor(), true)
	}
}



