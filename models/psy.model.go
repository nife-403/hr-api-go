package models

type PsyDoc struct {
	Name               string   `bson:"name"`
	Summary            string   `bson:"summary"`
	AddictionPotential string   `bson:"addictionPotential"`
	Toxicity           []string `bson:"toxicity"`
	CrossTolerances    []string `bson:"crossTolerances"`
	CommonNames        []string `bson:"commonNames"`
	Class              struct {
		Chemical     []string `bson:"chemical"`
		Psychoactive []string `bson:"psychoactive"`
	} `bson:"class"`
	Tolerance struct {
		Full string `bson:"full"`
		Half string `bson:"half"`
		Zero string `bson:"zero"`
	} `bson:"tolerance"`
	UncertainInteractions []struct {
		Name string `bson:"name"`
	} `bson:"uncertainInteractions"`
	UnsafeInteractions []struct {
		Name string `bson:"name"`
	} `bson:"unsafeInteractions"`
	DangerousInteractions []struct {
		Name string `bson:"name"`
	} `bson:"dangerousInteractions"`
	Roa struct {
		Oral struct {
			Name string `bson:"name"`
			Dose struct {
				Units     string  `bson:"units"`
				Threshold float64 `bson:"threshold"`
				Heavy     float64 `bson:"heavy"`
				Common    struct {
					Min float64 `bson:"min"`
					Max float64 `bson:"max"`
				} `bson:"common"`
				Light struct {
					Min float64 `bson:"min"`
					Max float64 `bson:"max"`
				} `bson:"light"`
				Strong struct {
					Min float64 `bson:"min"`
					Max float64 `bson:"max"`
				} `bson:"strong"`
			} `bson:"dose"`
			Duration struct {
				Afterglow any `bson:"afterglow"`
				Comeup    any `bson:"comeup"`
				Duration  any `bson:"duration"`
				Offset    struct {
					Min   float64 `bson:"min"`
					Max   float64 `bson:"max"`
					Units string  `bson:"units"`
				} `bson:"offset"`
				Onset struct {
					Min   float64 `bson:"min"`
					Max   float64 `bson:"max"`
					Units string  `bson:"units"`
				} `bson:"onset"`
				Peak struct {
					Min   float64 `bson:"min"`
					Max   float64 `bson:"max"`
					Units string  `bson:"units"`
				} `bson:"peak"`
				Total struct {
					Min   float64 `bson:"min"`
					Max   float64 `bson:"max"`
					Units string  `bson:"units"`
				} `bson:"total"`
			} `bson:"duration"`
			Bioavailability any `bson:"bioavailability,omitempty"`
		} `bson:"oral,omitempty"`
		Sublingual    any `bson:"sublingual,omitempty"`
		Buccal        any `bson:"buccal,omitempty"`
		Insufflated   any `bson:"insufflated,omitempty"`
		Rectal        any `bson:"rectal,omitempty"`
		Transdermal   any `bson:"transdermal,omitempty"`
		Subcutaneous  any `bson:"subcutaneous,omitempty"`
		Intramuscular any `bson:"intramuscular,omitempty"`
		Intravenous   any `bson:"intravenous,omitempty"`
		Smoked        any `bson:"smoked,omitempty"`
	} `bson:"roa"`
}
