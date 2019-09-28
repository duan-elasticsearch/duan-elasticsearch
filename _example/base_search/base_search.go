package main

import (
	"fmt"
	"os"

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
				Bool: &d_es.Bool {
					Must: &[]d_es.Subnode {
						d_es.Subnode {
							Match: &PasswdCrackType {
								Value: "123456789",
							},
						},
					},
				},
			},
		})),
		es.Search.WithPretty (),
	)
	if err != nil {
		panic (err)
	}

	fmt.Println (res)
}
