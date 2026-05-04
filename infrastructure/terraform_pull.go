package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"
)

const (
	repoURL    = "https://api.github.com/repos/hashicorp/terraform/releases/latest"
	installDir = "/usr/local/bin/terraform"
)

type Release struct {
	TagName string `json:"tag_name"`
	Assets  []struct {
		Name               string `json:"name"`
		BrowserDownloadURL string `json:"browser_download_url"`
	} `json:"assets"`
}

func main() {
	if runtime.GOOS != "linux" {
		fmt.Println("Erro: Este script suporta apenas sistemas Linux.")
		os.Exit(1)
	}

	fmt.Println("🔍 Buscando a versão mais recente do Terraform...")
	resp, err := http.Get(repoURL)
	if err != nil {
		fmt.Printf("Erro ao consultar API: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	var release Release
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		fmt.Printf("Erro ao decodificar JSON: %v\n", err)
		os.Exit(1)
	}

	arch := runtime.GOARCH
	suffix := fmt.Sprintf("linux_%s.zip", arch)
	var downloadURL string
	var fileName string

	for _, asset := range release.Assets {
		if strings.HasSuffix(asset.Name, suffix) {
			downloadURL = asset.BrowserDownloadURL
			fileName = asset.Name
			break
		}
	}

	if downloadURL == "" {
		fmt.Printf("Erro: Não foi encontrado um binário para a arquitetura %s\n", arch)
		os.Exit(1)
	}

	fmt.Printf("✅ Versão encontrada: %s\n", release.TagName)
	fmt.Printf("🚀 Baixando: %s\n", fileName)

	if err := downloadFile(fileName, downloadURL); err != nil {
		fmt.Printf("Erro no download: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("📦 Extraindo arquivo...")
	if err := extractAndInstall(fileName); err != nil {
		fmt.Printf("Erro na instalação: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✨ Terraform instalado com sucesso em", installDir)
}

func downloadFile(filepath string, url string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func extractAndInstall(fileName string) error {
	cmdUnzip := exec.Command("unzip", "-o", fileName)
	if err := cmdUnzip.Run(); err != nil {
		return fmt.Errorf("falha ao extrair zip: %v", err)
	}

	cmdMv := exec.Command("sudo", "mv", "terraform", installDir)
	if err := cmdMv.Run(); err != nil {
		return fmt.Errorf("falha ao mover para /usr/local/bin: %v", err)
	}

	cmdChmod := exec.Command("sudo", "chmod", "+x", installDir)
	if err := cmdChmod.Run(); err != nil {
		return err
	}

	os.Remove(fileName)
	return nil
}
