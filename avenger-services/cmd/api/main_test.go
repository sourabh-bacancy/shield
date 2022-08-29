package main

import (
	"avenger/cmd/api/handler/controller"
	"avenger/cmd/api/handler/controller/util"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddAvenger(t *testing.T) {
	var jsonStr = []byte(`{"name": "Captain America","abilities": "Ariel Combat","communication_type": "Pager"}`)

	req, err := http.NewRequest("POST", "/api/addAvenger", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.AddAvenger)
	handler.ServeHTTP(res, req)
	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response util.ResponseHandler
	err = json.Unmarshal(res.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), `{
            "status": true,
            "message": "Avenger added successfully",
            "data": null
        }`)
	}

	expected := `Avenger added successfully`
	if response.Message != expected && !response.Status {
		t.Errorf("handler returned unexpected message: got %v want %v",
			response.Message, expected)
	}
}

func TestAPIs(t *testing.T) {

	// Test AddAvenger ---------------------------------
	var jsonStr = []byte(`{
		"mission_name": "MissionX",
		"mission_details": "X Combat",
		"mission_status": "ASSIGNED",
		"assigned_avengers" : ["Captain America"] 
	}`)

	req, err := http.NewRequest("POST", "/api/assignMission", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.AssignMission)
	handler.ServeHTTP(res, req)
	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response util.ResponseHandler
	err = json.Unmarshal(res.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), `{
				"status": true,
				"message": "Mission created successfully",
				"data": null
			}`)
	}

	expected := `Success`
	if response.Message != expected && !response.Status {
		t.Errorf("handler returned unexpected message: got %v want %v",
			response.Message, expected)
	}
	fmt.Println("Pass : AddAvenger")

	// Test AllMissionsSatus ---------------------------------
	req2, err2 := http.NewRequest("GET", "/api/allMissionsSatus", nil)
	if err2 != nil {
		t.Fatal(err)
	}
	res2 := httptest.NewRecorder()
	handler2 := http.HandlerFunc(controller.AllMissionsSatus)
	handler2.ServeHTTP(res2, req2)
	if status := res2.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response2 util.ResponseHandler
	err = json.Unmarshal(res2.Body.Bytes(), &response2)
	if err != nil {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res2.Body.String(), `{
	    			"status": true,
	    			"message": "Success",
	    			"data": [
	    			    {
	    			        "mission_name": "MissionX",
							"mission_details": "X Combat",
							"mission_status": "ASSIGNED",
							"assigned_avengers" : ["Captain America"] 
	    			    }
	    			]
				}`)
	}

	expected2 := `Success`
	if response2.Message != expected && !response2.Status {
		t.Errorf("handler returned unexpected message: got %v want %v",
			response2.Message, expected2)
	}
	fmt.Println("Pass : AllMissionsSatus")

	// Test MissionDetails ---------------------------------
	req3, err3 := http.NewRequest("POST", "/api/missionDetails", bytes.NewBuffer(jsonStr))
	if err3 != nil {
		t.Fatal(err)
	}
	req3.Header.Set("Content-Type", "application/json")
	res3 := httptest.NewRecorder()
	handler3 := http.HandlerFunc(controller.MissionDetails)
	handler3.ServeHTTP(res3, req3)
	if status := res3.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response3 util.ResponseHandler
	err = json.Unmarshal(res3.Body.Bytes(), &response3)
	if err != nil {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), `{
				"status": true,
				"message": "Success",
				"data": {
					"mission_name": "MissionX",
					"mission_details": "X Combat",
					"mission_status": "ASSIGNED",
					"assigned_avengers" : ["Captain America"] 
				}
			}`)
	}

	expected3 := `Success`
	if response3.Message != expected && !response3.Status {
		t.Errorf("handler returned unexpected message: got %v want %v",
			response3.Message, expected3)
	}
	fmt.Println("Pass : MissionDetails")

	// Test UpdateMissionStatus ---------------------------------
	var jsonStr2 = []byte(`{
		"mission_name": "MissionX",
		"mission_status": "UNASSIGNED"
	}`)

	req4, err4 := http.NewRequest("POST", "/api/updateMissionStatus", bytes.NewBuffer(jsonStr2))
	if err4 != nil {
		t.Fatal(err)
	}
	req4.Header.Set("Content-Type", "application/json")
	res4 := httptest.NewRecorder()
	handler4 := http.HandlerFunc(controller.UpdateMissionStatus)
	handler4.ServeHTTP(res4, req4)
	if status := res4.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response4 util.ResponseHandler
	err = json.Unmarshal(res4.Body.Bytes(), &response4)
	if err != nil {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), `{
				"status": true,
				"message": "Mission status update from ASSIGNED to UNASSIGNED",
				"data": null
			}`)
	}

	expected4 := `"Mission status updated`
	if !strings.Contains(response4.Message, expected) && !response4.Status {
		t.Errorf("handler returned unexpected message: got %v want %v",
			response.Message, expected4)
	}
	fmt.Println("Pass : UpdateMissionStatus")

	// Test AllAvengersSatus ---------------------------------
	req5, err5 := http.NewRequest("GET", "/api/allAvengersSatus", nil)
	if err5 != nil {
		t.Fatal(err)
	}
	res5 := httptest.NewRecorder()
	handler5 := http.HandlerFunc(controller.AllAvengersSatus)
	handler5.ServeHTTP(res5, req5)
	if status := res5.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response5 util.ResponseHandler
	err = json.Unmarshal(res5.Body.Bytes(), &response5)
	if err != nil {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(),
			`{
				"status": true,
				"message": "Success",
				"data": [
					{
						"name": "Captain America",
						"availability": "Available",
						"assigned_missions": "Mission1,Mission2"
					}
				]
			}`)
	}

	expected5 := `Success`
	if response5.Message != expected5 && !response5.Status {
		t.Errorf("handler returned unexpected message: got %v want %v",
			response.Message, expected5)
	}
	fmt.Println("Pass : AllAvengersSatus")
}

// func TestAllMissionsSatus(t *testing.T) {
// 	// a, b := controller.XYZ()
// 	// fmt.Println(a, b)

// 	req, err := http.NewRequest("GET", "/api/allMissionsSatus", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	res := httptest.NewRecorder()
// 	handler := http.HandlerFunc(controller.AllMissionsSatus)
// 	handler.ServeHTTP(res, req)
// 	if status := res.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}

// 	var response util.ResponseHandler
// 	err = json.Unmarshal(res.Body.Bytes(), &response)
// 	if err != nil {
// 		t.Errorf("handler returned unexpected body: got %v want %v",
// 			res.Body.String(),
// 			`{
//     			"status": true,
//     			"message": "Success",
//     			"data": [
//     			    {
//     			        "mission_name": "Mission2",
//     			        "mission_details": "Jungle Combat",
//     			        "mission_status": "ASSIGNED",
//     			        "assigned_avengers": [
//     			            "Captain America"
//     			        ]
//     			    }
//     			]
// 			}`)
// 	}

// 	expected := `Success`
// 	if response.Message != expected && !response.Status {
// 		t.Errorf("handler returned unexpected message: got %v want %v",
// 			response.Message, expected)
// 	}
// }

func TestMissionDetails(t *testing.T) {
	var jsonStr = []byte(`{"mission_name": "Mission2"}`)

	req, err := http.NewRequest("POST", "/api/missionDetails", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.MissionDetails)
	handler.ServeHTTP(res, req)
	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response util.ResponseHandler
	err = json.Unmarshal(res.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), `{
				"status": true,
				"message": "Success",
				"data": {
					"mission_name": "Mission2",
					"mission_details": "Jungle Combat",
					"mission_status": "ASSIGNED",
					"assigned_avengers": [
						"Captain America"
					]
				}
			}`)
	}

	expected := `Success`
	if response.Message != expected && !response.Status {
		t.Errorf("handler returned unexpected message: got %v want %v",
			response.Message, expected)
	}
}

func TestUpdateMissionStatus(t *testing.T) {
	var jsonStr = []byte(`{
		"mission_name": "Mission3",
		"mission_status": "ASSIGNED"
	}`)

	req, err := http.NewRequest("POST", "/api/updateMissionStatus", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.UpdateMissionStatus)
	handler.ServeHTTP(res, req)
	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response util.ResponseHandler
	err = json.Unmarshal(res.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), `{
				"status": true,
				"message": "Mission status update from ASSIGNED to UNASSIGNED",
				"data": null
			}`)
	}

	expected := `"Mission status updated`
	if !strings.Contains(response.Message, expected) && !response.Status {
		t.Errorf("handler returned unexpected message: got %v want %v",
			response.Message, expected)
	}
}

func TestAllAvengersSatus(t *testing.T) {

	req, err := http.NewRequest("GET", "/api/allAvengersSatus", nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.AllAvengersSatus)
	handler.ServeHTTP(res, req)
	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response util.ResponseHandler
	err = json.Unmarshal(res.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(),
			`{
				"status": true,
				"message": "Success",
				"data": [
					{
						"name": "Captain America",
						"availability": "Available",
						"assigned_missions": "Mission1,Mission2"
					}
				]
			}`)
	}

	expected := `Success`
	if response.Message != expected && !response.Status {
		t.Errorf("handler returned unexpected message: got %v want %v",
			response.Message, expected)
	}
}
