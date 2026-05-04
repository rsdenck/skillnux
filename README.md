# NUX Skills Repository

![NUX Skills](https://img.shields.io/badge/NUX-Skills-orange?style=for-the-badge&labelColor=black)
![Scripts](https://img.shields.io/badge/scripts-183+-black?style=for-the-badge&labelColor=orange)
![Go](https://img.shields.io/badge/language-Go-orange?style=for-the-badge&labelColor=black)
![License](https://img.shields.io/badge/license-MIT-black?style=for-the-badge&labelColor=orange)

---

<p align="center">
  <img src="https://img.shields.io/badge/NUX_Skills-Repository-orange?style=for-the-badge&labelColor=black">
</p>

<h1 align="center">NUX Skills Repository</h1>

<p align="center">
  <b>Go scripts for automatic installation of 120+ Linux tools</b><br/>
  <i>Compatible with multiple Linux distributions</i>
</p>

---

## About This Repository

This repository contains all `.go` installation scripts for the **NUX CLI** skills system.

### Key Features

- **183 Go scripts** ready for deployment
- **Organized by categories** (infrastructure, automation, ci-cd, etc.)
- **Multi-distribution support** (apt, yum, dnf, pacman)
- **Organized branches**: `main` (stable) and `dev` (development)

---

## Repository Structure

```
skillnux/
├── infrastructure/     # Terraform, Pulumi, Packer, Crossplane
├── automation/         # Ansible, Puppet, Chef
├── ci-cd/             # CircleCI, Jenkins, GitLab CI
├── git/               # GitHub CLI, GitLab CLI
├── testing/           # k6, JMeter
├── kubernetes/        # kubectl, Helm, K9s, Kind
├── containers/        # Docker, Podman, Buildah
├── security/          # Nmap, Metasploit, John, Hashcat
├── monitoring/        # Prometheus, Grafana, Loki
├── databases/         # Redis, MySQL, PostgreSQL
├── cloud/             # AWS, Azure, GCloud
├── languages/         # Python, Go, Rust, Java
├── tools/             # 50+ utility tools
└── README.md
```

---

## How It Works

1. **User executes** in NUX CLI:
   ```bash
   nux skill install ansible
   ```

2. **NUX reads** the skill file `skills/ansible.md` which contains:
   ```markdown
   - **Repo:** https://github.com/rsdenck/skillnux/automation/ansible_pull.go
   ```

3. **NUX downloads** the `.go` script from this repository

4. **NUX executes** the script to install the tool

---

## Available Scripts

| Category | Scripts | Example |
|----------|---------|--------|
| Infrastructure | 4 | `terraform_pull.go`, `pulumi_pull.go` |
| Automation | 1 | `ansible_pull.go` |
| CI/CD | 2 | `circleci_pull.go`, `github_pull.go` |
| Git | 3 | `gh-cli_pull.go`, `gitlab-cli_pull.go` |
| Testing | 1 | `k6_pull.go` |
| Kubernetes | 8 | `kubectl_pull.go`, `helm_pull.go` |
| Containers | 4 | `docker_pull.go`, `podman_pull.go` |
| Security | 15+ | `nmap_pull.go`, `metasploit_pull.go` |
| Tools | 50+ | `htop_pull.go`, `bat_pull.go` |

---

## Script Standard

Each `.go` script follows this pattern:

```go
package main

import (
    "fmt"
    "os"
    "runtime"
    "os/exec"
)

func main() {
    if runtime.GOOS != "linux" {
        fmt.Println("Error: This script supports Linux only.")
        os.Exit(1)
    }
    
    // Detect package manager
    if hasCommand("apt") {
        installApt()
    } else if hasCommand("yum") || hasCommand("dnf") {
        installYum()
    } else if hasCommand("pacman") {
        installPacman()
    } else {
        fmt.Println("Error: Unsupported package manager")
        os.Exit(1)
    }
    
    fmt.Println("Success: Tool installed!")
}
```

---

## Branches

- `main` - Stable and tested scripts
- `dev` - New skills in development

---

## Contributing

1. Create branch from `dev`: `git checkout -b dev-new-skill`
2. Add the script in `category/new_skill_pull.go`
3. Open Pull Request to `dev`
4. After review, merge to `main`

---

## License

MIT License - see [LICENSE](LICENSE)

---

<p align="center">
  <img src="https://img.shields.io/badge/Part_of-NUX_Ecosystem-orange?style=for-the-badge&labelColor=black">
  <img src="https://img.shields.io/badge/For-Linux_Admins-black?style=for-the-badge&labelColor=orange">
</p>
