package entity

type Notification struct {
	Id int `json:"id"`
	Type int `json:"type"`
	GroupId int `json:"group_id"`
	SenderId int `json:"sender_id"`
	ReceiverID int `json:"receiver_id"`
	EventID int `json:"event_id"`
	Message string `json:"message"`
	Accepted bool `json:"accepted"`
}

const (
	//notification answer
	NotificationWithoutAnswer =  iota
	NotificationAccepted 
	NotificationNotAccepted 
)

const (
		//type of notification
		FollowingNotification = iota
		EventNotification 
		GroupInvitationNotification 
		GroupParticipationNotification 
)