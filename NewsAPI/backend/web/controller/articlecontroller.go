package mahasiswa

import (
    "context"
    "fmt"
    "SPENEWS/backend/config"
    "SPENEWS/backend/models"
    "log"
    "errors"
    "database/sql"
    "strconv"
)
 
const   table = "berita"
// ambil semua data dari database
func GetAll(ctx context.Context) ([]models.Article, error) {
 
    var articles []models.Article
 
    db, err := config.Mysql()
 
    if err != nil {
        log.Fatal("Gagal terhubung ke MySQL", err)
    }
 
    queryText := fmt.Sprintf("SELECT * FROM %v", table)
 
    rowQuery, err := db.QueryContext(ctx, queryText)
 
    if err != nil {
        log.Fatal(err)
    }
 
    for rowQuery.Next() {
        var article models.Article
 
        if err = rowQuery.Scan(
			&article.Id,
            &article.Nama,
            &article.Nim,
			&article.Email);
			err != nil && sql.ErrNoRows != nil{
            return nil, err
        } 
        articles = append(mahasiswas, mahasiswa)
    }
 
    return articles, nil
}

// get selected
func GetSelected(ctx context.Context) ([]models.Article, error) {
 
    var articles []models.Article
 
    db, err := config.Mysql()
 
    if err != nil {
        log.Fatal("Gagal terhubung ke MySQL", err)
    }
 
    queryText := fmt.Sprintf("SELECT * FROM %v", table)
 
    rowQuery, err := db.QueryContext(ctx, queryText)
 
    if err != nil {
        log.Fatal(err)
    }
 
    for rowQuery.Next() {
        var article models.Article
 
        if err = rowQuery.Scan(
			&article.Id,
            &article.Nama,
            &article.Nim,
			&article.Email);
			err != nil && sql.ErrNoRows != nil{
            return nil, err
        } 
        articles = append(mahasiswas, mahasiswa)
    }
 
    return articles, nil
}