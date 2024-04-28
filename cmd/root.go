package cmd

import (
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
	"fmt"
)

var rootCmd = &cobra.Command{
	Use:   "yourapp",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting TUI...")
		app := tview.NewApplication()

		// Create the main layout
		flex := tview.NewFlex().
			SetDirection(tview.FlexColumn)

		// Left box that takes full height
		leftBox := tview.NewBox().
			SetBorder(true).
			SetTitle("Left Box")

		// Right top box
		rightTopBox := tview.NewBox().
			SetBorder(true).
			SetTitle("Right Top Box")

		// Right bottom box
		rightBottomBox := tview.NewBox().
			SetBorder(true).
			SetTitle("Right Bottom Box")

		// Create a vertical layout for the right side
		rightFlex := tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(rightTopBox, 0, 1, false).
			AddItem(rightBottomBox, 0, 1, false)

		// Add the left box and the right layout to the main layout
		flex.AddItem(leftBox, 0, 1, false).
			AddItem(rightFlex, 0, 1, false)

		// Set the root widget and start the application
		if err := app.SetRoot(flex, true).Run(); err != nil {
			panic(err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func init() {
	// Here, you might include other initializations
}
