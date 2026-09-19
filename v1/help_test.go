package structparser

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestHelp(t *testing.T) {
	t.Run("Success: Format and print help table correctly", func(t *testing.T) {
		records := []help_record{
			{
				field:       "Host",
				args:        "host",
				env:         "APP_HOST",
				description: "Server host address",
				def:         "localhost",
				required:    false,
			},
			{
				field:       "ApiKey",
				args:        "key",
				env:         "APP_KEY",
				description: "Secret API token",
				def:         "",
				required:    true,
			},
		}

		// 1. Capture stdout output using an os.Pipe
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		// 2. Call the help function
		help(records)

		// 3. Restore stdout and close the pipe write-end
		_ = w.Close()
		os.Stdout = oldStdout

		// 4. Read captured output
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, r)
		output := buf.String()

		// 5. Assertions
		expectedKeywords := []string{
			"FIELD", "FLAG", "ENV VAR", "DEFAULT", "REQUIRED", "DESCRIPTION",
			"Host", "--host", "APP_HOST", "localhost", "false", "Server host address",
			"ApiKey", "--key", "APP_KEY", "true", "Secret API token",
		}

		for _, kw := range expectedKeywords {
			if !strings.Contains(output, kw) {
				t.Errorf("expected help output to contain %q, but got:\n%s", kw, output)
			}
		}
	})

	t.Run("Edge Case: Empty records slice produces no output", func(t *testing.T) {
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		help([]help_record{})

		_ = w.Close()
		os.Stdout = oldStdout

		var buf bytes.Buffer
		_, _ = io.Copy(&buf, r)

		if buf.Len() > 0 {
			t.Errorf("expected no output for empty help records, got %q", buf.String())
		}
	})
}
