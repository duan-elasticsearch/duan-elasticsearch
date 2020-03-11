package main

import (
	"fmt"
	// "os"
	"reflect"
	"encoding/json"
	"bytes"

	d_es "github.com/duan-elasticsearch/duan_elasticsearch/v5"
	// "github.com/elastic/go-elasticsearch/v5"
	// "github.com/elastic/go-elasticsearch/v5/esutil"
	// "github.com/elastic/go-elasticsearch/v5/estransport"
)

type PasswdCrackType struct {
	Hehe string `json:"hehe,omitempty"`
	Htype string `json:"h_type,omitempty" jpath:"h_type"`
	Key string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

var host = []string {
	"http://192.168.11.133:9200",
}

func main () {
	/*
	es, err := elasticsearch.NewClient (elasticsearch.Config {
		Logger: &estransport.TextLogger {
			Output: os.Stdout,
		},
		Addresses: host,
	})

	if err != nil {
		panic (err)
	}

	res, err := es.Search (
		es.Search.WithIndex ("password"),
		es.Search.WithSize (4),
		es.Search.WithBody (esutil.NewJSONReader (&d_es.DuanElasticsearch {
			Query: &d_es.Query {
				Wildcard: &PasswdCrackType {
					Value: "*34567*",
				},
			},
		})),
		// es.Search.WithPretty (),
	)
	if err != nil {
		panic (err)
	}
	*/

	fmt.Println (`{"took":24,"timed_out":false,"_shards":{"total":5,"successful":5,"failed":0},"hits":{"total":1,"max_score":0.2876821,"hits":[{"_index":"password","_type":"content","_id":"89ee81bf2e327ece744636ba1f247897dea8f6a7e7f0630a3878e084e1553416b478d451f00ef9cb5f61d1dd49cfbb816a841343a813951c93e6cd917cee26f5","_score":0.2876821,"_source":{"hehe":"hehehaha", "h_type":"13000","key":"123456","value":"$rar5$16$e26de2dc5bd24632a16221881fd57558$15$6684c306a4d54d590f67b2b19d2d67dc$8$f0ee46b40136d4eb"}}]}}`)
	res := bytes.NewReader ([]byte (`{"took":24,"timed_out":false,"_shards":{"total":5,"successful":5,"failed":0},"hits":{"total":1,"max_score":0.2876821,"hits":[{"_index":"password","_type":"content","_id":"89ee81bf2e327ece744636ba1f247897dea8f6a7e7f0630a3878e084e1553416b478d451f00ef9cb5f61d1dd49cfbb816a841343a813951c93e6cd917cee26f5","_score":0.2876821,"_source":{"hehe":"hehehaha","h_type":"13000","key":"123456","value":"$rar5$16$e26de2dc5bd24632a16221881fd57558$15$6684c306a4d54d590f67b2b19d2d67dc$8$f0ee46b40136d4eb"}}]}}`))
	fmt.Println (res)

	resObj := d_es.QueryResponse {}
	if err := json.NewDecoder (res).Decode (&resObj); err != nil {
		panic (err)
	}

	fmt.Println ("begin cover:", resObj.Hits.Hits[0].Source)
	resObj.CoverSource (reflect.TypeOf (PasswdCrackType{}))
	fmt.Println ("end cover:", resObj.Hits.Hits[0].Source)

	fmt.Println ("password:", resObj.Hits.Hits[0].Source)
	fmt.Println ("password:", resObj.Hits.Hits[0].Source.(*PasswdCrackType))
	fmt.Println ("password:", resObj.Hits.Hits[0].Source.(*PasswdCrackType).Key)
}
