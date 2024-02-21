package models

type UserDetails struct {
	Username string
	Token    string
	Domain      string
	JqlURL   string
}

type Fields struct {
	Summary string `json:"summary"`
}

type Issue struct {
	Key    string `json:"key"`
	Fields Fields `json:"fields"`
}

type Storage struct {
    Tickets []Issue `json:"tickets"`
    LastUpdated string `json:"lastUpdated"`
    RecentlyUsedTickets []Issue `json:"recentlyUsedTickets"`
    RecentCommitMessages []string `json:"recentCommitMessages"`
}

type Response struct {
	Expand     string  `json:"expand"`
	StartAt    int     `json:"startAt"`
	MaxResults int     `json:"maxResults"`
	Total      int     `json:"total"`
	Issues     []Issue `json:"issues"`
}
