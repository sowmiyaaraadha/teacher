package teacher

type Teacher struct {
	Id            string        `bson:"_id" json:"id, omitempty"`
	Name          string        `bson:"name" json:"name, omitempty"`
	SchoolName    string        `bson:"school" json:"school, omitempty"`
	TakingSubject string        `bson:"sub" json:"sub, omitempty"`
	HomeAddress   []HomeAddress `bson:"home" json:"home, omitempty"`
}

type HomeAddress struct {
	Number string `bson:"num" json:"number,omitempty"`
	City   string `bson:"city" json:"city,omitempty"`
	State  string `bson:"state" json:"state,omitempty"`

}
