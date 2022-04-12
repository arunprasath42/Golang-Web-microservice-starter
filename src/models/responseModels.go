package models

/****Add the Response JSON model structs here***/
type GetUsers struct {
	Unique_id int
	Name      string
	Email     string
	Location  string
}
