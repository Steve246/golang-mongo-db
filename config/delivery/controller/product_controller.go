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

	FindByIdUsecase usecase.GetByIdUseCase

	FindByCategoryUsecase usecase.GetByCategoryUseCase



}

func (pc *ProductController) FindCategory(ctx *gin.Context) {
	categoryProduct := ctx.Param("categoryProduct")

	product, err := pc.FindByCategoryUsecase.FindByCategory(categoryProduct)

	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": product})


}

func (pc *ProductController) FindPostById(ctx *gin.Context) {
	productID := ctx.Param("postId")

	// productID := "62cda38af994b6a69b95bf7c"
	//jalan kalau dikasih id manual

	post, err := pc.FindByIdUsecase.FindById(productID)

	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": post})
}


func (pc *ProductController) DeleteProduct (ctx *gin.Context) {
	productId := ctx.Param("postID")
	// productId := "62cd8c761988bbb77e85e3c5"
	//jalan kalau dikasih ide manual
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
	// productID := ctx.Param("postID")

	productID := "62cd8c6d1988bbb77e85e3c4"

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

	//nambain page juga sama limit

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
	DeleteProdukUsecase usecase.DeleteUseCase, 	FindByIdUsecase usecase.GetByIdUseCase,FindByCategoryUsecase usecase.GetByCategoryUseCase	) *ProductController {
	controller := ProductController{
		router: router,
		productUseCase: productUseCase,
		FindlimitUsecase: FindlimitUsecase,
		UpdateProdukUsecase: UpdateProdukUsecase,
		DeleteProdukUsecase: DeleteProdukUsecase,
		FindByIdUsecase: FindByIdUsecase,
		FindByCategoryUsecase: FindByCategoryUsecase,
	}

	//GET di param querry, POST baru pake body

	router.POST("/product", controller.registerNewProduct)

	router.GET("/productAll", controller.FindLimit)

	router.PATCH("/productUpdate", controller.UpdateProduct)

	router.DELETE("/productDelete/:postID", controller.DeleteProduct)

	router.GET("/productGetID/:postID", controller.FindPostById)

	router.GET("/productGetCategory/:categoryProduct", controller.FindCategory)


	return &controller
}