package teacher

import (
	"gopkg.in/mgo.v2"
	"log"
	"gopkg.in/mgo.v2/bson"
	"github.com/pkg/errors"
)


var Session *mgo.Session

func init() {
	var err error
	Session, err = mgo.Dial("localhost")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
}

func (t *Teacher) Insert() error {
	err := Session.DB("temp").C("Teachers").Insert(&t)
	if err != nil {
		err = errors.Wrapf(err, "failed to insert student: %+v", t)
	}
	return err
}

func GetAll() ([]Teacher, error) {
	tea := make([]Teacher, 0, 0)
	sess := Session.Copy()
	defer sess.Close()
	err := sess.DB("temp").C("Teachers").Find(nil).All(&tea)
	if err != nil {
		err = errors.Wrapf(err, "failed to get students %+v", tea)
		return tea, err
	}
	return tea, err
}

func GetOne(id string) (*Teacher, error) {
	t := new(Teacher)
	sess := Session.Copy()
	defer sess.Close()
	err := sess.DB("temp").C("Teachers" ).Find(bson.M{"_id":id}).One(t)
	if err != nil {
		err = errors.Wrapf(err, "failed to get that student %v",t)
		return nil, err
	}
	return t, err
}

func Delete(id string) error {
	sess := Session.Copy()
	return sess.DB("temp").C("Teachers").Remove(bson.M{"_id":id})
}

