package mysql

import "gorm.io/gorm"

type (
	RepositoryOutlet struct {
		*gorm.DB
	}
)

func NewOutletRepo(db *gorm.DB) *RepositoryOutlet {
	return &RepositoryOutlet{db}
}

func (r *RepositoryOutlet) Get(ID string) {

}
