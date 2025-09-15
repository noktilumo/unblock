package roblox

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"sync"
)

const (
	GETBLOCKED_API       = "https://apis.roblox.com/user-blocking-api/v1/users/get-blocked-users?cursor=&count=50"
	UNBLOCK_API          = "https://apis.roblox.com/user-blocking-api/v1/users/%v/unblock-user"
	ENV_COOKIE           = "COOKIE"
	ENV_X_CSRF_TOKEN     = "X_CSRF_TOKEN"
	GETBLOCKED_MAX_COUNT = 50
)

type ResponseData struct {
	BlockedUserIds []int `json:"blockedUserIds"`
}

func ReadEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	return os.Getenv(key)
}

func FetchAllBlockedUserIds() (blockedUsersId []int) {
	envCookie := ReadEnvVariable(ENV_COOKIE)
	client := &http.Client{}
	var allBlocked []int
	cursor := ""

	for {
		url := fmt.Sprintf("https://apis.roblox.com/user-blocking-api/v1/users/get-blocked-users?cursor=%s&count=50", cursor)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("Cookie", envCookie)
		req.Header.Set("Accept", "application/json, text/plain, */*")

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Fatalf("Unexpected status code: %d", resp.StatusCode)
		}

		var result struct {
			Data struct {
				BlockedUserIds []int  `json:"blockedUserIds"`
				Cursor         string `json:"cursor"`
			} `json:"data"`
		}

		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		allBlocked = append(allBlocked, result.Data.BlockedUserIds...)

		// Stop if fewer than 50 returned or no cursor
		if len(result.Data.BlockedUserIds) < GETBLOCKED_MAX_COUNT || result.Data.Cursor == "" {
			break
		}

		cursor = result.Data.Cursor
	}

	return allBlocked
}

func UnblockAllBlockedUsers() {
	cookie := ReadEnvVariable(ENV_COOKIE)
	csrf := ReadEnvVariable(ENV_X_CSRF_TOKEN)
	client := &http.Client{}

	blocked := FetchAllBlockedUserIds()
	var wg sync.WaitGroup

	for _, id := range blocked {
		wg.Add(1)
		go UnblockAsync(id, &wg, client, cookie, csrf)
	}

	wg.Wait()
	log.Println("Finished unblocking all users.")
}

func UnblockAsync(id int, wg *sync.WaitGroup, client *http.Client, cookie, csrf string) {
	defer wg.Done()

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf(UNBLOCK_API, id), nil)
	if err != nil {
		log.Printf("Error creating request for ID %d: %v\n", id, err)
		return
	}

	req.Header.Set("Cookie", cookie)
	req.Header.Set("x-csrf-token", csrf)

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request for ID %d: %v\n", id, err)
		return
	}
	defer resp.Body.Close()

	log.Printf("Unblocked user %d — Status: %d\n", id, resp.StatusCode)
}

func Unblock(id int) (statusCode int) {
	envCookie := ReadEnvVariable(ENV_COOKIE)
	EnvXcsrfToken := ReadEnvVariable(ENV_X_CSRF_TOKEN)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf(UNBLOCK_API, id), nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Cookie", envCookie)
	req.Header.Set("x-csrf-token", EnvXcsrfToken)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	return resp.StatusCode
}
