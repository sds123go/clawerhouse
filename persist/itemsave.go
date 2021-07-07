package persist

import (
	"context"
	"crawler/engine"
	"encoding/json"
	"errors"
	"log"

	elasticsearch "github.com/olivere/elastic/v7"
)

func ItemSaver() chan engine.Item {
	out := make(chan engine.Item)
	client, err := elasticsearch.NewClient(
		elasticsearch.SetSniff(false),
	)
	if err != nil {
		log.Fatalf("create elastic client failed with error:%s", err)
	}
	go func() {
		for {
			item := <-out
			log.Printf("Item Saver:Got item:%+v", item)
			err := save(client, item)
			if err != nil {
				log.Printf("Item Saver:error"+"saving item %v:%v", item, err)
			}
		}
	}()
	return out
}

type TempStruct struct {
	Url     string
	Type    string
	Id      string
	Payload string
}

func save(client *elasticsearch.Client, item engine.Item) (err error) {

	if item.Type == "" {
		return errors.New("type is nil")
	}

	if err != nil {
		return err
	}
	marshal, err := json.Marshal(item.Payload)
	if err != nil {
		return err
	}
	var temp = TempStruct{
		Url:     item.Url,
		Type:    item.Type,
		Id:      item.Id,
		Payload: string(marshal),
	}
	_, err = client.Index().
		Index("juli").
		BodyJson(temp).
		Do(context.Background())
	if err != nil {
		log.Fatalf("add data to elasticsearch with error:%s", err)
		return err
	}

	return err
}
