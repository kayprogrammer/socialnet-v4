package managers

import (
	"fmt"

	"github.com/kayprogrammer/socialnet-v4/ent"
	"github.com/google/uuid"
	"github.com/kayprogrammer/socialnet-v4/ent/file"


)

type FileManager struct {

}

func (obj FileManager) GetByID(client *ent.Client, id uuid.UUID) *ent.File {
	f, _ := client.File.
		Query().
		Where(file.ID(id)).
		Only(Ctx)
	return f
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

func (obj FileManager) Update(client *ent.Client, file *ent.File, resourceType string) *ent.File {
	f, _ := file.
		Update().
		SetResourceType(resourceType).
		Save(Ctx)
	return f
}