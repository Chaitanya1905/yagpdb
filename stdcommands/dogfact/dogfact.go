package dogfact

import (
	"github.com/jonas747/dcmd"
	"github.com/jonas747/yagpdb/commands"
	"math/rand"
)

var Command = &commands.YAGCommand{
	CmdCategory: commands.CategoryFun,
	Name:        "DogFact",
	Aliases:     []string{"dog", "dogfacts"},
	Description: "Dog Facts",

	RunFunc: func(data *dcmd.Data) (interface{}, error) {
		df := dogfacts[rand.Intn(len(dogfacts))]
		return df, nil
	},
}
