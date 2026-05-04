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

	fmt.Println("🔍 Instalando Ansible...")
	
	// Detecta package manager e instala
	cmd := exec.Command("which", "apt")
	if err := cmd.Run(); err == nil {
		exec.Command("sudo", "apt", "update").Run()
		exec.Command("sudo", "apt", "install", "-y", "ansible").Run()
	} else {
		cmd = exec.Command("which", "yum")
		if err := cmd.Run(); err == nil {
			exec.Command("sudo", "yum", "install", "-y", "ansible").Run()
		}
	}
	
	fmt.Println("✅ Ansible instalado com sucesso!")
	exec.Command("ansible", "--version").Stdout = os.Stdout
}
