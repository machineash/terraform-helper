# Terraform Helper (v2)

## Overview
tfhelper-v2 is a Go-based command-line tool built with **Cobra**.

It's the second project in a three-part series exploring Terraform, Go, and Azure integration.

The goal is to create a lightweight CLI that wraps Terraform workflows, starting with setup and validation, then expanding to ```plan```, ```apply```, and log management.

## Current Phase: Scaffold Complete
Phase 1 established the foundation: 
- Go module initialized and structured with Cobra
- Root command generated (cmd / root.go)
- Builds and runs successfully inside Azure Cloud Shell
- Repository connected to GitHub for version control

At this stage, the CLI runs and responds to help output, confirming the framework is wired correctly. 

## Next Steps
- Phase 2: Add first Terraform commands (plan, apply, destroy)
- Phase 3: Integrate Azure SDK for monitoring and API automation
- Write usage documentation and tests once commands are functional

## Project Context
This tool sits between two other learning projects:
1. Project 1 - Terraform + Azure Lab: declarative IaC foundation
2. Project 2 - Terraform Helper (this): Go CLI engineering layer
3. Project 3 - Azure API Utility: direct cloud API interaction and telemetry

Together, they demonstrate full-stack understanding of infrastructure creation, automation, and platform integration.

## License

MIT
