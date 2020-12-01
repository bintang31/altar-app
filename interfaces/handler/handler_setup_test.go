package interfaces

import "altar-app/utils/mock"

var (
	userApp    mock.UserAppInterface
	roleApp    mock.RoleAppInterface
	fakeUpload mock.UploadFileInterface
	fakeAuth   mock.AuthInterface
	fakeToken  mock.TokenInterface

	s  = NewUsers(&userApp, &fakeAuth, &fakeToken) //We use all mocked data here
	rs = NewRoles(&roleApp, &fakeAuth, &fakeToken)
	au = NewAuthenticate(&userApp, &fakeAuth, &fakeToken) //We use all mocked data here

)
