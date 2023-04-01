package token

import (
	"fmt"
    "log"
	"github.com/google/uuid"
	"sync"
	"time"
)

var tokenMap map[string]int64 = make(map[string]int64)
var mutex sync.Mutex

func GetToken() string {
	mutex.Lock()
	defer mutex.Unlock()
	token := fmt.Sprintf("%s", uuid.New())
	timestamp := time.Now().Local().Unix()
	tokenMap[token] = timestamp
	return token
}

func UpdateTokenTimestamp(token string)  {
	mutex.Lock()
	defer mutex.Unlock()
	if _, ok := tokenMap[token]; ok {
        tokenMap[token] = time.Now().Local().Unix()
	}
}

func DelToken(token string) {
	mutex.Lock()
	defer mutex.Unlock()
	delete(tokenMap, token)
}

func IsTokenTimeout(token string, duration int64) bool {
	mutex.Lock()
	defer mutex.Unlock()
	if timestamp, ok := tokenMap[token]; ok {
		return time.Now().Local().Unix()-timestamp > duration
	}
	return true
}

func ClearTimeoutToken(duration int64) {
	mutex.Lock()
	defer mutex.Unlock()
	timeoutToken := []string{}
	for k, v := range tokenMap {
		if time.Now().Local().Unix()-v > duration {
			timeoutToken = append(timeoutToken, k)
		}
	}

	for _, token := range timeoutToken {
        log.Println("delete timeout token: ", token)
		delete(tokenMap, token)
	}
}
