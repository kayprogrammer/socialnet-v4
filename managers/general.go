package managers

import (
	"context"

	"github.com/kayprogrammer/socialnet-v4/ent"
)

type SiteDetailManager struct {
}

var Ctx = context.Background()

func (obj SiteDetailManager) Get(client *ent.Client) *ent.SiteDetail {
	s, _ := client.SiteDetail.
		Query().
		First(Ctx)
	return s
}

func (obj SiteDetailManager) Create(client *ent.Client) *ent.SiteDetail {
	s := client.SiteDetail.
		Create().
		SaveX(Ctx)
	return s
}

func (obj SiteDetailManager) GetOrCreate (client *ent.Client) *ent.SiteDetail {
	sitedetail := obj.Get(client)
	if sitedetail == nil {
		sitedetail = obj.Create(client)
	}
	return sitedetail
}