package handlers

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"formative-13/models"
)

type BioskopHandler struct {
	DB *sql.DB
}

func NewBioskopHandler(db *sql.DB) *BioskopHandler {
	return &BioskopHandler{
		DB: db,
	}
}

// POST /bioskop
func (h *BioskopHandler) CreateBioskop(c *gin.Context) {
	var input models.BioskopInput

	// Membaca JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Format JSON tidak valid",
			"error":   err.Error(),
		})
		return
	}

	// Validasi Nama
	if strings.TrimSpace(input.Nama) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Nama bioskop tidak boleh kosong",
		})
		return
	}

	// Validasi Lokasi
	if strings.TrimSpace(input.Lokasi) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Lokasi bioskop tidak boleh kosong",
		})
		return
	}

	var id int

	query := `
		INSERT INTO bioskop (nama, lokasi, rating)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err := h.DB.QueryRow(
		query,
		input.Nama,
		input.Lokasi,
		input.Rating,
	).Scan(&id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal menambahkan bioskop",
			"error":   err.Error(),
		})
		return
	}

	bioskop := models.Bioskop{
		ID:     id,
		Nama:   input.Nama,
		Lokasi: input.Lokasi,
		Rating: input.Rating,
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Bioskop berhasil ditambahkan",
		"data":    bioskop,
	})
}

// GET /bioskop
func (h *BioskopHandler) GetAllBioskop(c *gin.Context) {
	rows, err := h.DB.Query(`
		SELECT id, nama, lokasi, rating
		FROM bioskop
		ORDER BY id
	`)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengambil data bioskop",
			"error":   err.Error(),
		})
		return
	}

	defer rows.Close()

	var bioskopList []models.Bioskop

	for rows.Next() {
		var bioskop models.Bioskop

		err := rows.Scan(
			&bioskop.ID,
			&bioskop.Nama,
			&bioskop.Lokasi,
			&bioskop.Rating,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Gagal membaca data bioskop",
				"error":   err.Error(),
			})
			return
		}

		bioskopList = append(bioskopList, bioskop)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Terjadi kesalahan saat mengambil data",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": bioskopList,
	})
}

// GET /bioskop/:id
func (h *BioskopHandler) GetBioskopByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID harus berupa angka",
		})
		return
	}

	var bioskop models.Bioskop

	query := `
		SELECT id, nama, lokasi, rating
		FROM bioskop
		WHERE id = $1
	`

	err = h.DB.QueryRow(query, id).Scan(
		&bioskop.ID,
		&bioskop.Nama,
		&bioskop.Lokasi,
		&bioskop.Rating,
	)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Bioskop tidak ditemukan",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengambil data bioskop",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": bioskop,
	})
}

// PUT /bioskop/:id
func (h *BioskopHandler) UpdateBioskop(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID harus berupa angka",
		})
		return
	}

	var input models.BioskopInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Format JSON tidak valid",
			"error":   err.Error(),
		})
		return
	}

	// Validasi Nama
	if strings.TrimSpace(input.Nama) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Nama bioskop tidak boleh kosong",
		})
		return
	}

	// Validasi Lokasi
	if strings.TrimSpace(input.Lokasi) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Lokasi bioskop tidak boleh kosong",
		})
		return
	}

	// Update data
	query := `
		UPDATE bioskop
		SET nama = $1,
		    lokasi = $2,
		    rating = $3
		WHERE id = $4
	`

	result, err := h.DB.Exec(
		query,
		input.Nama,
		input.Lokasi,
		input.Rating,
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal memperbarui bioskop",
			"error":   err.Error(),
		})
		return
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal membaca hasil update",
			"error":   err.Error(),
		})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Bioskop tidak ditemukan",
		})
		return
	}

	bioskop := models.Bioskop{
		ID:     id,
		Nama:   input.Nama,
		Lokasi: input.Lokasi,
		Rating: input.Rating,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Bioskop berhasil diperbarui",
		"data":    bioskop,
	})
}

// DELETE /bioskop/:id
func (h *BioskopHandler) DeleteBioskop(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID harus berupa angka",
		})
		return
	}

	result, err := h.DB.Exec(
		"DELETE FROM bioskop WHERE id = $1",
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal menghapus bioskop",
			"error":   err.Error(),
		})
		return
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal membaca hasil penghapusan",
			"error":   err.Error(),
		})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Bioskop tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Bioskop berhasil dihapus",
	})
}
