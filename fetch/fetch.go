package fetch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/EvaLLLLL/ghcld/types"
)

func FetchGithubCalendar(config *types.Config) *[]types.Week {
	apiURL := "https://api.github.com/graphql"

	query := `query ($userName: String!) {
  user(login: $userName) {
    contributionsCollection {
      contributionCalendar {
        totalContributions
        weeks {
          contributionDays {
            color
            contributionCount
            date
          }
        }
      }
    }
  }
}`

	variables := map[string]interface{}{
		"userName": config.USER_NAME,
	}

	reqBody := types.GraphQLRequest{
		Query:     query,
		Variables: variables,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatalf("Failed to marshal request: %v", err)
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.TOKEN)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to execute request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response: %v", err)
	}

	var graphqlResp types.GraphQLResponse
	if err := json.Unmarshal(body, &graphqlResp); err != nil {
		log.Fatalf("Failed to unmarshal response: %v", err)
	}

	if len(graphqlResp.Errors) > 0 {
		log.Fatalf("GraphQL errors: %v", graphqlResp.Errors)
	}

	var githubCalendarData types.GithubCalendarData

	err = json.Unmarshal(graphqlResp.Data, &githubCalendarData)

	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return &[]types.Week{}
	}

	weeks := githubCalendarData.User.ContributionsCollection.ContributionCalendar.Weeks

	return &weeks
}
