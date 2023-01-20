package  entities


type Agrist struct {
	ID uint `gorm:"primary_key"`
	Name string `json:"name"`
	Possession_id []Possession
}
type Possession struct {
	ID uint `gorm:"primary_key"`
	Agrist_id int `json:"id"`
	Farmland_id int `json:"farmland_id"`
	Farmlang Farmland
	
}
type Farmland struct {
	ID uint `gorm:"primary_key"`
	Name string `json:"name"`
}
type Temp_humi struct {
	Narmlang_id int `json:"farmland_id"`
	Nay string `json:"day"`
	Nemp float32 `json:"temp"`
	Humi float32 `json:"humi"`
}
type Planting struct {
	Farmlang_id int `json:"farmland_id"`
	Crops_id int `json:"crops_id"`
	Startday string `json:"startday"`
}
type Crops struct {
	Name string `json:"name"`
}
type Harvest struct {
	Machine_id string `json:"machine_id"`
	Crops_id int `json:"crops_id"`
	Amount int `json:"amount"`
	Photo_id int `json:"photo_id"`
}
type Unharvest struct {
	Machine_id string `json:"machine_id"`
	Crops_id int `json:"crops_id"`
	Amount int `json:"amount"`
	Photo_id int `json:"photo_id"`
}
type Photo struct {
	Num int `json:"num"`
	Url string `json:"url"`
}