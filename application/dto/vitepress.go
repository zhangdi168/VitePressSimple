package dto

type VitePressCreate struct {
	Title      string `json:"title,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Dir        string `json:"dir,omitempty"`
	DocsDir    string `json:"docsDir,omitempty"`
}
