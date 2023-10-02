package sensorinfo

//model related

// import "sample/pkg/repos/interfaces"

// "sample/pkg/base"

// "sample/pkg/module/token"
// "sample/pkg/repos/interfaces"
// "sample/pkg/repos/models"

type SensorConfigService struct {
	// base.Service
	// SensorConfigRepo interfaces.SensorConfigRepo
	// TokenService *token.TokenService
}

// func NewSensorConfigService(SensorConfigRepo interfaces.SensorConfigRepo) *SensorConfigService {
// 	var SensorConfigService SensorConfigService
// 	SensorConfigService.SensorConfigRepo = SensorConfigRepo
// 	// SensorConfigService.TokenService = tokenService
// 	return &SensorConfigService
// }

func NewSensorConfigService() *SensorConfigService {
	var SensorConfigService SensorConfigService
	// SensorConfigService.SensorConfigRepo = SensorConfigRepo
	// SensorConfigService.TokenService = tokenService
	return &SensorConfigService
}

func (s *SensorConfigService) GetServiceType() ([]string, error) {
	serviceType := []string{"ViewLink"}
	return serviceType, nil
}

// func (s *SensorConfigService) GetSensorList() (sensorconfig []models.SensorConfigModel, err error) {
// 	sensorConfigModel, err := s.SensorConfigRepo.GetSensorList()
// 	if err != nil {
// 		return sensorConfigModel, err
// 	}
// 	return sensorConfigModel, nil
// }

// func (s *SensorConfigService) GetBuildingList() (bool, err error) {
// 	// if err := dto.Check(); err != nil {
// 	// 	return "", s.InvalidArgument(err.Error())
// 	// }
// 	// SensorConfigModel, err := s.SensorConfigRepo.GetSensorConfigByAccount(dto.Account)

// 	// if err != nil {
// 	// 	return "", err
// 	// }

// 	// hashedPassword := h.GetSHA256HashCode(dto.Password)
// 	// if hashedPassword != SensorConfigModel.Password {
// 	// 	return "", s.InvalidArgument("password_not_match")
// 	// }
// 	// token, err = s.TokenService.Create(SensorConfigModel.SensorConfigId, SensorConfigModel.Name)
// 	// if err != nil {
// 	// 	return "", err
// 	// }
// 	return true, nil
// }

// func (s *SensorConfigService) GetSensorConfig() (err error) {
// 	// if err := dto.Check(); err != nil {
// 	// 	return s.InvalidArgument(err.Error())
// 	// }
// 	// if dto.ConfirmPassword != dto.NewPassword {
// 	// 	return s.InvalidArgument("confirm_password_not_match")
// 	// }

// 	// SensorConfigId := dto.SensorConfigId
// 	// SensorConfigModel, err := s.SensorConfigRepo.GetSensorConfig(SensorConfigId)

// 	// if err != nil {
// 	// 	return err
// 	// }
// 	// if SensorConfigModel.Id == 0 {
// 	// 	return s.InvalidArgument("no_matched_account")
// 	// }

// 	// if !s.confirmOldPwd(&dto.OldPassword, &SensorConfigModel.Password) {
// 	// 	return s.InvalidArgument("old_password_not_match")
// 	// }

// 	// hashNewPwd := h.GetSHA256HashCode(dto.NewPassword)

// 	// if err := s.SensorConfigRepo.ChangePwd(SensorConfigId, hashNewPwd); err != nil {
// 	// 	return true
// 	// }
// 	return nil
// }

// func (s *SensorConfigService) CreateSensorConfig() bool {
// 	// hashPassword := h.GetSHA256HashCode(*pwd)
// 	return true
// }

// func (s *SensorConfigService) DeleteSensorConfig() bool {
// 	// hashPassword := h.GetSHA256HashCode(*pwd)
// 	return true
// }

// func (s *SensorConfigService) UpdateSensorConfig() bool {

// 	return true
// }
