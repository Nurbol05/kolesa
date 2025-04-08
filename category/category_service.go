package category

type CategoryRepository interface {
	Create(category *Category) error
	GetAll() ([]Category, error)
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
