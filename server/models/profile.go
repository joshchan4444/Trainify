package models

type UserProfile struct {
	ProfileID   int  `json:"profile_id"`   // required
	Age         *int `json:"age"`          // optional, in years
	Weight      *int `json:"weight"`       // optional, in pounds
	Height      *int `json:"height"`       // optional, in centimeters
	MaxPushUps  *int `json:"max_push_ups"` // optional, in reps
	AvgPushUps  *int `json:"avg_push_ups"` // optional, in reps
	MaxPullUps  *int `json:"max_pull_ups"` // optional, in reps
	AvgPullUps  *int `json:"avg_pull_ups"` // optional, in reps
	MaxSquat    *int `json:"max_squat"`    // optional, in pounds
	AvgSquat    *int `json:"avg_squat"`    // optional, in pounds
	MaxBench    *int `json:"max_bench"`    // optional, in pounds
	AvgBench    *int `json:"avg_bench"`    // optional, in pounds
	CardioLevel *int `json:"cardio_level"` // optional, from 1 - 10
}
