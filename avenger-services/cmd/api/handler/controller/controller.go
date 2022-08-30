package controller

import (
	"avenger/cmd/api/handler/model"
	"avenger/cmd/api/handler/util"
	"encoding/json"
	"net/http"
	"strings"
)

var Avengers = make(map[string]model.Avenger)

// AddAvenger : To add new avengers with association
func AddAvenger(w http.ResponseWriter, r *http.Request) {

	var avenger model.Avenger
	err := json.NewDecoder(r.Body).Decode(&avenger)
	if err != nil {
		util.ErrorResponse(w, r, http.StatusBadRequest, nil, "JSON not supported")
		return
	}

	if avenger.Name == "" {
		util.ErrorResponse(w, r, http.StatusBadRequest, nil, "Name is required")
	}
	if avenger.Abilities == "" {
		util.ErrorResponse(w, r, http.StatusBadRequest, nil, "Abilities are required")
	}
	if avenger.CommsType == "" {
		util.ErrorResponse(w, r, http.StatusBadRequest, nil, "CommsType is required")
	}

	Avengers[avenger.Name] = avenger

	// NEXT we can add the details to a database from here once we got in request and process the data according

	util.SuccessResponse(w, r, http.StatusOK, nil, "Avenger added successfully")
}

var AvengersMission []model.Mission
var AvengersMissions = make(map[string]model.Mission)
var avengerMissionCount int
var avengerForMission string

// AssignMission : To assign new mission to Avenger(s)
func AssignMission(w http.ResponseWriter, r *http.Request) {

	var mission model.Mission

	err := json.NewDecoder(r.Body).Decode(&mission)
	if err != nil {
		util.ErrorResponse(w, r, http.StatusBadRequest, nil, "JSON not supported")
		return
	}

	if len(AvengersMission) > 0 {
		for _, avngr := range mission.AssignedAvengers {
			if _, ok := Avengers[avngr]; ok {
				avengerForMission = avngr
				avengerMissionCount++
			}
		}
		if avengerMissionCount >= 2 && (mission.Status == "ASSIGNED" || mission.Status == "assigned") {
			util.ErrorResponse(w, r, http.StatusBadRequest, nil, "Cannot assign new mission to "+avengerForMission+" as already two assigned.")
			// avengerForMission = ""
			// avengerMissionCount = 0
			return
		}
	}

	AvengersMission = append(AvengersMission, mission)
	AvengersMissions[mission.Name] = mission

	util.SuccessResponse(w, r, http.StatusOK, nil, "Mission created successfully")
}

// AllMissionsSatus : To get all missions status
func AllMissionsSatus(w http.ResponseWriter, r *http.Request) {
	if len(AvengersMissions) > 0 {
		var missions []model.Mission
		for _, mission := range AvengersMissions {
			missions = append(missions, mission)
		}
		util.SuccessResponse(w, r, http.StatusOK, missions, "Success")
		return
	}

	util.ErrorResponse(w, r, http.StatusNotFound, nil, "Missions records not found")
}

// MissionDetails : To get a mission details on Mission name
func MissionDetails(w http.ResponseWriter, r *http.Request) {
	var mission model.Mission

	err := json.NewDecoder(r.Body).Decode(&mission)
	if err != nil {
		util.ErrorResponse(w, r, http.StatusBadRequest, nil, "JSON not supported")
		return
	}

	if mission, ok := AvengersMissions[mission.Name]; ok {
		util.SuccessResponse(w, r, http.StatusOK, mission, "Success")
		return
	}

	util.ErrorResponse(w, r, http.StatusNotFound, nil, "Mission not avilable")
}

// UpdateMissionStatus : To get a mission details on Mission name
func UpdateMissionStatus(w http.ResponseWriter, r *http.Request) {
	var mission model.Mission

	err := json.NewDecoder(r.Body).Decode(&mission)
	if err != nil {
		util.ErrorResponse(w, r, http.StatusBadRequest, nil, "JSON not supported")
		return
	}

	if avngrMission, ok := AvengersMissions[mission.Name]; ok {
		previousMissionStatus := avngrMission.Status
		avngrMission.Status = mission.Status
		AvengersMissions[mission.Name] = avngrMission

		util.SuccessResponse(w, r, http.StatusOK, nil, "Mission status updated from "+previousMissionStatus+" to "+avngrMission.Status)
		return
	}

	util.ErrorResponse(w, r, http.StatusNotFound, nil, "Mission not avilable")
}

var AvengersDetails = make(map[string][]string)
var avngrMissionCount = make(map[string]int)

// AllAvengersSatus : To get all avengers status
func AllAvengersSatus(w http.ResponseWriter, r *http.Request) {
	if len(AvengersMissions) > 0 {
		for _, mission := range AvengersMissions {
			for _, avngr := range mission.AssignedAvengers {
				if _, ok := Avengers[avngr]; ok {
					AvengersDetails[avngr] = append(AvengersDetails[avngr], mission.Name)
					avngrMissionCount[avngr] += 1
				}
			}
		}

		var avngrsStatus []model.AvengersStatus
		for avngr, details := range AvengersDetails {
			var avngrStats model.AvengersStatus
			avngrStats.Name = avngr
			avngrStats.AssignedMissions = strings.Join(details, ",")
			if missionCount, ok := avngrMissionCount[avngr]; ok {
				if missionCount >= 2 {
					avngrStats.Status = "Unavailable"
				}
				avngrStats.Status = "Available"
			}
			avngrsStatus = append(avngrsStatus, avngrStats)
		}

		util.SuccessResponse(w, r, http.StatusOK, avngrsStatus, "Success")
		return
	}

	util.ErrorResponse(w, r, http.StatusNotFound, nil, "Missions records not found")
}
