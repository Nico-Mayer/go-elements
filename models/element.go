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
	NumberofNeutrons  int      `json:"numberOfNeutrons"`
	NumberofElectrons int      `json:"numberOfElectrons"`
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
	NumberOFIsotopes  *int     `json:"numberOfIsotopes"`
	YearOfDiscovery   *int     `json:"yearOfDiscovery"`
	SpecificHeat      *float64 `json:"specificHeat"`
	NumberofShells    int      `json:"numberOfShells"`
	NumberofValence   *int     `json:"numberOfValence"`
}

func ElementByAtomicNumber(atomicNumber int) Element {
	fmt.Println(atomicNumber)
	query := "SELECT * FROM elements WHERE AtomicNumber = $1"
	row := db.DB.QueryRow(query, atomicNumber)

	fmt.Println(row)

	var element Element
	err := row.Scan(&element.AtomicNumber, &element.Element, &element.Symbol, &element.AtomicMass, &element.NumberofNeutrons, &element.NumberofElectrons, &element.Period, &element.Group, &element.Phase, &element.Radioactive, &element.Natural, &element.Metal, &element.Nonmetal, &element.Metalloid, &element.Type, &element.AtomicRadius, &element.Electronegativity, &element.FirstIonization, &element.Density, &element.MeltingPoint, &element.BoilingPoint, &element.NumberOFIsotopes, &element.YearOfDiscovery, &element.SpecificHeat, &element.NumberofShells, &element.NumberofValence)

	if err != nil {
		return Element{}
	}

	return element
}
