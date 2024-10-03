package services

import (
	"fmt"
	"dope-recon/utils"
	"os/exec"
	"strings"
)

func FuzzDomains(domainList string, wordlist string) error {
	// Load domains from the domain list
	domains, err := utils.LoadDomains(domainList)
	if err != nil {
		return fmt.Errorf("error loading domain list: %v", err)
	}

	for _, domain := range domains {
		domain = strings.TrimSpace(domain)
		if domain == "" {
			continue
		}

		// Update resolvers
		utils.UpdateResolvers()

		// Execute PureDNS for each domain
		outputFile := fmt.Sprintf("fuzz-result-%s.txt", domain)
		err = runPureDNS(wordlist, domain, "resolvers.txt", "trusted_resolvers.txt", outputFile)
		if err != nil {
			return fmt.Errorf("error running PureDNS for domain %s: %v", domain, err)
		}
	}
	return nil
}

// Executes the PureDNS command
func runPureDNS(wordlist, domain, resolversFile, trustedResolversFile, outputFile string) error {
	tempOutputFile := fmt.Sprintf("temp-fuzz-result-%s.txt", domain)
	cmd := exec.Command("puredns", "bruteforce", wordlist, domain,
		"--resolvers", resolversFile,
		"--resolvers-trusted", trustedResolversFile,
		"--write", tempOutputFile,
		"--quiet",
	)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to run PureDNS: %v", err)
	}

	// Sort and deduplicate the output
	err = utils.UniqueLines(tempOutputFile, outputFile)
	if err != nil {
		return fmt.Errorf("error sorting output for domain %s: %v", domain, err)
	}

	fmt.Printf("Results saved to %s\n", outputFile)
	return nil
}
