package cmd

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/valyala/fasthttp"
)

var (
	serverPort int
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start a FastHTTP server",
	Run: func(cmd *cobra.Command, args []string) {
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

func requestHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "FastHTTP server is up & running")

	log.Info().
		Str("Method", string(ctx.Method())).
		Str("Request URI", string(ctx.RequestURI())).
		Str("Query string", ctx.QueryArgs().String()).
		Str("Headers", string(ctx.Request.Header.Header())).
		Msg("Inbound HTTP request")

	ctx.SetContentType("application/json; charset=utf8")

	// Set arbitrary headers
	ctx.Response.Header.Set("X-My-Header", "my-header-value")

}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().IntVar(&serverPort, "port", 8080, "Server port")
}
