package config

import (
	"fmt"
	"log"
	"os"

	discovery "github.com/gkarthiks/k8s-discovery"
)

type Secret struct {
	DevSparrowDBPostgres string
}

var (
	k8s *discovery.K8s
)

func GetEnv() (namespace string, err error) {
	// localServer := getDefaultServer()

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
	// secretName := "sparrow/devSparrowDBPostgres"
	// region := "us-east-1"
	// var secret Secret

	// switch namespace {
	// case "jx-production":
	// 	GetAwsSecret(secretName)
	// case "jx-staging":
	// 	GetAwsSecret(secretName)
	// default:
	// 	GetAwsSecret(secretName)
	// }

	// sess := session.Must(session.NewSession())

	// svc := secretsmanager.New(sess, aws.NewConfig().WithRegion(region))

	// result, err := svc.GetSecretValue(&secretsmanager.GetSecretValueInput{SecretId: &secretName})
	// if err != nil {
	// 	log.Fatal(err.Error())
	// 	return
	// }

	// json.Unmarshal([]byte(*result.SecretString), &secret)

	// keyvalue = secret.DevSparrowDBPostgres
	keyvalue = "postgresql://postgres:postgres@localhost:5432/sparrow_db?sslmode=disable"

	return
}

func getDefaultServer() string {
	if server := os.Getenv("DEFAULT_SERVER"); len(server) > 0 {
		return server
	}
	return "http://localhost:8080"
}
