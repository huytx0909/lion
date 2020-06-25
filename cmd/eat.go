package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var eatCmd = &cobra.Command{
	Use:   "eat",
	Short: "order lion to find its prey",
	Run:   hunt,
}

func hunt(command *cobra.Command, args []string) {
	author := viper.GetString("author")
	fmt.Println("The lion is looking for some snack! rawwww " + author)
}
