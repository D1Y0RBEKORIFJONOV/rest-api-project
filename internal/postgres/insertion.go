package postgres


func (db *DB)UserInsertInto(first_name,last_name,email,password string) (id int,err error){
	err = db.ConnectDB()
	if err != nil {
		return 0,err
	}

	query := `
	INSERT INTO users(first_name,last_name,email,password) 
	VALUES($1,$2,$3,$4) 
	RETURNING user_id ;
	`
	err = db.DB.QueryRow(query,first_name,last_name,email,password).Scan(&id)
	if err != nil {
		return 0,err
	}

	return id,err
}