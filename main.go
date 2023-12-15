package main

import (
	"context"
	"dev/github-fav-language/api"
	"dev/github-fav-language/clients/github"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// HTTP Server
	ctx := context.Background()

	gitHubClient := github.NewClient(http.DefaultClient, `https://api.github.com`) // TODO: Move url to config
	server := api.NewServer(gitHubClient)

	// Graceful shutdown
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-osSignals
		fmt.Fprint(os.Stdout, "🛑 Shutting down Server")
		server.Shutdown()
		fmt.Fprint(os.Stdout, "🛑 HTTP Server shutdown complete")
	}()

	fmt.Println("🚀 Starting HTTP Server")
	server.StartAndListen(ctx)
}
