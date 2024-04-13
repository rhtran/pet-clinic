package main

import (
	"context"
	"fmt"
	"github.com/rhtran/golang-petclinic-service/app"
	"github.com/rhtran/golang-petclinic-service/internal/api/health"
	"github.com/rhtran/golang-petclinic-service/internal/api/info"
	"github.com/rhtran/golang-petclinic-service/internal/api/owner"
	"github.com/rhtran/golang-petclinic-service/internal/api/pet"
	"github.com/rhtran/golang-petclinic-service/internal/api/vet"
	"github.com/rhtran/golang-petclinic-service/internal/api/visit"
	"github.com/rhtran/golang-petclinic-service/pkg/dbase"
	"github.com/rhtran/golang-petclinic-service/pkg/ds"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"os"

	"github.com/rhtran/golang-petclinic-service/middleware"

	"gorm.io/gorm"

	"github.com/gin-contrib/location"

	"github.com/gin-gonic/gin"
	"github.com/qiangxue/go-restful-api/pkg/log"
	"golang.org/x/sync/errgroup"
)

var (
	g      errgroup.Group
	pg     *gorm.DB
	logger log.Logger
	r      *gin.Engine
)

func init() {
	// load application configurations
	if err := app.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}
	logger = log.New().With(context.TODO(), "version", app.Config.AppInfo.Version)
	logger.Infof("pet-clinic starts")
	pg = dbase.PgConnect()

	//ctx, client, err := okta.NewClient(context.TODO(), okta.WithOrgUrl("https://dev-293522.okta.com/"), okta.WithToken("{apiToken}"))

	// Creates a router without any middleware by default
	r = gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default, gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())
	r.Use(location.Default())
	r.Use(middleware.SetRequestUUID())
}

// Component initialization
func init() {
	logger.Info("Component initialization starts")
	healthCheckService := health.NewHealthCheckService(logger)
	healthCheckRouter := health.NewHealthCheckRouter(healthCheckService, logger)

	infoService := info.NewInfoService(logger)
	ipService := info.NewIPService(logger)
	infoRouter := info.NewInfoRouter(logger, infoService, ipService)

	// Pet
	petRepository := pet.NewPetRepository(logger, pg)
	petService := pet.NewPetService(logger, petRepository)
	petRouter := pet.NewPetRouter(logger, petService)

	// Vet
	vetRepository := vet.NewVetRepository(logger, pg)
	vetService := vet.NewVetService(logger, vetRepository)
	vetRouter := vet.NewVetRouter(logger, vetService)

	// Owner
	ownerRepository := owner.NewOwnerRepository(logger, pg)
	ownerService := owner.NewOwnerService(logger, ownerRepository)
	ownerRouter := owner.NewOwnerRouter(logger, ownerService)

	// Visit
	visitRepository := visit.NewVisitRepository(logger, pg)
	visitService := visit.NewVisitService(logger, visitRepository)
	visitRouter := visit.NewVisitRouter(logger, visitService)

	//authenService := service.NewAuthenService(logger)

	// config prometheus & endpoint group
	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)
	home := r.Group("/")
	v1 := r.Group("/v1")
	//v1.Use(middleware.Authenticate(authenService))

	healthCheckRouter.HealthCheckRegister(home.Group("/health"))
	infoRouter.InfoRegister(home.Group("/info"))
	petRouter.PetRegister(v1.Group("/pet"))
	vetRouter.VetRegister(v1.Group("/vet"))
	ownerRouter.OwnerRegister(v1.Group("/owner"))
	visitRouter.VisitRegister(v1.Group("/visit"))
}

func main() {
	httpServer := ds.NewHttpServer(r)
	httpRouter := httpServer.HttpRouter()

	g.Go(func() error {
		return httpRouter.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		logger.Error(err)
		os.Exit(-1)
	}
}
