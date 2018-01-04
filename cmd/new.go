package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gopkg.in/mgo.v2/bson"
	"github.com/atotto/clipboard"
)

var flags = struct {
	Copy bool
}{}

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Generate a new Mongo ObjectID",
	Run: func(cmd *cobra.Command, args []string) {
		hex := bson.NewObjectId().Hex()
		fmt.Println(hex)

		if flags.Copy {
			err := clipboard.WriteAll(hex)
			if err != nil {
				fmt.Println("failed to copy to clipboard")
				return
			}

			fmt.Println("copied to clipboard")
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().BoolVarP(&flags.Copy, "copy", "c", false, "Copy to clipboard")
}
