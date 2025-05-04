package templating

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/ADM87/orion/system/templating/types"
	"gopkg.in/yaml.v3"
)

const (
	TemplatingSpecFileName = "spec.yaml"
)

var (
	ErrUnableToResolveSpecPath      = errors.New("unable to resolve spec path")
	ErrUnableToResolveTemplatePaths = errors.New("unable to resolve template paths")
	ErrUnableToReadSpec             = errors.New("unable to read spec file")
	ErrUnableToParseSpec            = errors.New("unable to parse spec file")
)

type ITemplateResolver interface {
	GetTemplatePaths() []string
	GetValues() map[string]any

	LoadSpec(path string) error
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
		p := filepath.Clean(filepath.Join(filepath.Dir(base), path))
		if filepath.Base(p) != "*" {
			resolved = append(resolved, p)
			continue
		}

		collectedPaths, err := resolveWildcardPath(p)
		if err != nil {
			return nil, err
		}

		resolved = append(resolved, collectedPaths...)
	}
	return resolved, nil
}

func resolveWildcardPath(path string) ([]string, error) {
	var paths []string
	return paths, filepath.WalkDir(path, func(p string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			paths = append(paths, p)
		}
		return nil
	})
}
