package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	QleetMkdocsConfig     = "mkdocs.yml"
	ThreeportMkdocsConfig = "threeport-user-docs/mkdocs.yml"
)

// MkDocs holds the content of the mkdocs.yml config file
type MkDocs struct {
	SiteName           string                 `yaml:"site_name"`
	SiteAuthor         string                 `yaml:"site_author"`
	RepoURL            string                 `yaml:"repo_url"`
	RepoName           string                 `yaml:"repo_name"`
	EditURI            string                 `yaml:"edit_uri"`
	Nav                []interface{}          `yaml:"nav"`
	Theme              map[string]interface{} `yaml:"theme"`
	Plugins            []interface{}          `yaml:"plugins"`
	Extra              map[string]interface{} `yaml:"extra"`
	MarkdownExtensions []interface{}          `yaml:"markdown_extensions"`
	Copyright          string                 `yaml:"copyright"`
}

// mergeMkdocsConfig nests the content under .nav in the Threeport docs mkdocs
// config under a key 'Threeport' in the Qleet docs .nav.
func mergeMkdocsConfig(config *Config) error {
	qleetData, err := ioutil.ReadFile(QleetMkdocsConfig)
	if err != nil {
		return fmt.Errorf("failed to read Qleet docs mkdocs config: %w", err)
	}

	var qleetMkDocs MkDocs
	err = yaml.Unmarshal(qleetData, &qleetMkDocs)
	if err != nil {
		return fmt.Errorf("failed to unmarshal Qleet docs mkdocs config: %w", err)
	}

	threeportData, err := ioutil.ReadFile(ThreeportMkdocsConfig)
	if err != nil {
		return fmt.Errorf("failed to read Threeport docs mkdocs config: %w", err)
	}

	var threeportMkDocs MkDocs
	err = yaml.Unmarshal(threeportData, &threeportMkDocs)
	if err != nil {
		return fmt.Errorf("failed to unmarshal Threeport docs mkdocs config: %w", err)
	}

	// update file paths to Threeport docs
	updatedThreeportNav := updateDocPath(threeportMkDocs.Nav, config.Exclude)

	// add or update the Threeport nav content to Qleet nav
	updateThreeportNav(&qleetMkDocs.Nav, updatedThreeportNav)

	file, err := os.Create(QleetMkdocsConfig)
	if err != nil {
		return fmt.Errorf("failed to open Qleet mkdocs config file to overwrite: %w", err)
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	defer encoder.Close()
	encoder.SetIndent(2)

	if err = encoder.Encode(&qleetMkDocs); err != nil {
		return fmt.Errorf("failed to write new content to Qleet mkdocs config file: %w", err)
	}

	return nil
}

// updateThreeportNav updates the Threeport nav if present or appends it if not
// in the Qleet docs nav.
func updateThreeportNav(qleetNav *[]interface{}, updatedThreeportNav []interface{}) {
	found := false
	for i, item := range *qleetNav {
		if navMap, ok := item.(map[string]interface{}); ok {
			if _, exists := navMap["Threeport"]; exists {
				(*qleetNav)[i] = map[string]interface{}{"Threeport": updatedThreeportNav}
				found = true
				break
			}
		}
	}
	if !found {
		*qleetNav = append(*qleetNav, map[string]interface{}{"Threeport": updatedThreeportNav})
	}
}

// updateDocPath updates the path in the nav slice. If the path matches any
// value in the excludePaths slice, it removes the item from nav instead of
// updating it.
func updateDocPath(nav []interface{}, excludePaths []string) []interface{} {
	var newNav []interface{}

	// remove the 'docs/threeport/' prefix from exclude paths
	for i, path := range excludePaths {
		excludePaths[i] = strings.TrimPrefix(path, "docs/threeport/")
	}

	for _, item := range nav {
		switch v := item.(type) {
		case map[string]interface{}:
			shouldRemove := false
			for key, val := range v {
				if valSlice, ok := val.([]interface{}); ok {
					v[key] = updateDocPath(valSlice, excludePaths)
				} else if strVal, ok := val.(string); ok {
					for _, removePath := range excludePaths {
						if strVal == removePath {
							shouldRemove = true
							break
						}
					}
					if !shouldRemove {
						v[key] = "threeport/" + strVal
					}
				}
			}
			if !shouldRemove {
				newNav = append(newNav, v)
			}
		case map[string][]interface{}:
			for key, valSlice := range v {
				v[key] = updateDocPath(valSlice, excludePaths)
			}
			newNav = append(newNav, v)
		case string:
			shouldRemove := false
			for _, removePath := range excludePaths {
				if v == removePath {
					shouldRemove = true
					break
				}
			}
			if !shouldRemove {
				newNav = append(newNav, "threeport/"+v)
			}
		default:
			newNav = append(newNav, item)
		}
	}

	return newNav
}
