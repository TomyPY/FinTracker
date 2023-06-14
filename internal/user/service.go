package user

type service struct {
	r Repository
}

type Service interface {
	GetAll() ([]User, error)
	Get(id int) (User, error)
	Create(user User) (int, error)
	Delete(id int) error
	Update(id int, user User) (int, error)
}

func NewService(r Repository) Service {
	return &service{
		r: r,
	}
}

func (s *service) GetAll() ([]User, error) {
	return nil, nil
}

func (s *service) Get(id int) (User, error) {
	return User{}, nil
}

func (s *service) Create(user User) (int, error) {
	return 0, nil
}

func (s *service) Delete(id int) error {
	return nil
}

func (s *service) Update(id int, user User) (int, error) {
	return 0, nil
}
