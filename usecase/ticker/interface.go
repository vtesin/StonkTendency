package ticker

import "github.com/vtesin/StonkTendency/entity"

//Reader interface
type Reader interface {
	Get(symbol string) (*entity.Ticker, error)
	Search(query string) ([]*entity.Ticker, error)
	List() ([]*entity.Ticker, error)
}

//Writer user writer
type Writer interface {
	Create(t *entity.Ticker) error
	Update(t *entity.Ticker) error
	Delete(symbol string) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	GetStonk(symbol string) (*entity.Ticker, error)
	SearchStonks(query string) ([]*entity.Ticker, error)
	ListStonks() ([]*entity.Ticker, error)
	CreateStonk(symbol string, sentiment float32) error
	UpdateStonk(t *entity.Ticker) error
	DeleteStonk(symbol string) error
}
