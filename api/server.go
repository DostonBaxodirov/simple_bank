package api

import (
	db "simpleBank/tutorial"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}
