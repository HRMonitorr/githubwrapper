package githubwrapper

type Commit struct {
	Type  string `json:"type"`
	Items struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Type        string `json:"type"`
		Properties  struct {
			URL         string `json:"url"`
			SHA         string `json:"sha"`
			NodeID      string `json:"node_id"`
			HTMLURL     string `json:"html_url"`
			CommentsURL string `json:"comments_url"`
			Commit      struct {
				URL    string `json:"url"`
				Author struct {
					Name  string `json:"name"`
					Email string `json:"email"`
					Date  string `json:"date"`
				} `json:"author"`
				Committer struct {
					Name  string `json:"name"`
					Email string `json:"email"`
					Date  string `json:"date"`
				} `json:"committer"`
				Message      string `json:"message"`
				CommentCount int    `json:"comment_count"`
				Tree         struct {
					SHA string `json:"sha"`
					URL string `json:"url"`
				} `json:"tree"`
				Verification struct {
					Verified  bool   `json:"verified"`
					Reason    string `json:"reason"`
					Payload   string `json:"payload"`
					Signature string `json:"signature"`
				} `json:"verification"`
			} `json:"commit"`
			Author struct {
				Name              string `json:"name"`
				Email             string `json:"email"`
				Login             string `json:"login"`
				ID                int    `json:"id"`
				NodeID            string `json:"node_id"`
				AvatarURL         string `json:"avatar_url"`
				GravatarID        string `json:"gravatar_id"`
				URL               string `json:"url"`
				HTMLURL           string `json:"html_url"`
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
				SiteAdmin         bool   `json:"site_admin"`
				StarredAt         string `json:"starred_at"`
			} `json:"author"`
			Committer struct {
				Name              string `json:"name"`
				Email             string `json:"email"`
				Login             string `json:"login"`
				ID                int    `json:"id"`
				NodeID            string `json:"node_id"`
				AvatarURL         string `json:"avatar_url"`
				GravatarID        string `json:"gravatar_id"`
				URL               string `json:"url"`
				HTMLURL           string `json:"html_url"`
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
				SiteAdmin         bool   `json:"site_admin"`
				StarredAt         string `json:"starred_at"`
			} `json:"committer"`
			Parents struct {
				Type  string `json:"type"`
				Items struct {
					Sha     string `json:"sha"`
					URL     string `json:"url"`
					HTMLURL string `json:"html_url"`
				} `json:"items"`
			} `json:"parents"`
			Stats struct {
				Type       string `json:"type"`
				Properties struct {
					Additions int `json:"additions"`
					Deletions int `json:"deletions"`
					Total     int `json:"total"`
				} `json:"properties"`
			} `json:"stats"`
			Files struct {
				Type  string `json:"type"`
				Items struct {
					Title       string `json:"title"`
					Description string `json:"description"`
					Type        string `json:"type"`
					Properties  struct {
						SHA              string `json:"sha"`
						Filename         string `json:"filename"`
						Status           string `json:"status"`
						Additions        int    `json:"additions"`
						Deletions        int    `json:"deletions"`
						Changes          int    `json:"changes"`
						BlobURL          string `json:"blob_url"`
						RawURL           string `json:"raw_url"`
						ContentsURL      string `json:"contents_url"`
						Patch            string `json:"patch"`
						PreviousFilename string `json:"previous_filename"`
					} `json:"properties"`
				} `json:"items"`
			} `json:"files"`
		} `json:"items"`
	} `json:"items"`
}

type Repositories struct {
	Name     *string `json:"name"`
	FullName *string `json:"full-name"`
	Homepage *string `json:"homepage"`
}
