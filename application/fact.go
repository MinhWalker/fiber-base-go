package application

import (
	"fiber-base-go/config"
	"fiber-base-go/domain"
	"fiber-base-go/infrastructure/persistence"
)

// GetAllFacts return all domain.news
func GetAllFacts() ([]domain.Fact, error) {
	conn, err := config.ConnectDb()
	if err != nil {
		return nil, err
	}
	sqlDB, _ := conn.DB()
	defer sqlDB.Close()

	repo := persistence.NewNewsRepositoryWithRDB(conn)

	return repo.GetAll()
}

// AddFact saves new Fact
func AddFact(p domain.Fact) error {
	conn, err := config.ConnectDb()
	if err != nil {
		return err
	}
	sqlDB, _ := conn.DB()
	defer sqlDB.Close()

	repo := persistence.NewNewsRepositoryWithRDB(conn)
	return repo.Create(&p)
}
