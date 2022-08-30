package model

type Avenger struct {
	Name      string `json:"name"`
	Abilities string `json:"abilities"`
	CommsType string `json:"communication_type"`
}

type Mission struct {
	Name             string   `json:"mission_name"`
	Description      string   `json:"mission_details"`
	Status           string   `json:"mission_status"`
	AssignedAvengers []string `json:"assigned_avengers"`
}

type AvengersStatus struct {
	Name             string `json:"name"`
	Status           string `json:"availability"`
	AssignedMissions string `json:"assigned_missions"`
}
