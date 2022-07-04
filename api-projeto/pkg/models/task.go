package models

type Task struct {
	Id_task   string `json:"id_task"`
	Name_task string `json:"name_task"`
	Time      string `json:"id_time"`
	Id_member string `json:"id_member"`
}