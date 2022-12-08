package twentyTwentyTwo

import (
	"github.com/spf13/cobra"
)

var twentyTwentyTwo = &cobra.Command{
	Use:   "2022",
	Short: "2022 questions",
}

// Load adds our root command to the command we are given
func Load(root *cobra.Command) {
	root.AddCommand(twentyTwentyTwo)
}
