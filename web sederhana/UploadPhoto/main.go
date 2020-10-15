package main

import (
    "database/sql"
    "io"
    "net/http"
    "os"

    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    _ "github.com/go-sql-driver/mysql"
    //pusher "github.com/pusher/pusher-http-go"
)

type Photo struct {
	ID  int64  `json:"id"`
	Src string `json:"src"`
}

type PhotoCollection struct {
	Photos []Photo `json:"items"`
}

// var client = pusher.Client{
// 	AppId:   "PUSHER_APP_ID",
// 	Key:     "PUSHER_APP_KEY",
// 	Secret:  "PUSHER_APP_SECRET",
// 	Cluster: "PUSHER_APP_CLUSTER",
// 	Secure:  true,
// }

func main() {
	db := dbConn()
	//migrateDatabase(db)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.File("/", "public/index.html")
	e.GET("/photos", getPhotos(db))
	e.POST("/photos", uploadPhoto(db))
	e.Static("/uploads", "public/uploads")

	e.Logger.Fatal(e.Start(":9000"))
}

func dbConn() (db *sql.DB) {

    dbDriver := "mysql"
    dbUser := "root"
    dbPass := ""
    dbName := "yii2basic"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())
    }
	return db
}

// func migrateDatabase(db *sql.DB) {
// 	sql := "CREATE TABLE IF NOT EXISTS photos(id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, src VARCHAR NOT NULL)"

// 	_, err := db.Exec(sql)
// 	if err != nil {
// 		panic(err)
// 	}
// }

func getPhotos(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		rows, err := db.Query("SELECT * FROM photos")
		if err != nil {
			panic(err)
		}

		defer rows.Close()

		result := PhotoCollection{}

		for rows.Next() {
			photo := Photo{}

			err2 := rows.Scan(&photo.ID, &photo.Src)
			if err2 != nil {
				panic(err2)
			}

			result.Photos = append(result.Photos, photo)
		}

		return c.JSON(http.StatusOK, result)
	}
}

func uploadPhoto(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}

		src, err := file.Open()
		if err != nil {
			return err
		}

		defer src.Close()

		filePath := "./public/Uploads/" + file.Filename
		fileSrc := "http://127.0.0.1:9000/uploads/" + file.Filename

		dst, err := os.Create(filePath)
		if err != nil {
			panic(err)
		}

		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			panic(err)
		}

		stmt, err := db.Prepare("INSERT INTO photos (src) VALUES(?)")
		if err != nil {
			panic(err)
		}

		defer stmt.Close()

		result, err := stmt.Exec(fileSrc)
		if err != nil {
			panic(err)
		}

		insertedId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		photo := Photo{
			Src: fileSrc,
			ID:  insertedId,
		}

		//client.Trigger("photo-stream", "new-photo", photo)
		return c.JSON(http.StatusOK, photo)
	}
}