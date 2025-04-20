package entity

type Groups []Group

type Group struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        int    `json:"type"`  // 0 - real group, 1 - chat group ,  2 - fake group
	Admin       int    `json:"admin"`
	MemberCount int    `json:"member_count"`
	PostCount   int    `json:"post_count"`
	Members     []int  `json:"members"`
	Posts       []Post `json:"posts"`
}

type Invitation struct {
	GroupId int `json:"group_id"`
	InviterID int `json:"inviter_id"`
	InvitedID int `json:"invited_id"`
}

const (
	RealGroup     = 0
	ChatGroup     = 1
	FakeGroup     = 2
)