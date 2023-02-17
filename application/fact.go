package application

import (
	"fiber-base-go/domain"
	"fiber-base-go/infrastructure/persistence"
	"gorm.io/gorm"
)

// GetAllStudents return all domain.news
func GetAllStudents(conn *gorm.DB) ([]domain.Student, error) {
	sqlDB, _ := conn.DB()
	defer sqlDB.Close()

	repo := persistence.NewNewsRepositoryWithRDB(conn)

	return repo.GetAll()
}

// AddStudent saves new Student
func AddStudent(conn *gorm.DB, p domain.Student) error {
	sqlDB, _ := conn.DB()
	defer sqlDB.Close()

	repo := persistence.NewNewsRepositoryWithRDB(conn)
	return repo.Create(&p)
}
