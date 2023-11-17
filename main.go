package main

import (

	// "k8s.io/cli-runtime/pkg/genericclioptions"
	// "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	// "k8s.io/cli-runtime/pkg/genericiooptions"

	"os"

	"github.com/toj900/grab/cmd"
	"k8s.io/klog/v2"
)

// type GrabFlags struct {
// 	Factory         cmdutil.Factory
// 	Selector        string
// 	AllNamespaces   bool
// 	FilenameOptions *resource.FilenameOptions
// 	genericiooptions.IOStreams
// }

// // NewGrabFlags returns a default GrabFlags
// func NewGrabFlags(f cmdutil.Factory, streams genericiooptions.IOStreams) *GrabFlags {
// 	return &GrabFlags{
// 		Factory:         f,
// 		FilenameOptions: &resource.FilenameOptions{},
// 		IOStreams:       streams,
// 	}
// }

// func defaultConfigFlags() *genericclioptions.ConfigFlags {
// 	return genericclioptions.NewConfigFlags(true).WithDeprecatedPasswordFlag().WithDiscoveryBurst(300).WithDiscoveryQPS(50.0)
// }

func main() {
	if err := cmd.Execute(); err != nil {
		klog.Error(err)
		os.Exit(1)
	}
}

// func main() {
// 	// ioStreams := genericiooptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}
// 	var builderArgs []string
// 	builderArgs = append(builderArgs, "gitrepositories.source.toolkit.fluxcd.io")
// 	builderArgs = append(builderArgs, "buckets.source.toolkit.fluxcd.io")

// 	defaultType := "crd"
// 	kubeConfigFlags := defaultConfigFlags()

// 	b := resource.NewBuilder(kubeConfigFlags).Unstructured()

// 	r := b.ResourceNames(defaultType, builderArgs...).
// 		ContinueOnError().
// 		Flatten().
// 		Do()

// 	err := r.Err()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	infos, err := r.Infos()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	test := &v1beta1.CustomResourceDefinition{}
// 	for i := range infos {
// 		data, err := getObject(infos[i].Object)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		yaml.Unmarshal(data, test)
// 		schema.ParseCRD(test)
// 	}
// }

// func getObject(obj runtime.Object) ([]byte, error) {
// 	var metadataAccessor = meta.NewAccessor()
// 	annots, err := metadataAccessor.Annotations(obj)
// 	if err != nil {
// 		return nil, err
// 	}

// 	original, ok := annots[v1.LastAppliedConfigAnnotation]
// 	if !ok {
// 		return nil, nil
// 	}
// 	return []byte(original), nil
// }
