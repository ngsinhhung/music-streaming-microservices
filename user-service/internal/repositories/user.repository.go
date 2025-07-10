package repositories

type IUserRepository interface {
	GetUserByEmail(email string) (bool, error)
	GetUserById(id int) (bool, error)
}

type userRepository struct {
}

func (u userRepository) GetUserByEmail(email string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) GetUserById(id int) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}
