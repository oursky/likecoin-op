package image

import (
	"os/exec"
)

func Convert(srcPath string, targetPath string) error {
	cmd := exec.Command("convert", srcPath, targetPath)
	_, err := cmd.Output()
	if err != nil {
		return err
	}
	return nil
}
