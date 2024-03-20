package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	SourceDocsDir          = "threeport-user-docs/docs"
	DestinationDocsDir     = "docs/threeport"
	SourceImgDir           = "threeport-user-docs/docs/img"
	DestinationImgDir      = "docs/img/threeport"
	TptctlStr              = "tptctl"
	QleetctlStr            = "qleetctl"
	ThreeportImgPathPrefix = "../img/"
	QleetImgPathPrefix     = "../../../img/threeport/"
)

// clearContent removes existing Threeport docs markdown and image files from
// Qleet docs in perparation for adding updated content.
func clearContent() error {
	for _, dir := range []string{DestinationDocsDir, DestinationImgDir} {
		d, err := os.Open(dir)
		if err != nil {
			return err
		}
		defer d.Close()

		names, err := d.Readdirnames(-1)
		if err != nil {
			return err
		}

		for _, name := range names {
			err = os.RemoveAll(filepath.Join(dir, name))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// copyFile copies a file from src to dst, replaces occurrences of 'tptctl'
// with 'qleetctl' and updates the image path in documents.
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	scanner := bufio.NewScanner(sourceFile)
	writer := bufio.NewWriter(destFile)

	for scanner.Scan() {
		// replace 'tptctl' with 'qleetctl'
		line := strings.Replace(scanner.Text(), TptctlStr, QleetctlStr, -1)
		// update image path
		line = strings.Replace(line, ThreeportImgPathPrefix, QleetImgPathPrefix, -1)
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return writer.Flush()
}

// copyMarkdownFiles copies .md files from SourceDocsDir to DestinationDocsDir.
// It also checks the exclude files from the config and skips if there's a match.
func copyMarkdownFiles(config *Config) error {
	err := filepath.Walk(SourceDocsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			relativePath, err := filepath.Rel(SourceDocsDir, path)
			if err != nil {
				return err
			}

			destPath := filepath.Join(DestinationDocsDir, relativePath)
			if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
				return err
			}

			exclude := false
			for _, excludeFile := range config.Exclude {
				if excludeFile == destPath {
					exclude = true
					break
				}
			}

			if !exclude {
				if err = copyFile(path, destPath); err != nil {
					return err
				}
				fmt.Printf("Copied and modified %s to %s\n", path, destPath)
			}
		}
		return nil
	})

	return err
}

// copyImgFile copies a file from src to dst. If src and dst files exist, and
// are the same, then return success.  Otherwise, attempt to create a copy of
// the src file.
func copyImgFile(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return &os.PathError{Op: "copy", Path: src, Err: err}
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}

// copyImgDir copies the entire image directory recursively.
func copyImgDir(src, dst string) error {
	var err error
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			if err = copyImgDir(srcPath, dstPath); err != nil {
				log.Println(err)
			}
		} else {
			if err = copyImgFile(srcPath, dstPath); err != nil {
				log.Println(err)
			}
		}
	}
	return nil
}
