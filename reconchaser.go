package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"path/filepath"
	"time"
)

const (
	ColorReset  = "\033[0m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorRed    = "\033[31m"
	ColorCyan   = "\033[36m"
	ColorBlue   = "\033[34m"
)

const banner = `
			                    __                         
	   ________  _________  ____  _____/ /_  ____ _________  _____
	  / ___/ _ \/ ___/ __ \/ __ \/ ___/ __ \/ __ \/ ___/ _ \/ ___/
	 / /  /  __/ /__/ /_/ / / / / /__/ / / / /_/ (__  )  __/ /    
	/_/   \___/\___/\____/_/ /_/\___/_/ /_/\__,_/____/\___/_/
		Built by sherwood chaser. 
`

func runCommand(cmd string, args []string) (string, error) {
	command := exec.Command(cmd, args...)
	output, err := command.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error executing command: %v", err)
	}
	return string(output), nil
}

func removeDuplicates(subdomains []string) []string {
	unique := make(map[string]struct{})
	var result []string

	for _, subdomain := range subdomains {
		if _, exists := unique[subdomain]; !exists {
			unique[subdomain] = struct{}{}
			result = append(result, subdomain)
		}
	}

	return result
}

func showLoadingAnimation(message string, done chan bool) {
	loadingChars := []string{"|", "/", "-", "\\"}
	i := 0
	for {
		select {
		case <-done:
			return
		default:
			fmt.Printf("\r%s [%s]", message, loadingChars[i])
			time.Sleep(100 * time.Millisecond)
			i++
			if i >= len(loadingChars) {
				i = 0
			}
		}
	}
}

func getSubdomainsUsingTool(target string, tool string, args []string, message string) ([]string, error) {
	done := make(chan bool)
	go showLoadingAnimation(message, done)
	output, err := runCommand(tool, args)
	if err != nil {
		return nil, err
	}
	done <- true
	subdomainCount := len(strings.Split(output, "\n")) - 1
	fmt.Printf("\r%s [\033[32m%d subdomains found\033[0m]\n", message, subdomainCount)

	subdomains := strings.Split(output, "\n")
	return subdomains, nil
}

func saveSubdomainsToFile(target string, subdomains []string) error {
	fileName := target + ".txt"
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, subdomain := range subdomains {
		if subdomain != "" && strings.HasPrefix(subdomain, "*.") {
			subdomain = subdomain[2:]
		}
		writer.WriteString(subdomain + "\n")
	}
	writer.Flush()
	absPath, err := filepath.Abs(fileName)
	if err != nil {
		return fmt.Errorf("error getting the absolute path: %v", err)
	}
	fmt.Printf("\033[32m✅ Subdomains saved to: %s\033[0m\n", absPath)

	return nil
}



func main() {
	fmt.Println("\033[36m\033[1m" + banner + "\033[0m")
	if len(os.Args) != 3 || os.Args[1] != "-t" {
		fmt.Println("\033[31mUsage: reconchaser -t <target.com>\033[0m")
		return
	}
	target := os.Args[2]
	tools := []struct {
		tool    string
		args    []string
		message string
	}{
    // Define the tools and their respective arguments and messages...
		{"subfinder", []string{"-d", target, "-silent"}, "\033[36m🔍 Getting subdomains using subfinder\033[0m"},
		{"findomain", []string{"-t", target, "-q"}, "\033[36m🔍 Getting subdomains using findomain\033[0m"},
		{"assetfinder", []string{"-subs-only", target}, "\033[36m🔍 Getting subdomains using assetfinder\033[0m"},
	}

	var allSubdomains []string
  
	for _, tool := range tools {
		subdomains, err := getSubdomainsUsingTool(target, tool.tool, tool.args, tool.message)
		if err != nil {
			fmt.Println("\033[31mError fetching subdomains using", tool.tool+":\033[0m", err)
			return
		}
		allSubdomains = append(allSubdomains, subdomains...)
	}
	uniqueSubdomains := removeDuplicates(allSubdomains)
	fmt.Printf("\033[32m✅ Found %d unique subdomains, happy hacking 😊 \033[0m\n", len(uniqueSubdomains))
	err := saveSubdomainsToFile(target, uniqueSubdomains)
	if err != nil {
		fmt.Println("\033[31mError saving subdomains to file:\033[0m", err)
	}
}
