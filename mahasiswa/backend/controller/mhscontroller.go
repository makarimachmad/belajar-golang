package mahasiswa

import (
    "context"
    "fmt"
    "belajar-golang/mahasiswa/backend/config"
    "belajar-golang/mahasiswa/backend/models"
    "log"
    "errors"
    "database/sql"
)
 
const   table = "biodata"
// ambil semua data dari database
func GetAll(ctx context.Context) ([]models.Mahasiswa, error) {
 
    var mahasiswas []models.Mahasiswa
 
    db, err := config.Mysql()
 
    if err != nil {
        log.Fatal("Cant connect to MySQL", err)
    }
 
    queryText := fmt.Sprintf("SELECT * FROM %v", table)
 
    rowQuery, err := db.QueryContext(ctx, queryText)
 
    if err != nil {
        log.Fatal(err)
    }
 
    for rowQuery.Next() {
        var mahasiswa models.Mahasiswa
 
        if err = rowQuery.Scan(
			&mahasiswa.Id,
            &mahasiswa.Nama,
            &mahasiswa.Nim,
			&mahasiswa.Email);
			err != nil && sql.ErrNoRows != nil{
            return nil, err
        } 
        mahasiswas = append(mahasiswas, mahasiswa)
    }
 
    return mahasiswas, nil
}

func Detail(ctx context.Context, mhs models.Mahasiswa) ([]models.Mahasiswa, error){
    
    var mahasiswas []models.Mahasiswa
 
    db, err := config.Mysql()
    
    if err != nil {
        log.Fatal("Gagal terhubung dengan MySQL lur..", err)
    }
    
    queryText := fmt.Sprintf("SELECT * FROM %v WHERE id = '%d'", table, mhs.Id)
 
    rowQuery, err := db.QueryContext(ctx, queryText)
    
    if err != nil {
        log.Fatal(err)
    }
    
    for rowQuery.Next() {
        var mahasiswa models.Mahasiswa

        if err = rowQuery.Scan(
            &mahasiswa.Id,
            &mahasiswa.Nama,
            &mahasiswa.Nim,
            &mahasiswa.Email);
            err != nil && sql.ErrNoRows != nil {
            return nil, err
        }
        mahasiswas = append(mahasiswas, mahasiswa)
    }
 
    return mahasiswas, nil
}

// menambahkan data ke database
func Tambah(ctx context.Context, mhs models.Mahasiswa) error {
	db, err := config.Mysql()

	if err != nil {
		log.Fatal("gagal terhubung ke MySQL", err)
	}

    queryText := fmt.Sprintf("INSERT INTO %v (nama, nim, email) values('%v','%v','%v')",
        table,
		mhs.Nama,
		mhs.Nim,
        mhs.Email)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// memperbarui data
func Update(ctx context.Context, mhs models.Mahasiswa) error {
 
    db, err := config.Mysql()
 
    if err != nil {
        log.Fatal("Gagal terhubung MySQL", err)
    }

    queryText := fmt.Sprintf("UPDATE %v set nama = '%s', nim ='%s', email = '%s' where id = '%d'",
        table,
        mhs.Nama,
        mhs.Nim,
        mhs.Email,
        mhs.Id,
    )
    fmt.Println(queryText)
 
    _, err = db.ExecContext(ctx, queryText)
 
    if err != nil {
        return err
    }
 
    return nil
}

// hapus data
func Hapus(ctx context.Context, mhs models.Mahasiswa) error {
 
    db, err := config.Mysql()
 
    if err != nil {
        log.Fatal("gagal terhubung ke MySQL", err)
    }
 
    queryText := fmt.Sprintf("DELETE FROM %v where id = '%d'", table, mhs.Id)
 
    s, err := db.ExecContext(ctx, queryText)
 
    if err != nil && err != sql.ErrNoRows {
        return err
    }
 
    check, err := s.RowsAffected()
    fmt.Println(check)
    if check == 0 {
        return errors.New("id tidak ada gan..")
    }
 
    return nil
}