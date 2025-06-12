package cmd

import (
	"github.com/spf13/cobra"

	"github.com/thatgurkangurk/mcpack/datapack"
	"github.com/thatgurkangurk/mcpack/utils"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise your pack",
	Run: func(cmd *cobra.Command, args []string) {
		packName, _ := cmd.Flags().GetString("name")
		packId, _ := cmd.Flags().GetString("id")
		outDir, _ := cmd.Flags().GetString("dir")
		forceClean, _ := cmd.Flags().GetBool("clean")

		if forceClean {
			utils.EmptyDirectory(outDir)
		}

		if !utils.IsSnakeCase(packId) {
			cmd.PrintErrln("the pack id has to be in snake case")
			return
		}

		err := datapack.CreatePack(packName, packId, outDir, forceClean)

		if err != nil {
			cmd.PrintErrf("something went wrong: %s", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().String("name", "", "the pack name")
	initCmd.Flags().String("id", "", "the pack id")
	initCmd.Flags().String("dir", "", "the directory to create the pack in (has to be empty, will be created)")
	initCmd.Flags().Bool("clean", false, "force the directory to be emptied")
	initCmd.MarkFlagRequired("name")
	initCmd.MarkFlagRequired("id")
	initCmd.MarkFlagRequired("dir")
	initCmd.MarkFlagDirname("dir")
}
