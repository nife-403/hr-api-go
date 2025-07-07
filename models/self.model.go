package models

import "go.mongodb.org/mongo-driver/v2/bson"

type SelfModel struct {
	ID                bson.ObjectID `bson:"_id,omitempty"`
	Name              string        `bson:"name"`
	ChemicalClass     string        `bson:"chemical_class"`
	PsychoactiveClass string        `bson:"psychoactive_class"`
	CommonNames       []string      `bson:"common_names"`
	Summary           string        `bson:"summary"`
	HistoryAndCulture string        `bson:"history_and_culture"`
	Chemistry         string        `bson:"chemistry"`
	HarmPotential     string        `bson:"harm_potential"`
	Withdrawal        string        `bson:"withdrawal"`
	Tolerance         struct {
		FullTolerance  string `bson:"full_tolerance"`
		CrossTolerance string `bson:"cross_tolerance"`
	} `bson:"tolerance"`
	Effects struct {
		Physical     []string `bson:"physical"`
		Visual       []any    `bson:"visual"`
		Cognitive    []string `bson:"cognitive"`
		Auditory     []any    `bson:"auditory"`
		MultiSensory []any    `bson:"multi_sensory"`
		After        []string `bson:"after"`
	} `bson:"effects"`
	References []struct {
		Source  string `bson:"source"`
		Section string `bson:"section"`
		URL     string `bson:"url"`
	} `bson:"references"`
	Dosage struct {
		Oral struct {
			ThresholdMg string `bson:"threshold_mg"`
			LightMg     string `bson:"light_mg"`
			CommonMg    string `bson:"common_mg"`
			StrongMg    string `bson:"strong_mg"`
			HeavyMg     string `bson:"heavy_mg"`
			Onset       string `bson:"onset"`
			ComeUp      string `bson:"come_up"`
			Peak        string `bson:"peak"`
			Duration    struct {
				Total        string `bson:"total"`
				Offset       string `bson:"offset"`
				Aftereffects string `bson:"aftereffects"`
			} `bson:"duration"`
		} `bson:"oral"`
		Insufflated struct {
			ThresholdMg string `bson:"threshold_mg"`
			LightMg     string `bson:"light_mg"`
			CommonMg    string `bson:"common_mg"`
			StrongMg    string `bson:"strong_mg"`
			HeavyMg     string `bson:"heavy_mg"`
			Onset       string `bson:"onset"`
			ComeUp      string `bson:"come_up"`
			Peak        string `bson:"peak"`
			Duration    struct {
				Total        string `bson:"total"`
				Offset       string `bson:"offset"`
				Aftereffects string `bson:"aftereffects"`
			} `bson:"duration"`
		} `bson:"insufflated"`
	} `bson:"dosage"`
	BioavailabilityLevels []any `bson:"bioavailability_levels"`
	V                     int   `bson:"__v,omitempty"`
}
