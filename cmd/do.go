package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/Athulus/tasks/db"
)

func init() {
	rootCmd.AddCommand(do)
}

var do = &cobra.Command{
	Use:   "do",
	Short: "complete a task from your list",
	Run: func(cmd *cobra.Command, args []string) {
		var toDelete []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println(err)
			}
			toDelete = append(toDelete, id)
		}
		err := db.DeleteTasks(toDelete)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("*task has beeen deleted*")
	},
}
