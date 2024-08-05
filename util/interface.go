package util

// TODO(bruce): delete file, just for practicing

type Store interface {
	GetUser() (string, error)

	GetValue() (float64, error)
}

type StoreImpl1 struct {
	user  string
	value float64
}

func NewStoreImpl1(user string, value float64) StoreImpl1 {
	return StoreImpl1{
		user:  user,
		value: value,
	}
}

func (s StoreImpl1) GetUser() (string, error) {
	return s.user, nil
}

func (s StoreImpl1) GetValue() (float64, error) {
	return s.value, nil
}

type StoreImpl2 bool

func (s StoreImpl2) GetUser() (string, error) {
	if s {
		return "yes", nil
	}

	return "n0", nil
}

func (s StoreImpl2) GetValue() (float64, error) {
	if s {
		return 1, nil
	}

	return -1, nil
}

type PointStore interface {
	GetPoint() (float64, error)

	CreatePoint() error
}

type UserStore interface {
	GetUser() (string, error)

	CreateUser() error
}

type MultiStore interface {
	PointStore
	UserStore
}
