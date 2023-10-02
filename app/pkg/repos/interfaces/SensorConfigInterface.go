package interfaces

type SensorConfigRepo interface {
	// 	GetAll() (member []models.MemberModel, err error)

	// 	Create(member models.MemberModel) (err error)

	// 	GetMember(memberId string) (member models.MemberModel, err error)

	// 	GetMemberByAccount(account string) (member models.MemberModel, err error)

	// 	ChangePwd(memberId string, newPwd string) (err error)

	// 	IsEmailExist(email string) bool
	// }
	GetServiceType()
	GetSensorList()
	GetBuildingList()
	GetSensorConfig()
	CreateSensorConfig()
	DeleteSensorConfig()
	UpdateSensorConfig()
}
