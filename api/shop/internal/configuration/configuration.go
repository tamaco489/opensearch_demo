package configuration

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/kelseyhightower/envconfig"
)

var globalConfig Config

type Config struct {
	API struct {
		Env         string `envconfig:"API_ENV" default:"dev"`
		ServiceName string `envconfig:"API_SERVICE_NAME" default:"shop-api"`
		EndPoint    string `envconfig:"API_ENDPOINT" default:"http://localhost:8080"`
	}
	OpenSearch struct {
		Username             string `envconfig:"OPENSEARCH_USERNAME"`
		Password             string `envconfig:"OPENSEARCH_PASSWORD"`
		EndPoint             string `envconfig:"OPENSEARCH_ENDPOINT"`
		InitialAdminPassword string `envconfig:"OPENSEARCH_INITIAL_ADMIN_PASSWORD"`
	}
	Logging   string `envconfig:"LOGGING" default:"off"`
	AWSConfig aws.Config
}

func Get() Config {
	return globalConfig
}

func Load(ctx context.Context) (Config, error) {
	envconfig.MustProcess("", &globalConfig)
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := loadAWSConf(ctx); err != nil {
		return globalConfig, err
	}
	// NOTE: secret情報を管理する必要になったら実装する
	// if err := loadSecrets(ctx, globalConfig); err != nil {
	// 	return globalConfig, err
	// }

	return globalConfig, nil
}
