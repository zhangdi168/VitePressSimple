package vptype

// SocialLink 表示社交媒体链接，包含图标、链接以及可选的ARIA标签
type SocialLink struct {
	// Icon 社交媒体链接图标
	Icon string `json:"icon"`

	// Link 链接地址
	Link string `json:"link"`

	// AriaLabel 可选的ARIA标签，用于辅助功能描述
	AriaLabel string `json:"ariaLabel,omitempty"`
}

// SocialLinkIcon 定义了社交媒体链接图标的类型，可以是预定义的字符串常量，
// 也可以是一个自定义SVG字符串
type SocialLinkIcon string // 实际使用时需要根据应用场景转换为具体类型

// 预定义的SocialLinkIcon枚举值
const (
	DiscordIcon   SocialLinkIcon = "discord"
	FacebookIcon  SocialLinkIcon = "facebook"
	GitHubIcon    SocialLinkIcon = "github"
	InstagramIcon SocialLinkIcon = "instagram"
	LinkedInIcon  SocialLinkIcon = "linkedin"
	MastodonIcon  SocialLinkIcon = "mastodon"
	NPMIcon       SocialLinkIcon = "npm"
	SlackIcon     SocialLinkIcon = "slack"
	TwitterIcon   SocialLinkIcon = "twitter"
	XIcon         SocialLinkIcon = "x"
	YoutubeIcon   SocialLinkIcon = "youtube"
)
