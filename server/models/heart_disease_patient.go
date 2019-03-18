package models

// refer to NOTES.md for description of values:
type HeartDiseasePatient struct {
	Age                   string
	Sex                   string
	ChestPainType         string
	RestingBloodPress     string
	SerumCholesterol      string
	FastingBP             string
	ExerciseInducedAngina string
	HasHeartDisease       string
}

func (p *HeartDiseasePatient) DataRow() []string {
	return []string{
		p.Age,
		p.Sex,
		p.ChestPainType,
		p.RestingBloodPress,
		p.SerumCholesterol,
		p.FastingBP,
		p.ExerciseInducedAngina,
		p.HasHeartDisease,
	}
}
