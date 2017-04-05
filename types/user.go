package types

// UserProfile is brief info of a user
// POST http://music.163.com/api/search/get?s=nickname&limit=100&type=1002
type UserProfile struct {
	AvatarUrl string `json:"avatarUrl"`
	Nickname  string `json:"nickname"`
	UserID    int    `json:"userId"`
	Followeds int    `json:"followeds"`
	Follows   int    `json:"follows"`
}
