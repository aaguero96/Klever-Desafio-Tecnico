package models

type Upvote struct {
	ID        string `bson:"_id,omitempty"`
	ServiceID string `bson:"service_id,omitempty"`
	UserID    string `bson:"user_id,omitempty"`
	Vote      string `bson:"vote,omitempty"`
	Comment   string `bson:"comment,omitempty"`
}
