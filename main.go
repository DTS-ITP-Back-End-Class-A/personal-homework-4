package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type Produk struct{
	NamaProduk string `json:"nama_produk"`
	JenisProduk string `json:"jenis_produk"`
	HargaProduk string `json:"nama_produk"`
	TempatPembelian string `json:"tempat_pembelian"`
	NomorBarcode string `json:"nomor_barcode"`
	TanggalPembelian string `json:"tanggal_pembelian"`
}

var signingKey = []byte("rakamin")

type UserRequest struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

// var produk = []Produk{
// 	{
// 		namaProduk: "mie goreng",
// 		jenisProduk: "makanan",
// 		hargaProduk: "2500",
// 		tempatPembelian: "indomaret",
// 		nomorBarcode: "09709735s",
// 		tanggalPembelian: "17 Agustus",
// 	},
// 	{
// 		namaProduk: "fanta",
// 		jenisProduk: "minuman",
// 		hargaProduk: "2050",
// 		tempatPembelian: "alfamaret",
// 		nomorBarcode: "025246735s",
// 		tanggalPembelian: "17 Agustus",
// 	},
// 	{
// 		namaProduk: "lefbouy",
// 		jenisProduk: "sabun",
// 		hargaProduk: "1450",
// 		tempatPembelian: "alfamaret",
// 		nomorBarcode: "02524373775",
// 		tanggalPembelian: "17 Agustus",
// 	},

// }

func main()  {
	
    app := fiber.New()

    app.Get("/get-product", func(c *fiber.Ctx) error {
		var produk = []Produk{
			{
				NamaProduk: "mie goreng",
				JenisProduk: "makanan",
				HargaProduk: "2500",
				TempatPembelian: "indomaret",
				NomorBarcode: "09709735s",
				TanggalPembelian: "17 Agustus",
			},
			{
				NamaProduk: "fanta",
				JenisProduk: "minuman",
				HargaProduk: "2050",
				TempatPembelian: "alfamaret",
				NomorBarcode: "025246735s",
				TanggalPembelian: "17 Agustus",
			},
			{
				NamaProduk: "lefbouy",
				JenisProduk: "sabun",
				HargaProduk: "1450",
				TempatPembelian: "alfamaret",
				NomorBarcode: "02524373775",
				TanggalPembelian: "17 Agustus",
			},
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{ 
			"message": "success", 
			"status": "OK", 
			"data": produk,
		})
	})
		

	apiGroup := app.Group("/api")
	apiGroup.Post("/login", func(c *fiber.Ctx) (err error) {
		var req UserRequest
		err = c.BodyParser(&req)
		if err != nil {
			log.Printf("Error in parsing the JSON request: %v.", err)
			return
		}

		if req.User != "admin" || req.Password != "4dm1n" {
			err = c.SendStatus(fiber.StatusUnauthorized)
			return
		}

		signJwt := jwt.New(jwt.SigningMethodHS256)

		claims := signJwt.Claims.(jwt.MapClaims)
		claims["name"] = "Admin"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		signJwt.Claims = claims

		token, err := signJwt.SignedString(signingKey)
		if err != nil {
			err = c.SendStatus(fiber.StatusInternalServerError)
			return
		}

		err = c.JSON(fiber.Map{"token": token})
		return
    })

    app.Listen(":3000")
}