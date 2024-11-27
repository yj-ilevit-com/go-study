package main

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected!")

	// Retrieve albums by John Coltrane
	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Albums found: %v\n", albums)

        albID, err := addAlbum(Album{ Title: "The Modern Sound of Betty Carter", Artist: "Betty Carter", Price: 49.99, }) 
        if err != nil { 
                log.Fatal(err) 
        } 

        fmt.Printf("ID of added album: %v\n", albID)
}

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func albumsByArtist(name string) ([]Album, error) {
	var albums []Album
	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()

	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	return albums, nil
}

func addAlbum(alb Album) (int64, error) { 
        result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price) 
        if err != nil { 
                return 0, fmt.Errorf("addAlbum: %v", err) 
        } 
        id, err := result.LastInsertId() 

        if err != nil { 
                return 0, fmt.Errorf("addAlbum: %v", err) 
        } 

        return id, nil 
}
