package dao

import (
	"distributedCrawfer/dao"
	"distributedCrawfer/engine"
	"log"

	"gopkg.in/olivere/elastic.v5"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := dao.Save(s.Index, s.Client, item)
	log.Printf("Item %v saved.", item)
	if err == nil {
		*result = "ok"
	} else {
		log.Printf("Error saving item: %v : %v", item, err)
	}
	return err
}
