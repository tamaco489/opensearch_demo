package open_search

import (
	"github.com/opensearch-project/opensearch-go/v4"
	"github.com/opensearch-project/opensearch-go/v4/opensearchapi"
	"github.com/tamaco489/opensearch_demo/api/shop/internal/configuration"

	requestsigner "github.com/opensearch-project/opensearch-go/v4/signer/awsv2"
)

// NewOpenSearchAPIClient:
//
// NOTE: https://opensearch.org/docs/latest/clients/go/
func NewOpenSearchAPIClient(cnf configuration.Config) (*opensearchapi.Client, error) {

	// Create an opensearch client and use the request-signer.
	client, err := opensearchapi.NewClient(
		opensearchapi.Config{
			Client: opensearch.Config{
				Addresses: []string{
					cnf.OpenSearch.EndPoint,
				},
				Username: cnf.OpenSearch.Username,
				Password: cnf.OpenSearch.Password,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// NewOpenSearchAPIClientWithSigner:
func NewOpenSearchAPIClientWithSigner(cnf configuration.Config) (*opensearchapi.Client, error) {

	// Create an AWS request Signer and load AWS configuration using default config folder or env vars.
	signer, err := requestsigner.NewSignerWithService(cnf.AWSConfig, "es") // Use "aoss" for Amazon OpenSearch Serverless
	if err != nil {
		return nil, err
	}
	// Create an opensearch client and use the request-signer.
	client, err := opensearchapi.NewClient(
		opensearchapi.Config{
			Client: opensearch.Config{
				Addresses: []string{
					configuration.Get().OpenSearch.EndPoint,
				},
				Signer: signer,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	return client, nil
}
