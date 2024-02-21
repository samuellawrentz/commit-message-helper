# Commit Message Helper

Commit Helper is a Go-based CLI application that helps you in formatting your git commit messages. It fetches issue tickets from JIRA, and guides you in structuring your commit message with ticket ID, commit type, and a brief note.

## Functionality

1. Fetches issue tickets from JIRA using given user credentials.
2. Prompts user to input ticket ID, commit type, and commit message.
3. Forms the commit message in required format: `"[TicketID] CommitType: CommitMessage"`.
4. Commits the changes with the formed commit message.

## Usage

1. Download the binary
2. Put it in bin folder
3. Use it by running `commit-helper` in the terminal

## Note

To tailor the application according to your needs, make adjustments in the `config` package. Update the user credentials, commit types, and other details accordingly.

Make use of this commit helper and ease your committing process! Happy Coding! 🚀
