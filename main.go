package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"log"
	"post-eac-friendspam/accounts"
)

const (
	ProxyPath    = "proxies"
	AccountsPath = "accounts"
)

func main() {
	// answers struct
	answers := struct {
		UID        string
		AvatarID   string
		Accounts   []string
		UseProxies bool
		UseAvatar  bool
	}{}

	// Define generally needed questions
	var qs = []*survey.Question{
		{
			Name: "uid",
			Prompt: &survey.Input{
				Message: "UserID?",
			},
			Validate: survey.Required,
		},
		{
			Name: "useAvatar",
			Prompt: &survey.Confirm{
				Message: "Do you want to use Avatar?",
			},
		},
		{
			Name: "useProxies",
			Prompt: &survey.Confirm{
				Message: "Do you want to use Proxies?",
			},
		},
	}

	// Ask the general questions
	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	qs = []*survey.Question{} // reset general questions slice

	// if UseAvatar is true, then ask for AvatarID
	if answers.UseAvatar {
		qs = append(qs, &survey.Question{
			Name: "avatarid",
			Prompt: &survey.Input{
				Message: "AvatarID?",
			},
			Validate: survey.Required,
		})
	}

	// Ask the AvatarID question if necessary
	if len(qs) > 0 {
		err := survey.Ask(qs, &answers)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	accounts.Load(AccountsPath)

	err = friendSpam(answers.UID, answers.AvatarID, accounts.Accounts, answers.UseProxies, answers.UseAvatar)
	if err != nil {
		log.Fatalln(err)
	}
}
