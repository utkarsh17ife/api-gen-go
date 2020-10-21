package template

func GetBabelConfig() (string, string, string) {
	return ".babelrc", `/`, `
{
  "presets": [
    [
      "@babel/preset-env",
      {
        "targets": {
          "node": "11.7"
        }
      }
    ]
  ]
}
`

}
