package dto

// TreeNode 文章列表属性结构
type TreeNode struct {
	Title    string      `json:"title"`
	Key      string      `json:"key"`
	Style    string      `json:"style,omitempty"`
	Class    string      `json:"class,omitempty"`
	Children []*TreeNode `json:"children,omitempty"`
}
type TreeNode1 struct {
	Title string `json:"title"`
	Key   string `json:"key"`
}
