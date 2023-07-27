package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func fetchNews(apiKey, category, language string) ([]byte, error) {
	url := fmt.Sprintf(
		"https://newsapi.org/v2/everything?q=%s&language=%s",
		category,
		language,
	)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Api-Key", apiKey)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("status code : %d", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	response.Body.Close()

	return body, nil
}

func printNews(articles []byte) {
	var data struct {
		Articles []struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		} `json:"articles"`
	}

	err := json.Unmarshal(articles, &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, article := range data.Articles {
		fmt.Println(article.Title)
		fmt.Println(article.Description)
		fmt.Println("======================================================================================")
	}
}

func main() {
	apiKey := ""

	fmt.Println("\n\nEnter Topic of News: ")
	var topic string
	fmt.Scanln(&topic)

	//fetch news
	tech_news, err := fetchNews(apiKey, topic, "en")
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
	}
	fmt.Printf("\n\n %s News:\n", topic)
	printNews(tech_news)

	//fetch Python news
	// py_news, err := fetchNews(apiKey, "python", "en")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("\n\n\n Python News:")
	// printNews(py_news)
}
