package jira

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/viper"
)

const (
	keyToken  = "jira.token"
	keyDomain = "jira.domain"
)

type Tracker struct{}

func (j Tracker) SetUpTracker() error {
	apiUrl := viper.GetString(keyDomain)
	if apiUrl == "" {
		prompt := &survey.Input{
			Message: "Please enter your Jira API URL:",
		}
		err := survey.AskOne(prompt, &apiUrl)
		if err != nil {
			return err
		}
		viper.Set(keyDomain, apiUrl)
	}
	token := viper.GetString(keyToken)
	if token == "" {
		prompt := &survey.Password{
			Message: "Please enter your Jira API token:",
		}
		err := survey.AskOne(prompt, &token)
		if err != nil {
			return err
		}
		viper.Set(keyToken, token)
	}
	return nil
}
