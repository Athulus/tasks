package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Athulus/tasks/db"
)

func init() {
	rootCmd.AddCommand(cmdList)
}

var cmdList = &cobra.Command{
	Use:   "list",
	Short: "list all of the tasks you currently have",
	Run:   list,
}

func list(*cobra.Command, []string) {
	tasks, err := db.GetTasks()
	if err != nil {
		fmt.Println(err)
	}
	for _, task := range tasks {
		fmt.Printf("%v: %v \n", task.ID, task.Value)
	}
}
