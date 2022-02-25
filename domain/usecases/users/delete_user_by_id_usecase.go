package users

type DeleteUserByIDUsecase interface {
	Delete(ID string) error
}
