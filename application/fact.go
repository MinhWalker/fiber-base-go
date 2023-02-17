package application

import (
	"fiber-base-go/config"
	"fiber-base-go/domain"
	"fiber-base-go/infrastructure/persistence"
)

// GetAllStudents return all domain.news
func GetAllStudents() ([]domain.Student, error) {
	conn, err := config.ConnectDb()
	if err != nil {
		return nil, err
	}
	sqlDB, _ := conn.DB()
	defer sqlDB.Close()

	repo := persistence.NewNewsRepositoryWithRDB(conn)

	return repo.GetAll()
}

// AddStudent saves new Student
func AddStudent(p domain.Student) error {
	conn, err := config.ConnectDb()
	if err != nil {
		return err
	}
	sqlDB, _ := conn.DB()
	defer sqlDB.Close()

	repo := persistence.NewNewsRepositoryWithRDB(conn)
	return repo.Create(&p)
}
