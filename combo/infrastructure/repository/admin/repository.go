package admin

import (
	ent "combo/ent"
)

type repo struct {
	entClient *ent.Client
}

func New(entClient *ent.Client) Admin {
	repo := repo{
		entClient: entClient,
	}

	var admin Admin = &repo

	return admin
}
