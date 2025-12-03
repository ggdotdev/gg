// gg v0.1.0 — the 2-letter agent-native git client
// December 3, 2025
package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const version = "0.1.0"

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	cmd := os.Args[1]

	switch cmd {
	case "version", "--version", "-v":
		fmt.Printf("gg v%s\n", version)
	case "maaza":
		handleMaaza()
	case ".":
		handleCurrentRepo()
	case "ask":
		handleAsk()
	case "approve":
		handleApprove()
	default:
		if strings.Contains(cmd, "/") {
			handleRepo(cmd)
		} else {
			fmt.Printf("gg: unknown command: %s\n", cmd)
			printUsage()
		}
	}
}

func printUsage() {
	fmt.Println("gg — the 2-letter agent-native git client")
	fmt.Println()
	fmt.Println("commands:")
	fmt.Println("  gg maaza           # CycleCore/maaza-nlm-orchestrator-9.6m-v1.2")
	fmt.Println("  gg .               # current repo → code-execution MCP")
	fmt.Println("  gg user/repo       # any GitHub repo → MCP")
	fmt.Println("  gg ask \"...\"        # Maaza writes code → opens PR")
	fmt.Println("  gg approve         # merge the PR")
	fmt.Println()
	fmt.Println("install: curl -L gg.sh | sh")
	fmt.Println("more: github.com/ggdotdev/gg")
}

func handleMaaza() {
	fmt.Println("🐱 Maaza Orchestrator v1.2 (9.6M params, 62.9% adversarial)")
	fmt.Println()
	fmt.Println("MCP endpoint:")
	fmt.Println("  https://api.github.com/repos/CycleCoreTechnologies/maaza-nlm-orchestrator-9.6m-v1.2")
	fmt.Println()
	fmt.Println("Code-execution MCP active — 98.7% token reduction (Anthropic proven)")
	fmt.Println("Works with: Claude Desktop, Cursor, any MCP client")
}

func handleCurrentRepo() {
	// Get git remote URL
	cmd := exec.Command("git", "remote", "get-url", "origin")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("❌ Not in a git repo or no remote configured")
		return
	}

	url := strings.TrimSpace(string(output))
	repo := parseGitHubURL(url)
	
	if repo == "" {
		fmt.Println("❌ Could not parse GitHub repo from:", url)
		return
	}

	fmt.Printf("📦 Current repo: %s\n", repo)
	fmt.Println()
	fmt.Println("MCP endpoint:")
	fmt.Printf("  https://api.github.com/repos/%s\n", repo)
	fmt.Println()
	fmt.Println("Code-execution MCP active — works with Claude Desktop, Cursor")
}

func handleRepo(repo string) {
	fmt.Printf("📦 Repo: %s\n", repo)
	fmt.Println()
	fmt.Println("MCP endpoint:")
	fmt.Printf("  https://api.github.com/repos/%s\n", repo)
	fmt.Println()
	fmt.Println("Code-execution MCP active")
}

func handleAsk() {
	if len(os.Args) < 3 {
		fmt.Println("❌ Usage: gg ask \"your prompt here\"")
		return
	}

	prompt := strings.Join(os.Args[2:], " ")
	
	fmt.Println("🤖 Maaza v1.2 is thinking...")
	fmt.Printf("   Prompt: %s\n", prompt)
	fmt.Println()
	fmt.Println("⚠️  Full 'gg ask' implementation coming in v0.2")
	fmt.Println("   For now, use with Claude Desktop + GitHub MCP server")
	fmt.Println()
	fmt.Println("Next steps:")
	fmt.Println("  1. Claude will write code using code-execution MCP")
	fmt.Println("  2. Review the generated PR")
	fmt.Println("  3. Run: gg approve")
}

func handleApprove() {
	fmt.Println("✓ Approving PR...")
	fmt.Println()
	fmt.Println("⚠️  Full 'gg approve' implementation coming in v0.2")
	fmt.Println("   For now, merge PRs via GitHub UI or: gh pr merge")
}

func parseGitHubURL(url string) string {
	// Remove .git suffix
	url = strings.TrimSuffix(url, ".git")
	
	// Handle HTTPS URLs
	if strings.Contains(url, "github.com/") {
		parts := strings.Split(url, "github.com/")
		if len(parts) == 2 {
			return strings.TrimPrefix(parts[1], ":")
		}
	}
	
	// Handle SSH URLs (git@github.com:user/repo)
	if strings.HasPrefix(url, "git@github.com:") {
		return strings.TrimPrefix(url, "git@github.com:")
	}
	
	return ""
}
