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
				"func cf{{ .dqid }}(in io.Reader, _w io.Writer) {",
				"\tout := bufio.NewWriter(_w)",
				"\tdefer out.Flush()",
				"\t",
				"}",
				"",
				"func main() { cf{{ .dqid }}(bufio.NewReader(os.Stdin), os.Stdout) }",
			},
			Test: []string{
				"// Generated by copypasta/template/generator_test.go",
				"package main",
				"",
				"import (",
				"\t\"testing\"",
				"",
				"\t\"github.com/EndlessCheng/codeforces-go/main/testutil\"",
				")",
				"",
				"func Test_cf{{ .dqid }}(t *testing.T) {",
				"\ttestCases := [][2]string{",
				"{{ .dexample }}",
				"\t}",
				"\ttestutil.AssertEqualStringCase(t, testCases, 0, cf{{ .dqid }})",
				"}",
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
