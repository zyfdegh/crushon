package types

// SearchUserResp is body of search result
// POST http://music.163.com/api/search/get?s=nickname&limit=10&type=1002
type SearchUserResp struct {
	Result Result `json:"result"`
	Code   int    `json:"code"`
}

type Result struct {
	UserProfileCount int           `json:"userprofileCount"`
	UserProfiles     []UserProfile `json:"userprofiles"`
}

type ErrResp struct {
	Msg  string `json:"msg"`
	Code int    `json:"int"`
}
