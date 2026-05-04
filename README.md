# NUX Skills Repository

![Skills](https://img.shields.io/badge/NUX-Skills-orange?style=for-the-badge&labelColor=black)
![Scripts](https://img.shields.io/badge/scripts-183+-black?style=for-the-badge&labelColor=orange)
![Go](https://img.shields.io/badge/language-Go-orange?style=for-the-badge&labelColor=black)
![License](https://img.shields.io/badge/license-MIT-black?style=for-the-badge&labelColor=orange)

---

<p align="center">
  <img src="https://img.shields.io/badge/NUX_Skills-Repository-orange?style=for-the-badge&labelColor=black">
</p>

<h1 align="center">🔧 Repositório de Scripts de Deploy para NUX CLI</h1>

<p align="center">
  <b>Scripts Go para instalação automática de +120 ferramentas Linux</b><br/>
  <i>Compatível com múltiplas distribuições</i>
</p>

---

## 📦 Sobre

Este repositório contém todos os scripts `.go` de instalação para as **skills** da CLI **NUX**.

### ✨ O que você encontra aqui

- 🚀 **183 scripts Go** prontos para deploy
- 🗂️ **Organização por categorias** (infrastructure, automation, ci-cd, etc.)
- 🐧 **Multi-distribuição** (apt, yum, dnf, pacman)
- 🔄 **Branches organizadas**: `main` (estável) e `dev` (desenvolvimento)

---

## 🗂️ Estrutura

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
├── tools/             # +50 ferramentas utilitárias
└── README.md
```

---

## 🚀 Como Funciona

1. **Usuário executa** no NUX:
   ```bash
   nux skill install ansible
   ```

2. **NUX lê** o arquivo `skills/ansible.md` que contém:
   ```markdown
   - **Repo:** https://github.com/rsdenck/skillnux/automation/ansible_pull.go
   ```

3. **NUX faz download** do script `.go`

4. **NUX executa** o script para instalar a ferramenta

---

## 📋 Scripts Disponíveis

| Categoria | Scripts | Exemplo |
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

## 🛠️ Padrão dos Scripts

Cada script `.go` segue o padrão:

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
        fmt.Println("Erro: Apenas Linux")
        os.Exit(1)
    }
    
    // Detecta package manager
    if hasCommand("apt") { installApt() }
    else if hasCommand("yum") { installYum() }
    else if hasCommand("pacman") { installPacman() }
    
    fmt.Println("✅ Instalado com sucesso!")
}
```

---

## 🌿 Branches

- `main` - Scripts estáveis e testados
- `dev` - Novas skills em desenvolvimento

---

## 🤝 Contribuindo

1. Crie branch na `dev`: `git checkout -b dev-nova-skill`
2. Adicione o script em `categoria/nova_skill_pull.go`
3. Abra Pull Request para `dev`
4. Após revisão, merge para `main`

---

## 📄 Licença

MIT License - veja [LICENSE](LICENSE)

---

<p align="center">
  <img src="https://img.shields.io/badge/Part_of-NUX_Ecosystem-orange?style=for-the-badge&labelColor=black">
  <img src="https://img.shields.io/badge/For-Linux_Admins-black?style=for-the-badge&labelColor=orange">
</p>
