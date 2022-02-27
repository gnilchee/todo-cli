package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"todo-cli/todo"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the todos",
	Long:  `Listing todo list.`,
	Run:   listRun,
}

func init() {
	rootCmd.AddCommand(listCmd)

}

func listRun(cmd *cobra.Command, args []string) {
	dataFile := viper.GetString("datafile")
	items, err := todo.ReadItems(dataFile)
	if err != nil {
		log.Printf("%v\n", err)
	}
	sort.Sort(todo.ByPri(items))
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	for _, i := range items {
		fmt.Fprintln(w, i.Label()+"\t"+i.PrettyDone()+"\t"+i.PrettyPriority()+"\t"+i.Text+"\t")
	}
	w.Flush()
}
