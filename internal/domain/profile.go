package domain

type StudentProfile struct {
	SlackUserName string `json:"slackUsername"`
	Backend       bool   `json:"backend"`
	Age           int64  `json:"age"`
	Bio           string `json:"bio"`
}
