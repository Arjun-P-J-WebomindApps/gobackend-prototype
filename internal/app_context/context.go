package app_context

import (
	"github.com/Arjun-P-J-WebomindApps/gobackend-prototype/internal/db"
	"github.com/Arjun-P-J-WebomindApps/gobackend-prototype/internal/search"
)

type Context struct {
	DB     *db.DBContext
	Search *search.TypesenseContext
}
