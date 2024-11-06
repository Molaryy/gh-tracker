package main

import (
	"runtime/debug"
	"time"
)

type GithubEvent struct {
    action string
    description string
    author string
    title string
    createdTime time.Time
    branchName string
    commitMessage string
}

func CreateGithubEvent(body map[string]interface{}, action string) *GithubEvent{
    description := ""
    author := ""
    title := ""
    var createdTime time.Time
    branchName := ""
    commitMessage := ""

    switch action {
        case "push":
            action = "Pushed"
            
    }

    return &GithubEvent {
        action: action,
        description: description,
        author: author,
        title: title,
        createdTime: createdTime,
        branchName: branchName,
        commitMessage: commitMessage,
    }
}
