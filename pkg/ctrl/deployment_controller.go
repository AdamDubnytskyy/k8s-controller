package ctrl

import (
	"context"
	"fmt"

	"github.com/AdamDubnytskyy/k8s-controller/pkg/config"
	"github.com/rs/zerolog/log"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

type DeploymentReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

var replicas int32 = 1

func (r *DeploymentReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log.Info().Msgf("Reconciling Deployment: %s/%s", req.Namespace, req.Name)
	var deployment appsv1.Deployment
	if err := r.Get(ctx, req.NamespacedName, &deployment); err != nil {
		log.Info().Msgf("Deployment %s/%s not found, assuming deleted", req.Namespace, req.Name)
		return ctrl.Result{}, nil
	}

	if deployment.Labels == nil || deployment.Labels["app"] != "nodejs" {
		log.Info().Msgf("Deployment %s/%s does not have app=nodejs label, skipping", req.Namespace, req.Name)
		return ctrl.Result{}, nil
	}

	// ensure otel configmap
	if err := r.ensureOTelConfigMap(ctx, req.Namespace); err != nil {
		log.Error().Err(err).Msgf("Failed to ensure OpenTelemetry ConfigMap exists")
		return ctrl.Result{}, err
	}

	log.Info().Msgf("Deployment %s/%s has app=nodejs label, adding OTEL environment variables and annotations for automatic instrumentation", req.Namespace, req.Name)

	// Create or update the original deployment with OpenTelemetry instrumentation
	if err := r.instrumentDeploymentWithOTel(ctx, &deployment); err != nil {
		log.Error().Err(err).Msgf("Failed to instrument deployment with OpenTelemetry")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *DeploymentReconciler) ensureOTelConfigMap(ctx context.Context, namespace string) error {
	configMapName := "otel-collector-config"

	var existingConfigMap corev1.ConfigMap
	err := r.Get(ctx, types.NamespacedName{
		Name:      configMapName,
		Namespace: namespace,
	}, &existingConfigMap)

	if err != nil && errors.IsNotFound(err) {
		// ConfigMap doesn't exist, create it
		configMap := &corev1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Name:      configMapName,
				Namespace: namespace,
				Labels: map[string]string{
					"component": "opentelemetry-collector",
				},
			},
			Data: map[string]string{
				config.OTelCollectorConfigFileName: config.OTelCollectorConfig,
			},
		}

		if err := r.Create(ctx, configMap); err != nil {
			return err
		}
		log.Info().Msgf("Created OpenTelemetry ConfigMap: %s/%s", namespace, configMapName)
	} else if err != nil {
		return err
	}

	return nil
}

// instrumentDeploymentWithOTel modifies the original deployment to add OpenTelemetry instrumentation
func (r *DeploymentReconciler) instrumentDeploymentWithOTel(ctx context.Context, deployment *appsv1.Deployment) error {
	// Check if already instrumented
	if deployment.Annotations != nil && deployment.Annotations["otel.instrumented"] == "true" {
		log.Info().Msgf("Deployment %s/%s already instrumented", deployment.Namespace, deployment.Name)
		return nil
	}

	// Create a copy for modification
	updatedDeployment := deployment.DeepCopy()

	// Add annotation to mark as instrumented
	if updatedDeployment.Annotations == nil {
		updatedDeployment.Annotations = make(map[string]string)
	}
	updatedDeployment.Annotations["otel.instrumented"] = "true"

	// Modify the main application container for OpenTelemetry instrumentation
	if len(updatedDeployment.Spec.Template.Spec.Containers) > 0 {
		appContainer := &updatedDeployment.Spec.Template.Spec.Containers[0]

		// Add OpenTelemetry environment variables for Node.js auto-instrumentation
		otelEnvVars := []corev1.EnvVar{
			{
				Name:  "NODE_OPTIONS",
				Value: "--require @opentelemetry/auto-instrumentations-node/register",
			},
			{
				Name:  "OTEL_SERVICE_NAME",
				Value: deployment.Name,
			},
			{
				Name:  "OTEL_EXPORTER_OTLP_ENDPOINT",
				Value: fmt.Sprintf("http://%s-otel-collector.%s.svc.cluster.local:4318", deployment.Name, deployment.Namespace),
			},
			{
				Name:  "OTEL_EXPORTER_OTLP_PROTOCOL",
				Value: "http/protobuf",
			},
			{
				Name:  "OTEL_RESOURCE_ATTRIBUTES",
				Value: fmt.Sprintf("service.name=%s,service.version=1.0.0", deployment.Name),
			},
			{
				Name:  "OTEL_PROPAGATORS",
				Value: "tracecontext,baggage,b3",
			},
		}

		// Append to existing env vars (avoid duplicates)
		existingEnvMap := make(map[string]bool)
		for _, env := range appContainer.Env {
			existingEnvMap[env.Name] = true
		}

		for _, newEnv := range otelEnvVars {
			if !existingEnvMap[newEnv.Name] {
				appContainer.Env = append(appContainer.Env, newEnv)
			}
		}
	}

	// Update the deployment
	if err := r.Update(ctx, updatedDeployment); err != nil {
		return fmt.Errorf("failed to update deployment with OpenTelemetry instrumentation: %w", err)
	}

	log.Info().Msgf("Successfully instrumented deployment %s/%s with OpenTelemetry", deployment.Namespace, deployment.Name)
	return nil
}

func AddDeploymentController(mgr manager.Manager) error {
	r := &DeploymentReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}
	return ctrl.NewControllerManagedBy(mgr).
		For(&appsv1.Deployment{}).
		WithOptions(controller.Options{MaxConcurrentReconciles: 1}).
		Complete(r)
}
