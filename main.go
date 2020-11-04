package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var (
	client        = &http.Client{}
	token         string
	amountRemoved int
)

type relationship struct {
	ID string `json:"id"`
}

func main() {
	fmt.Print("Enter your token:")
	fmt.Scan(&token)
	if !verifyToken() {
		fmt.Println("Enter your correct token")
		time.Sleep(time.Duration(5) * time.Second)
		os.Exit(0)
	}
	res := sendrequest("GET", "https://discordapp.com/api/v8/users/@me/relationships")
	var decoded []relationship
	json.Unmarshal([]byte(res), &decoded)
	for i := 0; i < len(decoded); i++ {
		sendrequest("DELETE", fmt.Sprintf("https://discord.com/api/v6/users/@me/relationships/%s", decoded[i].ID))
		amountRemoved++
	}
	fmt.Println(fmt.Sprintf("Removed %d friend(s)", amountRemoved))
}

func verifyToken() bool {
	res := sendrequest("GET", "https://discordapp.com/api/v8/users/@me")
	var decoded map[string]string
	json.Unmarshal([]byte(res), &decoded)
	if len(decoded["username"]) > 0 {
		return true
	}
	return false
}

func sendrequest(method string, url string) []byte {
	req, _ := http.NewRequest(method, url, nil)
	req.Header.Set("Authorization", token)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.130 Safari/537.36")
	res, _ := client.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	return body
}
