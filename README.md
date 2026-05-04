# NUX Skills Repository

Repositório central de scripts de instalação para a CLI **NUX** (https://github.com/rsdenck/nux).

## Estrutura

```
skillnux/
├── automation/      # Ansible, Puppet, Chef, etc.
├── ci-cd/          # CircleCI, Jenkins, GitLab CI, etc.
├── git/            # GitHub CLI, GitLab CLI, etc.
├── infrastructure/  # Terraform, Pulumi, etc.
├── testing/        # k6, JMeter, etc.
└── README.md
```

## Padrão de Nomenclatura

Cada skill segue o padrão: `{skill}_pull.go`

Exemplo: `ansible` → `ansible_pull.go`

## Branches

- `main`: Scripts estáveis e testados
- `dev`: Novas skills em desenvolvimento

## Como funciona

1. Usuário executa: `nux skill install ansible`
2. NUX lê `skills/ansible.md` que contém:
   ```
   - **Repo:** https://github.com/rsdenck/skillnux/automation/ansible_pull.go
   ```
3. NUX faz download do script `.go`
4. NUX executa o script para instalar a skill

## Requisitos dos Scripts

- Compatível com múltiplas distribuições Linux (apt, yum, dnf, pacman)
- Detecta arquitetura automaticamente (amd64, arm64, etc.)
- Tratamento de erros adequado
- Output claro com emojis indicativos

## Contribuindo

1. Crie nova skill na branch `dev`
2. Teste em múltiplas distribuições
3. Abra Pull Request para `main`

## Licença

MIT
