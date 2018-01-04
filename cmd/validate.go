package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gopkg.in/mgo.v2/bson"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Check that a hex is a valid ObjectID",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("hex to validate required")
			return
		}

		if len(args) > 1 {
			fmt.Println("too many arguments")
			return
		}

		hex := args[0]

		if ok := bson.IsObjectIdHex(hex); ok {
			fmt.Println("valid")
		} else {
			fmt.Println("invalid")
		}
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// validateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// validateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
