package create_account

type Account struct {
	ID       uint64 `db:"id"`
	Email    string `db:"email"`
	Name     string `db:"name"`
	Password string `db:"password"`
}
