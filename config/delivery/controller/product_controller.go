package controller

import (
	"golang-mongodb/model"
	"golang-mongodb/usecase"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	router *gin.Engine
	productUseCase usecase.ProductRegistrationUseCase

	FindlimitUsecase usecase.PaginationUseCase
	UpdateProdukUsecase usecase.UpdateProductUsecase

	DeleteProdukUsecase usecase.DeleteUseCase

}

func (pc *ProductController) DeleteProduct (ctx *gin.Context) {
	productId := ctx.Param("postID")
	// productId := "62cd8c761988bbb77e85e3c5"
	err := pc.DeleteProdukUsecase.Delete(productId)

	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)

}

func (pc *ProductController) UpdateProduct(ctx *gin.Context) {
	productID := ctx.Param("postID")

	// produkId := "62cd8c761988bbb77e85e3c5"

	var produk *model.Product

	if err := ctx.ShouldBindJSON(&produk); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	updatedPost, err := pc.UpdateProdukUsecase.Update(productID, produk)

	// updatedPost, err := pc.postService.UpdatePost(postId, post)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedPost})

	
}

func (pc *ProductController) FindLimit(ctx *gin.Context) {
	//bikin model baru buat ambil nilai int

	var newLimit model.FindLimit

	err := ctx.ShouldBindJSON(&newLimit)

	if err != nil {
		log.Println(err.Error())
		return
	}

	listProduct, err := pc.FindlimitUsecase.FindLimit(&newLimit)

	if err != nil {
		log.Println(err)
		return 
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": listProduct,
	})

	
}


func (pc *ProductController) registerNewProduct(ctx *gin.Context) {
	var newProduct model.Product

	err := ctx.ShouldBindJSON(&newProduct)

	if err != nil {
		log.Println(err.Error())
		return
	}
	
	err = pc.productUseCase.Register(&newProduct)

	if err != nil {
		log.Println(err)
		return 
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": newProduct,
	})
}


func NewProductController(router *gin.Engine, productUseCase usecase.ProductRegistrationUseCase, FindlimitUsecase usecase.PaginationUseCase, UpdateProdukUsecase usecase.UpdateProductUsecase, 
	DeleteProdukUsecase usecase.DeleteUseCase) *ProductController {
	controller := ProductController{
		router: router,
		productUseCase: productUseCase,
		FindlimitUsecase: FindlimitUsecase,
		UpdateProdukUsecase: UpdateProdukUsecase,
		DeleteProdukUsecase: DeleteProdukUsecase,
	}

	router.POST("/product", controller.registerNewProduct)

	router.GET("/productAll", controller.FindLimit)

	router.PATCH("/productUpdate/:postID", controller.UpdateProduct)

	router.DELETE("/productDelete/:postID", controller.DeleteProduct)


	return &controller
}