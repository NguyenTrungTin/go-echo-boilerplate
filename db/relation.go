package db

func Relation() {

	db.Model(&PasswordReset{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

}
