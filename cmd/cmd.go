package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wildanfaz/cinema-ticket/configs"
	"github.com/wildanfaz/cinema-ticket/internal/routers"
)

func InitCmd(ctx context.Context) {
	var rootCmd = &cobra.Command{
		Short: "Cinema Ticket",
	}

	rootCmd.AddCommand(startEchoApp)
	rootCmd.AddCommand(setAdmin)
	rootCmd.AddCommand(addUserBalance)

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		panic(err)
	}
}

var startEchoApp = &cobra.Command{
	Use:   "start",
	Short: "Start the application",
	Run: func(cmd *cobra.Command, args []string) {
		routers.InitEchoRouter()
	},
}

var setAdmin = &cobra.Command{
	Use:   "set-admin [email] [database-url]",
	Short: "Set role to admin",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		email := args[0]
		databaseUrl := args[1]
		db := configs.InitPostgreSQL(databaseUrl)

		_, err := db.Exec(context.Background(), "UPDATE users SET role = 'admin' WHERE email = $1", email)
		if err != nil {
			panic(err)
		}

		fmt.Println("set-admin success")
	},
}

var addUserBalance = &cobra.Command{
	Use:   "add-balance [email] [amount] [database-url]",
	Short: "Add balance to user",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		email := args[0]
		amount := args[1]
		databaseUrl := args[2]
		db := configs.InitPostgreSQL(databaseUrl)

		_, err := db.Exec(context.Background(), "UPDATE users SET balance = balance + $1 WHERE email = $2", amount, email)
		if err != nil {
			panic(err)
		}

		fmt.Println("add-balance success")
	},
}
