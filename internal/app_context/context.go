package app_context

import "github.com/Arjun-P-J-WebomindApps/gobackend-prototype/internal/db"

type Context struct {
	DB *db.DBContext
}
