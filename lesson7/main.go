package main

import (
	"fmt"
	"strings"
)

type Playable interface {
	TeamName() string
}

type Family struct {
	Name            string
	NumberOfMembers int
	Assets          *Assets
}

func NewFamily(name string, numberOfMembers int, assets *Assets) *Family {
	return &Family{
		Name:            name,
		NumberOfMembers: numberOfMembers,
		Assets:          assets,
	}
}

func (f *Family) TeamName() string {
	return f.Name
}

type Assets struct {
	NumberOfCars   int
	NumberOfHouses int
}

func NewAssets(numberOfCars int, numberOfHouses int) *Assets {
	return &Assets{
		NumberOfCars:   numberOfCars,
		NumberOfHouses: numberOfHouses,
	}
}

type Company struct {
	Name              string
	NumberOfEmployees int
	Industry          string
}

func NewCompany(name string, numberOfEmployees int, industry string) *Company {
	return &Company{
		Name:              name,
		NumberOfEmployees: numberOfEmployees,
		Industry:          industry,
	}
}

func (c *Company) TeamName() string {
	return c.Name
}

type Game struct {
	Teams []Playable
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) AddTeam(team Playable) {
	g.Teams = append(g.Teams, team)
}

func (g *Game) Announce() {
	names := make([]string, 0, len(g.Teams))
	for _, t := range g.Teams {
		names = append(names, t.TeamName())
	}
	fmt.Printf("Today we have %d team(s): %s\n", len(g.Teams), strings.Join(names, ", "))
}

func main() {
	simpsons := NewFamily("Simpsons", 5, NewAssets(2, 1))

	smiths := NewFamily("Smiths", 5, NewAssets(3, 1))

	dunderMifflin := NewCompany("dunder mifflin", 1000, "Paper company")

	quiz := NewGame()
	quiz.AddTeam(simpsons)
	quiz.AddTeam(smiths)
	quiz.AddTeam(dunderMifflin)
	quiz.Announce()
}
