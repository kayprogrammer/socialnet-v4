package managers

import (
	"context"

	"fmt"

	"github.com/kayprogrammer/socialnet-v4/ent"
)

type SiteDetailManager struct {
}

var Ctx = context.Background()

func (obj SiteDetailManager) Get(client *ent.Client) (*ent.SiteDetail, error) {
	s, err := client.SiteDetail.
		Query().
		First(Ctx)
	if err != nil {
		fmt.Printf("failed querying site details: %v\n", err)
		return nil, nil
	}
	return s, nil
}

func (obj SiteDetailManager) Create(client *ent.Client) (*ent.SiteDetail, error) {
	s, err := client.SiteDetail.
		Create().
		Save(Ctx)
	if err != nil {
		fmt.Printf("failed creating site details: %v\n", err)
		return nil, nil
	}
	return s, nil
}

func (obj SiteDetailManager) GetOrCreate (client *ent.Client) *ent.SiteDetail {
	sitedetail, _ := obj.Get(client)
	if sitedetail == nil {
		sitedetail, _ = obj.Create(client)
	}
	return sitedetail
}