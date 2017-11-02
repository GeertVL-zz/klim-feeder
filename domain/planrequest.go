package domain

import (
	"github.com/satori/go.uuid"
	"time"
)

type PlanRequestType int

const (
	WorksiteAnnouncement PlanRequestType = 1 + iota
	ConsultPlannedWorks
	EasementRequest
	ListOfRelevantUtil
	Fetrapi
	Powalco
)

type WorkType int

const (
	InstallCablePipe WorkType = 1 + iota
	SoilStudy
	Roadworks
	FoundationWorks
	ExtRenovationBuildings
	OtherUnder
	OtherAbove
	NewBuilding
	AgriWorks
	Demolition
	PreciseCablePipe
	RailwayWorks
	CivilEngStruct
	HydraulicWorks
	TempOccupation
	CivilEngMaint
	MastsPylonsTurbines
	UndergroundChamber
)

type ExecutionMethod int

const (
	MechanicalIncDrill ExecutionMethod = 1 + iota
	MechanicalHeavy
	Manual
	MechanicalLight
	SpecialConstruction
)

type PlanRequestStatus int

const (
	Created PlanRequestStatus = 1 + iota
	WorkZoneDrawn
	Canceled
	Completed
	Published
	Abandoned
	Finished
)


type PlanRequest struct {
	UserId uuid.UUID
	MprId uuid.UUID
	CustomerId string
	ExternalPlanRequestId string
	PlanRequestType PlanRequestType
	WorkType WorkType
	ExecutionMethod ExecutionMethod
	AnnouncerIdentification string
	StartDate time.Time
	EndDate time.Time
	WorkDescription string
	Location string
	CenterLong float32
	CenterLat float32
	Status PlanRequestStatus
	Archived bool
	CreatedAt time.Time
	IntersectedAt time.Time
}

func (pr *PlanRequest) Import(data []string) {

}
