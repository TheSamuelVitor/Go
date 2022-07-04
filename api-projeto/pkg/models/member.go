package models

type Member struct {
	Id_member   string `json:"id_member"`
	Name_member string `json:"name_member"`
	Age         string `json:"age"`
	Team        string `json:"name_team"`
}