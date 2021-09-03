package setup

import (
	"github.com/Dasha-Kinsely/mega-narrator/routers"
)

func (server *Server) InitializeRoutes() {
	base := server.Router.Group("/api")
	// authentication
	routers.AuthRoutes(base.Group("/auth"))
	// user creations
	routers.CreationRoutes(base.Group("/creation"))
	// routes that take care of small and simple functions
	routers.MiscRoutes(base.Group("/misc"))
}