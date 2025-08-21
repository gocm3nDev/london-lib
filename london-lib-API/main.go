package main

import (
	"database/sql"          // SQL veritabanı işlemleri
	"errors"                // Hata oluşturma
	"fmt"                   // Ekrana yazdırma, string formatlama
	"log"                   // Loglama
	"london-lib-API/models" // Projedeki veri modelleri
	"net/http"              // HTTP işlemleri
	"strconv"               // String - sayı dönüşümleri

	"github.com/gin-contrib/cors" // CORS ayarları
	"github.com/gin-gonic/gin"    // Web/API framework
	_ "github.com/lib/pq"         // PostgreSQL sürücüsü
)

func ConnectDB() *sql.DB {
	const ( // Bağlantı parametreleri
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "427973aefe@root"
		dbname   = "dbLondonLib"
	)

	// Parametreleri string olarak gönder
	psqlInfo := fmt.Sprintf("host= %s port= %d user= %s password= %s dbname= %s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo) // Verilen parametreler ile bağlanmaya çalış
	if err != nil {                           // Bağlanamazsa hata döndür
		log.Fatal("Connection failed")
	}

	err = db.Ping() // Bağlantının doğru şekilde sağlanıp sağlanmadığını kontrol etmek içi ping at
	if err != nil {
		log.Fatal("Ping failed")
	}

	fmt.Println("Connection established successfully to postgres.")

	return db
}

func getAllAuthors(ctx *gin.Context) {
	db := ConnectDB()
	defer db.Close()

	rows, err := db.Query(`SELECT "ID", "name" FROM public."Authors"`)

	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()

	var authors []models.Author

	for rows.Next() {
		var author models.Author
		err = rows.Scan(&author.ID, &author.Name)

		if err != nil {
			fmt.Println("an error throwed while reading rows")
			continue
		}

		authors = append(authors, author)
	}

	ctx.JSON(http.StatusOK, authors)
}

func getAllBooks(ctx *gin.Context) {
	db := ConnectDB() // Veritabanına bağlan
	defer db.Close()  // Fonksiyon bitince veritabanı bağlantısını kapat

	// Veritabanına sorguyu gönder
	rows, err := db.Query(`SELECT "ID","bookName", author, quantity, description, published, page, "bookLanguage" FROM public."Books"`)

	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "500.html", gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var books []models.Book

	for rows.Next() { // Gelen verileri books dizisine ata
		var bk models.Book
		err := rows.Scan(&bk.ID, &bk.BookName, &bk.Author, &bk.Quantity, &bk.Description, &bk.Published, &bk.Page, &bk.Language)

		if err != nil {
			fmt.Println("an error throwed while reading rows")
			continue
		}

		books = append(books, bk)
	}

	ctx.JSON(http.StatusOK, books) // kitapları döndür
}

func getBookById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	ID, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "500.html", gin.H{"message": "invalid id input"})
		return
	}

	db := ConnectDB()
	defer db.Close()

	var bk models.Book

	err = db.QueryRow(`SELECT "ID","bookName", author, quantity, description, published, page, "bookLanguage" FROM public."Books" WHERE "ID" = $1`, ID).Scan(&bk.ID, &bk.BookName, &bk.Author, &bk.Quantity, &bk.Description, &bk.Published, &bk.Page, &bk.Language)

	if errors.Is(err, sql.ErrNoRows) {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "no matching user"})
		return
	} else if err != nil {
		ctx.HTML(500, "500.html", gin.H{"message": "sql query error"})
		return
	}

	ctx.JSON(http.StatusOK, bk)
}

func removeBookById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	ID, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "500.html", gin.H{"message": "invalid id input"})
		return
	}

	db := ConnectDB()
	defer db.Close()

	row, err := db.Exec(`DELETE FROM public."Books" WHERE ID=$1`, ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "delete failed"})
		return
	}

	rowsAffected, err := row.RowsAffected()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not get affected rows"})
		return
	}

	if rowsAffected > 0 {
		ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%d row(s) affected", rowsAffected)})
	} else if rowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "no record found to delete"})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "customer deleted successfully"})
}

func insertLanguage(ctx *gin.Context) {
	db := ConnectDB()
	defer db.Close()

	var language models.Language

	err := ctx.ShouldBind(&language)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "500.html", gin.H{"message": err.Error()})
		return
	}

	fmt.Printf("%v\n", language.Name)

	result, err := db.Exec(`Insert INTO public."Languages" ("name") values ($1)`, language.Name)

	if err != nil {
		log.Println("Insert Error", err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	fmt.Println("Inserted rows", rowsAffected)

	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%d row(s) inserted", rowsAffected)})
}

func insertAuthor(ctx *gin.Context) {
	db := ConnectDB()
	defer db.Close()

	var author models.Author

	err := ctx.ShouldBind(&author)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "500.html", gin.H{"message": err.Error()})
		return
	}

	fmt.Printf("%v\n", author.Name)

	result, err := db.Exec(`INSERT INTO public."Authors" ("name") values ($1)`, author.Name)

	if err != nil {
		log.Println("Insert Error", err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	fmt.Println("Inserted rows", rowsAffected)

	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%d row(s) inserted", rowsAffected)})
}

func insertBook(ctx *gin.Context) {
	db := ConnectDB()
	defer db.Close()

	var book models.Book

	err := ctx.ShouldBind(&book)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "500.html", gin.H{"message": err.Error()})
		return
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v, %v\n",
		book.BookName, book.Author, book.Quantity,
		book.Description, book.Published, book.Page, book.Language)

	result, err := db.Exec(`INSERT INTO public."Books" ("bookName", "author", "quantity", "description", "published", "page", "bookLanguage") values 
	($1, $2, $3, $4, $5, $6, $7)`, book.BookName, book.Author, book.Quantity, book.Description, book.Published, book.Page, book.Language)

	if err != nil {
		log.Println("Insert Error", err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	fmt.Println("Inserted rows", rowsAffected)

	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%d row(s) inserted", rowsAffected)})
}

func main() {
	router := gin.Default() // set router

	router.Use(cors.Default())

	router.StaticFile("/favicon.ico", "./public/favicon.ico")

	router.LoadHTMLGlob("template/*") // show template directory to gin

	router.NoRoute(func(ctx *gin.Context) { // 404 routing
		ctx.HTML(404, "404.html", gin.H{
			"message": "page not found",
		})
	})

	router.GET("/books", getAllBooks)
	router.GET("/books/:id", getBookById)
	router.POST("/books", insertBook)
	router.POST("/authors", insertAuthor)
	router.POST("/languages", insertLanguage)
	router.DELETE("/books/:id", removeBookById)

	router.GET("/authors", getAllAuthors)

	err := router.Run(":8080")
	if err != nil {
		return
	} // http://localhost:8080
}
