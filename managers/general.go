package managers

import (
	"context"

	"fmt"

	"github.com/kayprogrammer/socialnet-v4/ent"
)

type SiteDetailManager struct {
}

func (obj SiteDetailManager) Get(client *ent.Client) (*ent.SiteDetail, error) {
	s, err := client.SiteDetail.
		Query().
		First(context.Background())
	if err != nil {
		fmt.Printf("failed querying site details: %v\n", err)
		return nil, nil
	}
	return s, nil
}

func (obj SiteDetailManager) Create(client *ent.Client) (*ent.SiteDetail, error) {
	s, err := client.SiteDetail.
		Create().
		Save(context.Background())
	if err != nil {
		fmt.Printf("failed creating site details: %v\n", err)
		return nil, nil
	}
	return s, nil
}
