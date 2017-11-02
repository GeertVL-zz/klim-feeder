package domain

import "strconv"

type ActorState int

const (
	Unverified ActorState = 1 + iota
	Verified
	Banned
)

type Role int

const (
	MainPlanRequestor Role = 1 + iota
	PlanRequestor
	Member
	Owner
	MainAdmin
	Admin
	FunctionalManager
	TechnicalManager
	M2M
)

type Language int

const (
	Nl Language = 1 + iota
	Fr
	De
	En
)

type ActorType int

const (
	Contractor = 1 + iota
	SubContractor
	StudyAgency
	Architect
	Founder
	MunicipalAdministration
	PrivatePerson
	Notary
	Other
)

type Actor struct {
	Email string
	Role Role
	FirstName string
	LastName string
	Company string
	Country string
	PostalCode string
	Municipality string
	Street string
	HouseNumber string
	PhoneNumber string
	Fax string
	Fetrapi bool
	Website string
	Language Language
	Type ActorType
	State ActorState
}

func (a *Actor) Import(data []string) {
	a.Email = data[0]
	role, _ := strconv.Atoi(data[1])
	a.Role = Role(role)
	a.FirstName = data[2]
	a.LastName = data[3]
	a.Company = data[4]
	a.Country = data[5]
	a.PostalCode = data[6]
	a.Municipality = data[7]
	a.Street = data[8]
	a.HouseNumber = data[9]
	a.PhoneNumber = data[10]
	a.Fax = data[11]
	fetrapi, _ := strconv.ParseBool(data[12])
	a.Fetrapi = fetrapi
	a.Website = data[13]
	lang, _ := strconv.Atoi(data[14])
	a.Language = Language(lang)
	actortype, _ := strconv.Atoi(data[15])
	a.Type = ActorType(actortype)
	state, _ := strconv.Atoi(data[16])
	a.State = ActorState(state)
}
