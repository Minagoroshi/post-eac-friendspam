package vrchat

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// LoginResponse is the response from the login request
type LoginResponse struct {
	ID                 string   `json:"id"`
	Username           string   `json:"username"`
	DisplayName        string   `json:"displayName"`
	UserIcon           string   `json:"userIcon"`
	Bio                string   `json:"bio"`
	BioLinks           []string `json:"bioLinks"`
	ProfilePicOverride string   `json:"profilePicOverride"`
	StatusDescription  string   `json:"statusDescription"`
	PastDisplayNames   []struct {
		DisplayName string    `json:"displayName"`
		UpdatedAt   time.Time `json:"updated_at"`
	} `json:"pastDisplayNames"`
	HasEmail                       bool     `json:"hasEmail"`
	HasPendingEmail                bool     `json:"hasPendingEmail"`
	ObfuscatedEmail                string   `json:"obfuscatedEmail"`
	ObfuscatedPendingEmail         string   `json:"obfuscatedPendingEmail"`
	EmailVerified                  bool     `json:"emailVerified"`
	HasBirthday                    bool     `json:"hasBirthday"`
	Unsubscribe                    bool     `json:"unsubscribe"`
	StatusHistory                  []string `json:"statusHistory"`
	StatusFirstTime                bool     `json:"statusFirstTime"`
	Friends                        []string `json:"friends"`
	FriendGroupNames               []string `json:"friendGroupNames"`
	CurrentAvatarImageURL          string   `json:"currentAvatarImageUrl"`
	CurrentAvatarThumbnailImageURL string   `json:"currentAvatarThumbnailImageUrl"`
	CurrentAvatar                  string   `json:"currentAvatar"`
	CurrentAvatarAssetURL          string   `json:"currentAvatarAssetUrl"`
	AcceptedTOSVersion             float64  `json:"acceptedTOSVersion"`
	SteamID                        string   `json:"steamId"`
	SteamDetails                   struct {
	} `json:"steamDetails"`
	OculusID              string    `json:"oculusId"`
	HasLoggedInFromClient bool      `json:"hasLoggedInFromClient"`
	HomeLocation          string    `json:"homeLocation"`
	TwoFactorAuthEnabled  bool      `json:"twoFactorAuthEnabled"`
	State                 string    `json:"state"`
	Tags                  []string  `json:"tags"`
	DeveloperType         string    `json:"developerType"`
	LastLogin             time.Time `json:"last_login"`
	LastPlatform          string    `json:"last_platform"`
	AllowAvatarCopying    bool      `json:"allowAvatarCopying"`
	Status                string    `json:"status"`
	DateJoined            string    `json:"date_joined"`
	IsFriend              bool      `json:"isFriend"`
	FriendKey             string    `json:"friendKey"`
	FallbackAvatar        string    `json:"fallbackAvatar"`
	AccountDeletionDate   string    `json:"accountDeletionDate"`
	OnlineFriends         []string  `json:"onlineFriends"`
	ActiveFriends         []string  `json:"activeFriends"`
	OfflineFriends        []string  `json:"offlineFriends"`
}

func Login(account string, proxyClient *http.Client) (string, error) {

	username := strings.Split(account, ":")[0]
	password := strings.Split(account, ":")[1]
	code := strings.Split(account, ":")[2]

	//encode the account into base64
	account = base64.StdEncoding.EncodeToString([]byte(username + ":" + password))

	url := "https://vrchat.com/api/1/auth/user"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", errors.New("[login] failed to create request")
	}

	req.Header.Add("Authorization", "Basic "+account)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/112.0")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-GPC", "1")
	req.Header.Add("Cookie", "twoFactorAuth="+code)

	resp, err := proxyClient.Do(req)
	if err != nil {
		return "", errors.New("[login] failed to send request" + err.Error())
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("[login] failed to read response")
	}

	if !strings.Contains(string(body), "currentAvatar") {
		fmt.Println(string(body))
		return "", errors.New("[login] failed to login")
	}

	authToken := resp.Cookies()[0].Value

	return authToken, nil
}
