package main

import (
	"api_service/internal/server"
	"context"
	"fmt"
	"os"
)

func main() {
	ctx := context.Background()
	if err := server.Run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
