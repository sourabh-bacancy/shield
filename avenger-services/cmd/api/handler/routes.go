package handler

import (
	"avenger/cmd/api/handler/controller"

	"github.com/gorilla/mux"
)

// func AddRouteHandlers(router *httprouter.Router) {
// 	router.POST("/api/addAvenger", controller.AddAvenger)
// 	router.POST("/api/assignMission", controller.AssignMission)
// 	router.GET("/api/allMissionsSatus", controller.AllMissionsSatus)
// 	router.POST("/api/missionDetails", controller.MissionDetails)
// 	router.POST("/api/updateMissionStatus", controller.UpdateMissionStatus)
// 	router.GET("/api/allAvengersSatus", controller.AllAvengersSatus)

// }

func AddRouteHandlers(router *mux.Router) {
	router.HandleFunc("/api/addAvenger", controller.AddAvenger).Methods("POST")
	router.HandleFunc("/api/assignMission", controller.AssignMission).Methods("POST")
	router.HandleFunc("/api/allMissionsSatus", controller.AllMissionsSatus).Methods("GET")
	router.HandleFunc("/api/missionDetails", controller.MissionDetails).Methods("POST")
	router.HandleFunc("/api/updateMissionStatus", controller.UpdateMissionStatus).Methods("POST")
	router.HandleFunc("/api/allAvengersSatus", controller.AllAvengersSatus).Methods("GET")
}
