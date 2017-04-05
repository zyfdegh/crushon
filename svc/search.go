package svc

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zyfdegh/crushon/types"
	"github.com/zyfdegh/crushon/util"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	// ErrInvalidSearchResult returned when user search result is not valid
	ErrInvalidSearchResult = errors.New("invalid search result")
	// ErrUserNotFound returned if nickname dose not match(ignore case)
	// with any search result
	ErrUserNotFound = errors.New("user not found")
)

// Call NeteaseCloudMusic API and get profiles of matched users
// http://music.163.com/api/search/get?s=nickname&limit=10&type=1002
func SearchUser(nickname string, limit int) (users *[]types.UserProfile, err error) {
	var apiURL = "http://music.163.com/api/search/get"

	apiURL, _ = util.SetQuery(apiURL, "s", nickname)
	apiURL, _ = util.SetQuery(apiURL, "limit", fmt.Sprintf("%d", limit))
	apiURL, _ = util.SetQuery(apiURL, "type", "1002")

	log.Printf("POST %s \n", apiURL)
	resp, err := http.Post(apiURL, "application/json", nil)
	if err != nil {
		log.Printf("call search user api error: %v\n", err)
		return
	}

	// decode response
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("read resp body error: %v\n", err)
		return
	}

	// unmarshal as SearchUserResp
	var searchResp = types.SearchUserResp{}
	err = json.Unmarshal(data, &searchResp)
	if err != nil {
		log.Printf("unmarshal resp to SearchUserResp error: %v\n", err)

		// unmarshal as ErrResp
		var errResp = types.ErrResp{}
		if err2 := json.Unmarshal(data, &errResp); err2 != nil {
			log.Printf("unmarshal resp to ErrResp error: %v\n", err2)
			log.Printf("raw response: %s\n", string(data))
			return
		}
		log.Printf("search user failed, code: %d, msg: %s\n",
			errResp.Code, errResp.Msg)
		return
	}

	// fmt.Printf("%+v\n", searchResp)

	if searchResp.Code != 200 {
		log.Printf("invalid search result, code %d\n", searchResp.Code)
		err = ErrInvalidSearchResult
		return
	}
	users = &searchResp.Result.UserProfiles
	return
}

func GetUserID(nickname string) (userID int, err error) {
	users, err := SearchUser(nickname, 10)
	if err != nil {
		log.Printf("search user error: %v\n", err)
		return
	}

	for _, user := range *users {
		if strings.EqualFold(user.Nickname, nickname) {
			// found
			userID = user.UserID
			return
		}
	}
	err = ErrUserNotFound
	log.Println("user not found")
	return
}
