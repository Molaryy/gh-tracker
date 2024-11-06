package main

import (
	"io"
	"log"
	"net/http"
	"encoding/json"
)

const (
    port = "8080"
)

func jsonToMap(body string) map[string]interface{}  {
    result := make(map[string]interface{})
	json.Unmarshal([]byte(body), &result)
	return result
}

func handleGithubSubscribedEvent(w http.ResponseWriter, r *http.Request) {
    bytes, err := io.ReadAll(r.Body)
    
    if err != nil {
        log.Fatal("Error at reading body: ", err)
        return
    }
    body := string(bytes)
    bodyMap := jsonToMap(body)
    gh := CreateGithubEvent(bodyMap)
    SendMessage(gh)
}

func main() {
    http.HandleFunc("/github/event", handleGithubSubscribedEvent)
    http.ListenAndServe(":" + port, nil)
}
