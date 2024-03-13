package vitepress

import (
	"testing"
	"wailstemplate/application/vitepress/tsparser"
)

func TestParseTsArray(t *testing.T) {
	s := `import { defineConfig } from "vitepress"; export const vpSimpleNav = [
  {
    text: 'Home',
    link: '/',
  },
  {
    text: 'About',
    link: '/about',
  },
  {
    text: 'Contact',
    link: '/contact',
  },
]`

	println(tsparser.GetTsDataContent(s, "[", "]", "export"))

}
