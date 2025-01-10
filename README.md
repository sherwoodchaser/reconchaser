# ReconChaser

ReconChaser is a powerful subdomain discovery automation tool designed for reconnaissance and bug bounty hunting. It automatically gathers subdomains for a given domain and saves the results to a file, while removing duplicates for a clean output. The tool is easy to extend with additional discovery tools in the future.

## Features

- Gathers subdomains for a target domain.
- Displays a real-time animated indicator while gathering subdomains.
- Removes duplicates and saves the results to a text file.
- Easily extensible to add new subdomain discovery tools.
- Outputs unique subdomains with the total count and message.

## Installation

1. **Clone this repository**:
    ```bash
    git clone https://github.com/sherwoodchaser/reconchaser.git
    cd reconchaser
    ```

2. **Ensure you have Go installed**:
    - You can download and install Go from the official site: https://golang.org/dl/.
    - Verify your Go installation by running:
      ```bash
      go version
      ```

3. **Install the required tools**:
    - Please ensure you have the required subdomain discovery tools installed (subfinder, findomain).
    - NOTE: you can also add your own tools to source code or replace exsiting ones (look for this part in source code) :
      ```go
      	tools := []struct {
      		tool    string
      		args    []string
      		message string
      	}{
          // add your discovery tools here as a list...
      		{"subfinder", []string{"-d", target, "-silent"}, "\033[36m🔍 Getting subdomains using subfinder\033[0m"},
      		{"findomain", []string{"-t", target, "-q"}, "\033[36m🔍 Getting subdomains using findomain\033[0m"},
      		{"assetfinder", []string{"-subs-only", target}, "\033[36m🔍 Getting subdomains using assetfinder\033[0m"},
      	}
      ```

## Usage

1. **Build the Go tool**:
    ```bash
    go run reconchaser.go -t target.com
    ```
    Replace `target.com` with your desired domain.

## Example Output

```bash
                                        __                         
       ________  _________  ____  _____/ /_  ____ _________  _____
      / ___/ _ \/ ___/ __ \/ __ \/ ___/ __ \/ __ \/ ___/ _ \/ ___/
     / /  /  __/ /__/ /_/ / / / / /__/ / / / /_/ (__  )  __/ /    
    /_/   \___/\___/\____/_/ /_/\___/_/ /_/\__,_/____/\___/_/     
                Built by sherwood chaser.
                      
🔍 Getting subdomains using subfinder [3738 subdomains found]
🔍 Getting subdomains using findomain [3582 subdomains found]
🔍 Getting subdomains using assetfinder [510 subdomains found]
✅ Found 3873 unique subdomains, happy hacking 😊
✅ Subdomains saved to: /home/kali/Desktop/target.com.txt

```
