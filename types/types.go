package types

import "encoding/json"

type GraphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables,omitempty"`
}

type GraphQLResponse struct {
	Data   json.RawMessage `json:"data"`
	Errors []interface{}   `json:"errors,omitempty"`
}

type ContributionDay struct {
	Color             string `json:"color"`
	ContributionCount int    `json:"contributionCount"`
	Date              string `json:"date"`
}

type Week struct {
	ContributionDays []ContributionDay `json:"contributionDays"`
}

type ContributionCalendar struct {
	TotalContributions int    `json:"totalContributions"`
	Weeks              []Week `json:"weeks"`
}

type ContributionsCollection struct {
	ContributionCalendar ContributionCalendar `json:"contributionCalendar"`
}

type User struct {
	ContributionsCollection ContributionsCollection `json:"contributionsCollection"`
}

type GithubCalendarData struct {
	User User `json:"user"`
}

type Config struct {
	USER_NAME string
	TOKEN     string
	SYMBOL    string
}
