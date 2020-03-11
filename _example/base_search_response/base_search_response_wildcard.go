package main

import (
	"fmt"
	"os"
	"reflect"
	"encoding/json"

	d_es "github.com/duan-elasticsearch/duan_elasticsearch/v5"
	"github.com/elastic/go-elasticsearch/v5"
	"github.com/elastic/go-elasticsearch/v5/esutil"
	"github.com/elastic/go-elasticsearch/v5/estransport"
)

type PasswdCrackType struct {
	HType string `json:"h_type,omitempty"`
	Key string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

var host = []string {
	"http://192.168.11.133:9200",
}

func main () {
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

	fmt.Println (res)

	resObj := d_es.QueryResponse {}
	if err := json.NewDecoder (res.Body).Decode (&resObj); err != nil {
		panic (err)
	}

	resObj.CoverSource (reflect.TypeOf (PasswdCrackType{}))
	fmt.Println (resObj)

	fmt.Println ("password:", resObj.Hits.Hits[0].Source.(*PasswdCrackType).Key)
}
