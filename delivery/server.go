package delivery

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/config"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/delivery/controller"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/repository"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/usecase"
)

type Server struct {
	vehicleUC  usecase.VehicleUseCase
	brandUC    usecase.BrandUseCase
	customerUC usecase.CustomerUseCase
	engine     *gin.Engine
	host       string
}

func (s *Server) Run() {
	s.initController()
	err := s.engine.Run(s.host)
	if err != nil {
		panic(err)
	}
}

func (s *Server) initController() {
	controller.NewVehicleController(s.engine, s.vehicleUC)
	controller.NewBrandController(s.engine, s.brandUC)
	controller.NewCustomerController(s.engine, s.customerUC)
}

func NewServer() *Server {
	c, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}

	dbConn, _ := config.NewDbConnection(c)
	db := dbConn.Conn()

	r := gin.Default()
	vehicleRepo := repository.NewVehicleRepository(db)
	brandRepo := repository.NewBrandRepository(db)
	customerRepo := repository.NewCustomerRepository(db)
	brandUC := usecase.NewBrandUseCase(brandRepo)
	vehilceUC := usecase.NewVehicleUseCase(vehicleRepo, brandUC)
	customerUC := usecase.NewCustomerUseCase(customerRepo)
	host := fmt.Sprintf("%s:%s", c.ApiHost, c.ApiPort)
	return &Server{
		vehicleUC:  vehilceUC,
		brandUC:    brandUC,
		customerUC: customerUC,
		engine:     r,
		host:       host}
}
