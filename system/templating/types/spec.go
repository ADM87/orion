package types

type SpecYaml struct {
	TemplatePaths []string       `yaml:"templates"`
	Values        map[string]any `yaml:"values"`
}
