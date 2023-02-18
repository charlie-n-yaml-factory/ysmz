package services

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/charlie-n-yaml-factory/ysmz/ysmz-backend/config"
	"github.com/charlie-n-yaml-factory/ysmz/ysmz-backend/databases"
)

type OAuthGoogleTokens struct {
	AccessToken string
	IDToken     string
}

var StatePrefix = "OAuthGoogleState_"

func GetOAuthGoogleTokens(code, state string, conf *config.ConfigStruct, redisDB databases.RedisInterface) (*OAuthGoogleTokens, error) {
	// check if state is valid
	savedStateValue, err := redisDB.Get(StatePrefix + state[:len(state)/2])
	if savedStateValue != state[len(state)/2:] {
		return nil, errors.New("google oauth invalid state")
	}

	const rootURl = "https://oauth2.googleapis.com/token"

	values := url.Values{}
	values.Add("grant_type", "authorization_code")
	values.Add("code", code)
	values.Add("client_id", conf.OAuthGoogleClientID)
	values.Add("client_secret", conf.OAuthGoogleClientSecret)
	values.Add("redirect_uri", conf.OAuthGoogleRedirectURL)

	query := values.Encode()

	req, err := http.NewRequest("POST", rootURl, bytes.NewBufferString(query))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := http.Client{
		Timeout: time.Second * 5,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("could not retrieve google oauth tokens")
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var GoogleOauthTokenRes map[string]interface{}

	if err := json.Unmarshal(resBody, &GoogleOauthTokenRes); err != nil {
		return nil, err
	}

	tokenBody := &OAuthGoogleTokens{
		AccessToken: GoogleOauthTokenRes["access_token"].(string),
		IDToken:     GoogleOauthTokenRes["id_token"].(string),
	}

	return tokenBody, nil
}

func generateRandomState() string {
	// Generate a random byte sequence with 1024 bytes
	randomBytes := make([]byte, 1024)
	if _, err := rand.Read(randomBytes); err != nil {
		panic(err)
	}

	// Compute SHA256 hash of the random byte sequence
	hash := sha256.Sum256(randomBytes)

	// Convert hash to hexadecimal string
	hashString := hex.EncodeToString(hash[:])

	return hashString
}

func SaveOAuthGoogleState(redisDB databases.RedisInterface) (string, error) {
	state := generateRandomState()

	key := StatePrefix + state[:len(state)/2]
	value := state[len(state)/2:]

	redisDB.Set(key, value, 300*time.Second)

	return state, nil
}

func GetOAuthGoogleUserInfo(accessToken string) (map[string]interface{}, error) {
	const rootURl = "https://www.googleapis.com/oauth2/v2/userinfo"

	req, err := http.NewRequest("GET", rootURl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := http.Client{
		Timeout: time.Second * 5,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("could not retrieve google oauth user info")
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var userInfo map[string]interface{}

	if err := json.Unmarshal(resBody, &userInfo); err != nil {
		return nil, err
	}

	return userInfo, nil
}
