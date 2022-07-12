package delivery

import (
	"fmt"
	"golang-mongodb/config"
	"golang-mongodb/config/delivery/controller"
	"golang-mongodb/manager"
	"log"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	useCaseManager manager.UseCaseManager
	engine *gin.Engine
	host string 
}

func (a *appServer) initHandlers() {
	controller.NewProductController(a.engine, a.useCaseManager.ProductRegistrationUseCase(), a.useCaseManager.PaginationUseCase(), a.useCaseManager.UpdateProductUseCase(), a.useCaseManager.DeleteProductUsecase(), a.useCaseManager.GetByIdUseCase(), a.useCaseManager.GetByCategoryUsecase())
}

func (a *appServer) Run() {
	a.initHandlers()
	err := a.engine.Run(a.host)

	if err != nil {
		log.Println(err)
		return 
	}
}

func NewServer() *appServer {
	r := gin.Default()
	c := config.NewConfig()
	infraManager := manager.NewInfraManager(c)
	repoManager := manager.NewRepositoryManager(infraManager)
	useCaseManager := manager.NewUseCaseManager(repoManager)

	host := fmt.Sprintf("%s:%s", c.ApiHost, c.ApiPort)

	return &appServer{
		useCaseManager: useCaseManager,
		engine: r,
		host: host,
	}
}