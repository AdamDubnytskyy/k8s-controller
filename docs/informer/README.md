## [Package informers integration](https://pkg.go.dev/k8s.io/client-go/informers)

#### Import packages

- [k8s.io/client-go/informers](https://pkg.go.dev/k8s.io/client-go/informers) 

    Package informers provides generated informers for Kubernetes APIs.

- [k8s.io/apimachinery](https://pkg.go.dev/k8s.io/apimachinery) 

    This library is a shared dependency for servers and clients to work with Kubernetes API infrastructure without direct type dependencies. Its first consumers are k8s.io/kubernetes, k8s.io/client-go, and k8s.io/apiserver.

- [k8s.io/client-go/kubernetes](https://pkg.go.dev/k8s.io/client-go/kubernetes)
    
    Package kubernetes holds packages which implement a clientset for Kubernetes APIs.

- [k8s.io/client-go/tools/cache](https://pkg.go.dev/k8s.io/client-go/tools/cache) 

    Package cache is a client-side caching mechanism. It is useful for reducing the number of server calls you'd otherwise need to make.

E.g.:
```sh
import (
    ...
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/apimachinery/pkg/fields"
    "k8s.io/client-go/informers"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/cache"
```

#### Implement StartDeploymentInformer function to start Deployment Informer

##### To create informer factory we leverage [func NewSharedInformerFactoryWithOptions](https://pkg.go.dev/k8s.io/client-go/informers#NewSharedInformerFactoryWithOptions)

Required arguments:
- [kubernetes clientset](https://pkg.go.dev/k8s.io/client-go@v0.33.2/kubernetes#Clientset)
- resyncPeriod is the time interval at which the informer should resync with the API server. The Informer relists and reprocesses all objects from the cache.
- specify namespace for informer to operate within
- options to filter resources within namespace
```sh
factory := informers.NewSharedInformerFactoryWithOptions(
    clientset,
    resyncTime,
    informers.WithNamespace(namespace),
    informers.WithTweakListOptions(func(options *metav1.ListOptions) {
        options.FieldSelector = fields.Everything().String()
    }),
)

informer := factory.Apps().V1().Deployments().Informer()
```

#### Now that we’ve created a new informer that will watch for changes in Deployment resources, we need to register handler functions for add, update and delete events. 

This is done via the informer.AddEventHandler method:
```sh
...

informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
    AddFunc: func(obj interface{}) {
        log.Info().Msgf("Deployment added: %s", getDeploymentName(obj))
    },
    UpdateFunc: func(oldObj, newObj interface{}) {
        log.Info().Msgf("Deployment updated: %s", getDeploymentName(newObj))
    },
    DeleteFunc: func(obj interface{}) {
        log.Info().Msgf("Deployment deleted: %s", getDeploymentName(obj))
    },
})
```

#### Integration part between Informer and FastHTTP server

Informer is launched as a new goroutine.

Clientset can be configured, either by pasing a `kubeconfig` file or `in-cluster` mode by leveraging [go-client rest package](https://pkg.go.dev/k8s.io/client-go/rest).
```sh
clientset, err := getServerKubeClient(serverKubeconfig, serverInCluster)
...
ctx := context.Background()
go informer.StartDeploymentInformer(ctx, clientset, namespace)

...

func getServerKubeClient(kubeconfigPath string, inCluster bool) (*kubernetes.Clientset, error) {
	...
	if inCluster {
		config, err = rest.InClusterConfig()
	} else {
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	}
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}
```

## [Envtest setup](https://github.com/kubernetes-sigs/controller-runtime/tree/main/tools/setup-envtest)

This is a small tool that manages binaries for envtest. It can be used to download new binaries, list currently installed and available ones, and clean up versions.

To use it, just go-install it with Golang 1.24+ (it's a separate, self-contained module):
```sh
go install sigs.k8s.io/controller-runtime/tools/setup-envtest@latest
```

#### download the latest 1.33.0 envtest
```sh
setup-envtest use -p path 1.33.0
```

#### Set KUBEBUILDER_ASSETS environment variable
```sh
export KUBEBUILDER_ASSETS="$HOME/.local/share/kubebuilder-envtest/k8s/1.33.0-linux-amd64"
```

#### Ensure the env var $KUBEBUILDER_ASSETS is set

```sh
setup-envtest use -i --use-env
```


#### Run tests
```sh
make test
```