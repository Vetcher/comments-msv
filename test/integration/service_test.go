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

type ResponseComment struct {
	Data *Comment `json:"data"`
	Err  string   `json:"error"`
}

const CommentText = "Testing Comment"

var testCommentID uint = 5

func TestHTTPPostComment(t *testing.T) {
	comment := Comment{
		Text:     CommentText,
		AuthorID: 1,
	}
	dest := "http://localhost:8081/comment/post"
	b, err := json.Marshal(&comment)
	if err != nil {
		t.Fatalf("Marshal to JSON error: %v", err)
	}
	buf := bytes.NewBuffer(b)
	resp, err := http.Post(dest, "application/json", buf)
	if err != nil {
		t.Fatalf("can't do request: %v", err)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("reading response Body error: %v", err)
	}
	var c ResponseComment
	err = json.Unmarshal(respBody, &c)
	if err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	if c.Err != "" {
		t.Fatal(c.Err)
	}
	if !(c.Data.AuthorID == 1 && c.Data.Text == CommentText) {
		t.Fatalf("\nHave: %v\nExpected: %v", c, ResponseComment{
			Data: &Comment{
				AuthorID: 1,
				Text:     CommentText,
			},
		})
	}
	testCommentID = c.Data.ID
}

func TestHTTPGetComment(t *testing.T) {
	id := struct {
		ID uint `json:"id"`
	}{ID: testCommentID}
	dest := "http://localhost:8081/comment/get"
	b, err := json.Marshal(&id)
	if err != nil {
		t.Fatalf("Marshal to JSON error: %v", err)
	}
	buf := bytes.NewBuffer(b)
	resp, err := http.Post(dest, "application/json", buf)
	if err != nil {
		t.Fatalf("can't do request: %v", err)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("reading response Body error: %v", err)
	}
	var c ResponseComment
	err = json.Unmarshal(respBody, &c)
	if err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	if c.Err != "" {
		t.Fatal(c.Err)
	}
	if !(c.Data.AuthorID == 1 && c.Data.Text == CommentText && c.Data.ID == testCommentID) {
		t.Fatalf("Have: %v\nExpected: %v", c, ResponseComment{
			Data: &Comment{
				AuthorID: 1,
				Text:     CommentText,
				ID:       testCommentID,
			},
		})
	}
}

func TestHTTPGetCommentsForSpecificAuthor(t *testing.T) {
	id := struct {
		ID uint `json:"id"`
	}{ID: 1}
	dest := "http://localhost:8081/comments/author"
	b, err := json.Marshal(&id)
	if err != nil {
		t.Fatalf("Marshal to JSON error: %v", err)
	}
	buf := bytes.NewBuffer(b)
	resp, err := http.Post(dest, "application/json", buf)
	if err != nil {
		t.Fatalf("can't do request: %v", err)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("reading response Body error: %v", err)
	}
	var comments struct {
		Data []*Comment `json:"data"`
		Err  string     `json:"error"`
	}
	err = json.Unmarshal(respBody, &comments)
	if err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}

	if len(comments.Data) == 0 || comments.Err != "" {
		t.Fatalf("Have: %v\n`Data` should not be empty and `Err` should be empty", comments)
	}
}

func TestHTTPDeleteComment(t *testing.T) {
	id := struct {
		ID uint `json:"id"`
	}{ID: testCommentID}
	dest := "http://localhost:8081/comment/del"
	b, err := json.Marshal(&id)
	if err != nil {
		t.Fatalf("Marshal to JSON error: %v", err)
	}
	buf := bytes.NewBuffer(b)
	resp, err := http.Post(dest, "application/json", buf)
	if err != nil {
		t.Fatalf("can't do request: %v", err)
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
	// try to Get
	dest = "http://localhost:8081/comment/get"
	b, err = json.Marshal(&id)
	if err != nil {
		t.Fatalf("Marshal to JSON error: %v", err)
	}
	buf = bytes.NewBuffer(b)
	resp, err = http.Post(dest, "application/json", buf)
	if err != nil {
		t.Fatalf("can't do request: %v", err)
	}
	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("reading response Body error: %v", err)
	}
	var respC ResponseComment
	err = json.Unmarshal(respBody, &respC)
	if err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	if respC.Err == "database error: record not found" && respC.Data != nil {
		t.Fatalf("Have: %v\nExpected: %v", c, ResponseComment{
			Data: nil,
			Err:  "database error: record not found",
		})
	}
}
