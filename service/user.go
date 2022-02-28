package service

import (
    // "fmt"
    "reflect"

    "around/backend"
    "around/constants"
    "around/model"

    "github.com/olivere/elastic/v7"
)

func CheckUser(username, password string) (bool, error) {

	query := elastic.NewBoolQuery() // select * from xxx where username = xxx and password = xxx
	query.Must(elastic.NewTermQuery("username", username))
	query.Must(elastic.NewTermQuery("password", password))

	searchResult, err := backend.ESBackend.ReadFromES(query, constants.USER_INDEX)
	if err != nil {
		return false, err
	}

	// if searchResult.TotalHits() > 0 {
	// 	return true, nil
	// }
	// return false, nil

	var utype model.User
	for _, item := range searchResult.Each(reflect.TypeOf(utype)) {
		u := item.(model.User)
		if u.Password == password {
			return true, nil
		}
	}
	return false, nil
	
}

// note: error: db error
func AddUser(user *model.User) (bool, error) {

	query := elastic.NewTermQuery("username", user.Username)
	searchResult, err := backend.ESBackend.ReadFromES(query, constants.USER_INDEX)
	if err != nil {
		return false, err
	}

	if searchResult.TotalHits() > 0 {
		return false, nil
	}

	// mysql: insert duplicate?
	err = backend.ESBackend.SaveToES(user, constants.USER_INDEX, user.Username)
	if err != nil {
		return false, err
	}

	return true, nil
}