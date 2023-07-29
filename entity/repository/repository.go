package repositoryInterface

type Repository[T interface{}] interface {
	Save(obj *T) error
	FindBy(nameField string, value []interface{}) (*T, error)
	FindById(id uint) (*T, error)
	FindAll() ([]*T, error)
	Delete(obj *T) error
	DeleteBy(nameField string, value []interface{}) error
	DeleteById(id uint) error
}
