package template

func GetGitIgnore() (string, string, string) {
	return ".gitignore", `/`, `
node_modules
`

}
