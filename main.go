package main

import (
	db "easy-gin/database"
	router "easy-gin/router"
)

func main() {
	defer db.SqlDB.Close()
	router.R.Run(":8000")
}
