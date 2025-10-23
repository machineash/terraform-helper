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

var destroyCmd = &cobra.Command{
    Use:     "destroy",
    Short: "Runs 'terraform destroy -auto-approve' to tear down infrastructure and log results",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Running Terraform destroy...")

        logDir := "logs"
        os.MkdirAll(logDir, 0755)
        logFile := filepath.Join(logDir, fmt.Sprintf("destroy-%s.log", time.Now().Format("20060102-150405")))
        file, _ := os.Create(logFile)
        defer file.Close()

        command := exec.Command("terraform", "destroy", "-auto-approve")
        command.Stdout = io.MultiWriter(os.Stdout, file)
        command.Stderr = io.MultiWriter(os.Stderr, file)

        if err := command.Run(); err != nil {
            fmt.Fprintf(os.Stderr, "Error running terraform destroy: %v\n", err)
            os.Exit(1)
        }

        fmt.Printf("Terraform destroy completed successfully! Log saved to %s\n", logFile)
    },
}

func init() {
    rootCmd.AddCommand(destroyCmd)
}