package persist

import (
	"log"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"crawler/engine"
	"errors"
)

func ItemSaver(index string) (chan engine.Item, error){
	client, err := elastic.NewClient(
		// must turn off in docker
		elastic.SetSniff(false),
		elastic.SetURL("http://192.168.99.100:9200"),
	)

	if err != nil{
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for{
			item := <- out
			log.Printf("Item Saver got Item%d: %v", itemCount, item)
			itemCount++

			err := Save(client, item, index)

			if err != nil{
				log.Printf("Itemsaver error when saving item %v: %v", item, err)
			}
		}
	}()

	return out, nil
}

func Save(client *elastic.Client, item engine.Item, index string) error{


	if item.Type == ""{
		return errors.New("must supply Type")
	}


	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)

	if item.Id != ""{
		indexService.Id(item.Id)
	}


	_, err := indexService.
		Do(context.Background())

	if err != nil{
		return err
	}

	//fmt.Printf("%+v", response)
	return nil
}
