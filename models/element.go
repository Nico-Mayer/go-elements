package models

import (
	"fmt"

	"github.com/nico-mayer/go-elements/db"
)

type Element struct {
	AtomicNumber      int      `json:"atomicNumber"`
	Element           string   `json:"element"`
	Symbol            string   `json:"symbol"`
	AtomicMass        float64  `json:"atomicMass"`
	NumberOfNeutrons  int      `json:"numberOfNeutrons"`
	NumberOfElectrons int      `json:"numberOfElectrons"`
	Period            int      `json:"period"`
	Group             int      `json:"group"`
	Phase             string   `json:"phase"`
	Radioactive       bool     `json:"radioactive"`
	Natural           bool     `json:"natural"`
	Metal             bool     `json:"metal"`
	Nonmetal          bool     `json:"nonmetal"`
	Metalloid         bool     `json:"metalloid"`
	Type              *string  `json:"type"`
	AtomicRadius      *float64 `json:"atomicRadius"`
	Electronegativity *float64 `json:"electronegativity"`
	FirstIonization   *float64 `json:"firstIonization"`
	Density           *float64 `json:"density"`
	MeltingPoint      *float64 `json:"meltingPoint"`
	BoilingPoint      *float64 `json:"boilingPoint"`
	NumberOfIsotopes  *int     `json:"numberOfIsotopes"`
	YearOfDiscovery   *int     `json:"yearOfDiscovery"`
	SpecificHeat      *float64 `json:"specificHeat"`
	NumberOfShells    int      `json:"numberOfShells"`
	NumberOfValence   *int     `json:"numberOfValence"`
}

func ElementByAtomicNumber(atomicNumber int) Element {
	query := `SELECT * FROM elements WHERE atomicnumber = $1`
	row := db.DB.QueryRow(query, atomicNumber)

	var element Element
	err := row.Scan(
		&element.AtomicNumber,
		&element.Element,
		&element.Symbol,
		&element.AtomicMass,
		&element.NumberOfNeutrons,
		&element.NumberOfElectrons,
		&element.Period,
		&element.Group,
		&element.Phase,
		&element.Radioactive,
		&element.Natural,
		&element.Metal,
		&element.Nonmetal,
		&element.Metalloid,
		&element.Type,
		&element.AtomicRadius,
		&element.Electronegativity,
		&element.FirstIonization,
		&element.Density,
		&element.MeltingPoint,
		&element.BoilingPoint,
		&element.NumberOfIsotopes,
		&element.YearOfDiscovery,
		&element.SpecificHeat,
		&element.NumberOfShells,
		&element.NumberOfValence,
	)

	if err != nil {
		fmt.Println(err)
		return Element{}
	}

	return element
}
