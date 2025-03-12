/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	// "time"

	"github.com/go-resty/resty/v2"
	"github.com/phenriqx/github-cli/cmd/helpers"
	"github.com/spf13/cobra"
)

func GetUserActivity(username string) {
	github_token, _ := helpers.GetGithubToken()
	client := resty.New()
	client.SetHeader("Authorization", "token "+github_token)

	request_url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	response, err := client.R().Get(request_url)
	if err != nil {
		fmt.Println("Error getting user activity: ", err)
		return
	}

	var activity []helpers.UserActivity
	err = json.Unmarshal(response.Body(), &activity)
	if err != nil {
		fmt.Println("Error unmarshalling response: ", err)
		return
	}

	for _, data := range activity[:5] {
		formattedDate := helpers.ParseDate(data.CreatedAt)
		if data.Type == "PushEvent" {
			fmt.Printf("%s pushed %d commit(s) to %s on %s\n", username, data.Payload.Size, data.Repo.Name, formattedDate)
		}

	}
}

// activityCmd represents the activity command
var activityCmd = &cobra.Command{
	Use:   "activity <username>",
	Short: "Get activity information from a certain user.",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		GetUserActivity(username)
	},
}

func init() {
	rootCmd.AddCommand(activityCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// activityCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// activityCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
