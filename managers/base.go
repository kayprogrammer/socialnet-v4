package managers

import (
	"github.com/google/uuid"
	"github.com/kayprogrammer/socialnet-v4/ent"
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

func (obj FileManager) Create(client *ent.Client, fileType string) *ent.File {
	file, _ := client.File.
		Create().
		SetResourceType(fileType).
		Save(Ctx)
	return file
}

func (obj FileManager) Update(client *ent.Client, file *ent.File, resourceType string) *ent.File {
	f, _ := file.
		Update().
		SetResourceType(resourceType).
		Save(Ctx)
	return f
}

func (obj FileManager) UpdateOrCreate(client *ent.Client, fileObj *ent.File, resourceType string) *ent.File {
	if fileObj != nil {
		fileObj = client.File.UpdateOneID(fileObj.ID).SetResourceType(resourceType).SaveX(Ctx)
	} else {
		fileObj = obj.Create(client, resourceType)
	}
	return fileObj
}
