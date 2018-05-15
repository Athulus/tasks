package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/Athulus/tasks/db"
)

func init() {
	rootCmd.AddCommand(add)
}

var add = &cobra.Command{
	Use:   "add",
	Short: "Add a task to your list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		id, err := db.AddTask(task)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("added new task, `%s` with index %d \n", task, id)
	},
}
