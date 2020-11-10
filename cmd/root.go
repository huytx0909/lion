package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "lion",
		Short: "lion is a testing project",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yml")
	rootCmd.PersistentFlags().StringP("author", "a", "huytx", "author name")
	_ = viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.SetDefault("author", "Default author value-huytx")

	rootCmd.AddCommand(eatCmd)
	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(httpCmd)
	rootCmd.AddCommand(grpcCmd)
	rootCmd.AddCommand(redisCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
