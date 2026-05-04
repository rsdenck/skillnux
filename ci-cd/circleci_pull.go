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

	fmt.Println("🔍 Instalando CircleCI CLI...")

	cmd := exec.Command("which", "curl")
	if err := cmd.Run(); err != nil {
		exec.Command("sudo", "apt", "update").Run()
		exec.Command("sudo", "apt", "install", "-y", "curl").Run()
	}

	installCmd := exec.Command("bash", "-c", "curl -fLSs https://raw.githubusercontent.com/CircleCI-Public/circleci-cli/main/install.sh | bash")
	if err := installCmd.Run(); err != nil {
		fmt.Printf("Erro na instalação: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✨ CircleCI CLI instalado com sucesso!")
	exec.Command("circleci", "version").Stdout = os.Stdout
}
