package teacher

import "testing"

func TestTeachersHandler(t *testing.T) {

	// 1. Empty db. Test for no content returned
	// 2. Set up data for teacher -> if same data is returned from api call
	// 3. Delete the data you set up
}

func TestCreateTeacherHandler(t *testing.T) {

	//1. Particular data is not present
}

func TestTeacherHandler(t *testing.T) {

	//1. Data cannot be created
	//2. Indicates duplicate data
	//3. Data is not inserted in the db
}

func TestDeleteTeacherHandler(t *testing.T) {

	//1. Data is not found/already deleted
}
