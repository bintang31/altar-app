package repository

import (
	"altar-app/domain/entity"
)

//PelangganRepository : Pelanggan collection of methods that the infrastructure
type PelangganRepository interface {
	GetPelanggans() ([]entity.Pelanggan, error)
}
