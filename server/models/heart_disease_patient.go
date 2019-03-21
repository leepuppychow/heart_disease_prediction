package models

// refer to NOTES.md for description of values:
type HeartDiseasePatient struct {
	Age                      string
	Sex                      string
	ChestPainType            string
	RestingBloodPress        string
	SerumCholesterol         string
	FastingBP                string
	RestECG                  string
	MaxHR                    string
	ExerciseInducedAngina    string
	STDepressionWithExercise string
	SlopeSTSegment           string
	NumberOfVesselsFlouro    string
	Thal                     string
	HasHeartDisease          string
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
