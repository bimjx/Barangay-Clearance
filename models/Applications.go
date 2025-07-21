package models

import "github.com/uadmin/uadmin"

type Applications struct {
	uadmin.Model
	Fullname   string
	Email      string
	Address    string
	Purpose    Purpose
	ValidID    string `uadmin:"file"`
	isApproved bool
}

type Purpose int

func (Purpose) Employment() Purpose {
	return 1
}

func (Purpose) Business() Purpose {
	return 2
}

func (Purpose) SchoolRequirements() Purpose {
	return 3
}

func (Purpose) Other() Purpose {
	return 4
}
