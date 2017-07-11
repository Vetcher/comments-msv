package service_test

import (
	"net/http"
	"testing"

	"bytes"
	"encoding/json"

	"fmt"
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

func HTTPGet(body interface{}, dest string) ([]byte, error) {
	client := &http.Client{}
	b, err := json.Marshal(&body)
	if err != nil {
		return nil, fmt.Errorf("Marshal to JSON error: %v", err)
	}
	buf := bytes.NewBuffer(b)
	req, err := http.NewRequest("GET", dest, buf)
	if err != nil {
		return nil, fmt.Errorf("can't make request: %v", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("can't send request: %v", err)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response Body error: %v", err)
	}
	return respBody, nil
}

func TestPostComment(t *testing.T) {
	comment := Comment{
		Text:     CommentText,
		AuthorID: 1,
	}
	respBody, err := HTTPGet(comment, "http://localhost:8081/comment/post")
	if err != nil {
		t.Fatal(err)
	}
	var c ResponseComment
	err = json.Unmarshal(respBody, &c)
	if err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	if !(c.Data.AuthorID == 1 && c.Data.Text == CommentText) {
		t.Fatalf("Have: %v\nExpected: %v", c, ResponseComment{
			Data: &Comment{
				AuthorID: 1,
				Text:     CommentText,
			},
		})
	}
}

func TestGetComment(t *testing.T) {
	id := struct {
		ID uint `json:"id"`
	}{ID: 1}
	respBody, err := HTTPGet(id, "http://localhost:8081/comment/get")
	if err != nil {
		t.Fatal(err)
	}
	var c ResponseComment
	err = json.Unmarshal(respBody, &c)
	if err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}

	if !(c.Data.AuthorID == 1 && c.Data.Text == CommentText && c.Data.ID == 1) {
		t.Fatalf("Have: %v\nExpected: %v", c, ResponseComment{
			Data: &Comment{
				AuthorID: 1,
				Text:     CommentText,
				ID:       1,
			},
		})
	}
}

func TestGetCommentsForSpecificAuthor(t *testing.T) {
	id := struct {
		ID uint `json:"id"`
	}{ID: 1}
	respBody, err := HTTPGet(id, "http://localhost:8081/comments/author")
	if err != nil {
		t.Fatal(err)
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

func TestDeleteComment(t *testing.T) {
	id := struct {
		ID uint `json:"id"`
	}{ID: 1}
	respBody, err := HTTPGet(id, "http://localhost:8081/comment/del")
	if err != nil {
		t.Fatal(err)
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
	respBody, err = HTTPGet(id, "http://localhost:8081/comment/get")
	if err != nil {
		t.Fatal(err)
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
