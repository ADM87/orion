package templating

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/ADM87/orion/system/templating/types"
	"gopkg.in/yaml.v3"
)

const (
	TemplatingSpecFileName = "spec.yaml" // Default spec file name with extension
	TemplatingSpecFileExt  = ".yaml"     // Spec file extension
)

var (
	ErrUnableToResolveSpecPath      = errors.New("unable to resolve spec path")
	ErrUnableToResolveTemplatePaths = errors.New("unable to resolve template paths")
	ErrUnableToReadSpec             = errors.New("unable to read spec file")
	ErrUnableToParseSpec            = errors.New("unable to parse spec file")
)

type ITemplateResolver interface {
	GetTemplatePaths() []string // GetTemplatePaths returns the template paths from the spec file
	GetValues() map[string]any  // GetValues returns the values from the spec file
	LoadSpec(path string) error // LoadSpec loads the spec file from the given path
}

type templateResolver struct {
	*types.SpecYaml
}

func NewTemplateResolver() ITemplateResolver {
	return &templateResolver{
		SpecYaml: &types.SpecYaml{},
	}
}

func (tr *templateResolver) GetTemplatePaths() []string {
	return tr.TemplatePaths
}

func (tr *templateResolver) GetValues() map[string]any {
	return tr.Values
}

func (tr *templateResolver) LoadSpec(path string) error {
	path, err := resolveSpecPath(path)
	if err != nil {
		return errors.Join(ErrUnableToResolveSpecPath, err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return errors.Join(ErrUnableToReadSpec, err)
	}
	println(string(data))
	if err := yaml.Unmarshal(data, tr); err != nil {
		return errors.Join(ErrUnableToParseSpec, err)
	}

	if tr.TemplatePaths, err = resolveTemplatePaths(path, tr.TemplatePaths); err != nil {
		return errors.Join(ErrUnableToResolveTemplatePaths, err)
	}
	return nil
}

// ======================================================================
// Helper functions
// ======================================================================

func resolveSpecPath(path string) (string, error) {
	info, err := os.Stat(path)
	if err != nil {
		return "", err
	}

	if !info.IsDir() {
		return path, nil
	}
	return filepath.Join(path, TemplatingSpecFileName), nil
}

func resolveTemplatePaths(base string, paths []string) ([]string, error) {
	var resolved []string
	for _, path := range paths {
		cleanedPath := filepath.Clean(filepath.Join(filepath.Dir(base), path))

		info, err := os.Stat(cleanedPath)
		if err != nil {
			return nil, err
		}

		if info.IsDir() || filepath.Base(cleanedPath) != "*" {
			resolved = append(resolved, cleanedPath)
			continue
		}

		collectedPaths, err := resolveWildcardPath(filepath.Dir(cleanedPath))
		if err != nil {
			return nil, err
		}

		resolved = append(resolved, collectedPaths...)
	}
	return resolved, nil
}

func resolveWildcardPath(path string) ([]string, error) {
	var paths []string
	return paths, filepath.WalkDir(path, func(current string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !entry.IsDir() && filepath.Ext(current) == TemplatingSpecFileExt {
			paths = append(paths, current)
		}
		return nil
	})
}
