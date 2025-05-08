package api

import (
	"sistema_artigos_bpk/internal/app/handler"
	authRepository "sistema_artigos_bpk/internal/app/repository/auth"
	eventoRepository "sistema_artigos_bpk/internal/app/repository/evento"
	usuarioRepository "sistema_artigos_bpk/internal/app/repository/usuario"
	"sistema_artigos_bpk/internal/app/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB ,jwtKey []byte) {
    apiGroup := e.Group("/api")

    registerAuthRoutes(apiGroup,db)                                       
    registerEventoRoutes(apiGroup,db)

}



func registerAuthRoutes(g *echo.Group, db *gorm.DB ) {
    groupAuth := g.Group("/auth")
    g.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
         
            

        
            // Use as variáveis locais para criar os repositórios e serviços
            repository := authRepository.NewAuthRepository(db,c.Logger())
            repositoryJwt := authRepository.NewJwtRepository()
            service := service.NewAuthService(repository,repositoryJwt, c.Logger())

            h := handler.NewAuthHandler(repository, repositoryJwt, service)
            
            groupAuth.POST("/criar", h.CreateAuth)
            groupAuth.POST("/login", h.Login)

            return next(c)
        }
    })
}



func registerEventoRoutes(g *echo.Group, db *gorm.DB ) {
    groupAuth := g.Group("/evento")
    g.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
         
            

        
            // Use as variáveis locais para criar os repositórios e serviços
            repository := eventoRepository.NewEventoRepository(db,c.Logger())
            repositoryJwt := authRepository.NewJwtRepository()
            usuarioRep := usuarioRepository.NewUsuarioRepository(db,c.Logger())
            service := service.NewEventoService(repository,repositoryJwt, usuarioRep ,c.Logger())

            h := handler.NewEventoHandler(repository, service)
            
            groupAuth.POST("/criar", h.CreateEvento)
            groupAuth.POST("/:id", h.UploadPDF)
            groupAuth.GET("/", h.AllEvents)
            groupAuth.PUT("/", h.EditEvento)

            return next(c)
        }
    })
}