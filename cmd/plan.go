package cmd

import (
    "fmt"
    "io"
    "os"
    "os/exec"
    "path/filepath"
    "time"

    "github.com/spf13/cobra"
)

// planCmd represents the 'plan' command
var planCmd = &cobra.Command{
    Use:     "plan",
    Short:   "Runs 'terraform plan' to preview infrastructure changes",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Running Terraform plan...")

        // Create log directory if it doesn't exist
        logDir := "logs"
        if err := os.MkdirAll(logDir, 0755); err != nil {
            fmt.Fprintf(os.Stderr, "Error creating logs directory: %v\n", err)
            os.Exit(1)
        }

        // Create a log file with timestamp
        logFile := filepath.Join(logDir, fmt.Sprintf("plan-%s.log", time.Now().Format("20060102-150405")))
        file, err := os.Create(logFile)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error creating log file: %v\n", err)
            os.Exit(1)
        }
        defer file.Close()

        // Run Terraform plan and stream output to both console and file
        command := exec.Command("terraform", "plan")
        command.Stdout = io.MultiWriter(os.Stdout, file)
        command.Stderr = io.MultiWriter(os.Stderr, file)

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