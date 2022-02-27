package cmd

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"todo-cli/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:     "done [task number]",
	Short:   "Mark task as done",
	Run:     doneRun,
}

func init() {
	rootCmd.AddCommand(doneCmd)

}

func doneRun(cmd *cobra.Command, args []string) {
	dataFile := viper.GetString("datafile")
	items, err := todo.ReadItems(dataFile)
	if err != nil {
		log.Fatalf("Read items : %v\n", err)
	}
	i, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(args[0], "is not a valid label", err)
	}
	if i > 0 && i <= len(items) {
		items[i-1].Done = true
		fmt.Printf("%s %v\n", items[i-1].Text, "marked done")
		sort.Sort(todo.ByPri(items))
		err = todo.SaveItems(dataFile, items)
		if err != nil {
			log.Fatalf("Save items : %v\ns", err)
		}
	} else {
		log.Println(i, "doesnt match any item")
	}
}

