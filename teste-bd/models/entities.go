package models

type ToDo struct {
	id          int64  `json: "id"`
	title       string `json: "tile"`
	description string `json: "description"`
	done        bool   `json: "done"`
}
