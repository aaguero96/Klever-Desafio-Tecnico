package models

type Service struct {
	ID   string `bson:"_id,omitempty"`
	Name string `bson:"name,omitempty"`
	Site string `bson:"site,omitempty"`
}
