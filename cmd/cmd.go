package cmd

import (
	"log"
	"os"

	"github.com/nk521/pehredar/pehredar"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pehredar",
	Short: "Watchdog for unknown USBs!",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Starting Pehredar ... ")
	},
}

var databaseCmd = &cobra.Command{
	Use:   "db",
	Short: "Show and edit Pehredar's database",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var randomFillDBCmd = &cobra.Command{
	Use: "random",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Filling up random stuff in database ...")
		pdb := pehredar.NewPehredarDatabase()
		pdb.Clear()
		pdb = pdb.Refresh()
		pdb.RandomFill()
	},
}

var clearDBCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear Pehredar's database",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Clearing database ...")
		pdb := pehredar.NewPehredarDatabase()
		pdb.Clear()
		log.Println("Cleared database!")
	},
}

var viewDBCmd = &cobra.Command{
	Use:   "view",
	Short: "View Pehredar's database",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Viewing database")
	},
}

var addDBCmd = &cobra.Command{
	Use:   "add",
	Short: "Add all the current attached USBs to Pehredar's database",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Adding to database")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func init() {
	databaseCmd.AddCommand(clearDBCmd)
	databaseCmd.AddCommand(viewDBCmd)
	databaseCmd.AddCommand(addDBCmd)
	databaseCmd.AddCommand(randomFillDBCmd)

	rootCmd.AddCommand(databaseCmd)
}
