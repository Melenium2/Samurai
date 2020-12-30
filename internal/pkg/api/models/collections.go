package models

type Collection interface {
	Get() []string
}

type Categories []string

func (c Categories) Get() []string {
	if len(c) == 0 {
		return CategoriesGoogle
	}
	return c
}

// CategoriesGoogle this is Collection which contains available
// to request categories
var CategoriesGoogle = Categories{
	"apps_topselling_free",
	"apps_topgrossing",
	"apps_movers_shakers",
	"apps_topselling_paid",
}

// CategoriesIos Collection which contains available to request categories
var CategoriesIos = Categories {
	"newapplications",
	"newfreeapplications",
	"newpaidapplications",
	"topfreeapplications",
	"topgrossingapplications",
	"toppaidapplications",
}


