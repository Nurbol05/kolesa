package category

type CategoryRepository interface {
	Create(category *Category) error
	GetAll() ([]Category, error)
	Update(id int, newName string) error
	Delete(id int) error
}
type CategoryService struct {
	repo CategoryRepository
}

func NewCategoryService(repo CategoryRepository) *CategoryService {
	return &CategoryService{repo}
}

func (s *CategoryService) Create(category *Category) error {
	return s.repo.Create(category)
}

func (s *CategoryService) GetAll() ([]Category, error) {
	return s.repo.GetAll()
}

func (s *CategoryService) Update(id int, newName string) error {
	return s.repo.Update(id, newName)
}

func (s *CategoryService) Delete(id int) error {
	return s.repo.Delete(id)
}
