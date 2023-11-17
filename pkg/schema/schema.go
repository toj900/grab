package schema

import (
	"bytes"
	"fmt"
	"log"
	"runtime"
	"strings"

	// "gopkg.in/yaml.v2"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/meta"
	k8sruntime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/resource"
)

type customSchema struct {
	key          string
	value        string
	valueType    string
	description  string
	indent       int
	required     bool
	childElement bool
}

func defaultConfigFlags() *genericclioptions.ConfigFlags {
	return genericclioptions.NewConfigFlags(true).WithDeprecatedPasswordFlag().WithDiscoveryBurst(300).WithDiscoveryQPS(50.0)
}

func GrabCRD(builderArgs []string) {
	// var builderArgs []string
	// builderArgs = append(builderArgs, "gitrepositories.source.toolkit.fluxcd.io")
	// builderArgs = append(builderArgs, "buckets.source.toolkit.fluxcd.io")

	defaultType := "crd"
	kubeConfigFlags := defaultConfigFlags()

	b := resource.NewBuilder(kubeConfigFlags).Unstructured()

	r := b.ResourceNames(defaultType, builderArgs...).
		ContinueOnError().
		Flatten().
		Do()

	err := r.Err()
	if err != nil {
		log.Fatal(err)
	}

	infos, err := r.Infos()
	if err != nil {
		log.Fatal(err)
	}
	test := &v1beta1.CustomResourceDefinition{}
	for i := range infos {
		data, err := getObject(infos[i].Object)
		if err != nil {
			log.Fatal(err)
		}
		yaml.Unmarshal(data, test)
		ParseCRD(test)
	}
}

func ParseCRD(crd *v1beta1.CustomResourceDefinition) {
	group := crd.Spec.Group
	name := crd.Spec.Names.Kind
	for _, version := range crd.Spec.Versions {
		buf := &bytes.Buffer{}
		buf.Write([]byte(fmt.Sprintf("apiVersion: %s/%s\n", group, version.Name) +
			fmt.Sprintf("kind: %s\n", name) +
			fmt.Sprintf("metadata:\n  name: name\n  namespace: namespace\n") +
			fmt.Sprintf("spec:\n"),
		))

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

	suffix := []byte("- ")
	calldepth := callDepth() - 6
	required := make(map[string]bool)
	for _, r := range schema.Required {
		required[r] = true
	}

	for k, v := range schema.Properties {
		c := customSchema{
			key:          k,
			valueType:    v.Type,
			description:  v.Description,
			indent:       calldepth,
			required:     false,
			childElement: false,
		}

		// Check if required
		if calldepth == 1 && required[k] {
			c.required = true
		}

		// Check if parent is an array
		if bytes.HasSuffix(buf.Bytes(), suffix) {
			c.childElement = true
		}
		buf.Write(c.formatSchema())

		// Parse nested items
		if v.Items != nil {
			parseProperties(v.Items.Schema, buf)
		}

		// Parse additional items
		if v.AdditionalItems != nil {
			parseProperties(v.AdditionalItems.Schema, buf)
		}

		// Parse additional properties
		if v.AdditionalProperties != nil {
			parseProperties(v.AdditionalProperties.Schema, buf)
		}

		// Main recurcive loop
		parseProperties(&v, buf)
	}
}

// formatSchema formats the customSchema
func (c *customSchema) formatSchema() []byte {
	value := map[string]string{
		"string":  "\"string\"",
		"integer": "1",
		"boolean": "true",
		"object":  "",
		"array":   "",
	}

	description := strings.ReplaceAll(c.description, "\n", "") + "\n"
	indent := strings.Repeat("  ", c.indent)

	// Check if schema has default value
	if c.value != "" {
		value[c.valueType] = c.value
	}

	// Comment non required keys
	if !c.required {
		indent = indent + "# "
	}

	// Add hyphen for arrays
	if c.valueType == "array" {
		description = description + indent + "- "
	}

	// Remove indent for array elements
	if c.childElement {
		indent = ""
	}

	return []byte(fmt.Sprintf("%s: %s # %s", indent+c.key, value[c.valueType], description))
}

// Calldepth checks the depth of a function call
func callDepth() int {
	pc := make([]uintptr, 100)
	return runtime.Callers(6, pc)
}

func getObject(obj k8sruntime.Object) ([]byte, error) {
	var metadataAccessor = meta.NewAccessor()
	annots, err := metadataAccessor.Annotations(obj)
	if err != nil {
		return nil, err
	}

	original, ok := annots[v1.LastAppliedConfigAnnotation]
	if !ok {
		return nil, nil
	}
	return []byte(original), nil
}
