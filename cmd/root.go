package cmd

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "core-api",
	Short: "Backend API for BWA News",
	Long:  "Backend API for BWA News",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Run(startCmd, nil)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .env)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	// Load .env file using godotenv
	envFile := ".env"
	if cfgFile != "" {
		envFile = cfgFile
	}

	// Load the .env file into environment variables
	if err := godotenv.Load(envFile); err != nil {
		fmt.Printf("Warning: Error loading %s file: %v\n", envFile, err)
	}

	// Enable automatic environment variable reading
	viper.AutomaticEnv()

	// Bind environment variables explicitly
	viper.BindEnv("app.port", "APP_PORT")
	viper.BindEnv("app.env", "APP_ENV")
	viper.BindEnv("jwt.secret_key", "JWT_SECRET_KEY")
	viper.BindEnv("jwt.issuer", "JWT_ISSUER")

	viper.BindEnv("database.host", "DATABASE_HOST")
	viper.BindEnv("database.port", "DATABASE_PORT")
	viper.BindEnv("database.user", "DATABASE_USER")
	viper.BindEnv("database.password", "DATABASE_PASSWORD")
	viper.BindEnv("database.name", "DATABASE_NAME")
	viper.BindEnv("database.max_open", "DATABASE_MAX_OPEN_CONNECTION")
	viper.BindEnv("database.max_idle", "DATABASE_MAX_IDLE_CONNECTION")
}
