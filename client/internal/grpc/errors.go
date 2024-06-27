package grpc

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/status"
)

func formatDetails(details []any) string {
	if len(details) == 0 {
		return ""
	}

	var sb strings.Builder
	sb.WriteString("\nDetails:")

	for _, detail := range details {
		sb.WriteString(fmt.Sprintf("\n%v", detail))
	}

	return sb.String()
}

func ErrorMessage(err error) (msg string, ok bool) {
	s, ok := status.FromError(err)
	if !ok {
		return fmt.Sprintf("%s", err), false
	}

	code := s.Code()
	desc := code.String()
	message := s.Message()
	details := formatDetails(s.Details())

	return fmt.Sprintf(`Operation resulted in "%s" error (code=%d): %s%s`, desc, code, message, details), true
}
