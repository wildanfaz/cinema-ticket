package main

import (
	"context"

	_ "github.com/joho/godotenv/autoload"
	"github.com/wildanfaz/cinema-ticket/cmd"
)

func main() {
	cmd.InitCmd(context.Background())
}
