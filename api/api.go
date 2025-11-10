package api

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	client_url   string
	client_token string
}

func NewClient(input_url, input_token string) *Client {
	return &Client{
		client_url:   input_url,
		client_token: input_token,
	}
}

func (c Client) post(ctx context.Context, action string, params map[string]interface{}) string {
	formData := url.Values{
		"api_key": {c.client_token},
		"action":  {action},
	}

	for key, value := range params {
		// Убедитесь, что мы не перезаписываем action
		if key == "action" {
			continue // пропускаем, если кто-то случайно передал action в params
		}

		switch v := value.(type) {
		case string:
			formData.Add(key, v)
		case int:
			formData.Add(key, strconv.Itoa(v))
		case int64:
			formData.Add(key, strconv.FormatInt(v, 10))
		case float64:
			formData.Add(key, strconv.FormatFloat(v, 'f', -1, 64))
		case bool:
			formData.Add(key, strconv.FormatBool(v))
		default:
			fmt.Printf("Привожу к строке значение: %v так как не понял что это за тип", v)
			formData.Add(key, fmt.Sprintf("%v", v))
		}
	}

	resp, err := http.PostForm(c.client_url, formData)
	if err != nil {
		log.Fatalf("Ошибка %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Ошибка %v", err)
	}
	bodyString := string(body)
	return bodyString
}
