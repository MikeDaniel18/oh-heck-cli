package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"oh-heck/configs"
	"strconv"
	"strings"
)

func AskBashQuestionApi(question string) (*Completion, *ErrorResponse) {

	form := url.Values{}
	form.Add("question", question)

	req, err := http.NewRequest("POST", fmt.Sprintf("%v/bash/question", configs.GetApiURL()), strings.NewReader(form.Encode()))

	var completion *Completion
	var errorResponse *ErrorResponse

	if err != nil {
		fmt.Println("Something went wrong, please ensure you have updated oh-heck and have a valid API Key")
		log.Fatalln(err)
		return completion, errorResponse
	}

	cfg := configs.ReadConfig()

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", cfg.ApiKey))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("Something went wrong, please ensure you have updated oh-heck and have a valid API Key")
		log.Fatalln(err)
		return completion, errorResponse
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Something went wrong, please ensure you have updated oh-heck and have a valid API Key")
		log.Fatalln(err)
		return completion, errorResponse
	}

	statusCode, err := statusCode(body)

	if err != nil {
		fmt.Println("Something went wrong, please ensure you have updated oh-heck and have a valid API Key")
		log.Fatalln(err)
		return completion, errorResponse
	}

	if statusCode >= 200 && statusCode < 300 {
		// Parse the successful response
		var completionResponse CompletionResponse
		json.Unmarshal(body, &completionResponse)
		completion = &completionResponse.Completion
	} else {
		json.Unmarshal(body, &errorResponse)
	}

	return completion, errorResponse
}

func SetQuestionResponse(completion Completion, accepted bool) {
	form := url.Values{}

	var formAccepted string

	if accepted {
		formAccepted = "true"
	} else {
		formAccepted = "false"
	}
	form.Add("accepted", formAccepted)

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%v/bash/question/%v", configs.GetApiURL(), completion.Id), strings.NewReader(form.Encode()))

	if err != nil {
		return
	}

	cfg := configs.ReadConfig()

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", cfg.ApiKey))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return
	}

	statusCode, err := statusCode(body)

	if err != nil {
		return
	}

	if statusCode >= 200 && statusCode < 300 {
		// Parse the successful response
	} else {
	}
}

func statusCode(body []byte) (int, error) {
	var result map[string]interface{}
	json.Unmarshal(body, &result)

	if status, ok := result["status"]; ok {
		statusCode, err := strconv.Atoi(fmt.Sprintf("%v", status))
		if err != nil {
			return 0, err
		} else {
			return statusCode, nil
		}
	}

	return 0, errors.New("Something went wrong")
}
