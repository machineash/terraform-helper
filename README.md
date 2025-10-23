# Terraform Helper

## Overview
terraform-helper is a Go-based command-line tool built with **Cobra** to streamline common Terraform workflows.

It's the second project in a three-part series focused on infrastructure automation with Terraform, Go, and Azure SDKs.

This phase established the core command framework and full Terraform lifecycle control with logging for each run to simplify troubleshooting and auditing.

## Current Features (Phase 2 Complete)

Commands:
- ```terraform-helper plan```: Runs ```terraform plan``` to preview changes.
- ```terraform-helper apply```: Runs ```terraform apply -auto-approve``` to deploy infrastructure.
- ```terraform-helper destroy```: Runs ```terraform destroy -auto-approve``` to remove infrastructure.

Each command streams Terraform output to the console and writes a timestamped log file under /logs. 

## Next Steps (Phase 3 Preview)
- Integrate the Azure Go SDK to query or validate deployed resources.
- Add environment validation before Terraform runs.
- Improve log formatting with JSON output and error levels.

## Project Context
This tool is part of a three-project sequence:
1. Project 1 - Terraform + Azure Lab: foundational IaC setup (RG, VNet, Subnet).
2. Project 2 - Terraform Helper (this): Go CLI automation layer for Terraform.
3. Project 3 - Azure API Utility: direct Azure SDK and API interaction for telemetry and validation.

Together, they demonstrate full-stack understanding of infrastructure creation, automation, and platform integration.

## License

MIT
