package vrchat

import (
	"errors"
	"io"
	"net/http"
)

func ChangeAvatar(avatarid, token, tfa string, proxyClient *http.Client) (string, error) {

	// proxy client

	url := "https://vrchat.com/api/1/avatars/" + avatarid + "/select"
	method := "PUT"

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", errors.New("[avatar] Error creating request: " + err.Error())
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
		return "", errors.New("[avatar] Error sending request")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("[avatar] Error reading response")
	}

	return string(body), nil
}
