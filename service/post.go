package service

import (
	"mime/multipart"
	"reflect"

	"around/backend"
	"around/constants"
	"around/model"

	"github.com/olivere/elastic/v7"
)



func SearchPostByUser(user string) ([]model.Post, error) {

	termQuery := elastic.NewTermQuery("user", user)
	searchResult, err := backend.ESBackend.ReadFromES(termQuery, constants.POST_INDEX)

	if err != nil{
		return nil, err
	}

	var ptype model.Post
	var posts []model.Post
	for _, item := range searchResult.Each(reflect.TypeOf(ptype)) {
		p := item.(model.Post)  // type cast: item is an interface{}
		posts = append(posts,p)
	}

	return posts, nil
}

func SearchPostByKeywords(keywords string) ([]model.Post, error) {

	
	query := elastic.NewMatchQuery("message", keywords)
	// keyword relationship and or ?
	query.Operator("AND")

	// if offer no keywords
	if keywords == "" {
		query.ZeroTermsQuery("all")
	}

	searchResult, err := backend.ESBackend.ReadFromES(query, constants.POST_INDEX)

	if err != nil{
		return nil, err
	}

	var ptype model.Post
	var posts []model.Post
	for _, item := range searchResult.Each(reflect.TypeOf(ptype)) {
		p := item.(model.Post)  // type cast: item is an interface{}
		posts = append(posts,p)
	}

	return posts, nil

}


func SavePost(post *model.Post, file multipart.File) error {
	// use post id as the file name
	mediaLink, err := backend.GCSBackend.SaveToGCS(file, post.Id)
	if err != nil{
		return err
	}

	post.Url = mediaLink
	err = backend.ESBackend.SaveToES(post, constants.POST_INDEX, post.Id)
	return err

}