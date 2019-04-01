package models

// refer to NOTES.md for description of values:
type HeartDiseasePatient struct {
	Age                      string `json:"age"`
	Sex                      string `json:"sex"`
	ChestPainType            string `json:"cp"`
	RestingBloodPress        string `json:"trestbps"`
	SerumCholesterol         string `json:"chol"`
	FastingBP                string `json:"fbs"`
	RestECG                  string `json:"restecg"`
	MaxHR                    string `json:"thalach"`
	ExerciseInducedAngina    string `json:"exang"`
	STDepressionWithExercise string `json:"oldpeak"`
	SlopeSTSegment           string `json:"slope"`
	NumberOfVesselsFlouro    string `json:"ca"`
	Thal                     string `json:"thal"`
	HasHeartDisease          string `json:"hasHeartDisease"`
}

func (p *HeartDiseasePatient) DataRow() []string {
	return []string{
		p.Age,
		p.Sex,
		p.ChestPainType,
		p.RestingBloodPress,
		p.SerumCholesterol,
		p.FastingBP,
		p.RestECG,
		p.MaxHR,
		p.ExerciseInducedAngina,
		p.STDepressionWithExercise,
		p.SlopeSTSegment,
		p.NumberOfVesselsFlouro,
		p.Thal,
		p.HasHeartDisease,
	}
}
