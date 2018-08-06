package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/govinda-attal/dlt-pfm/internal/pfmapi"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
	"github.com/urfave/negroni"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pfmapi",
	Short: "Starts microservice",
	Run:   startServer,
}

func startServer(cmd *cobra.Command, args []string) {
	assetSrv := pfmapi.NewAssetSrv()
	ah := pfmapi.NewAssetHandler(assetSrv)
	r := mux.NewRouter()

	r.PathPrefix("/api/").Handler(http.StripPrefix("/api", http.FileServer(http.Dir("./api"))))

	r.HandleFunc("/assets", pfmapi.WrapperHandler(ah.FetchAsset)).Methods("GET")
	r.HandleFunc("/assets", pfmapi.WrapperHandler(ah.AddAsset)).Methods("POST")

	h := cors.Default().Handler(r)
	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(h)

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      n,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("pfm-api server shutdown ...")
	os.Exit(0)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
}

func main() {
	Execute()
}
