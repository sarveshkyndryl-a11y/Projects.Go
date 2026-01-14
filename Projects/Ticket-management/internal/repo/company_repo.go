package repo

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"ticket/internal/db"
	"ticket/internal/models"
)



func CreateCompany(ctx context.Context, company *models.Company) error {
	query := `
		INSERT INTO companies (
			name, address, contact_email, contact_phone
		)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at
	`

	row := db.DB.QueryRow(ctx, query,
		company.Name,
		company.Address,
		company.Contact_email,
		company.Contact_phone,
	)

return row.Scan(&company.ID, &company.CreatedAt)

}


func GetAllCompanies(ctx context.Context) ([]models.Company, error) {
	query := `
		SELECT 
			id, name, address, contact_email, contact_phone, created_at
		FROM companies
		ORDER BY created_at DESC
	`

	rows, err := db.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var companies []models.Company

	for rows.Next() {
		var c models.Company
		if err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.Address,
			&c.Contact_email,
			&c.Contact_phone,
			&c.CreatedAt,
		); err != nil {
			return nil, err
		}
		companies = append(companies, c)
	}

	return companies, nil
}

func DeleteCompany(ctx context.Context, id string) error {
	_, err := db.DB.Exec(
		ctx,
		`DELETE FROM companies WHERE id = $1`,
		id,
	)
	return err
}

func PatchCompany(ctx context.Context, id string, updates map[string]any) error {
	allowed := map[string]bool{
		"name":          true,
		"address":       true,
		"contact_email": true,
		"contact_phone": true,
	}

	set := []string{}
	args := []any{}
	i := 1

	for field, value := range updates {
		if allowed[field] {
			set = append(set, fmt.Sprintf("%s = $%d", field, i))
			args = append(args, value)
			i++
		}
	}

	if len(set) == 0 {
		return errors.New("no valid fields to update")
	}

	query := fmt.Sprintf(
		"UPDATE companies SET %s WHERE id = $%d",
		strings.Join(set, ", "),
		i,
	)

	args = append(args, id)

	_, err := db.DB.Exec(ctx, query, args...)
	return err
}

func UpdateCompany(ctx context.Context, c *models.Company) error {
	query := `
		UPDATE companies
		SET name = $1,
		    address = $2,
		    contact_email = $3,
		    contact_phone = $4
		WHERE id = $5
	`
	_, err := db.DB.Exec(
		ctx,
		query,
		c.Name,
		c.Address,
		c.Contact_email,
		c.Contact_phone,
		c.ID,
	)
	return err
}
