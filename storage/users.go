package storage

import "github.com/jmoiron/sqlx"

type (
	UsersStorage interface {
		Create(User) error
		ByID(User) (User, error)
		Update(User) error
		Delete(User) error
	}

	Users struct {
		*sqlx.DB
	}

	User struct {
		ID    int64  `json:"id"`
		State string `json:"state"`
	}
)

func (db *Users) Create(u User) error {
	const q = "INSERT INTO users (id, state) VALUES ($1, $2)"
	_, err := db.Exec(q, u.ID, u.State)

	return err

}

func (db *Users) ByID(u User) (user User, _ error) {
	const q = "SELECT * FROM users WHERE id = $1"
	return user, db.Get(&user, q, u.ID)
}

func (db *Users) Update(u User) error {
	const q = "UPDATE users SET state = $1 WHERE id = $2"
	_, err := db.Exec(q, u.State, u.ID)
	return err
}

func (db *Users) Delete(u User) error {
	const q = "DELETE FROM users WHERE id = $1"
	_, err := db.Exec(q, u.ID)
	return err
}
