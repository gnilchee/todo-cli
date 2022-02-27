package cmd

import (
	"log"
	"fmt"

	"todo-cli/todo"
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
)

var priority int

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [task quoted]",
	Short: "Add a new todo",
	Long:  `Add will create a new todo item to the list.`,
	Run:   addRun,
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority : 1,2,3")
}

func addRun(cmd *cobra.Command, args []string) {
	dataFile := viper.GetString("datafile")
	items, err := todo.ReadItems(dataFile)
	if err != nil {
		log.Printf("%v\n", err)
	}
	for _, x := range args {
		item := todo.Item{Text: x}
		item.SetPriority(priority)
		items = append(items, item)
	}
	fmt.Printf("%#v\n", items)
	err = todo.SaveItems(dataFile, items)
	if err != nil {
		log.Printf("%v\n", err)
	}
}

