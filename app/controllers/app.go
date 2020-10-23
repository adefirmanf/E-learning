package controllers

import (
	"e-learning/app/models"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	var materials = []*models.Material{
		&models.Material{
			MaterialID: 1,
			Name:       "Reinforcement Learning: AI Flight with Unity ML Agents",
			ImgURL:     "https://img-a.udemycdn.com/course/480x270/2369480_c479_5.jpg?yiz8qXCA9RwjttLmtQ79OBcJUUMIGDyg2hHdv4WncHOSP6CU8-vXS1954ZZhQSg4RJQr_moLeCCZjdR58kHidWXT3hOAypl7sepJOTgHqV7R-q0g9tT3ASBN1dn2thFz",
			Author:     "John Doe",
		},
		&models.Material{
			MaterialID: 2,
			Name:       "Illustration: Clip Studio Paint for iPad",
			ImgURL:     "https://img-a.udemycdn.com/course/480x270/2497762_52f1_4.jpg?AxJc6_VRhlOA3Rq_DGMDRuVCNuA4x-jInCL_gjFpmz0eeOaExtKqd0jkbaurgac1U5k5EvWYPqiytuabkt3j38wTz6t8C6OIIZi4HTYzMMeVlpwbr5yTjyp-X1KtNlOD",
			Author:     "Alice",
		},
		&models.Material{
			MaterialID: 2,
			Name:       "Photography",
		},
		&models.Material{
			MaterialID: 2,
			Name:       "Photography",
		},
		&models.Material{
			MaterialID: 2,
			Name:       "Photography",
		},
		&models.Material{
			MaterialID: 2,
			Name:       "Photography",
		},
		&models.Material{
			MaterialID: 2,
			Name:       "Photography",
		},
		&models.Material{
			MaterialID: 2,
			Name:       "Photography",
		},
	}

	return c.Render(materials)
}

func (c App) Login() revel.Result {
	return c.Render()
}

func (c App) Dashboard() revel.Result {
	return c.Render()
}

func (c App) Material(id int) revel.Result {
	return c.Render()
}
