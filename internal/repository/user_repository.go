package repository

import (
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2"
	"wms-api-v2/internal/types"
)

type UserRepository interface {
	FindAll(ctx *fiber.Ctx, tx *sql.Tx) ([]*types.FindUser, error)
	FindByID(ctx *fiber.Ctx, tx *sql.Tx, id int64) (*types.FindUser, error)
	Save(ctx *fiber.Ctx, tx *sql.Tx, user *types.UserRequest) (int, error)
	Update(ctx *fiber.Ctx, tx *sql.Tx, user *types.UserUpdateRequest) error
}

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (u *UserRepositoryImpl) FindAll(ctx *fiber.Ctx, tx *sql.Tx) ([]*types.FindUser, error) {
	rows, err := tx.QueryContext(ctx.Context(), "SELECT id, name, email, no_hp, hak_akses, created_at FROM users")
	if err != nil {
		return nil, errors.New("Error query process")
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	var users []*types.FindUser
	for rows.Next() {
		var user types.FindUser
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.NoHp,
			&user.HakAkses,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, errors.New("Error query process")
		}
		users = append(users, &user)
	}
	return users, nil
}

func (u *UserRepositoryImpl) FindByID(ctx *fiber.Ctx, tx *sql.Tx, id int64) (*types.FindUser, error) {
	rows, err := tx.QueryContext(ctx.Context(), "SELECT id, name, email, no_hp, hak_akses, created_at FROM users WHERE id=?", id)
	if err != nil {
		return nil, errors.New("Error query process")
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	var user types.FindUser
	if rows.Next() {
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.NoHp,
			&user.HakAkses,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, errors.New("Error query process")
		}
		return &user, nil
	} else {
		return nil, errors.New("User not found")
	}
}

func (u *UserRepositoryImpl) Save(ctx *fiber.Ctx, tx *sql.Tx, user *types.UserRequest) (int, error) {
	SQL := "INSERT INTO users (name, email, no_hp, hak_akses, password) VALUES (?,?,?,?,?)"
	result, err := tx.ExecContext(ctx.Context(), SQL, user.Name, user.Email, user.NoHp, user.HakAkses, user.Password)
	if err != nil {
		return 0, errors.New("Error query process")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, errors.New("Error saat insert user")
	}

	return int(id), nil
}

func (u *UserRepositoryImpl) Update(ctx *fiber.Ctx, tx *sql.Tx, user *types.UserUpdateRequest) error {
	SQL := "UPDATE users SET name=?, email=?, hak_akses=?, password=? WHERE id=?"
	_, err := tx.ExecContext(ctx.Context(), SQL, user.Id)
	if err != nil {
		return errors.New("Error query process")
	}

	return nil
}
