package models

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"sync"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// News struct
type News struct {
	ID        bson.ObjectId `bson:"_id"`
	Title     string
	Image     string
	Detail    string
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

var (
	newsStorage []News
	mutexNews   sync.RWMutex
)

func generateID() string {
	buf := make([]byte, 16)
	rand.Read(buf)
	return base64.StdEncoding.EncodeToString(buf)
}

// CreateNews to mongodb
func CreateNews(news News) error {
	news.ID = bson.NewObjectId()
	news.CreatedAt = time.Now()
	news.UpdatedAt = news.CreatedAt

	s := mongoSession.Copy()
	defer s.Close()
	err := s.DB(database).C("news").Insert(&news)

	if err != nil {
		return err
	}
	return nil
}

// ListNews from mongodb
func ListNews() ([]*News, error) {
	s := mongoSession.Copy()
	defer s.Close()
	var news []*News
	err := s.DB(database).C("news").Find(nil).All(&news)

	if err != nil {
		return nil, err
	}

	return news, nil
}

// GetNews retrives news from database
func GetNews(id string) (*News, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, fmt.Errorf("Invalid id")
	}
	objectID := bson.ObjectIdHex(id)
	s := mongoSession.Copy()
	defer s.Close()
	var n News
	err := s.DB(database).C("news").FindId(objectID).One(&n)
	if err != nil {
		return nil, err
	}
	return &n, nil
}

// DeleteNews by id
func DeleteNews(id string) error {
	objectID := bson.ObjectId(id)
	if !objectID.Valid() {
		return fmt.Errorf("Invalid id")
	}
	s := mongoSession.Copy()
	defer s.Close()
	err := s.DB(database).C("news").RemoveId(objectID)
	if err != nil {
		return err
	}
	return nil
}
