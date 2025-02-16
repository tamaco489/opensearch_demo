package open_search

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/opensearch-project/opensearch-go/v4"
	"github.com/opensearch-project/opensearch-go/v4/opensearchapi"
	"github.com/tamaco489/elasticsearch_demo/api/shop/internal/configuration"

	requestsigner "github.com/opensearch-project/opensearch-go/v4/signer/awsv2"
)

// NewOpenSearchAPIClient:
//
// NOTE: https://opensearch.org/docs/latest/clients/go/
func NewOpenSearchAPIClient(awsCfg aws.Config) (*opensearchapi.Client, error) {
	// Create an opensearch client and use the request-signer.
	client, err := opensearchapi.NewClient(
		opensearchapi.Config{
			Client: opensearch.Config{
				Addresses: []string{
					configuration.Get().OpenSearch.EndPoint,
				},
			},
		},
	)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// NewOpenSearchAPIClientWithSigner:
func NewOpenSearchAPIClientWithSigner(awsCfg aws.Config) (*opensearchapi.Client, error) {

	var endpoint = ""

	switch configuration.Get().API.Env {
	case "dev":
		endpoint = "https://localhost:9200"

	case "stg":
		//

	default:
		//
	}

	// Create an AWS request Signer and load AWS configuration using default config folder or env vars.
	signer, err := requestsigner.NewSignerWithService(awsCfg, "es") // Use "aoss" for Amazon OpenSearch Serverless
	if err != nil {
		return nil, err
	}
	// Create an opensearch client and use the request-signer.
	client, err := opensearchapi.NewClient(
		opensearchapi.Config{
			Client: opensearch.Config{
				Addresses: []string{endpoint},
				Signer:    signer,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	return client, nil
}
