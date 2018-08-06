package persist

import (
	"log"
	"gopkg.in/olivere/elastic.v5"
	"context"
)

func ItemSaver() chan interface{}{
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for{
			item := <- out
			log.Printf("Item Saver got Item%d: %v", itemCount, item)
			itemCount++

			save(item)
		}
	}()
	return out
}

func save(item interface{}) (id string, err error){
	client, err := elastic.NewClient(
		// must turn off in docker
		elastic.SetSniff(false),
	)

	if err != nil{
		return "", err
	}

	response, err := client.Index().
		Index("dating_profile").
		Type("zhenai").
		BodyJson(item).
		Do(context.Background())

	if err != nil{
		return "", err
	}

	//fmt.Printf("%+v", response)
	return response.Id, nil
}
