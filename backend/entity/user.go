package entity

type User struct {
	ID             uint   `json:"id,omitempty"`
	Nickname       string `json:"nickname,omitempty"`
	Email          string `json:"email,omitempty"`
	Password       string `json:"password,omitempty"`
	Avatar         []byte `json:"avatar,omitempty"`
	First          string `json:"first,omitempty"`
	Last           string `json:"last,omitempty"`
	DateOfBirth    string `json:"date_of_birth,omitempty"`
	AboutMe        string `json:"about_me,omitempty"`
	Status         uint   `json:"status,omitempty"`
	ProfileOwner   bool   `json:"profile_owner,omitempty"`
	FollowingState uint   `json:"following,omitempty"` // 1 follower, 2 following, 0 none
	FollowersCount uint   `json:"followers_count,omitempty"`
	FollowingCount uint   `json:"following_count,omitempty"`
}

type Follows struct {
	Followers []User
	Following []User
}

type Credentials struct {
	Username string `json:"nickname"`
	Password string `json:"password"`
}

const (
	PrivateUser = 0
	PublicUser  = 1
)
