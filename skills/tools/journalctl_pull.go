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

    skillName := "journalctl"
    
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
