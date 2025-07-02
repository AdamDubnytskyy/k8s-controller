package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/AdamDubnytskyy/k8s-controller/pkg/ctrl"
	"github.com/AdamDubnytskyy/k8s-controller/pkg/informer"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/valyala/fasthttp"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	ctrlruntime "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var (
	serverPort       int
	serverKubeconfig string
	serverInCluster  bool
	labelSelectors   []string
	labelSelector    string
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start a FastHTTP server and deployment informer",
	Run: func(cmd *cobra.Command, args []string) {
		clientset, err := getServerKubeClient(serverKubeconfig, serverInCluster)
		if err != nil {
			log.Error().Err(err).Msg("Failed to create Kubernetes client")
			os.Exit(1)
		}
		ctx := context.Background()
		labelSelector := strings.Join(labelSelectors, ",")
		go informer.StartDeploymentInformer(ctx, clientset, namespace, labelSelector)

		mgr, err := ctrlruntime.NewManager(ctrlruntime.GetConfigOrDie(), manager.Options{})
		if err != nil {
			log.Error().Err(err).Msg("Failed to create controller-runtime manager")
			os.Exit(1)
		}
		if err := ctrl.AddDeploymentController(mgr); err != nil {
			log.Error().Err(err).Msg("Failed to add deployment controller")
			os.Exit(1)
		}
		go func() {
			log.Info().Msg("Starting controller-runtime manager...")
			if err := mgr.Start(cmd.Context()); err != nil {
				log.Error().Err(err).Msg("Manager exited with error")
				os.Exit(1)
			}
		}()

		handler := requestHandler
		addr := fmt.Sprintf(":%d", serverPort)
		fmt.Println(addr)
		log.Info().Msgf("FastHTTP server started on %s port", addr)

		if err := fasthttp.ListenAndServe(addr, handler); err != nil {
			log.Error().Err(err).Msg("Error starting FastHTTP server")
			os.Exit(1)
		}
	},
}

func getServerKubeClient(kubeconfigPath string, inCluster bool) (*kubernetes.Clientset, error) {
	var config *rest.Config
	var err error
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

func requestHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/deployments":
		log.Info().Msg("Deployments request received")
		ctx.Response.Header.Set("Content-Type", "application/json")
		deployments := informer.GetDeploymentNames(namespace)
		log.Info().Msgf("Deployments: %v", deployments)
		// ctx.SetStatusCode(200)
		ctx.Write([]byte("["))
		for i, name := range deployments {
			ctx.WriteString("\"")
			ctx.WriteString(name)
			ctx.WriteString("\"")
			if i < len(deployments)-1 {
				ctx.WriteString(",")
			}
		}
		ctx.Write([]byte("]"))
		return
	default:
		log.Info().Msg("Default request received")
		fmt.Fprintf(ctx, "Hello from FastHTTP!")
	}

	log.Info().
		Str("Method", string(ctx.Method())).
		Str("Request URI", string(ctx.RequestURI())).
		Str("Query string", ctx.QueryArgs().String()).
		Str("Headers", string(ctx.Request.Header.Header())).
		Msg("Inbound HTTP request")

	ctx.SetContentType("application/json; charset=utf8")
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().IntVar(&serverPort, "port", 8080, "Server port")
	serverCmd.Flags().StringVar(&serverKubeconfig, "kubeconfig", "", "Path to the kubeconfig file")
	serverCmd.Flags().BoolVar(&serverInCluster, "in-cluster", false, "Use in-cluster Kubernetes config")
	serverCmd.Flags().StringVar(&namespace, "namespace", "", "Specify kubernetes namespace")
	serverCmd.Flags().StringSliceVar(&labelSelectors, "labelSelectors", []string{}, "Specify kubernetes label selectors. E.g.: labelName=labelValue")
}
