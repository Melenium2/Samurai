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

var CategoriesGoogle = Categories{
	"apps_topselling_free",
	"apps_topgrossing",
	"apps_movers_shakers",
	"apps_topselling_paid",
}

var CategoriesIos = Categories {
	"newapplications",
	"newfreeapplications",
	"newpaidapplications",
	"topfreeapplications",
	"topgrossingapplications",
	"toppaidapplications",
}


