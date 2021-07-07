package persist

import (
	"context"
	"crawler/engine"
	"errors"
	"log"

	elasticsearch "github.com/olivere/elastic/v7"
)

func ItemSaver() chan engine.Item {
	out := make(chan engine.Item)
	go func() {
		for {
			item := <-out
			log.Printf("Item Saver:Got item:%v", item)
			err := save(item)
			if err != nil {
				log.Printf("Item Saver:error"+"saving item %v:%v", item, err)
			}
		}
	}()
	return out
}
func save(item engine.Item) (err error) {
	client, err := elasticsearch.NewClient(
		elasticsearch.SetSniff(false),
	)
	if err != nil {
		log.Fatalf("create elastic client failed with error:%s", err)
		return err
	}
	if item.Type == "" {
		return errors.New("type is nil")
	}
	_, err = client.Index().
		Index("dating_profile").
		Id(item.Id).
		BodyJson(item).
		Do(context.Background())
	if err != nil {
		log.Fatalf("add data to elasticsearch with error:%s", err)
		return err
	}

	return err
}
