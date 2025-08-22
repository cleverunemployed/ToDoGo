package models

type Task struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Comleted bool   `json:"comleted"`
}

type SchemaTask struct {
	Title string `json:"title"`
}
