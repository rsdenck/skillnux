#!/bin/bash
create_go() {
    local name="$1"
    local category="$2"
    local file="skills/${category}/${name}_pull.go"
    mkdir -p "skills/$category"
    
    cat > "$file" << GOEOF
package main

import (
    "fmt"
    "os"
    "runtime"
    "os/exec"
)

func main() {
    if runtime.GOOS != "linux" {
        fmt.Println("Erro: Este script suporta apenas sistemas Linux.")
        os.Exit(1)
    }

    skillName := "${name}"
    
    fmt.Printf("🔍 Instalando %s...\n", skillName)
    
    if hasCommand("apt") {
        installApt(skillName)
    } else if hasCommand("yum") || hasCommand("dnf") {
        installYum(skillName)
    } else if hasCommand("pacman") {
        installPacman(skillName)
    } else {
        fmt.Println("Erro: Package manager não suportado")
        os.Exit(1)
    }
    
    fmt.Printf("✅ %s instalado com sucesso!\n", skillName)
    exec.Command(skillName, "--version").Stdout = os.Stdout
}

func hasCommand(cmd string) bool {
    return exec.Command("which", cmd).Run() == nil
}

func installApt(pkg string) {
    exec.Command("sudo", "apt", "update").Run()
    exec.Command("sudo", "apt", "install", "-y", pkg).Run()
}

func installYum(pkg string) {
    exec.Command("sudo", "yum", "install", "-y", pkg).Run()
}

func installPacman(pkg string) {
    exec.Command("sudo", "pacman", "-S", "--noconfirm", pkg).Run()
}
GOEOF
    echo "Criado: $file"
}

# Categorias
declare -A categories=(
    ["7zip"]="tools" ["age"]="security" ["aider"]="tools" 
    ["aircrack-ng"]="security" ["ansible"]="automation" ["argocd"]="kubernetes"
    ["aws"]="cloud" ["azure"]="cloud" ["bash"]="languages"
    ["bat"]="tools" ["borg"]="tools" ["bpftrace"]="tools"
    ["btop"]="monitoring" ["buildah"]="containers" ["bun"]="languages"
    ["calcurse"]="tools" ["cargo"]="languages" ["ceph"]="storage"
    ["certbot"]="security" ["circleci"]="ci-cd" ["claude"]="tools"
    ["cmus"]="tools" ["composer"]="languages" ["consul"]="web"
    ["cosign"]="security" ["crossplane"]="infrastructure" ["curl"]="tools"
    ["cursor-cli"]="tools" ["cve"]="security" ["dig"]="tools"
    ["discord"]="tools" ["dive"]="containers" ["docker-compose"]="containers"
    ["docker"]="containers" ["doctl"]="cloud" ["dotnet"]="languages"
    ["duf"]="tools" ["enum4linux"]="security" ["etcdctl"]="databases"
    ["ethtool"]="tools" ["eza"]="tools" ["falco"]="security"
    ["fd"]="tools" ["ffuf"]="security" ["fish"]="languages"
    ["fluent-bit"]="monitoring" ["fzf"]="tools" ["gcloud"]="cloud"
    ["gcx"]="tools" ["gem"]="languages" ["geoip"]="tools"
    ["github-cli"]="git" ["gitlab-cli"]="git" ["git"]="git"
    ["golang"]="languages" ["gpg"]="security" ["gradle"]="languages"
    ["grafana-cli"]="monitoring" ["hashcat"]="security" ["helm"]="kubernetes"
    ["htop"]="tools" ["httpie"]="tools" ["hydra"]="security"
    ["iftop"]="tools" ["iperf3"]="tools" ["ipmitool"]="tools"
    ["jaeger"]="monitoring" ["java"]="languages" ["john"]="security"
    ["journalctl"]="tools" ["jq"]="tools" ["jruby"]="languages"
    ["k3d"]="kubernetes" ["k6"]="testing" ["k9s"]="kubernetes"
    ["kind"]="kubernetes" ["kubectl"]="kubernetes" ["kubectx"]="kubernetes"
    ["kubens"]="kubernetes" ["kustomize"]="kubernetes"
    ["linode"]="cloud" ["llama-cpp"]="tools" ["lnav"]="tools"
    ["loki"]="monitoring" ["lynis"]="security" ["masscan"]="security"
    ["maven"]="languages" ["metasploit"]="security" ["micro"]="tools"
    ["midnight-commander"]="tools" ["minikube"]="kubernetes" ["mosh"]="tools"
    ["mtr"]="tools" ["mysql"]="databases" ["nats"]="databases"
    ["ncdu"]="tools" ["neovim"]="tools" ["nerdctl"]="containers"
    ["netcat"]="tools" ["newsboat"]="tools" ["nikto"]="security"
    ["nmap"]="security" ["nodejs"]="languages" ["node"]="languages"
    ["nomad"]="orchestration" ["npm"]="languages" ["nslookup"]="tools"
    ["oci"]="cloud" ["ollama"]="tools" ["openai"]="tools"
    ["openclaw"]="tools" ["openssl"]="security" ["opentelemetry"]="monitoring"
    ["packer"]="infrastructure" ["pandoc"]="tools" ["pass"]="security"
    ["php"]="languages" ["pip"]="languages" ["pnpm"]="languages"
    ["podman"]="containers" ["postman-cli"]="testing" ["promtool"]="monitoring"
    ["proton"]="tools" ["proxmox"]="virtualization" ["psql"]="databases"
    ["pulumi"]="infrastructure" ["python"]="languages" ["qoder"]="tools"
    ["rabbitmq-admin"]="databases" ["rclone"]="tools" ["redis"]="databases"
    ["restic"]="tools" ["ripgrep"]="tools" ["rsync"]="tools"
    ["ruby"]="languages" ["rust"]="languages" ["scp"]="tools"
    ["screen"]="tools" ["sftp"]="tools" ["sgpt"]="tools"
    ["slack"]="tools" ["snmpwalk"]="tools" ["socat"]="tools"
    ["sqlmap"]="security" ["ssh"]="tools" ["stern"]="tools"
    ["syncthing"]="tools" ["sysstat"]="monitoring" ["tailscale"]="tools"
    ["tar"]="tools" ["taskwarrior"]="tools" ["tcpdump"]="tools"
    ["telegram"]="tools" ["tempo"]="monitoring" ["terraform"]="infrastructure"
    ["tldr"]="tools" ["tmux"]="tools" ["todotxt"]="tools"
    ["tofu"]="infrastructure" ["traceroute"]="tools" ["trivy"]="security"
    ["unzip"]="tools" ["uv"]="languages" ["vault"]="security"
    ["vector"]="monitoring" ["virt"]="virtualization" ["websocat"]="tools"
    ["wget"]="tools" ["whois"]="tools" ["wireguard"]="tools"
    ["wpscan"]="security" ["xh"]="tools" ["yq"]="tools"
    ["ytdlp"]="tools" ["zip"]="tools" ["zsh"]="languages"
)

for skill in "${!categories[@]}"; do
    create_go "$skill" "${categories[$skill]}"
done
