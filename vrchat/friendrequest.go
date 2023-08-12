package vrchat

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"time"
)

type FriendRequestResponse struct {
	ID             string    `json:"id"`
	SenderUserID   string    `json:"senderUserId"`
	SenderUsername string    `json:"senderUsername"`
	Type           string    `json:"type"`
	Message        string    `json:"message"`
	Details        string    `json:"details"`
	Seen           bool      `json:"seen"`
	CreatedAt      time.Time `json:"created_at"`
}

// FriendRequest is a http Request to send a friend request to another user, and takes a userID as a parameter
// also takes an auth token as a cookie
// returns a FriendRequestResponse
func FriendRequest(userid, token, tfa string, proxyClient *http.Client) (string, error) {

	url := "https://vrchat.com/api/1/user/" + userid + "/friendRequest"
	method := "POST"

	payload := "{\"params\":{\"userId\":\" " + userid + " \"}}"

	req, err := http.NewRequest(method, url, strings.NewReader(payload))
	if err != nil {
		return "", errors.New("[friendrequest] Error creating request")
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/112.0")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-GPC", "1")
	req.Header.Add("Cookie", "twoFactorAuth="+tfa+"; auth="+token)

	resp, err := proxyClient.Do(req)
	if err != nil {
		return "", errors.New("[friendrequest] Error sending request")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("[friendrequest] Error reading response")
	}

	return string(body), nil
}
