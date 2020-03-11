package main

import (
	"fmt"
	"os"

	d_es "github.com/duan-elasticsearch/duan_elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	"github.com/elastic/go-elasticsearch/v7/estransport"
)

type TestType struct {
	FilePath string `json:"file_path,omitempty"`
	Content string `json:"content,omitempty"`
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
		es.Search.WithIndex ("000ab9f4d33e47cabc79c818f2d0b1b4760ac4ba9a0a8cc90b4f709fd31d1b1c3192e1290920629b96c18eb847432d05c9b512ed751f4a889f5582b37eb008ee_1"),
		es.Search.WithSize (4),
		es.Search.WithBody (esutil.NewJSONReader (&d_es.DuanElasticsearch {
			Query: &d_es.Query {
				Bool: &d_es.Bool {
					Must: &[]d_es.Subnode {
						d_es.Subnode {
							Match: &TestType {
								Content: "123",
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
