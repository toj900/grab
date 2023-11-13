package main

import (

	// "k8s.io/cli-runtime/pkg/genericclioptions"
	// "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	// "k8s.io/cli-runtime/pkg/genericiooptions"

	"log"

	"github.com/toj900/grab/pkg/schema"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/cli-runtime/pkg/genericclioptions"

	// "k8s.io/cli-runtime/pkg/genericiooptions"
	"k8s.io/cli-runtime/pkg/resource"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubectl/pkg/util"
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

func defaultConfigFlags() *genericclioptions.ConfigFlags {
	return genericclioptions.NewConfigFlags(true).WithDeprecatedPasswordFlag().WithDiscoveryBurst(300).WithDiscoveryQPS(50.0)
}

func main() {
	// ioStreams := genericiooptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}

	loot := genericclioptions.NewConfigFlags(true)
	lootFlags := cmdutil.NewMatchVersionFlags(loot)
	cmdutil.NewFactory(lootFlags)

	kubeConfigFlags := defaultConfigFlags()
	matchVersionKubeConfigFlags := cmdutil.NewMatchVersionFlags(kubeConfigFlags)

	f := cmdutil.NewFactory(matchVersionKubeConfigFlags)
	// o := NewGrabFlags(f, ioStreams)
	bro := resource.FilenameOptions{}
	bro.Filenames = []string{}

	cmdNamespace, _, err := f.ToRawKubeConfigLoader().Namespace()
	// cmdNamespace, enforceNamespace, err := f.ToRawKubeConfigLoader().Namespace()
	if err != nil {
		log.Fatal(err)
	}

	b := f.NewBuilder().Unstructured()

	r := b.NamespaceParam(cmdNamespace).DefaultNamespace().
		ResourceTypeOrNameArgs(true, "crd").
		// ResourceTypeOrNameArgs(true, "crd", "buckets.source.toolkit.fluxcd.io").
		Subresource("").
		ContinueOnError().
		Flatten().
		Do()
	err = r.Err()
	if err != nil {
		log.Fatal(err)
	}

	infos, err := r.Infos()
	if err != nil {
		log.Fatal(err)
	}

	test := &v1beta1.CustomResourceDefinition{}
	for i := range infos {
		data, err := util.GetOriginalConfiguration(infos[i].Object)
		if err != nil {
			log.Fatal(err)
		}
		yaml.Unmarshal(data, test)
		schema.ParseCRD(test)
	}
}
