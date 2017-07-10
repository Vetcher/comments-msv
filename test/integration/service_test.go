package service_test

import (
	"net/http"
	"testing"

	"bytes"
	"encoding/json"

	"io/ioutil"
)

type Comment struct {
	ID       uint   `json:"id,omitempty"`
	Text     string `json:"text"`
	AuthorID uint   `json:"author_id"`
}

type Response struct {
	Data interface{} `json:"data"`
	Err  string      `json:"error"`
}

func TestPostComment(t *testing.T) {
	client := &http.Client{}
	const CommentText = "Testing Comment"
	comment := Comment{
		Text:     CommentText,
		AuthorID: 1,
	}
	b, err := json.Marshal(&comment)
	if err != nil {
		t.Fatalf("Marshal to JSON error: %v", err)
	}
	buf := bytes.NewBuffer(b)
	req, err := http.NewRequest("GET", "http://localhost:8081/comment/post", buf)
	if err != nil {
		t.Fatalf("can't make request: %v", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("can't send request: %v", err)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("reading response Body error: %v", err)
	}
	var c Response
	err = json.Unmarshal(respBody, &c)
	if err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	if !(c.Data.(Comment).AuthorID == 1 && c.Data.(Comment).Text == CommentText) {
		t.Fatalf("Have: %v\nExpected: %v", c, Response{
			Data: Comment{
				AuthorID: 1,
				Text:     CommentText,
			},
		})
	}
}

func TestGetComment(t *testing.T) {
	client := &http.Client{}
	const CommentText = "Testing Comment"
	comment := struct {
		ID uint `json:"id"`
	}{ID: 1}
	b, err := json.Marshal(&comment)
	if err != nil {
		t.Fatalf("Marshal to JSON error: %v", err)
	}
	buf := bytes.NewBuffer(b)
	req, err := http.NewRequest("GET", "http://localhost:8081/comment/get", buf)
	if err != nil {
		t.Fatalf("can't make request: %v", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("can't send request: %v", err)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("reading response Body error: %v", err)
	}
	var c struct {
		Data Comment `json:"data"`
		Err  string  `json:"error"`
	}
	err = json.Unmarshal(respBody, &c)
	if err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	/*if !(c.Data.AuthorID == 1 && c.Data.Text == CommentText && c.Data.ID == 1) {
		t.Fatalf("Have: %v\nExpected: %v", c, struct {
			Data Comment `json:"data"`
		}{Data: Comment{Text: CommentText, AuthorID: 1, ID: 1}})
	}*/

	if !(c.Data.AuthorID == 1 && c.Data.Text == CommentText && c.Data.ID == 1) {
		t.Fatalf("Have: %v\nExpected: %v", c, Response{
			Data: Comment{
				AuthorID: 1,
				Text:     CommentText,
				ID:       1,
			},
		})
	}
}

func TestDeleteComment(t *testing.T) {
	client := &http.Client{}
	comment := struct {
		ID uint `json:"id"`
	}{ID: 1}
	b, err := json.Marshal(&comment)
	if err != nil {
		t.Fatalf("Marshal to JSON error: %v", err)
	}
	buf := bytes.NewBuffer(b)
	req, err := http.NewRequest("GET", "http://localhost:8081/comment/del", buf)
	if err != nil {
		t.Fatalf("can't make request: %v", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("can't send request: %v", err)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("reading response Body error: %v", err)
	}
	var c struct {
		Data bool `json:"data"`
	}
	err = json.Unmarshal(respBody, &c)
	if err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	if !(c.Data) {
		t.Fatalf("Have: %v\nExpected: %v", c, struct {
			Data bool
		}{true})
	}
}

func TestDeleteCommentAgain(t *testing.T) {
	client := &http.Client{}
	comment := struct {
		ID uint `json:"id"`
	}{ID: 1}
	b, err := json.Marshal(&comment)
	if err != nil {
		t.Fatalf("Marshal to JSON error: %v", err)
	}
	buf := bytes.NewBuffer(b)
	req, err := http.NewRequest("GET", "http://localhost:8081/comment/del", buf)
	if err != nil {
		t.Fatalf("can't make request: %v", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("can't send request: %v", err)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("reading response Body error: %v", err)
	}
	var c struct {
		Data bool   `json:"data"`
		Err  string `json:"error,omitempty"`
	}
	err = json.Unmarshal(respBody, &c)
	if err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	if c.Data {
		t.Fatalf("Have: %v\nExpected: %v", c, struct {
			Data bool
			Err  string
		}{false, ""})
	}
}
