package inmem

import (
	"e-learning/app/models"
)

// SeederMaterial .
func SeederMaterial() []*models.Material {
	return []*models.Material{
		&models.Material{
			MaterialID:  1,
			Name:        "Reinforcement Learning: AI Flight with Unity ML Agents",
			Description: "In sunt anim exercitation officia consequat dolor ex aliquip. Ut id cupidatat cupidatat labore pariatur tempor. Labore deserunt magna enim minim ut irure officia sunt. Minim incididunt laborum non deserunt. Cupidatat pariatur nostrud eiusmod consectetur ullamco tempor minim labore excepteur.",
			ImgURL:      "https://img-a.udemycdn.com/course/480x270/2369480_c479_5.jpg?yiz8qXCA9RwjttLmtQ79OBcJUUMIGDyg2hHdv4WncHOSP6CU8-vXS1954ZZhQSg4RJQr_moLeCCZjdR58kHidWXT3hOAypl7sepJOTgHqV7R-q0g9tT3ASBN1dn2thFz",
			Author:      "John Doe",
		},
		&models.Material{
			MaterialID:  2,
			Name:        "Illustration: Clip Studio Paint for iPad",
			Description: "In sunt anim exercitation officia consequat dolor ex aliquip. Ut id cupidatat cupidatat labore pariatur tempor. Labore deserunt magna enim minim ut irure officia sunt. Minim incididunt laborum non deserunt. Cupidatat pariatur nostrud eiusmod consectetur ullamco tempor minim labore excepteur.",
			ImgURL:      "https://img-a.udemycdn.com/course/480x270/2497762_52f1_4.jpg?AxJc6_VRhlOA3Rq_DGMDRuVCNuA4x-jInCL_gjFpmz0eeOaExtKqd0jkbaurgac1U5k5EvWYPqiytuabkt3j38wTz6t8C6OIIZi4HTYzMMeVlpwbr5yTjyp-X1KtNlOD",
			Author:      "Alice",
		},
	}
}
