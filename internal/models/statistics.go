package models

type FacultyGender struct {
	Year   int `json:"year" db:"year"`
	Male   int `json:"male" db:"male"`
	Female int `json:"female" db:"female"`
}

type FacultyProfession struct {
	ProfessionID   int    `json:"profession_id" db:"profession_id"`
	ProfessionName string `json:"profession_name" db:"profession_name"`
	Year           int    `json:"year" db:"year"`
	Count          int    `json:"count" db:"count"`
}

type FacultyAge struct {
	Age    int `json:"age" db:"age"`
	Male   int `json:"male" db:"male"`
	Female int `json:"female" db:"female"`
}

type FacultyRegion struct {
	RegionID   int    `json:"region_id" db:"region_id"`
	RegionName string `json:"region_name" db:"region_name"`
	Year       int    `json:"year" db:"year"`
	Count      int    `json:"count" db:"count"`
}
