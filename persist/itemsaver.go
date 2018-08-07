package persist

import (
	"log"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"crawler/engine"
	"errors"
)

func ItemSaver() chan engine.Item{
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for{
			item := <- out
			log.Printf("Item Saver got Item%d: %v", itemCount, item)
			itemCount++

			err := save(item)

			if err != nil{
				log.Printf("Itemsaver error when saving item %v: %v", item, err)
			}
		}
	}()
	return out
}

func save(item engine.Item) error{
	client, err := elastic.NewClient(
		// must turn off in docker
		elastic.SetSniff(false),
	)

	if err != nil{
		return err
	}

	if item.Type == ""{
		return errors.New("must supply Type")
	}


	indexService := client.Index().
		Index("dating_profile").
		Type(item.Type).
		BodyJson(item)

	if item.Id != ""{
		indexService.Id(item.Id)
	}


	_, err = indexService.
		Do(context.Background())

	if err != nil{
		return err
	}

	//fmt.Printf("%+v", response)
	return nil
}
