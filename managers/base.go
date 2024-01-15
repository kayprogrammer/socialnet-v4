package managers

import (
	"fmt"

	"github.com/kayprogrammer/socialnet-v4/ent"
)

type FileManager struct {

}

func (obj FileManager) Create(client *ent.Client, fileType *string) (*ent.File, error) {
	u, err := client.File.
		Create().
		SetResourceType(*fileType).
		Save(Ctx)
	if err != nil {
		fmt.Printf("failed creating file object: %v\n", err)
		return nil, nil
	}
	return u, nil
}