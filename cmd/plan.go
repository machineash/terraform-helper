package cmd

import (
    "fmt"
    "os"
    "os/exec"

    "github.com/spf13/cobra"
)

// planCmd represents the 'plan' command
var planCmd = &cobra.Command{
    Use:     "plan",
    Short:   "Runs 'terraform plan' to preview infrastructure changes",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Running Terraform plan...")

        command := exec.Command("terraform", "plan")
        command.Stdout = os.Stdout
        command.Stderr = os.Stderr

        if err := command.Run(); err != nil {
            fmt.Fprintf(os.Stderr, "Error running terraform plan: %v\n", err)
            os.Exit(1)
        }

        fmt.Println("Terraform plan completed successfully.")
    },
}

func init() {
    rootCmd.AddCommand(planCmd)
}