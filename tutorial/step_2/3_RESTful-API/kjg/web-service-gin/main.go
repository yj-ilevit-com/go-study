package main

// 필요한 패키지 참조
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 데이터 구조 및 초기 데이터
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// 앨범 데이터
var albums = []album{ 
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99}, 
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99}, 
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99}, 
}

// GET /albums 핸들러
func getAlbums(c *gin.Context) { // Gin 컨텍스트를 매개변수로 받음
    c.IndentedJSON(http.StatusOK, albums) // JSON 형식으로 데이터 반환
}

// POST /albums 핸들러
func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
			return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// GET /albums/:id 핸들러
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
			if a.ID == id {
					c.IndentedJSON(http.StatusOK, a)
					return
			}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// main 함수 - 애플리케이션 실행 진입점
func main() {
    router := gin.Default() // Gin 라우터 초기화
    router.GET("/albums/:id", getAlbumByID) // GET /albums/:id 라우팅

    router.Run("localhost:8080") // Gin 라우터 시작
}