package vitepress

import (
	"testing"
	"wailstemplate/application/vitepress/tsparser"
)

func TestParseTsArray(t *testing.T) {
	s := `export const vpSimpleNav = [
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

	println(tsparser.GetTsDataContent(s, "[", "]"))

}
