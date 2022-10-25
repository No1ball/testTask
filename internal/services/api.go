package services

type ApiService struct {
	followersService FollowersService
	templatesService TemplatesService
}

func NewApiService(followService FollowersService, templatesService TemplatesService) *ApiService {
	return &ApiService{followersService: followService, templatesService: templatesService}
}
