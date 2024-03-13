package vptype

import "regexp"

// NavItemWithLink 结构体代表一个具有链接的导航项。
// Text: 导航项的文本。
// Link: 导航项链接的URL。
// Items: 子导航项的数组，可以是多个 NavItem 实例。
// ActiveMatch: 一个正则表达式，用于检查当前导航项是否处于激活状态。
// Rel: 链接的rel属性。
// Target: 链接的目标窗口属性。
type NavItemWithLink struct {
	Text        string            `json:"text"`
	Link        string            `json:"link,omitempty"`
	Items       []NavItemWithLink `json:"items,omitempty"`
	ActiveMatch *regexp.Regexp    `json:"active_match,omitempty"`
	Rel         string            `json:"rel,omitempty"`
	Target      string            `json:"target,omitempty"`
}

// NavItemChildren 结构体代表含有子项的导航项。
// Text: 导航项的文本。
// Items: 子导航项的数组，每个子项是一个 NavItemWithLink 实例。
type NavItemChildren struct {
	Text  string
	Items []NavItemWithLink
}

// NavItemWithChildren 结构体代表一个含有子项的导航项，它的子项可以是 NavItemWithLink 或 NavItemChildren。
// Text: 导航项的文本。
// Items: 子导航项的数组，可以包含多种类型的导航项。
// ActiveMatch: 一个正则表达式，用于检查当前导航项是否处于激活状态。
type NavItemWithChildren struct {
	Text        string
	Items       []NavItemWithLink // Items 包含的元素类型可以是 NavItemWithLink 或 NavItemChildren。
	ActiveMatch string
}
