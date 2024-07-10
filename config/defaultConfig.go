package config

import (
	"os"
)

var defaultConfig = Config{
	Templates: Templates{
		Codeforces: Template{
			Code: []string{
				"package main",
				"",
				"import (",
				"\t\"bufio\"",
				"\t. \"fmt\"",
				"\t\"io\"",
				"\t\"os\"",
				")",
				"",
				"func cf%s(in io.Reader, _w io.Writer) {",
				"\tout := bufio.NewWriter(_w)",
				"\tdefer out.Flush()",
				"\t",
				"}",
				"",
				"func main() { cf%s(bufio.NewReader(os.Stdin), os.Stdout) }",
			},
		},
	},
}

// writeDefaultConfig 写入默认配置
func writeDefaultConfig(configFilePath string) {
	file, err := os.Create(configFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 写入默认配置
	err = WriteConfig(file, defaultConfig)
	if err != nil {
		panic(err)
	}
}
