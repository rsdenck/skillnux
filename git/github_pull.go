package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	if runtime.GOOS != "linux" {
		fmt.Println("Erro: Este script suporta apenas sistemas Linux.")
		os.Exit(1)
	}

	fmt.Println("🔍 Instalando GitHub CLI...")
	
	cmd := exec.Command("which", "apt")
	if err := cmd.Run(); err == nil {
		exec.Command("bash", "-c", "type -p curl >/dev/null || (sudo apt update && sudo apt install curl -y)").Run()
		exec.Command("bash", "-c", "curl -fsSL https://cli.github.com/packages/githubcli-archive-keyring.gpg -o /usr/share/keyrings/githubcli-archive-keyring.gpg").Run()
		exec.Command("bash", "-c", "echo \"deb [signed-by=/usr/share/keyrings/githubcli-archive-keyring.gpg] https://cli.github.com/packages stable main\" | sudo tee /etc/apt/sources.list.d/github-cli.list > /dev/null").Run()
		exec.Command("sudo", "apt", "update").Run()
		exec.Command("sudo", "apt", "install", "gh", "-y").Run()
	}
	
	fmt.Println("✅ GitHub CLI instalado com sucesso!")
	exec.Command("gh", "--version").Stdout = os.Stdout
}
