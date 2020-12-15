package services

// User user info
type User struct {
	Nickname        string  `json:"nickname"`
	Avatar          string  `json:"avatar"`
	TodayStudyTime  int     `json:"today_study_time"`
	StudySerialDays int     `json:"study_serial_days"`
	ISV             int     `JSON:"is_v"`
	VIPUser         VIPUser `json:"vip_user"`
	ISTeacher       int     `json:"is_teacher"`
	UIDHazy         string  `json:"uid_hazy"`
}

// VIPUser vip info
type VIPUser struct {
	Info string `json:"info"`
	Stat int    `json:"stat"`
}
