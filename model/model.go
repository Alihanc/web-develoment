package model

import (
	"time"

	"github.com/Alihanc/web-develoment/helper"
)

type Bunker struct {
	ID             string    `json:"id "`
	NAME           string    `json:"name "`
	CATEGORY       string    `json:"category "`
	PLAQUE         string    `json:"plaque "`
	MAINTANCEDATE  string    `json :" maintance_date`
	MAINTANCE_3600 time.Time `json:"maintance_3600`
	MAINTANCE_7200 time.Time `json:"maintance_7200`
}

var InitialData = []Bunker{
	{
		ID:             "1",
		NAME:           "God of War",
		CATEGORY:       "Open World",
		PLAQUE:         "34 ZPA 45",
		MAINTANCEDATE:  "13-12-2022",
		MAINTANCE_3600: helper.Maintance_3600,
		MAINTANCE_7200: helper.Maintance_7200,
	},
	{
		ID:             "2",
		NAME:           "World War",
		CATEGORY:       "Fantasy",
		PLAQUE:         "59 KPL 45",
		MAINTANCEDATE:  "15-11-2022",
		MAINTANCE_3600: helper.Maintance_3600,
		MAINTANCE_7200: helper.Maintance_7200,
	},
	{
		ID:             "3",
		NAME:           "BattleFiled",
		CATEGORY:       "Shooter",
		PLAQUE:         "59 KMT 67",
		MAINTANCEDATE:  "05-12-2022",
		MAINTANCE_3600: helper.Maintance_3600,
		MAINTANCE_7200: helper.Maintance_7200,
	},
}
