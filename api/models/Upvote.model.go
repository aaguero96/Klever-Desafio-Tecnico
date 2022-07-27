package models

type Upvote struct {
	Id        string `bson:"_id,omitempty"`
	ServiceId string `bson:"serviceId,omitempty"`
	UserId    string `bson:"userId,omitempty"`
	Vote      string `bson:"vote,omitempty"`
	Comment   string `bson:"comment,omitempty"`
}
