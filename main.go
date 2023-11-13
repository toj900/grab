package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"runtime"
	"sort"
	"strings"

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
	parseCRD(test)
}

func parseCRD(crd *v1beta1.CustomResourceDefinition) {
	group := crd.Spec.Group
	name := crd.Spec.Names.Kind
	for _, version := range crd.Spec.Versions {
		buf := &bytes.Buffer{}
		buf.Write([]byte(fmt.Sprintf("apiVersion: %s/%s\n", group, version.Name)))
		buf.Write([]byte(fmt.Sprintf("kind: %s\n", name)))
		buf.Write([]byte(fmt.Sprintf("metadata:\n  name: name\n  namespace: namespace\n")))
		buf.Write([]byte(fmt.Sprintf("spec: \n")))
		properties := version.Schema.OpenAPIV3Schema.Properties["spec"]
		parseProperties(&properties, buf)
		buf.Write([]byte(fmt.Sprintf("---")))
		fmt.Println(buf)
	}
}

// Recursive function to parse CRD schema
func parseProperties(schema *v1beta1.JSONSchemaProps, buf *bytes.Buffer) {
	if len(schema.Properties) == 0 {
		return
	}

	// Ensure ordering
	var schemaProperties []string
	for k := range schema.Properties {
		schemaProperties = append(schemaProperties, k)
	}
	sort.Strings(schemaProperties)

	// for k := range schema.Properties {
	for _, k := range schemaProperties {
		indent := strings.Repeat("  ", callDepth())
		required := false
		for _, r := range schema.Required {
			if callDepth() == 1 && k == r {
				required = true
			}
		}
		nestedSchema := schema.Properties[k]

		// Check if parent object is an array
		if bytes.HasSuffix(buf.Bytes(), []byte("- ")) {
			indent = ""
		} else if !required {
			indent = strings.Repeat("  ", callDepth()) + "# "
		}

		formatedSchema := formatSchema(indent, k, nestedSchema.Type, nestedSchema.Description)
		buf.Write([]byte(formatedSchema))

		// Parse nested Items
		if nestedSchema.Items != nil {
			parseProperties(nestedSchema.Items.Schema, buf)
		}

		// Parse nested Additional
		if nestedSchema.AdditionalItems != nil {
			parseProperties(nestedSchema.AdditionalItems.Schema, buf)
		}

		// Parse Additional Properties
		if nestedSchema.AdditionalProperties != nil {
			parseProperties(nestedSchema.AdditionalProperties.Schema, buf)
		}

		// Main recursive loop
		parseProperties(&nestedSchema, buf)
	}
}

func callDepth() int {
	pc := make([]uintptr, 100)
	return runtime.Callers(6, pc)
}

func formatSchema(indent string, key string, valueType string, description string) string {
	description = strings.ReplaceAll(description, "\n", "")

	switch valueType {
	case "array":
		return fmt.Sprintf("%s: # %s \n%s - ", indent+key, description, indent)

	case "object":
		return fmt.Sprintf("%s: # %s \n", indent+key, description)
	}
	return fmt.Sprintf("%s: %s # %s \n", indent+key, valueType, description)
}
