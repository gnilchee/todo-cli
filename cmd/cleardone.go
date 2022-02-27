package cmd

import (
	"log"

	"todo-cli/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// clearDoneCmd represents the done command
var clearDoneCmd = &cobra.Command{
	Use:     "cleardone",
	Short:   "Remove all tasks marked done",
	Run:     clearDoneRun,
}

func init() {
	rootCmd.AddCommand(clearDoneCmd)

}

func clearDoneRun(cmd *cobra.Command, args []string) {
	dataFile := viper.GetString("datafile")
	items, err := todo.ReadItems(dataFile)
	if err != nil {
		log.Fatalf("Read items : %v\n", err)
	}
	// Clear items with Done set to true
	for i := len(items) - 1; i >= 0; i-- {
		v := items[i]
		if v.Done {
			items = append(items[:i], items[i+1:]...)
		}
	}
	err = todo.SaveItems(dataFile, items)
	if err != nil {
		log.Fatalf("Save items : %v\ns", err)
	}
}

