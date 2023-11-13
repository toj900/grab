package main

import (
	"log"
	"os"

	"github.com/toj900/grab/pkg/schema"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/util/yaml"
)

func main() {

	// content, err := os.ReadFile("/home/tjmarsh/projects/grab/example/oci-repo.yaml")
	content, err := os.ReadFile("/home/tjmarsh/projects/grab/example/bucket.yaml")
	if err != nil {
		log.Fatal(err)
	}

	test := &v1beta1.CustomResourceDefinition{}
	if err := yaml.Unmarshal(content, test); err != nil {
		log.Fatal(err)
	}
	schema.ParseCRD(test)
}
