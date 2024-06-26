package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func YesNo(question string) bool {
	reader := bufio.NewReader(os.Stdin)
	if reader == nil {
		return false
	}

	fmt.Fprintf(os.Stderr, "%s [Yes/no]: ", question)
	response, err := reader.ReadString('\n')
	if err != nil {
		return false
	}

	response = strings.ToLower(strings.TrimSpace(response))
	if response == "yes" || response == "y" || response == "" {
		return true
	}

	return false
}
