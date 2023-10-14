package main

import (
	"github.com/Shawket4/FlavourFusion/Controllers"
	"github.com/Shawket4/FlavourFusion/Models"
)

func main()  {
	Models.Setup()
	Controllers.Setup()
}
