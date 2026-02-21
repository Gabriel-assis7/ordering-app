package elasticsearch

import (
	"fmt"

	"emperror.dev/errors"
	"github.com/elastic/go-elasticsearch/v8"
)

func NewElasticSearchClient(opt *ElasticSearchOptions) (*elasticsearch.Client, error) {
	esCfg, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{opt.Url},
	})
	if err != nil {
		return nil, errors.WrapIf(
			err,
			fmt.Sprintf("failed to create Elasticsearch client on URL %s", opt.Url),
		)
	}

	return esCfg, nil
}
