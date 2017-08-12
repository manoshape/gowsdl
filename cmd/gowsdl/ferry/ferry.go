package ferry

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type Season string

const (
	SeasonSpring Season = "Spring"

	SeasonSummer Season = "Summer"

	SeasonFall Season = "Fall"

	SeasonWinter Season = "Winter"
)

type AdjustmentType string

const (
	AdjustmentTypeAddition AdjustmentType = "Addition"

	AdjustmentTypeCancellation AdjustmentType = "Cancellation"
)

type Direction string

const (
	DirectionWestbound Direction = "Westbound"

	DirectionEastbound Direction = "Eastbound"
)

type TimeType string

const (
	TimeTypeDeparture TimeType = "Departure"

	TimeTypeArrival TimeType = "Arrival"
)

type LoadIndicator string

const (
	LoadIndicatorPassenger LoadIndicator = "Passenger"

	LoadIndicatorVehicle LoadIndicator = "Vehicle"

	LoadIndicatorBoth LoadIndicator = "Both"
)

type GetActiveScheduledSeasons struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetActiveScheduledSeasons"`
}

type GetActiveScheduledSeasonsResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetActiveScheduledSeasonsResponse"`

	GetActiveScheduledSeasonsResult *ArrayOfSchedBriefResponse `xml:"GetActiveScheduledSeasonsResult,omitempty"`
}

type GetAllAlerts struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetAllAlerts"`
}

type GetAllAlertsResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetAllAlertsResponse"`

	GetAllAlertsResult *ArrayOfAlertResponse `xml:"GetAllAlertsResult,omitempty"`
}

type GetAllRouteDetails struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetAllRouteDetails"`

	Request *TripDateMsg `xml:"request,omitempty"`
}

type GetAllRouteDetailsResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetAllRouteDetailsResponse"`

	GetAllRouteDetailsResult *ArrayOfRouteResponse `xml:"GetAllRouteDetailsResult,omitempty"`
}

type GetAllRoutes struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetAllRoutes"`

	Request *TripDateMsg `xml:"request,omitempty"`
}

type GetAllRoutesResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetAllRoutesResponse"`

	GetAllRoutesResult *ArrayOfRouteBriefResponse `xml:"GetAllRoutesResult,omitempty"`
}

type GetAllRoutesHavingServiceDisruptions struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetAllRoutesHavingServiceDisruptions"`

	Request *TripDateMsg `xml:"request,omitempty"`
}

type GetAllRoutesHavingServiceDisruptionsResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetAllRoutesHavingServiceDisruptionsResponse"`

	GetAllRoutesHavingServiceDisruptionsResult *ArrayOfRouteBriefResponse `xml:"GetAllRoutesHavingServiceDisruptionsResult,omitempty"`
}

type GetAllSchedRoutes struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetAllSchedRoutes"`
}

type GetAllSchedRoutesResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetAllSchedRoutesResponse"`

	GetAllSchedRoutesResult *ArrayOfSchedRouteBriefResponse `xml:"GetAllSchedRoutesResult,omitempty"`
}

type GetAllTerminals struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetAllTerminals"`

	Request *TripDateMsg `xml:"request,omitempty"`
}

type GetAllTerminalsResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetAllTerminalsResponse"`

	GetAllTerminalsResult *ArrayOfTerminalResponse `xml:"GetAllTerminalsResult,omitempty"`
}

type GetAllTerminalsAndMates struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetAllTerminalsAndMates"`

	Request *TripDateMsg `xml:"request,omitempty"`
}

type GetAllTerminalsAndMatesResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetAllTerminalsAndMatesResponse"`

	GetAllTerminalsAndMatesResult *ArrayOfTerminalComboResponse `xml:"GetAllTerminalsAndMatesResult,omitempty"`
}

type GetAllTimeAdj struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetAllTimeAdj"`
}

type GetAllTimeAdjResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetAllTimeAdjResponse"`

	GetAllTimeAdjResult *ArrayOfSchedTimeAdjResponse `xml:"GetAllTimeAdjResult,omitempty"`
}

type GetCacheFlushDate struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetCacheFlushDate"`
}

type GetCacheFlushDateResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetCacheFlushDateResponse"`

	GetCacheFlushDateResult time.Time `xml:"GetCacheFlushDateResult,omitempty"`
}

type GetRouteDetail struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetRouteDetail"`

	Request *RouteMsg `xml:"request,omitempty"`
}

type GetRouteDetailResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetRouteDetailResponse"`

	GetRouteDetailResult *RouteResponse `xml:"GetRouteDetailResult,omitempty"`
}

type GetRouteDetailsByTerminalCombo struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetRouteDetailsByTerminalCombo"`

	Request *TerminalComboMsg `xml:"request,omitempty"`
}

type GetRouteDetailsByTerminalComboResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetRouteDetailsByTerminalComboResponse"`

	GetRouteDetailsByTerminalComboResult *ArrayOfRouteResponse `xml:"GetRouteDetailsByTerminalComboResult,omitempty"`
}

type GetRoutesByTerminalCombo struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetRoutesByTerminalCombo"`

	Request *TerminalComboMsg `xml:"request,omitempty"`
}

type GetRoutesByTerminalComboResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetRoutesByTerminalComboResponse"`

	GetRoutesByTerminalComboResult *ArrayOfRouteBriefResponse `xml:"GetRoutesByTerminalComboResult,omitempty"`
}

type GetSchedRoutesByScheduledSeason struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetSchedRoutesByScheduledSeason"`

	Request *SchedMsg `xml:"request,omitempty"`
}

type GetSchedRoutesByScheduledSeasonResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetSchedRoutesByScheduledSeasonResponse"`

	GetSchedRoutesByScheduledSeasonResult *ArrayOfSchedRouteBriefResponse `xml:"GetSchedRoutesByScheduledSeasonResult,omitempty"`
}

type GetSchedSailingsBySchedRoute struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetSchedSailingsBySchedRoute"`

	Request *SchedRouteMsg `xml:"request,omitempty"`
}

type GetSchedSailingsBySchedRouteResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetSchedSailingsBySchedRouteResponse"`

	GetSchedSailingsBySchedRouteResult *ArrayOfSchedSailingResponse `xml:"GetSchedSailingsBySchedRouteResult,omitempty"`
}

type GetScheduleByRoute struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetScheduleByRoute"`

	Request *RouteMsg `xml:"request,omitempty"`
}

type GetScheduleByRouteResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetScheduleByRouteResponse"`

	GetScheduleByRouteResult *SchedResponse `xml:"GetScheduleByRouteResult,omitempty"`
}

type GetScheduleByTerminalCombo struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetScheduleByTerminalCombo"`

	Request *TerminalComboMsg `xml:"request,omitempty"`
}

type GetScheduleByTerminalComboResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetScheduleByTerminalComboResponse"`

	GetScheduleByTerminalComboResult *SchedResponse `xml:"GetScheduleByTerminalComboResult,omitempty"`
}

type GetTerminalMates struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetTerminalMates"`

	Request *TerminalMsg `xml:"request,omitempty"`
}

type GetTerminalMatesResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetTerminalMatesResponse"`

	GetTerminalMatesResult *ArrayOfTerminalResponse `xml:"GetTerminalMatesResult,omitempty"`
}

type GetTimeAdjByRoute struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetTimeAdjByRoute"`

	Request *RouteBriefMsg `xml:"request,omitempty"`
}

type GetTimeAdjByRouteResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetTimeAdjByRouteResponse"`

	GetTimeAdjByRouteResult *ArrayOfSchedTimeAdjResponse `xml:"GetTimeAdjByRouteResult,omitempty"`
}

type GetTimeAdjBySchedRoute struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetTimeAdjBySchedRoute"`

	Request *SchedRouteMsg `xml:"request,omitempty"`
}

type GetTimeAdjBySchedRouteResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetTimeAdjBySchedRouteResponse"`

	GetTimeAdjBySchedRouteResult *ArrayOfSchedTimeAdjResponse `xml:"GetTimeAdjBySchedRouteResult,omitempty"`
}

type GetTodaysScheduleByRoute struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetTodaysScheduleByRoute"`

	Request *RouteTodayMsg `xml:"request,omitempty"`
}

type GetTodaysScheduleByRouteResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetTodaysScheduleByRouteResponse"`

	GetTodaysScheduleByRouteResult *SchedResponse `xml:"GetTodaysScheduleByRouteResult,omitempty"`
}

type GetTodaysScheduleByTerminalCombo struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetTodaysScheduleByTerminalCombo"`

	Request *TerminalComboTodayMsg `xml:"request,omitempty"`
}

type GetTodaysScheduleByTerminalComboResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetTodaysScheduleByTerminalComboResponse"`

	GetTodaysScheduleByTerminalComboResult *SchedResponse `xml:"GetTodaysScheduleByTerminalComboResult,omitempty"`
}

type GetValidDateRange struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetValidDateRange"`
}

type GetValidDateRangeResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ GetValidDateRangeResponse"`

	GetValidDateRangeResult *ValidDateRangeResponse `xml:"GetValidDateRangeResult,omitempty"`
}

type ArrayOfSchedBriefResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ ArrayOfSchedBriefResponse"`

	SchedBriefResponse []*SchedBriefResponse `xml:"SchedBriefResponse,omitempty"`
}

type SchedBriefResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ SchedBriefResponse"`

	ScheduleID     int32     `xml:"ScheduleID,omitempty"`
	ScheduleName   string    `xml:"ScheduleName,omitempty"`
	ScheduleSeason *Season   `xml:"ScheduleSeason,omitempty"`
	SchedulePDFUrl string    `xml:"SchedulePDFUrl,omitempty"`
	ScheduleStart  time.Time `xml:"ScheduleStart,omitempty"`
	ScheduleEnd    time.Time `xml:"ScheduleEnd,omitempty"`
}

type APIAccessHeader struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ APIAccessHeader"`

	APIAccessCode string `xml:"APIAccessCode,omitempty"`
}

type ArrayOfAlertResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ ArrayOfAlertResponse"`

	AlertResponse []*AlertResponse `xml:"AlertResponse,omitempty"`
}

type AlertResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ AlertResponse"`

	BulletinID            int32       `xml:"BulletinID,omitempty"`
	BulletinFlag          bool        `xml:"BulletinFlag,omitempty"`
	BulletinText          string      `xml:"BulletinText,omitempty"`
	CommunicationFlag     bool        `xml:"CommunicationFlag,omitempty"`
	CommunicationText     string      `xml:"CommunicationText,omitempty"`
	RouteAlertFlag        bool        `xml:"RouteAlertFlag,omitempty"`
	RouteAlertText        string      `xml:"RouteAlertText,omitempty"`
	HomepageAlertText     string      `xml:"HomepageAlertText,omitempty"`
	PublishDate           time.Time   `xml:"PublishDate,omitempty"`
	DisruptionDescription string      `xml:"DisruptionDescription,omitempty"`
	AllRoutesFlag         bool        `xml:"AllRoutesFlag,omitempty"`
	SortSeq               int32       `xml:"SortSeq,omitempty"`
	AlertTypeID           int32       `xml:"AlertTypeID,omitempty"`
	AlertType             string      `xml:"AlertType,omitempty"`
	AlertFullTitle        string      `xml:"AlertFullTitle,omitempty"`
	AffectedRouteIDs      *ArrayOfInt `xml:"AffectedRouteIDs,omitempty"`
}

type ArrayOfInt struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ ArrayOfInt"`

	Int []int32 `xml:"int,omitempty"`
}

type TripDateMsg struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ TripDateMsg"`

	TripDate time.Time `xml:"TripDate,omitempty"`
}

type ArrayOfRouteResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ ArrayOfRouteResponse"`

	RouteResponse []*RouteResponse `xml:"RouteResponse,omitempty"`
}

type RouteResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ RouteResponse"`

	RouteID            int32              `xml:"RouteID,omitempty"`
	RouteAbbrev        string             `xml:"RouteAbbrev,omitempty"`
	Description        string             `xml:"Description,omitempty"`
	RegionID           int32              `xml:"RegionID,omitempty"`
	VesselWatchID      int32              `xml:"VesselWatchID,omitempty"`
	ReservationFlag    bool               `xml:"ReservationFlag,omitempty"`
	InternationalFlag  bool               `xml:"InternationalFlag,omitempty"`
	PassengerOnlyFlag  bool               `xml:"PassengerOnlyFlag,omitempty"`
	CrossingTime       string             `xml:"CrossingTime,omitempty"`
	AdaNotes           string             `xml:"AdaNotes,omitempty"`
	GeneralRouteNotes  string             `xml:"GeneralRouteNotes,omitempty"`
	SeasonalRouteNotes string             `xml:"SeasonalRouteNotes,omitempty"`
	Alerts             *ArrayOfRouteAlert `xml:"Alerts,omitempty"`
}

type ArrayOfRouteAlert struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ ArrayOfRouteAlert"`

	RouteAlert []*RouteAlert `xml:"RouteAlert,omitempty"`
}

type RouteAlert struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ RouteAlert"`

	BulletinID            int32     `xml:"BulletinID,omitempty"`
	BulletinFlag          bool      `xml:"BulletinFlag,omitempty"`
	CommunicationFlag     bool      `xml:"CommunicationFlag,omitempty"`
	PublishDate           time.Time `xml:"PublishDate,omitempty"`
	AlertDescription      string    `xml:"AlertDescription,omitempty"`
	DisruptionDescription string    `xml:"DisruptionDescription,omitempty"`
	AlertFullTitle        string    `xml:"AlertFullTitle,omitempty"`
	AlertFullText         string    `xml:"AlertFullText,omitempty"`
}

type ArrayOfRouteBriefResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ ArrayOfRouteBriefResponse"`

	RouteBriefResponse []*RouteBriefResponse `xml:"RouteBriefResponse,omitempty"`
}

type RouteBriefResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ RouteBriefResponse"`

	RouteID            int32                   `xml:"RouteID,omitempty"`
	RouteAbbrev        string                  `xml:"RouteAbbrev,omitempty"`
	Description        string                  `xml:"Description,omitempty"`
	RegionID           int32                   `xml:"RegionID,omitempty"`
	ServiceDisruptions *ArrayOfRouteBriefAlert `xml:"ServiceDisruptions,omitempty"`
}

type ArrayOfRouteBriefAlert struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ ArrayOfRouteBriefAlert"`

	RouteBriefAlert []*RouteBriefAlert `xml:"RouteBriefAlert,omitempty"`
}

type RouteBriefAlert struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ RouteBriefAlert"`

	BulletinID            int32     `xml:"BulletinID,omitempty"`
	BulletinFlag          bool      `xml:"BulletinFlag,omitempty"`
	PublishDate           time.Time `xml:"PublishDate,omitempty"`
	DisruptionDescription string    `xml:"DisruptionDescription,omitempty"`
}

type ArrayOfSchedRouteBriefResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ ArrayOfSchedRouteBriefResponse"`

	SchedRouteBriefResponse []*SchedRouteBriefResponse `xml:"SchedRouteBriefResponse,omitempty"`
}

type SchedRouteBriefResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ SchedRouteBriefResponse"`

	ScheduleID         int32                   `xml:"ScheduleID,omitempty"`
	SchedRouteID       int32                   `xml:"SchedRouteID,omitempty"`
	ContingencyOnly    bool                    `xml:"ContingencyOnly,omitempty"`
	RouteID            int32                   `xml:"RouteID,omitempty"`
	RouteAbbrev        string                  `xml:"RouteAbbrev,omitempty"`
	Description        string                  `xml:"Description,omitempty"`
	SeasonalRouteNotes string                  `xml:"SeasonalRouteNotes,omitempty"`
	RegionID           int32                   `xml:"RegionID,omitempty"`
	ServiceDisruptions *ArrayOfRouteBriefAlert `xml:"ServiceDisruptions,omitempty"`
	ContingencyAdj     *ArrayOfSchedRouteAdj   `xml:"ContingencyAdj,omitempty"`
}

type ArrayOfSchedRouteAdj struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ ArrayOfSchedRouteAdj"`

	SchedRouteAdj []*SchedRouteAdj `xml:"SchedRouteAdj,omitempty"`
}

type SchedRouteAdj struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ SchedRouteAdj"`

	DateFrom               time.Time       `xml:"DateFrom,omitempty"`
	DateThru               time.Time       `xml:"DateThru,omitempty"`
	EventID                int32           `xml:"EventID,omitempty"`
	EventDescription       string          `xml:"EventDescription,omitempty"`
	AdjType                *AdjustmentType `xml:"AdjType,omitempty"`
	ReplacedBySchedRouteID int32           `xml:"ReplacedBySchedRouteID,omitempty"`
}

type ArrayOfTerminalResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ ArrayOfTerminalResponse"`

	TerminalResponse []*TerminalResponse `xml:"TerminalResponse,omitempty"`
}

type TerminalResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ TerminalResponse"`

	TerminalID  int32  `xml:"TerminalID,omitempty"`
	Description string `xml:"Description,omitempty"`
}

type ArrayOfTerminalComboResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ ArrayOfTerminalComboResponse"`

	TerminalComboResponse []*TerminalComboResponse `xml:"TerminalComboResponse,omitempty"`
}

type TerminalComboResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ TerminalComboResponse"`

	DepartingTerminalID  int32  `xml:"DepartingTerminalID,omitempty"`
	DepartingDescription string `xml:"DepartingDescription,omitempty"`
	ArrivingTerminalID   int32  `xml:"ArrivingTerminalID,omitempty"`
	ArrivingDescription  string `xml:"ArrivingDescription,omitempty"`
}

type ArrayOfSchedTimeAdjResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ ArrayOfSchedTimeAdjResponse"`

	SchedTimeAdjResponse []*SchedTimeAdjResponse `xml:"SchedTimeAdjResponse,omitempty"`
}

type SchedTimeAdjResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ SchedTimeAdjResponse"`

	ScheduleID               int32                   `xml:"ScheduleID,omitempty"`
	SchedRouteID             int32                   `xml:"SchedRouteID,omitempty"`
	RouteID                  int32                   `xml:"RouteID,omitempty"`
	RouteDescription         string                  `xml:"RouteDescription,omitempty"`
	RouteSortSeq             int32                   `xml:"RouteSortSeq,omitempty"`
	SailingID                int32                   `xml:"SailingID,omitempty"`
	SailingDescription       string                  `xml:"SailingDescription,omitempty"`
	ActiveSailingDateRange   *SchedSailingDateRange  `xml:"ActiveSailingDateRange,omitempty"`
	SailingDir               *Direction              `xml:"SailingDir,omitempty"`
	JourneyID                int32                   `xml:"JourneyID,omitempty"`
	VesselID                 int32                   `xml:"VesselID,omitempty"`
	VesselName               string                  `xml:"VesselName,omitempty"`
	VesselHandicapAccessible bool                    `xml:"VesselHandicapAccessible,omitempty"`
	VesselPositionNum        int32                   `xml:"VesselPositionNum,omitempty"`
	JourneyTerminalID        int32                   `xml:"JourneyTerminalID,omitempty"`
	TerminalID               int32                   `xml:"TerminalID,omitempty"`
	TerminalDescription      string                  `xml:"TerminalDescription,omitempty"`
	TerminalBriefDescription string                  `xml:"TerminalBriefDescription,omitempty"`
	TimeToAdj                time.Time               `xml:"TimeToAdj,omitempty"`
	AdjDateFrom              time.Time               `xml:"AdjDateFrom,omitempty"`
	AdjDateThru              time.Time               `xml:"AdjDateThru,omitempty"`
	TidalAdj                 bool                    `xml:"TidalAdj,omitempty"`
	EventID                  int32                   `xml:"EventID,omitempty"`
	EventDescription         string                  `xml:"EventDescription,omitempty"`
	DepArrIndicator          *TimeType               `xml:"DepArrIndicator,omitempty"`
	AdjType                  *AdjustmentType         `xml:"AdjType,omitempty"`
	Annotations              *ArrayOfSchedAnnotation `xml:"Annotations,omitempty"`
}

type SchedSailingDateRange struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ SchedSailingDateRange"`

	DateFrom         time.Time `xml:"DateFrom,omitempty"`
	DateThru         time.Time `xml:"DateThru,omitempty"`
	EventID          int32     `xml:"EventID,omitempty"`
	EventDescription string    `xml:"EventDescription,omitempty"`
}

type ArrayOfSchedAnnotation struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ ArrayOfSchedAnnotation"`

	SchedAnnotation []*SchedAnnotation `xml:"SchedAnnotation,omitempty"`
}

type SchedAnnotation struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ SchedAnnotation"`

	AnnotationID         int32  `xml:"AnnotationID,omitempty"`
	AnnotationText       string `xml:"AnnotationText,omitempty"`
	AnnotationIVRText    string `xml:"AnnotationIVRText,omitempty"`
	AdjustedCrossingTime int32  `xml:"AdjustedCrossingTime,omitempty"`
	AnnotationImg        string `xml:"AnnotationImg,omitempty"`
	TypeDescription      string `xml:"TypeDescription,omitempty"`
	SortSeq              int32  `xml:"SortSeq,omitempty"`
}

type RouteMsg struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ RouteMsg"`

	TripDate time.Time `xml:"TripDate,omitempty"`
	RouteID  int32     `xml:"RouteID,omitempty"`
}

type TerminalComboMsg struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ TerminalComboMsg"`

	TripDate            time.Time `xml:"TripDate,omitempty"`
	DepartingTerminalID int32     `xml:"DepartingTerminalID,omitempty"`
	ArrivingTerminalID  int32     `xml:"ArrivingTerminalID,omitempty"`
}

type SchedMsg struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ SchedMsg"`

	ScheduleID int32 `xml:"ScheduleID,omitempty"`
}

type SchedRouteMsg struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ SchedRouteMsg"`

	SchedRouteID int32 `xml:"SchedRouteID,omitempty"`
}

type ArrayOfSchedSailingResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ ArrayOfSchedSailingResponse"`

	SchedSailingResponse []*SchedSailingResponse `xml:"SchedSailingResponse,omitempty"`
}

type SchedSailingResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ SchedSailingResponse"`

	ScheduleID         int32                         `xml:"ScheduleID,omitempty"`
	SchedRouteID       int32                         `xml:"SchedRouteID,omitempty"`
	RouteID            int32                         `xml:"RouteID,omitempty"`
	SailingID          int32                         `xml:"SailingID,omitempty"`
	SailingDescription string                        `xml:"SailingDescription,omitempty"`
	SailingNotes       string                        `xml:"SailingNotes,omitempty"`
	DisplayColNum      int32                         `xml:"DisplayColNum,omitempty"`
	SailingDir         *Direction                    `xml:"SailingDir,omitempty"`
	DayOpDescription   string                        `xml:"DayOpDescription,omitempty"`
	DayOpUseForHoliday bool                          `xml:"DayOpUseForHoliday,omitempty"`
	ActiveDateRanges   *ArrayOfSchedSailingDateRange `xml:"ActiveDateRanges,omitempty"`
	Journs             *ArrayOfSchedJourn            `xml:"Journs,omitempty"`
}

type ArrayOfSchedSailingDateRange struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ ArrayOfSchedSailingDateRange"`

	SchedSailingDateRange []*SchedSailingDateRange `xml:"SchedSailingDateRange,omitempty"`
}

type ArrayOfSchedJourn struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ ArrayOfSchedJourn"`

	SchedJourn []*SchedJourn `xml:"SchedJourn,omitempty"`
}

type SchedJourn struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ SchedJourn"`

	JourneyID                int32                     `xml:"JourneyID,omitempty"`
	ReservationInd           bool                      `xml:"ReservationInd,omitempty"`
	InternationalInd         bool                      `xml:"InternationalInd,omitempty"`
	InterislandInd           bool                      `xml:"InterislandInd,omitempty"`
	VesselID                 int32                     `xml:"VesselID,omitempty"`
	VesselName               string                    `xml:"VesselName,omitempty"`
	VesselHandicapAccessible bool                      `xml:"VesselHandicapAccessible,omitempty"`
	VesselPositionNum        int32                     `xml:"VesselPositionNum,omitempty"`
	TerminalTimes            *ArrayOfSchedTimeTerminal `xml:"TerminalTimes,omitempty"`
}

type ArrayOfSchedTimeTerminal struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ ArrayOfSchedTimeTerminal"`

	SchedTimeTerminal []*SchedTimeTerminal `xml:"SchedTimeTerminal,omitempty"`
}

type SchedTimeTerminal struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ SchedTimeTerminal"`

	JourneyTerminalID        int32                   `xml:"JourneyTerminalID,omitempty"`
	TerminalID               int32                   `xml:"TerminalID,omitempty"`
	TerminalDescription      string                  `xml:"TerminalDescription,omitempty"`
	TerminalBriefDescription string                  `xml:"TerminalBriefDescription,omitempty"`
	Time                     time.Time               `xml:"Time,omitempty"`
	DepArrIndicator          *TimeType               `xml:"DepArrIndicator,omitempty"`
	IsNA                     bool                    `xml:"IsNA,omitempty"`
	Annotations              *ArrayOfSchedAnnotation `xml:"Annotations,omitempty"`
}

type SchedResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ SchedResponse"`

	ScheduleID     int32                      `xml:"ScheduleID,omitempty"`
	ScheduleName   string                     `xml:"ScheduleName,omitempty"`
	ScheduleSeason *Season                    `xml:"ScheduleSeason,omitempty"`
	SchedulePDFUrl string                     `xml:"SchedulePDFUrl,omitempty"`
	ScheduleStart  time.Time                  `xml:"ScheduleStart,omitempty"`
	ScheduleEnd    time.Time                  `xml:"ScheduleEnd,omitempty"`
	AllRoutes      *ArrayOfInt                `xml:"AllRoutes,omitempty"`
	TerminalCombos *ArrayOfSchedTerminalCombo `xml:"TerminalCombos,omitempty"`
}

type ArrayOfSchedTerminalCombo struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ ArrayOfSchedTerminalCombo"`

	SchedTerminalCombo []*SchedTerminalCombo `xml:"SchedTerminalCombo,omitempty"`
}

type SchedTerminalCombo struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ SchedTerminalCombo"`

	DepartingTerminalID   int32             `xml:"DepartingTerminalID,omitempty"`
	DepartingTerminalName string            `xml:"DepartingTerminalName,omitempty"`
	ArrivingTerminalID    int32             `xml:"ArrivingTerminalID,omitempty"`
	ArrivingTerminalName  string            `xml:"ArrivingTerminalName,omitempty"`
	SailingNotes          string            `xml:"SailingNotes,omitempty"`
	Annotations           *ArrayOfString    `xml:"Annotations,omitempty"`
	Times                 *ArrayOfSchedTime `xml:"Times,omitempty"`
}

type ArrayOfString struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ ArrayOfString"`

	String []string `xml:"string,omitempty"`
}

type ArrayOfSchedTime struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ ArrayOfSchedTime"`

	SchedTime []*SchedTime `xml:"SchedTime,omitempty"`
}

type SchedTime struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ SchedTime"`

	DepartingTime            time.Time      `xml:"DepartingTime,omitempty"`
	ArrivingTime             time.Time      `xml:"ArrivingTime,omitempty"`
	LoadingRule              *LoadIndicator `xml:"LoadingRule,omitempty"`
	VesselID                 int32          `xml:"VesselID,omitempty"`
	VesselName               string         `xml:"VesselName,omitempty"`
	VesselHandicapAccessible bool           `xml:"VesselHandicapAccessible,omitempty"`
	VesselPositionNum        int32          `xml:"VesselPositionNum,omitempty"`
	Routes                   *ArrayOfInt    `xml:"Routes,omitempty"`
	AnnotationIndexes        *ArrayOfInt    `xml:"AnnotationIndexes,omitempty"`
}

type TerminalMsg struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ TerminalMsg"`

	TripDate   time.Time `xml:"TripDate,omitempty"`
	TerminalID int32     `xml:"TerminalID,omitempty"`
}

type RouteBriefMsg struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ RouteBriefMsg"`

	RouteID int32 `xml:"RouteID,omitempty"`
}

type RouteTodayMsg struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ RouteTodayMsg"`

	RouteID            int32 `xml:"RouteID,omitempty"`
	OnlyRemainingTimes bool  `xml:"OnlyRemainingTimes,omitempty"`
}

type TerminalComboTodayMsg struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ TerminalComboTodayMsg"`

	DepartingTerminalID int32 `xml:"DepartingTerminalID,omitempty"`
	ArrivingTerminalID  int32 `xml:"ArrivingTerminalID,omitempty"`
	OnlyRemainingTimes  bool  `xml:"OnlyRemainingTimes,omitempty"`
}

type ValidDateRangeResponse struct {
	XMLName xml.Name `xml:"http://www.wsdot.wa.gov/ferries/schedule/ ValidDateRangeResponse"`

	DateFrom time.Time `xml:"DateFrom,omitempty"`
	DateThru time.Time `xml:"DateThru,omitempty"`
}

type WSF_x0020_ScheduleSoap struct {
	client *SOAPClient
}

func NewWSF_x0020_ScheduleSoap(url string, tls bool, auth *BasicAuth) *WSF_x0020_ScheduleSoap {
	if url == "" {
		url = "http://b2b.wsdot.wa.gov/ferries/schedule/Default.asmx"
	}
	client := NewSOAPClient(url, tls, auth)

	return &WSF_x0020_ScheduleSoap{
		client: client,
	}
}

/* Provides a brief summary of all scheduled sailing seasons that are currently active / available. */
func (service *WSF_x0020_ScheduleSoap) GetActiveScheduledSeasons(request *GetActiveScheduledSeasons) (*GetActiveScheduledSeasonsResponse, error) {
	response := new(GetActiveScheduledSeasonsResponse)
	err := service.client.Call("http://www.wsdot.wa.gov/ferries/schedule/GetActiveScheduledSeasons", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Retrieves all published alerts. */
func (service *WSF_x0020_ScheduleSoap) GetAllAlerts(request *GetAllAlerts) (*GetAllAlertsResponse, error) {
	response := new(GetAllAlertsResponse)
	err := service.client.Call("http://www.wsdot.wa.gov/ferries/schedule/GetAllAlerts", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Provides detailed information for all available routes pertaining to a particular date. */
func (service *WSF_x0020_ScheduleSoap) GetAllRouteDetails(request *GetAllRouteDetails) (*GetAllRouteDetailsResponse, error) {
	response := new(GetAllRouteDetailsResponse)
	err := service.client.Call("http://www.wsdot.wa.gov/ferries/schedule/GetAllRouteDetails", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Provides all available routes for a particular date. */
func (service *WSF_x0020_ScheduleSoap) GetAllRoutes(request *GetAllRoutes) (*GetAllRoutesResponse, error) {
	response := new(GetAllRoutesResponse)
	err := service.client.Call("http://www.wsdot.wa.gov/ferries/schedule/GetAllRoutes", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Provides all available routes for a particular date where one or more service disruptions are present. */
func (service *WSF_x0020_ScheduleSoap) GetAllRoutesHavingServiceDisruptions(request *GetAllRoutesHavingServiceDisruptions) (*GetAllRoutesHavingServiceDisruptionsResponse, error) {
	response := new(GetAllRoutesHavingServiceDisruptionsResponse)
	err := service.client.Call("http://www.wsdot.wa.gov/ferries/schedule/GetAllRoutesHavingServiceDisruptions", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Retrieves the scheduled route(s) for all seasons that are currently active / available. */
func (service *WSF_x0020_ScheduleSoap) GetAllSchedRoutes(request *GetAllSchedRoutes) (*GetAllSchedRoutesResponse, error) {
	response := new(GetAllSchedRoutesResponse)
	err := service.client.Call("http://www.wsdot.wa.gov/ferries/schedule/GetAllSchedRoutes", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Provides all available terminals for a particular date. */
func (service *WSF_x0020_ScheduleSoap) GetAllTerminals(request *GetAllTerminals) (*GetAllTerminalsResponse, error) {
	response := new(GetAllTerminalsResponse)
	err := service.client.Call("http://www.wsdot.wa.gov/ferries/schedule/GetAllTerminals", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* For a given date, retrieves all available terminal combinations. */
func (service *WSF_x0020_ScheduleSoap) GetAllTerminalsAndMates(request *GetAllTerminalsAndMates) (*GetAllTerminalsAndMatesResponse, error) {
	response := new(GetAllTerminalsAndMatesResponse)
	err := service.client.Call("http://www.wsdot.wa.gov/ferries/schedule/GetAllTerminalsAndMates", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Provides a list of all individual time adjustments (additions or cancellations) that are currently active / available. */
func (service *WSF_x0020_ScheduleSoap) GetAllTimeAdj(request *GetAllTimeAdj) (*GetAllTimeAdjResponse, error) {
	response := new(GetAllTimeAdjResponse)
	err := service.client.Call("http://www.wsdot.wa.gov/ferries/schedule/GetAllTimeAdj", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Most web methods in this service are cached.  If you are also using caching in your user interface, it may be helpful to know the date and time that the cache was last flushed in this web service. */
func (service *WSF_x0020_ScheduleSoap) GetCacheFlushDate(request *GetCacheFlushDate) (*GetCacheFlushDateResponse, error) {
	response := new(GetCacheFlushDateResponse)
	err := service.client.Call("http://www.wsdot.wa.gov/ferries/schedule/GetCacheFlushDate", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Retrieves detailed information pertaining to a scheduled route. */
func (service *WSF_x0020_ScheduleSoap) GetRouteDetail(request *GetRouteDetail) (*GetRouteDetailResponse, error) {
	response := new(GetRouteDetailResponse)
	err := service.client.Call("http://www.wsdot.wa.gov/ferries/schedule/GetRouteDetail", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Retrieves detailed information for scheduled routes that are associated with a particular terminal combination. */
func (service *WSF_x0020_ScheduleSoap) GetRouteDetailsByTerminalCombo(request *GetRouteDetailsByTerminalCombo) (*GetRouteDetailsByTerminalComboResponse, error) {
	response := new(GetRouteDetailsByTerminalComboResponse)
	err := service.client.Call("http://www.wsdot.wa.gov/ferries/schedule/GetRouteDetailsByTerminalCombo", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Retrieves route(s) for a particular date and terminal combination. */
func (service *WSF_x0020_ScheduleSoap) GetRoutesByTerminalCombo(request *GetRoutesByTerminalCombo) (*GetRoutesByTerminalComboResponse, error) {
	response := new(GetRoutesByTerminalComboResponse)
	err := service.client.Call("http://www.wsdot.wa.gov/ferries/schedule/GetRoutesByTerminalCombo", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Retrieves scheduled route(s) for a particular active season. */
func (service *WSF_x0020_ScheduleSoap) GetSchedRoutesByScheduledSeason(request *GetSchedRoutesByScheduledSeason) (*GetSchedRoutesByScheduledSeasonResponse, error) {
	response := new(GetSchedRoutesByScheduledSeasonResponse)
	err := service.client.Call("http://www.wsdot.wa.gov/ferries/schedule/GetSchedRoutesByScheduledSeason", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Retrieves sailings and departure/arrival times that correspond with a particular scheduled route. */
func (service *WSF_x0020_ScheduleSoap) GetSchedSailingsBySchedRoute(request *GetSchedSailingsBySchedRoute) (*GetSchedSailingsBySchedRouteResponse, error) {
	response := new(GetSchedSailingsBySchedRouteResponse)
	err := service.client.Call("http://www.wsdot.wa.gov/ferries/schedule/GetSchedSailingsBySchedRoute", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Retrieves sailing times associated with a specific route for a particular date. */
func (service *WSF_x0020_ScheduleSoap) GetScheduleByRoute(request *GetScheduleByRoute) (*GetScheduleByRouteResponse, error) {
	response := new(GetScheduleByRouteResponse)
	err := service.client.Call("http://www.wsdot.wa.gov/ferries/schedule/GetScheduleByRoute", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Retrieves sailing times associated with a specific departing / arriving terminal combination for a particular date. */
func (service *WSF_x0020_ScheduleSoap) GetScheduleByTerminalCombo(request *GetScheduleByTerminalCombo) (*GetScheduleByTerminalComboResponse, error) {
	response := new(GetScheduleByTerminalComboResponse)
	err := service.client.Call("http://www.wsdot.wa.gov/ferries/schedule/GetScheduleByTerminalCombo", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Provides all available terminals that correspond to a given terminal for a particular date. */
func (service *WSF_x0020_ScheduleSoap) GetTerminalMates(request *GetTerminalMates) (*GetTerminalMatesResponse, error) {
	response := new(GetTerminalMatesResponse)
	err := service.client.Call("http://www.wsdot.wa.gov/ferries/schedule/GetTerminalMates", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Provides a list of individual time adjustments (additions or cancellations) for a particular route. */
func (service *WSF_x0020_ScheduleSoap) GetTimeAdjByRoute(request *GetTimeAdjByRoute) (*GetTimeAdjByRouteResponse, error) {
	response := new(GetTimeAdjByRouteResponse)
	err := service.client.Call("http://www.wsdot.wa.gov/ferries/schedule/GetTimeAdjByRoute", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Provides a list of individual time adjustments (additions or cancellations) for a particular scheduled route. */
func (service *WSF_x0020_ScheduleSoap) GetTimeAdjBySchedRoute(request *GetTimeAdjBySchedRoute) (*GetTimeAdjBySchedRouteResponse, error) {
	response := new(GetTimeAdjBySchedRouteResponse)
	err := service.client.Call("http://www.wsdot.wa.gov/ferries/schedule/GetTimeAdjBySchedRoute", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Retrieves sailing times associated with a specific route for the current date.  User may specify if only the times for the remainder of this sailing date are required. */
func (service *WSF_x0020_ScheduleSoap) GetTodaysScheduleByRoute(request *GetTodaysScheduleByRoute) (*GetTodaysScheduleByRouteResponse, error) {
	response := new(GetTodaysScheduleByRouteResponse)
	err := service.client.Call("http://www.wsdot.wa.gov/ferries/schedule/GetTodaysScheduleByRoute", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Retrieves sailing times associated with a specific departing / arriving terminal combination for the current date.  User may specify if only the times for the remainder of this sailing date are required. */
func (service *WSF_x0020_ScheduleSoap) GetTodaysScheduleByTerminalCombo(request *GetTodaysScheduleByTerminalCombo) (*GetTodaysScheduleByTerminalComboResponse, error) {
	response := new(GetTodaysScheduleByTerminalComboResponse)
	err := service.client.Call("http://www.wsdot.wa.gov/ferries/schedule/GetTodaysScheduleByTerminalCombo", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Reveals a valid date range for retrieving schedule data.  This begins with today's date and extends to the end of the most recently posted schedule. */
func (service *WSF_x0020_ScheduleSoap) GetValidDateRange(request *GetValidDateRange) (*GetValidDateRangeResponse, error) {
	response := new(GetValidDateRangeResponse)
	err := service.client.Call("http://www.wsdot.wa.gov/ferries/schedule/GetValidDateRange", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type WSF_x0020_ScheduleHttpGet struct {
	client *SOAPClient
}

func NewWSF_x0020_ScheduleHttpGet(url string, tls bool, auth *BasicAuth) *WSF_x0020_ScheduleHttpGet {
	if url == "" {
		url = ""
	}
	client := NewSOAPClient(url, tls, auth)

	return &WSF_x0020_ScheduleHttpGet{
		client: client,
	}
}

/* Provides a brief summary of all scheduled sailing seasons that are currently active / available. */
func (service *WSF_x0020_ScheduleHttpGet) GetActiveScheduledSeasons() (*ArrayOfSchedBriefResponse, error) {
	response := new(ArrayOfSchedBriefResponse)
	err := service.client.Call("", nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Retrieves all published alerts. */
func (service *WSF_x0020_ScheduleHttpGet) GetAllAlerts() (*ArrayOfAlertResponse, error) {
	response := new(ArrayOfAlertResponse)
	err := service.client.Call("", nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Retrieves the scheduled route(s) for all seasons that are currently active / available. */
func (service *WSF_x0020_ScheduleHttpGet) GetAllSchedRoutes() (*ArrayOfSchedRouteBriefResponse, error) {
	response := new(ArrayOfSchedRouteBriefResponse)
	err := service.client.Call("", nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Provides a list of all individual time adjustments (additions or cancellations) that are currently active / available. */
func (service *WSF_x0020_ScheduleHttpGet) GetAllTimeAdj() (*ArrayOfSchedTimeAdjResponse, error) {
	response := new(ArrayOfSchedTimeAdjResponse)
	err := service.client.Call("", nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Most web methods in this service are cached.  If you are also using caching in your user interface, it may be helpful to know the date and time that the cache was last flushed in this web service. */
//func (service *WSF_x0020_ScheduleHttpGet) GetCacheFlushDate() (*DateTime, error) {
//	response := new(DateTime)
//	err := service.client.Call("", nil, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}

/* Reveals a valid date range for retrieving schedule data.  This begins with today's date and extends to the end of the most recently posted schedule. */
func (service *WSF_x0020_ScheduleHttpGet) GetValidDateRange() (*ValidDateRangeResponse, error) {
	response := new(ValidDateRangeResponse)
	err := service.client.Call("", nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type WSF_x0020_ScheduleHttpPost struct {
	client *SOAPClient
}

func NewWSF_x0020_ScheduleHttpPost(url string, tls bool, auth *BasicAuth) *WSF_x0020_ScheduleHttpPost {
	if url == "" {
		url = ""
	}
	client := NewSOAPClient(url, tls, auth)

	return &WSF_x0020_ScheduleHttpPost{
		client: client,
	}
}

/* Provides a brief summary of all scheduled sailing seasons that are currently active / available. */
func (service *WSF_x0020_ScheduleHttpPost) GetActiveScheduledSeasons() (*ArrayOfSchedBriefResponse, error) {
	response := new(ArrayOfSchedBriefResponse)
	err := service.client.Call("", nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Retrieves all published alerts. */
func (service *WSF_x0020_ScheduleHttpPost) GetAllAlerts() (*ArrayOfAlertResponse, error) {
	response := new(ArrayOfAlertResponse)
	err := service.client.Call("", nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Retrieves the scheduled route(s) for all seasons that are currently active / available. */
func (service *WSF_x0020_ScheduleHttpPost) GetAllSchedRoutes() (*ArrayOfSchedRouteBriefResponse, error) {
	response := new(ArrayOfSchedRouteBriefResponse)
	err := service.client.Call("", nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Provides a list of all individual time adjustments (additions or cancellations) that are currently active / available. */
func (service *WSF_x0020_ScheduleHttpPost) GetAllTimeAdj() (*ArrayOfSchedTimeAdjResponse, error) {
	response := new(ArrayOfSchedTimeAdjResponse)
	err := service.client.Call("", nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Most web methods in this service are cached.  If you are also using caching in your user interface, it may be helpful to know the date and time that the cache was last flushed in this web service. */
//func (service *WSF_x0020_ScheduleHttpPost) GetCacheFlushDate() (*DateTime, error) {
//	response := new(DateTime)
//	err := service.client.Call("", nil, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}

/* Reveals a valid date range for retrieving schedule data.  This begins with today's date and extends to the end of the most recently posted schedule. */
func (service *WSF_x0020_ScheduleHttpPost) GetValidDateRange() (*ValidDateRangeResponse, error) {
	response := new(ValidDateRangeResponse)
	err := service.client.Call("", nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

var timeout = time.Duration(30 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`

	Body SOAPBody
}

type SOAPHeader struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`

	Header interface{}
}

type SOAPBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`

	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

	Code   string `xml:"faultcode,omitempty"`
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`
	Detail string `xml:"detail,omitempty"`
}

type BasicAuth struct {
	Login    string
	Password string
}

type SOAPClient struct {
	url  string
	tls  bool
	auth *BasicAuth
}

func (b *SOAPBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

func (f *SOAPFault) Error() string {
	return f.String
}

func NewSOAPClient(url string, tls bool, auth *BasicAuth) *SOAPClient {
	return &SOAPClient{
		url:  url,
		tls:  tls,
		auth: auth,
	}
}

func (s *SOAPClient) Call(soapAction string, request, response interface{}) error {
	envelope := SOAPEnvelope{
	//Header:        SoapHeader{},
	}

	envelope.Body.Content = request
	buffer := new(bytes.Buffer)

	encoder := xml.NewEncoder(buffer)
	//encoder.Indent("  ", "    ")

	err := encoder.Encode(envelope)
	if err == nil {
		err = encoder.Flush()
	}

	log.Println(buffer.String())
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", s.url, buffer)
	if s.auth != nil {
		req.SetBasicAuth(s.auth.Login, s.auth.Password)
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	if soapAction != "" {
		req.Header.Add("SOAPAction", soapAction)
	}

	req.Header.Set("User-Agent", "gowsdl/0.1")
	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: s.tls,
		},
		Dial: dialTimeout,
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	rawbody, err := ioutil.ReadAll(res.Body)
	if len(rawbody) == 0 {
		log.Println("empty response")
		return nil
	}

	log.Println(string(rawbody))
	respEnvelope := new(SOAPEnvelope)
	respEnvelope.Body = SOAPBody{Content: response}
	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fault
	}

	return nil
}
