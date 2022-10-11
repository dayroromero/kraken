package config

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	discovery "github.com/gkarthiks/k8s-discovery"
)

type Secret struct {
	DevSparrowDBPostgres string
}

var (
	k8s *discovery.K8s
)

func GetEnv() (namespace string, err error) {
	k8s, err = discovery.NewK8s()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	namespace, err = k8s.GetNamespace()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	fmt.Printf("Running in namespace: %s\n", namespace)

	return
}

func GetAwsSecret(namespace string) (keyvalue string, err error) {
	secretName := "sparrow/devSparrowDBPostgres"
	region := "us-east-1"
	var secret Secret

	switch namespace {
	case "jx-production":
		GetAwsSecret(secretName)
	case "jx-staging":
		GetAwsSecret(secretName)
	default:
		GetAwsSecret(secretName)
	}

	sess := session.Must(session.NewSession())

	svc := secretsmanager.New(sess, aws.NewConfig().WithRegion(region))

	result, err := svc.GetSecretValue(&secretsmanager.GetSecretValueInput{SecretId: &secretName})
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	json.Unmarshal([]byte(*result.SecretString), &secret)

	keyvalue = secret.DevSparrowDBPostgres

	return
}
