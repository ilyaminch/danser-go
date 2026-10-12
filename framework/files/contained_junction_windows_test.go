package files

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// Directory junctions exercise Windows reparse-point containment without
// requiring the privilege needed to create file symbolic links.
func TestContainedDirectoryJunction(t *testing.T) {
	base := t.TempDir()
	root := filepath.Join(base, "map")
	outside := filepath.Join(base, "outside")
	for _, dir := range []string{root, outside} {
		if err := os.Mkdir(dir, 0700); err != nil {
			t.Fatal(err)
		}
	}
	if err := os.WriteFile(filepath.Join(outside, "background.png"), []byte("dummy"), 0600); err != nil {
		t.Fatal(err)
	}
	link := filepath.Join(root, "external")
	if output, err := exec.Command("cmd.exe", "/c", "mklink", "/J", link, outside).CombinedOutput(); err != nil {
		t.Fatalf("cannot create fixture junction: %v: %s", err, output)
	}
	if f, err := OpenContained(root, filepath.Join(link, "background.png")); err == nil {
		f.Close()
		t.Fatal("outside junction accepted")
	}
}
