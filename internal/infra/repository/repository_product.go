package repository

import (
	"database/sql"

	"gituhb.com/gothello/live-full-cycle-kafka/internal/entity"
)

type RepositoryProductMySql struct {
	DB *sql.DB
}

func NewRespositoryProductMySql(db *sql.DB) *RepositoryProductMySql {
	return &RepositoryProductMySql{
		DB: db,
	}
}

func (m *RepositoryProductMySql) GetAll() ([]*entity.Product, error) {

	rows, err := m.DB.Query("select id, name, price from products")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []*entity.Product

	for rows.Next() {
		var prc entity.Product

		if err := rows.Scan(&prc.ID, &prc.Name, &prc.Price); err != nil {
			return res, err
		}

		res = append(res, &prc)
	}

	return res, nil
}

func (m *RepositoryProductMySql) Insert(prc *entity.Product) error {
	_, err := m.DB.Exec("insert into products(id, name, price) values  (?,?,?)",
		prc.ID, prc.Name, prc.Price)
	if err != nil {
		return err
	}

	return nil
}
