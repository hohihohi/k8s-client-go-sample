package main

import (
	"fmt"
	"log"

	"os"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func loadConfig(masterURL string, path string) (*rest.Config, error) {
	if masterURL != "" {
		if path != "" {
			return clientcmd.BuildConfigFromFlags(masterURL, path)
		}
		homePath := os.Getenv("HOME")
		fmt.Println(homePath)
		return clientcmd.BuildConfigFromFlags(masterURL, homePath+"/.kube/config")
	}
	return rest.InClusterConfig()
}

func main() {
	config, err := loadConfig("0.0.0.0", "")
	if err != nil {
		log.Fatal(err.Error())
		panic(err.Error())
	}
	fmt.Println(config.Host)
	fmt.Println(config.AcceptContentTypes)
}
