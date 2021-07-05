package persist

import (
	"log"

	elasticsearch "github.com/elastic/go-elasticsearch"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		for {
			item := <-out
			log.Printf("Item Saver:Got item:%v", item)
			save(item)
		}
	}()
	return out
}
func save(item interface{}) {
	elasticsearch.NewClient()
}
