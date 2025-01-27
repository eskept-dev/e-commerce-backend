package location

// Sample API to get province, district, ward
// API: GET https://provinces.open-api.vn/api?depth=3
// {
// 	"name": "Tỉnh Bắc Kạn",
// 	"code": 6,
// 	"codename": "tinh_bac_kan",
// 	"division_type": "tỉnh",
// 	"phone_code": 209,
// 	"districts": [
// 		{
// 			"name": "Thành Phố Bắc Kạn",
// 			"code": 58,
// 			"codename": "thanh_pho_bac_kan",
// 			"division_type": "huyện",
// 			"short_codename": "bac_kan",
// 			"wards": [
// 				{
// 					"name": "Phường Nguyễn Thị Minh Khai",
// 					"code": 1834,
// 					"codename": "phuong_nguyen_thi_minh_khai",
// 					"division_type": "xã",
// 					"short_codename": "nguyen_thi_minh_khai"
// 				}
// 			]
// 		}
// 	]
// }

import (
	"encoding/json"
	"eskept/internal/app/context"
	"eskept/internal/constants/enums"
	"eskept/internal/models"
	"eskept/internal/repositories"
	"io/ioutil"
	"log"
	"net/http"
)

// Location represents the structure of location data
// including provinces, districts, and wards
type Location struct {
	Name         string     `json:"name"`
	Code         int        `json:"code"`
	Codename     string     `json:"codename"`
	DivisionType string     `json:"division_type"`
	PhoneCode    int        `json:"phone_code"`
	Districts    []District `json:"districts"`
}

type District struct {
	Name          string `json:"name"`
	Code          int    `json:"code"`
	Codename      string `json:"codename"`
	DivisionType  string `json:"division_type"`
	ShortCodename string `json:"short_codename"`
	Wards         []Ward `json:"wards"`
}

type Ward struct {
	Name          string `json:"name"`
	Code          int    `json:"code"`
	Codename      string `json:"codename"`
	DivisionType  string `json:"division_type"`
	ShortCodename string `json:"short_codename"`
}

func storeFlattenedLocations(appCtx *context.AppContext, locations []Location) error {
	locationRepo := repositories.NewLocationRepository(appCtx)

	for _, province := range locations {
		// Store province
		provinceEntry := models.Location{
			Name:           province.Name,
			Type:           enums.LocationProvince,
			ParentCodeName: "",
		}
		err := locationRepo.Create(&provinceEntry)
		if err != nil {
			log.Println("Error storing province:", err)
			return err
		}

		for _, district := range province.Districts {
			// Store district
			districtEntry := models.Location{
				Name:           district.Name,
				Type:           enums.LocationDistrict,
				ParentCodeName: province.Codename,
			}
			err := locationRepo.Create(&districtEntry)
			if err != nil {
				log.Println("Error storing district:", err)
				return err
			}

			for _, ward := range district.Wards {
				// Store ward
				wardEntry := models.Location{
					Name:           ward.Name,
					Type:           enums.LocationWard,
					ParentCodeName: district.Codename,
				}
				err := locationRepo.Create(&wardEntry)
				if err != nil {
					log.Println("Error storing ward:", err)
					return err
				}

				log.Println("Stored ward:", wardEntry)
			}
			log.Println("Stored district:", districtEntry)
		}
		log.Println("Stored province:", provinceEntry)
	}

	return nil
}

func fetchAndStoreLocations(appCtx *context.AppContext) {
	resp, err := http.Get("https://provinces.open-api.vn/api?depth=3")
	if err != nil {
		log.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return
	}

	var locations []Location
	err = json.Unmarshal(body, &locations)
	if err != nil {
		log.Println("Error parsing JSON:", err)
		return
	}

	log.Println("Fetched locations:", len(locations))

	// Store data in the database
	err = storeFlattenedLocations(appCtx, locations)
	if err != nil {
		log.Println("Error storing locations:", err)
		return
	}
	log.Println("Locations stored successfully.")
}

func InitLocation(appCtx *context.AppContext) {
	fetchAndStoreLocations(appCtx)
}
