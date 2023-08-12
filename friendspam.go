package main

import (
	"fmt"
	"github.com/gookit/color"
	SyzProxy "github.com/minagoroshi/syzproxy"
	"log"
	"net/http"
	"post-eac-friendspam/vrchat"
	"regexp"
	"strings"
	"time"
)

var (
	UidRegex      = regexp.MustCompile("usr_[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}")
	AvatarIdRegex = regexp.MustCompile("avtr_[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}")
)

func friendSpam(uid string, avatarid string, accounts []string, useProxies, useAvatar bool) error {

	var client *http.Client
	PManager := &SyzProxy.ProxyManager{}

	if !useProxies {
		client = &http.Client{Timeout: 10 * time.Second}
	} else {
		num, err := PManager.LoadFromFile(ProxyPath, "http")
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Using", num, "Proxies")
	}

	for i, account := range accounts {
		if useProxies {
			transport, err := PManager.GetRandomTransport()
			if err != nil {
				log.Fatal(err)
			}
			client = SyzProxy.ClientFromTransport(transport)
		}

	login:
		var loginRetries int
		token, err := vrchat.Login(account, client)
		if err != nil {
			color.Error.Println("Error: " + err.Error())
			if loginRetries < 3 {
				color.Red.Println("Retrying...")
				loginRetries++
				goto login
			} else {
				color.Red.Println("Login Failed")
				continue
			}
		}

		if useAvatar {
			if AvatarIdRegex.MatchString(avatarid) {
			avatar:
				var avatarRetries int
				code := strings.Split(account, ":")[2]
				body, err := vrchat.ChangeAvatar(avatarid, token, code, client)
				if err != nil {
					color.Error.Println("Error: " + err.Error())
					if avatarRetries < 3 {
						color.Error.Println("Retrying...")
						avatarRetries++
						goto avatar
					} else {
						color.Error.Println("Avatar Change Failed")
						continue
					}
				}

				if strings.Contains(body, "currentAvatar") {
					color.HiGreen.Println("Avatar Changed!", i+1, "of", len(accounts))
				} else {
					color.Red.Println("Avatar Change Failed :(", i+1, "of", len(accounts))
				}
			}
		}

		var friendRequestRetries int
	friendRequest:
		code := strings.Split(account, ":")[2]
		respBody, err := vrchat.FriendRequest(uid, token, code, client)
		if err != nil {
			color.Error.Println("Error: " + err.Error())
			if friendRequestRetries < 3 {
				color.Error.Println("Retrying...")
				friendRequestRetries++
				goto friendRequest
			} else {
				color.Error.Println("Friend Request Failed")
				continue
			}
		}

		if strings.Contains(respBody, "senderUserId") {
			color.HiGreen.Println("Friend Request Sent!", i+1, "of", len(accounts))
		} else {
			color.Red.Println("Friend Request Failed :(", i+1, "of", len(accounts))
			fmt.Println(respBody)
		}
		time.Sleep(5 * time.Second)
	}
	return nil
}
