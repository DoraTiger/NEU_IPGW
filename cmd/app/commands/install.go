package commands

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"

	"github.com/DoraTiger/NEU_IPGW/version"
)

const (
	installBinaryName   = "NEU_IPGW"
	userDataSubdir      = ".local/share/doratiger/neu_ipgw"
	userBinSubdir       = ".local/bin"
	systemDefaultBindir = "/usr/local/bin"
)

var installSystem bool

var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install the NEU_IPGW binary",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !isInstallSupportedOS() {
			return fmt.Errorf("install command is not supported on %s", runtime.GOOS)
		}
		if installSystem {
			return installToSystem()
		}
		return installToUser()
	},
}

func init() {
	InstallCmd.Flags().BoolVar(&installSystem, "system", false, "install to system directory (requires sudo)")
}

func isInstallSupportedOS() bool {
	switch runtime.GOOS {
	case "linux", "freebsd", "darwin":
		return true
	default:
		return false
	}
}

func installToUser() error {
	src, err := resolveExecutable()
	if err != nil {
		return err
	}
	targetDir, err := userVersionedDir()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(targetDir, 0o755); err != nil {
		return err
	}
	targetBinary := filepath.Join(targetDir, binaryFileName())
	if err := copyBinary(src, targetBinary); err != nil {
		return err
	}
	if err := os.Chmod(targetBinary, 0o755); err != nil {
		return err
	}
	binDir, err := userBinDir()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(binDir, 0o755); err != nil {
		return err
	}
	linkPath := filepath.Join(binDir, binaryFileName())
	if err := prepareLinkPath(linkPath); err != nil {
		return err
	}
	if err := os.Symlink(targetBinary, linkPath); err != nil {
		return err
	}
	fmt.Printf("Installed to %s\n", targetBinary)
	fmt.Printf("Linked at %s\n", linkPath)
	fmt.Printf("Ensure %s is in your PATH\n", binDir)
	return nil
}

func installToSystem() error {
	src, err := resolveExecutable()
	if err != nil {
		return err
	}
	binDir := systemDefaultBindir
	if err := os.MkdirAll(binDir, 0o755); err != nil {
		return err
	}
	targetBinary := filepath.Join(binDir, binaryFileName())
	if err := copyBinary(src, targetBinary); err != nil {
		return err
	}
	if err := os.Chmod(targetBinary, 0o755); err != nil {
		return err
	}
	fmt.Printf("Installed to %s\n", targetBinary)
	return nil
}

func resolveExecutable() (string, error) {
	path, err := os.Executable()
	if err != nil {
		return "", err
	}
	resolved, err := filepath.EvalSymlinks(path)
	if err != nil {
		return "", err
	}
	return resolved, nil
}

func userVersionedDir() (string, error) {
	base, err := userInstallBase()
	if err != nil {
		return "", err
	}
	version := version.BuildVersion
	return filepath.Join(base, version), nil
}

func userInstallBase() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, userDataSubdir), nil
}

func userBinDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, userBinSubdir), nil
}

func binaryFileName() string {
	if runtime.GOOS == "windows" {
		return installBinaryName + ".exe"
	}
	return installBinaryName
}

func copyBinary(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return err
	}
	return dstFile.Sync()
}

func prepareLinkPath(path string) error {
	if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}
