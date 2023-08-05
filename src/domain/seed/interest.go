package seed

import (
	"github.com/abisatria1/simple-dating-app/src/domain/entity"
	"gorm.io/gorm"
)

type SeedManager interface {
	DoSeeding() error
	Register(SeederFunc Seeder)
}

type Options struct {
	DB *gorm.DB
}

type GormSeeder struct {
	DB      *gorm.DB
	Seeders []Seeder
}

type Seeder func(db *gorm.DB) error

func NewGormSeeder(o *Options) SeedManager {
	mo := &GormSeeder{
		DB:      o.DB,
		Seeders: []Seeder{},
	}

	mo.Register(InterestSeeder)

	return mo
}

func (s *GormSeeder) DoSeeding() (err error) {
	for _, f := range s.Seeders {
		err = f(s.DB)
		if err != nil {
			return
		}
	}
	return
}

func (s *GormSeeder) Register(f Seeder) {
	s.Seeders = append(s.Seeders, f)
}

func InterestSeeder(db *gorm.DB) (err error) {
	interests := []entity.Interest{
		{
			Name: "Dog",
		},
		{
			Name: "Cat",
		},
		{
			Name: "Sport",
		},
		{
			Name: "Drink",
		},
		{
			Name: "Travel",
		},
		{
			Name: "Badminton",
		},
		{
			Name: "Soccer",
		},
		{
			Name: "Makeup",
		},
		{
			Name: "Photography",
		},
	}
	err = db.Create(&interests).Error
	return
}
