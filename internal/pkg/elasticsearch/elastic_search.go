package elasticsearch

import (
	"emperror.dev/errors"
	"github.com/elastic/go-elasticsearch/v8"
)

func NewElasticSearchClient(opt *ElasticSearchOptions) (*elasticsearch.Client, error) {
	esCfg, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{opt.Url},
	})
	if err != nil {
		return nil, errors.WrapIf(err, "failed to create elasticsearch client")
	}

	return esCfg, nil
}
