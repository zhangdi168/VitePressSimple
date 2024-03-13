package vptype

// SidebarItem 表示侧边栏的一个单项
type SidebarItem struct {
	// 文本标签
	Text string `json:"text,omitempty"`

	// 链接地址
	Link string `json:"link,omitempty"`

	// 子项列表
	Items []SidebarItem `json:"items,omitempty"`

	// 是否默认折叠。如果未指定，则该组不可折叠。
	// 如果为true，则该组可折叠且默认折叠
	// 如果为false，则该组可折叠但默认展开
	Collapsed bool `json:"collapsed,omitempty"`

	// 子项的基础路径
	Base string `json:"base,omitempty"`

	// 自定义显示在上一页/下一页页脚的文字
	DocFooterText string `json:"docFooterText,omitempty"`

	Rel    string `json:"rel,omitempty"`
	Target string `json:"target,omitempty"`
}
