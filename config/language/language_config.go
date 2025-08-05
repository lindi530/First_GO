package language

type LanguageConfig struct {
	CompileCmd  []string `mapstructure:"compile_cmd"`
	RunCmd      []string `mapstructure:"run_cmd"`
	SourceFile  string   `mapstructure:"source_file"`
	Executable  string   `mapstructure:"executable"`
	OutputFile  string   `mapstructure:"output_file"`
	Language    string   `mapstructure:"language"`
	NeedCompile bool     `mapstructure:"need_compile"`
}
