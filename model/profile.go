package model

type Profile struct {
	BasicIfo    BasicIfo
	SaleIfo     SaleIfo
	BuildIfo    BuildIfo
	PropertyIfo PropertyIfo
}
type BasicIfo struct {
	Name         string
	AvgPrice     string
	TotalPrice   string
	PropertyType string
	Address      string
}
type SaleIfo struct {
	OpenTime string
	HandTime string
}
type BuildIfo struct {
	Developer string
	PeriodInt string
	Greenery  string
	BuildType string
	PlotRatio string
}
type PropertyIfo struct{
	Company string
	PropertyPrice string
	WaterGas string
}