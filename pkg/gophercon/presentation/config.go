package presentation

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/KathurimaKimathi/gophercon-demo/pkg/gophercon/application/common/helpers"
	"github.com/KathurimaKimathi/gophercon-demo/pkg/gophercon/infrastructure"
	"github.com/KathurimaKimathi/gophercon-demo/pkg/gophercon/infrastructure/datastore"
	"github.com/KathurimaKimathi/gophercon-demo/pkg/gophercon/infrastructure/datastore/psql"
	"github.com/KathurimaKimathi/gophercon-demo/pkg/gophercon/infrastructure/services/upload"
	"github.com/KathurimaKimathi/gophercon-demo/pkg/gophercon/presentation/rest"
	"github.com/KathurimaKimathi/gophercon-demo/pkg/gophercon/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
)

// GopherconAllowedOrigins is a list of CORS origins allowed to interact with this service
var GopherconAllowedOrigins = []string{
	"http://localhost:7777",
}

// GopherconAllowedHeaders is a list of CORS allowed headers for the clinical service
var GopherconAllowedHeaders = []string{
	"Accept",
	"Accept-Charset",
	"Accept-Language",
	"Accept-Encoding",
	"Origin",
	"Host",
	"User-Agent",
	"Content-Length",
	"Content-Type",
	"Authorization",
	"X-Authorization",
}

// Compile the regex patterns into a slice of *regexp.Regexp
func compilePatterns(patterns []string) []*regexp.Regexp {
	var compiledPatterns []*regexp.Regexp

	for _, pattern := range patterns {
		compiledPattern := regexp.MustCompile(pattern)
		compiledPatterns = append(compiledPatterns, compiledPattern)
	}

	return compiledPatterns
}

// Check if the origin is allowed by matching it against the compiled regex patterns
func isAllowedOrigin(origin string, compiledPatterns []*regexp.Regexp) bool {
	for _, pattern := range compiledPatterns {
		if pattern.MatchString(origin) {
			return true
		}
	}

	return false
}

// PrepareServer sets up the HTTP server
func PrepareAndStartServer(
	ctx context.Context, port int) {
	r := gin.Default()

	logger := initializeLogger()

	ds := datastore.NewDbService()

	err := psql.RunMigrations()
	if err != nil {
		fmt.Printf("unable to run migrations with error %v", err)
		panic(err)
	}

	uploadMedia := upload.NewServiceUpload(ctx)

	infra := infrastructure.NewInfrastructureInitializer(ds, uploadMedia)
	useCases := usecase.NewUseCasesInitializer(*infra)

	SetupRoutes(r, logger, useCases)

	addr := fmt.Sprintf(":%d", port)

	if err := r.Run(addr); err != nil {
		helpers.LogStartupError(ctx, err)
	}
}

func SetupRoutes(r *gin.Engine, log *slog.Logger, usecases usecase.IUsecase) {
	compiledPatterns := compilePatterns(GopherconAllowedOrigins)

	r.Use(cors.New(cors.Config{
		AllowWildcard:    true,
		AllowMethods:     []string{http.MethodPut, http.MethodPatch, http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodOptions},
		ExposeHeaders:    []string{"Content-Length", "Link"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			// Specific localhost origins
			if origin == "http://localhost:7777" {
				return true
			}

			allowed := isAllowedOrigin(origin, compiledPatterns)
			return allowed
		},
		MaxAge:          12 * time.Hour,
		AllowWebSockets: true,
	}))

	r.Use(sloggin.NewWithConfig(log, sloggin.Config{
		WithRequestBody:  true,
		WithResponseBody: true,
	}))
	r.Use(gin.Recovery())

	handlers := rest.NewPresentationHandlers(usecases)

	// graphQL := r.Group("/graphql")
	// graphQL.Use(rest.AuthenticationGinMiddleware(cacheStore, *authclient))
	// graphQL.Use(rest.TenantIdentifierExtractionMiddleware(infra.FHIR))
	// graphQL.Any("", GQLHandler(usecases))

	// Unauthenticated routes
	// ide := r.Group("/ide")
	// ide.Any("", playgroundHandler())

	apis := r.Group("/api")

	v1 := apis.Group("/v1")

	user := v1.Group("/user")
	user.POST("", handlers.HandleCreateUser())
}

// GQLHandler sets up a GraphQL resolver
// func GQLHandler(service usecases.Interactor) gin.HandlerFunc {
// 	resolver, err := graph.NewResolver(service)
// 	if err != nil {
// 		log.Panicf("failed to start graph resolver: %s", err)
// 	}

// 	server := handler.NewDefaultServer(
// 		generated.NewExecutableSchema(
// 			generated.Config{
// 				Resolvers: resolver,
// 			},
// 		),
// 	)

// 	return func(ctx *gin.Context) {
// 		server.ServeHTTP(ctx.Writer, ctx.Request)
// 	}
// }

// func playgroundHandler() gin.HandlerFunc {
// 	h := playground.Handler("GraphQL IDE", "/graphql")

// 	return func(c *gin.Context) {
// 		h.ServeHTTP(c.Writer, c.Request)
// 	}
// }

func initializeLogger() *slog.Logger {
	handler := slog.Handler(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))

	attrs := []slog.Attr{
		{
			Key: "service", Value: slog.StringValue("gophercon-cms"),
		},
	}

	handler.WithAttrs(attrs)

	log := slog.New(handler)
	slog.SetDefault(log)

	return log
}
