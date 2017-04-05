package main

import (
	"github.com/zyfdegh/crushon/env"
	"github.com/zyfdegh/crushon/svc"
	"log"
)

func main() {
	nickname, err := env.NICKNAME.ToString()
	if err != nil {
		if err == env.ErrNotSet {
			log.Printf("please set env NICKNAME, error: %v\n", err)
		}
		log.Printf("unknown error: %v\n", err)
		return
	}

	log.Printf("nickname: %s\n", nickname)

	userID, err := svc.GetUserID(nickname)
	if err != nil {
		log.Printf("get user ID error: %v\n", err)
		return
	}
	log.Printf("user ID: %d\n", userID)
	return
}
