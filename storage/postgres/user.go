package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"task/api/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(ctx context.Context, req *models.CreateUser) (*models.User, error) {
	var (
		userId = uuid.New().String()
		query  = `
			INSERT INTO "user"(
				"id",
				"full_name",
				"nick_name",
				"photo",
				"birthday",
				"location",
				"updated_at"
			) VALUES($1, $2, $3, $4, $5, $6, NOW())`
	)

	_, err := r.db.Exec(ctx,
		query,
		userId,
		req.FullName,
		req.NickName,
		req.Photo,
		req.Birthday,
		req.Location,
	)
	if err != nil {
		return nil, err
	}

	return r.GetByID(ctx, &models.UserPrimaryKey{Id: userId})
}

func (r *userRepo) GetByID(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error) {
	var (
		query = `
			SELECT
				"id",
				"full_name",
				"nick_name",
				"photo",
				"birthday",
				"location",
				"created_at",
				"updated_at"
			FROM "user"
			WHERE id = $1
		`
	)

	var (
		Id        sql.NullString
		FullName  sql.NullString
		NickName  sql.NullString
		Photo     sql.NullString
		Birthday  sql.NullString
		Location  sql.NullString
		CreatedAt sql.NullString
		UpdatedAt sql.NullString
	)

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&Id,
		&FullName,
		&NickName,
		&Photo,
		&Birthday,
		&Location,
		&CreatedAt,
		&UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.User{
		Id:        Id.String,
		FullName:  FullName.String,
		NickName:  NickName.String,
		Photo:     Photo.String,
		Birthday:  Birthday.String,
		Location:  Location.String,
		CreatedAt: CreatedAt.String,
		UpdatedAt: UpdatedAt.String,
	}, nil
}

func (r *userRepo) GetList(ctx context.Context, req *models.GetListUserRequest) (*models.GetListUserResponse, error) {
	var (
		resp   models.GetListUserResponse
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		sort   = " ORDER BY full_name ASC, id DESC"
	)

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	var (
		query = `
			SELECT
				COUNT(*) OVER(),
				"id",
				"full_name",
				"nick_name",
				"photo",
				"birthday",
				"location",
				"created_at",
				"updated_at"
			FROM "user"
		`
	)

	query += where + sort + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			Id        sql.NullString
			FullName  sql.NullString
			NickName  sql.NullString
			Photo     sql.NullString
			Birthday  sql.NullString
			Location  sql.NullString
			CreatedAt sql.NullString
			UpdatedAt sql.NullString
		)

		err = rows.Scan(
			&resp.Count,
			&Id,
			&FullName,
			&NickName,
			&Photo,
			&Birthday,
			&Location,
			&CreatedAt,
			&UpdatedAt,
		)

		user := models.User{
			Id:        Id.String,
			FullName:  FullName.String,
			NickName:  NickName.String,
			Photo:     Photo.String,
			Birthday:  Birthday.String,
			Location:  Location.String,
			CreatedAt: CreatedAt.String,
			UpdatedAt: UpdatedAt.String,
		}

		if err != nil {
			return nil, err
		}

		resp.Users = append(resp.Users, &user)
	}

	return &resp, nil
}

func (r *userRepo) Update(ctx context.Context, req *models.UpdateUser) (int64, error) {
	query := `
		UPDATE "user"
			SET
				"full_name" = $2,
				"nick_name" = $3,
				"photo" = $4,
				"birthday" = $5,
				"location" = $6,
				"updated_at" = NOW()
		WHERE id = $1
	`

	rowsAffected, err := r.db.Exec(ctx,
		query,
		req.Id,
		req.FullName,
		req.NickName,
		req.Photo,
		req.Birthday,
		req.Location,
	)

	if err != nil {
		return 0, err
	}
	return rowsAffected.RowsAffected(), nil
}

func (r *userRepo) Delete(ctx context.Context, req *models.UserPrimaryKey) error {
	_, err := r.db.Exec(ctx, `DELETE FROM "user" WHERE id = $1`, req.Id)
	return err
}

// func (r *userRepo) GetFields(ctx context.Context, req *models.GetListUserRequest) (resp *models.GetListUserResponse, err error) {
// 	query := fmt.Sprintf(`SELECT full_name, nick_name, %s FROM "user"`, req.Fields)
// 	rows, err := r.db.Query(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	for rows.Next() {
// 		var (
// 			FullName  sql.NullString
// 			NickName  sql.NullString
// 			Photo     sql.NullString
// 			Birthday  sql.NullString
// 			Location  sql.NullString
// 			CreatedAt sql.NullString
// 			UpdatedAt sql.NullString
// 			field     sql.NullString
// 		)
// 		switch req.Fields {
// 		case "photo":
// 			Photo = field
// 		}
// 		err = rows.Scan(
// 			&FullName,
// 			&NickName,
// 			&field,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}
// 		resp := &models.User{
// 			FullName: FullName.String,
// 			NickName: NickName.String,

// 		}
// 		return resp, nil
// 	}
// 	return nil, nil
// }
