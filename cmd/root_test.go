package cmd

import (
	"os/exec"
	"strings"
	"testing"
)

func TestHelpListsSubcommands(t *testing.T) {
	out, err := exec.Command("go", "run", "../main.go", "--help").CombinedOutput()
	if err != nil {
		t.Fatalf("help failed: %v\n%s", err, out)
	}
	want := []string{"init", "add", "commit", "log", "checkout"}
	for _, w := range want {
		if !strings.Contains(string(out), w) {
			t.Errorf("expected %q in help output", w)
		}
	}
}
