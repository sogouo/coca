package cmd

import (
	. "coca/src/adapter/sql"
	"github.com/spf13/cobra"
)

var sqlCmd *cobra.Command = &cobra.Command{
	Use:   "sql",
	Short: "scan sql",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		path := cmd.Flag("path").Value.String()
		if path != "" {
			app := new(SqlIdentifierApp)
			app.AnalysisPath(path)
		}
	},
}
