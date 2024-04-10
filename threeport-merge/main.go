package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config configures operations for threeport-merge
type Config struct {
	Exclude []string `yaml:"exclude"`
}

func main() {
	// flags
	configFile := flag.String("config-file", "", "path to threeport-merge config")
	flag.Parse()

	// load config
	var config Config
	if configFile != nil {
		if err := loadConfig(*configFile, &config); err != nil {
			fmt.Printf("Error: failed to load config file: %v\n", err)
		}
	}

	// empty existing Threeport content directories from Qleet docs
	if err := clearContent(); err != nil {
		fmt.Printf("Error: failed to empty existing image and markdown content: %w", err)
	}

	// copy threeport markdown docs
	if err := copyMarkdownFiles(&config); err != nil {
		fmt.Printf("Error: failed to copy Threeport documents: %v\n", err)
		os.Exit(1)
	}

	// copy threeport images
	if err := copyImgDir(SourceImgDir, DestinationImgDir); err != nil {
		fmt.Printf("Error: failed to copy images: %v\n", err)
		os.Exit(1)
	}

	// update Qleet mkdocs config to add Threeport docs content
	if err := mergeMkdocsConfig(&config); err != nil {
		fmt.Printf("Error: failed to merge Qleet and Threeport mkdocs configs: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Threeport docs merge complete")
}

func loadConfig(configFile string, config *Config) error {
	content, err := os.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	if err := yaml.Unmarshal(content, &config); err != nil {
		return fmt.Errorf("failed to unmarshal yaml from config file: %w", err)
	}

	return nil
}
