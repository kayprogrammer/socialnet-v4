package schemas

type SiteDetail struct {
	Name 		string 		`json:"name" example:"SocialNet"`
	Email 		string 		`json:"email" example:"johndoe@email.com"`
	Phone 		string		`json:"phone" example:"+2348133831036"`
	Address 	string		`json:"address" example:"234, Lagos, Nigeria"`
	Fb 			string		`json:"fb" example:"https://facebook.com"`
	Tw 			string		`json:"tw" example:"https://twitter.com"`
	Wh 			string		`json:"wh" example:"https://wa.me/2348133831036"`
	Ig 			string		`json:"ig" example:"https://instagram.com"`
}

type SiteDetailResponseSchema struct {
	ResponseSchema
	Data			SiteDetail		`json:"data"`
}

// func (obj SiteDetailResponseSchema) Init() SiteDetailResponseSchema {
// 	var originalSiteDetailStruct *ent.SiteDetail
// 	var targetSiteDetailStruct SiteDetail

// 	temporaryVariable, _ := json.Marshal(originalSiteDetailStruct)
// 	vv := json.Unmarshal(temporaryVariable, &targetSiteDetailStruct) 
// 	fmt.Println(vv)
	
// }