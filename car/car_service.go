package car

type CarRepository interface {
	Create(car *Car) error
	GetAll() ([]Car, error)
	GetByID(id int) (*Car, error)
	Update(car *Car) error
	Delete(id int) error
}

type CarService struct {
	repo CarRepository
}

func NewCarService(repo CarRepository) *CarService {
	return &CarService{repo}
}

func (s *CarService) Create(car *Car) error {
	return s.repo.Create(car)
}

func (s *CarService) GetAll() ([]Car, error) {
	return s.repo.GetAll()
}

func (s *CarService) GetByID(id int) (*Car, error) {
	return s.repo.GetByID(id)
}

func (s *CarService) Update(car *Car) error {
	return s.repo.Update(car)
}

func (s *CarService) Delete(id int) error {
	return s.repo.Delete(id)
}
