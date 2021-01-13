package projectpsqlrepository

import (
	"database/sql"
	"fmt"

	"github.com/P44elovod/task-management-app/domain"
)

type psqlProjectRepository struct {
	Conn *sql.DB
}

func NewPsqlProjectleRepository(Conn *sql.DB) domain.ProjectRepository {
	return &psqlProjectRepository{Conn}
}

func (p *psqlProjectRepository) Fetch() {
	fmt.Println("psqlProjectRepository fetch")
}
