package models

import "encoding/json"

type Project struct {
	Project_id        string          `json:"project_id"`
	Title             string          `json:"title"`
	Description       string          `json:"description"`
	Date_completed    string          `json:"date_completed"`
	Technologies_used json.RawMessage `json:"technologies_used"`
	Image_url         string          `json:"image_url"`
	Link              string          `json:"link"`
	User_id           string          `json:"user_id"`
}
