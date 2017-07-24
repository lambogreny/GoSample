package playground

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
	"time"
)

type News struct {
	ID        string
	Title     string
	Image     string
	Detail    string
	CreatedAt time.Time
	UpdatedAt time.Time
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

func CreateNews(news News) {
	news.ID = generateID()
	news.CreatedAt = time.Now()
	news.UpdatedAt = news.CreatedAt
	mutexNews.Lock()
	defer mutexNews.Unlock()
	newsStorage = append(newsStorage, news)
}

func ListNews() []*News {
	mutexNews.RLock()
	defer mutexNews.RUnlock()
	r := make([]*News, len(newsStorage))
	for i := range newsStorage {
		n := newsStorage[i]
		r[i] = &n
	}
	return r
}

func GetNews(id string) *News {
	mutexNews.RLock()
	defer mutexNews.RUnlock()
	for _, news := range newsStorage {
		if news.ID == id {
			n := news
			return &n
		}
	}
	return nil
}

func DeleteNews(id string) {
	mutexNews.Lock()
	defer mutexNews.Unlock()
	for i, news := range newsStorage {
		if news.ID == id {
			newsStorage = append(newsStorage[:i], newsStorage[i+1:]...)
			return
		}
	}
}
