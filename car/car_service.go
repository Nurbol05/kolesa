package car

type CarRepository interface {
	Create(car *Car) error
	GetAll(params GetCarsParams) ([]Car, error) // жаңартылды
	GetByID(id int) (*Car, error)
	Update(car *Car) error
	Delete(id int) error
}

type GetCarsParams struct {
	Limit  int
	Page   int
	Filter string
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

func (s *CarService) GetAll(params GetCarsParams) ([]Car, error) {
	return s.repo.GetAll(params)
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
