package models

type Upvote struct {
	ID        string `bson:"_id,omitempty"`
	ServiceID string `bson:"serviceId,omitempty"`
	UserID    string `bson:"userId,omitempty"`
	Vote      string `bson:"vote,omitempty"`
	Comment   string `bson:"comment,omitempty"`
}
