package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Post struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type Comment struct {
	ID     int    `json:"id"`
	Body   string `json:"body"`
	PostID int    `json:"postId"`
}

type Profile struct {
	Name string `json:"name"`
}

func main() {

	httpClient := http.DefaultClient

	posts, err := getPosts(context.TODO(), httpClient)

	if err != nil {
		panic(err)
	}

	fmt.Println(posts)

	/*post, err := getPost(context.TODO(), httpClient, "1")

	if err != nil {
		panic(err)
	}

	jsonPost, err := json.Marshal(post)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonPost))*/
}

func getPosts(ctx context.Context, httpClient *http.Client) (string, error) {

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:3000/posts", nil)

	if err != nil {
		return "", err
	}

	response, err := httpClient.Do(request)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		bytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return "", err
		}
		return "", errors.New(fmt.Sprintf("Status: [%v], Body: %v", response.Status, string(bytes)))
	}

	responseBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	return string(responseBytes), nil
}

func getPost(ctx context.Context, httpClient *http.Client, id string) (*Post, error) {

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("http://localhost:3000/posts/%s", id), nil)

	if err != nil {
		return nil, err
	}

	response, err := httpClient.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		bytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(fmt.Sprintf("Status: [%v], Body: %v", response.Status, string(bytes)))
	}

	var result Post
	err = json.NewDecoder(response.Body).Decode(&result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
