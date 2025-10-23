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

var applyCmd = &cobra.Command{
    Use:     "apply",
    Short:   "Runs 'terraform apply -auto-approve' to deploy infrastructure and log results",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Running Terraform apply...")

        logDir := "logs"
        os.MkdirAll(logDir, 0755)
        logFile := filepath.Join(logDir, fmt.Sprintf("apply-%s.log", time.Now().Format("20060102-150405")))
        file, _ := os.Create(logFile)
        defer file.Close()

        command := exec.Command("terraform", "apply", "-auto-approve")
        command.Stdout = io.MultiWriter(os.Stdout, file)
        command.Stderr = io.MultiWriter(os.Stderr, file)

        if err := command.Run(); err != nil {
            fmt.Fprintf(os.Stderr, "Error running terraform apply: %v\n", err)
            os.Exit(1)
        }

        fmt.Printf("Terraform apply completed successfully! Log saved to %s\n", logFile)
    },
}

func init() {
    rootCmd.AddCommand(applyCmd)
}