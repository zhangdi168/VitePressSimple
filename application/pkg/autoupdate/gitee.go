package autoupdate

import "time"

//https://gitee.com/api/v5/repos/zhangdi168/zhijian-chat/releases

type GiteeRelease []struct {
	ID              int64     `json:"id"`
	TagName         string    `json:"tag_name"`
	TargetCommitish string    `json:"target_commitish"`
	Prelease        bool      `json:"prerelease"`
	Name            string    `json:"name"`
	Body            string    `json:"body"`
	Author          Author    `json:"author"`
	CreatedAt       time.Time `json:"created_at"`
	Assets          []Asset   `json:"assets"`
}

type Author struct {
	ID                int64  `json:"id"`
	Login             string `json:"login"`
	Name              string `json:"name"`
	AvatarURL         string `json:"avatar_url"`
	URL               string `json:"url"`
	HtmlURL           string `json:"html_url"`
	Remark            string `json:"remark"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
}

type Asset struct {
	BrowserDownloadURL string `json:"browser_download_url"`
	// 根据实际JSON内容添加额外字段（例如上面的示例中包含"name"字段）
	//xxx.zip
	Name string `json:"name,omitempty"`
}
