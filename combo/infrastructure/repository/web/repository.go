package web

import (
	ent "combo/ent"
)

type repo struct {
	entClient *ent.Client
}

func New(entClient *ent.Client) Web {
	repo := repo{
		entClient: entClient,
	}

	var web Web = &repo

	return web
}
