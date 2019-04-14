package models

type DescriptiveStats struct {
	Count int     `json:"count"`
	Mean  float64 `json:"mean"`
	Std   float64 `json:"std"`
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
	Q1    float64 `json:"25%"`
	Q2    float64 `json:"50%"`
	Q3    float64 `json:"75%"`
}
