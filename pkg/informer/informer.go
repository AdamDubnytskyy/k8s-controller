package informer

import (
	"context"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

const resyncTime = 30 * time.Second

var deploymentInformer cache.SharedIndexInformer

func StartDeploymentInformer(ctx context.Context, clientset *kubernetes.Clientset, namespace string, labelSelectors string) {
	log.Info().Msgf("StartDeploymentInformer.namespace: %s", namespace)
	log.Info().Msgf("StartDeploymentInformer.labelSelector: %s", labelSelectors)

	factory := informers.NewSharedInformerFactoryWithOptions(
		clientset,
		resyncTime,
		informers.WithNamespace(namespace),
		informers.WithTweakListOptions(func(options *metav1.ListOptions) {
			options.FieldSelector = fields.Everything().String()
			if labelSelectors != "" {
				options.LabelSelector = labelSelectors
			}
		}),
	)
	deploymentInformer = factory.Apps().V1().Deployments().Informer()

	deploymentInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
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

	log.Info().Msg("Starting deployment informer...")
	factory.Start(ctx.Done())
	for t, ok := range factory.WaitForCacheSync(ctx.Done()) {
		if !ok {
			log.Error().Msgf("Failed to sync informer for %v", t)
			os.Exit(1)
		}
	}
	log.Info().Msg("Deployment informer cache synced. Watching for events...")
	<-ctx.Done() // Block until context is cancelled
}

func GetDeploymentNames(namespace string) []string {
	var names []string
	log.Info().Msgf("GetDeploymentNames method: %s", namespace)
	log.Info().Msgf("deploymentInformer: %s", deploymentInformer)
	if deploymentInformer == nil {
		log.Info().Msgf("deploymentInformer is nil")
		return names
	}
	for _, obj := range deploymentInformer.GetStore().List() {
		if d, ok := obj.(*appsv1.Deployment); ok {
			// Filter by namespace
			if d.Namespace == namespace {
				names = append(names, d.Name)
			}
		}
	}
	return names
}

func getDeploymentName(obj any) string {
	if d, ok := obj.(metav1.Object); ok {
		return d.GetName()
	}
	return "unknown"
}
