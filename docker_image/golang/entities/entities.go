package entities



type Agrist struct {
	ID           uint   `gorm:"primaryKey"`
	UserName     string `gorm:"column:username"`
	FarmlandID   uint   `gorm:"column:id"`
	FarmlandName string `gorm:"column:farmlandname"`
}

type Possession struct {
	ID         string `gorm:"primarykey"`
	AgristID   int    `gorm:"foreignKey:agrist_id"`
	FarmlandID string `gorm:"foreignKey:farmland_id"`
}
type Farmland struct {
	Name       string `gorm:"column:farmlandname"`
	CropsName  string `gorm:"column:cropsname"`
	MachineNum string `gorm:"columun:machine_num"`
}
type Temp_humi struct {
	Narmlang_id int     `json:"farmland_id"`
	Date        string  `json:"day"`
	Nemp        float32 `json:"temp"`
	Humi        float32 `json:"humi"`
}
type Planting struct {
	Farmlang_id int    `json:"farmland_id"`
	Crops_id    int    `json:"crops_id"`
	Startday    string `json:"startday"`
}
type Crops struct {
	Name string `json:"name"`
}
type Harvest struct {
	Machine_id string `json:"machine_id"`
	Crops_id   int    `json:"crops_id"`
	Amount     int    `json:"amount"`
	Photo_id   int    `json:"photo_id"`
}
type Unharvest struct {
	Machine_id string `json:"machine_id"`
	Crops_id   int    `json:"crops_id"`
	Amount     int    `json:"amount"`
	Photo_id   int    `json:"photo_id"`
}
type Photo struct {
	Num int    `json:"num"`
	Url string `json:"url"`
}

// Json用
type Temp struct {
	Temp float32 `gorm:"column:temp"`
	Humi float32 `gorm:"column:humi"`
}
type DayPhoto struct {
	Url string `gorm:"column:url"`
}
