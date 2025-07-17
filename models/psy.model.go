package models

import _ "go.mongodb.org/mongo-driver/v2/bson"

type PsyDoc struct {
	Name               string   `bson:"name,omitempty"`
	Summary            string   `bson:"summary,omitempty"`
	AddictionPotential string   `bson:"addictionPotential,omitempty"`
	Toxicity           []string `bson:"toxicity,omitempty"`
	CrossTolerances    []string `bson:"crossTolerances,omitempty"`
	CommonNames        []string `bson:"commonNames,omitempty"`
	Class              struct {
		Chemical     []string `bson:"chemical,omitempty"`
		Psychoactive []string `bson:"psychoactive,omitempty"`
	} `bson:"class,omitempty"`
	Tolerance struct {
		Full string `bson:"full,omitempty"`
		Half string `bson:"half,omitempty"`
		Zero string `bson:"zero,omitempty"`
	} `bson:"tolerance,omitempty"`
	UncertainInteractions []struct {
		Name string `bson:"name,omitempty"`
	} `bson:"uncertainInteractions,omitempty"`
	UnsafeInteractions []struct {
		Name string `bson:"name,omitempty"`
	} `bson:"unsafeInteractions,omitempty"`
	DangerousInteractions []struct {
		Name string `bson:"name,omitempty"`
	} `bson:"dangerousInteractions,omitempty"`
	Roa struct {
		Oral *struct {
			Name string `bson:"name,omitempty"`
			Dose struct {
				Units     string  `bson:"units,omitempty"`
				Threshold float64 `bson:"threshold,omitempty"`
				Light     struct {
					Min float64 `bson:"min,omitempty"`
					Max float64 `bson:"max,omitempty"`
				} `bson:"light,omitempty"`
				Common struct {
					Min float64 `bson:"min,omitempty"`
					Max float64 `bson:"max,omitempty"`
				} `bson:"common,omitempty"`
				Strong struct {
					Min float64 `bson:"min,omitempty"`
					Max float64 `bson:"max,omitempty"`
				} `bson:"strong,omitempty"`
				Heavy float64 `bson:"heavy,omitempty"`
			} `bson:"dose,omitempty"`
			Duration struct {
				Onset struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"onset,omitempty"`
				Comeup any `bson:"comeup,omitempty"`
				Peak   struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"peak,omitempty"`
				Offset struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"offset,omitempty"`
				Afterglow any `bson:"afterglow,omitempty"`
				Total     struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"total,omitempty"`
				Duration any `bson:"duration,omitempty"`
			} `bson:"duration,omitempty"`
			Bioavailability any `bson:"bioavailability,omitempty"`
		} `bson:"oral,omitempty"`
		Sublingual *struct {
			Name string `bson:"name,omitempty"`
			Dose struct {
				Units     string  `bson:"units,omitempty"`
				Threshold float64 `bson:"threshold,omitempty"`
				Light     struct {
					Min float64 `bson:"min,omitempty"`
					Max float64 `bson:"max,omitempty"`
				} `bson:"light,omitempty"`
				Common struct {
					Min float64 `bson:"min,omitempty"`
					Max float64 `bson:"max,omitempty"`
				} `bson:"common,omitempty"`
				Strong struct {
					Min float64 `bson:"min,omitempty"`
					Max float64 `bson:"max,omitempty"`
				} `bson:"strong,omitempty"`
				Heavy float64 `bson:"heavy,omitempty"`
			} `bson:"dose,omitempty"`
			Duration struct {
				Onset struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"onset,omitempty"`
				Comeup any `bson:"comeup,omitempty"`
				Peak   struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"peak,omitempty"`
				Offset struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"offset,omitempty"`
				Afterglow any `bson:"afterglow,omitempty"`
				Total     struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"total,omitempty"`
				Duration any `bson:"duration,omitempty"`
			} `bson:"duration,omitempty"`
			Bioavailability any `bson:"bioavailability,omitempty"`
		} `bson:"sublingual,omitempty"`
		Buccal *struct {
			Name string `bson:"name,omitempty"`
			Dose struct {
				Units     string  `bson:"units,omitempty"`
				Threshold float64 `bson:"threshold,omitempty"`
				Light     struct {
					Min float64 `bson:"min,omitempty"`
					Max float64 `bson:"max,omitempty"`
				} `bson:"light,omitempty"`
				Common struct {
					Min float64 `bson:"min,omitempty"`
					Max float64 `bson:"max,omitempty"`
				} `bson:"common,omitempty"`
				Strong struct {
					Min float64 `bson:"min,omitempty"`
					Max float64 `bson:"max,omitempty"`
				} `bson:"strong,omitempty"`
				Heavy float64 `bson:"heavy,omitempty"`
			} `bson:"dose,omitempty"`
			Duration struct {
				Onset struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"onset,omitempty"`
				Comeup any `bson:"comeup,omitempty"`
				Peak   struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"peak,omitempty"`
				Offset struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"offset,omitempty"`
				Afterglow any `bson:"afterglow,omitempty"`
				Total     struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"total,omitempty"`
				Duration any `bson:"duration,omitempty"`
			} `bson:"duration,omitempty"`
			Bioavailability any `bson:"bioavailability,omitempty"`
		} `bson:"buccal,omitempty"`
		Insufflated *struct {
			Name string `bson:"name,omitempty"`
			Dose struct {
				Units     string  `bson:"units,omitempty"`
				Threshold float64 `bson:"threshold,omitempty"`
				Light     struct {
					Min float64 `bson:"min,omitempty"`
					Max float64 `bson:"max,omitempty"`
				} `bson:"light,omitempty"`
				Common struct {
					Min float64 `bson:"min,omitempty"`
					Max float64 `bson:"max,omitempty"`
				} `bson:"common,omitempty"`
				Strong struct {
					Min float64 `bson:"min,omitempty"`
					Max float64 `bson:"max,omitempty"`
				} `bson:"strong,omitempty"`
				Heavy float64 `bson:"heavy,omitempty"`
			} `bson:"dose,omitempty"`
			Duration struct {
				Onset struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"onset,omitempty"`
				Comeup any `bson:"comeup,omitempty"`
				Peak   struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"peak,omitempty"`
				Offset struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"offset,omitempty"`
				Afterglow any `bson:"afterglow,omitempty"`
				Total     struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"total,omitempty"`
				Duration any `bson:"duration,omitempty"`
			} `bson:"duration,omitempty"`
			Bioavailability any `bson:"bioavailability,omitempty"`
		} `bson:"insufflated,omitempty"`
		Rectal *struct {
			Name string `bson:"name,omitempty"`
			Dose struct {
				Units     string  `bson:"units,omitempty"`
				Threshold float64 `bson:"threshold,omitempty"`
				Light     struct {
					Min float64 `bson:"min,omitempty"`
					Max float64 `bson:"max,omitempty"`
				} `bson:"light,omitempty"`
				Common struct {
					Min float64 `bson:"min,omitempty"`
					Max float64 `bson:"max,omitempty"`
				} `bson:"common,omitempty"`
				Strong struct {
					Min float64 `bson:"min,omitempty"`
					Max float64 `bson:"max,omitempty"`
				} `bson:"strong,omitempty"`
				Heavy float64 `bson:"heavy,omitempty"`
			} `bson:"dose,omitempty"`
			Duration struct {
				Onset struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"onset,omitempty"`
				Comeup any `bson:"comeup,omitempty"`
				Peak   struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"peak,omitempty"`
				Offset struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"offset,omitempty"`
				Afterglow any `bson:"afterglow,omitempty"`
				Total     struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"total,omitempty"`
				Duration any `bson:"duration,omitempty"`
			} `bson:"duration,omitempty"`
			Bioavailability any `bson:"bioavailability,omitempty"`
		} `bson:"rectal,omitempty"`
		Transdermal *struct {
			Name string `bson:"name,omitempty"`
			Dose struct {
				Units     string  `bson:"units,omitempty"`
				Threshold float64 `bson:"threshold,omitempty"`
				Light     struct {
					Min float64 `bson:"min,omitempty"`
					Max float64 `bson:"max,omitempty"`
				} `bson:"light,omitempty"`
				Common struct {
					Min float64 `bson:"min,omitempty"`
					Max float64 `bson:"max,omitempty"`
				} `bson:"common,omitempty"`
				Strong struct {
					Min float64 `bson:"min,omitempty"`
					Max float64 `bson:"max,omitempty"`
				} `bson:"strong,omitempty"`
				Heavy float64 `bson:"heavy,omitempty"`
			} `bson:"dose,omitempty"`
			Duration struct {
				Onset struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"onset,omitempty"`
				Comeup any `bson:"comeup,omitempty"`
				Peak   struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"peak,omitempty"`
				Offset struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"offset,omitempty"`
				Afterglow any `bson:"afterglow,omitempty"`
				Total     struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"total,omitempty"`
				Duration any `bson:"duration,omitempty"`
			} `bson:"duration,omitempty"`
			Bioavailability any `bson:"bioavailability,omitempty"`
		} `bson:"transdermal,omitempty"`
		Subcutaneuos *struct {
			Name string `bson:"name,omitempty"`
			Dose struct {
				Units     string  `bson:"units,omitempty"`
				Threshold float64 `bson:"threshold,omitempty"`
				Light     struct {
					Min float64 `bson:"min,omitempty"`
					Max float64 `bson:"max,omitempty"`
				} `bson:"light,omitempty"`
				Common struct {
					Min float64 `bson:"min,omitempty"`
					Max float64 `bson:"max,omitempty"`
				} `bson:"common,omitempty"`
				Strong struct {
					Min float64 `bson:"min,omitempty"`
					Max float64 `bson:"max,omitempty"`
				} `bson:"strong,omitempty"`
				Heavy float64 `bson:"heavy,omitempty"`
			} `bson:"dose,omitempty"`
			Duration struct {
				Onset struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"onset,omitempty"`
				Comeup any `bson:"comeup,omitempty"`
				Peak   struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"peak,omitempty"`
				Offset struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"offset,omitempty"`
				Afterglow any `bson:"afterglow,omitempty"`
				Total     struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"total,omitempty"`
				Duration any `bson:"duration,omitempty"`
			} `bson:"duration,omitempty"`
			Bioavailability any `bson:"bioavailability,omitempty"`
		} `bson:"subcutaneuos,omitempty"`
		Intramuscular *struct {
			Name string `bson:"name,omitempty"`
			Dose struct {
				Units     string  `bson:"units,omitempty"`
				Threshold float64 `bson:"threshold,omitempty"`
				Light     struct {
					Min float64 `bson:"min,omitempty"`
					Max float64 `bson:"max,omitempty"`
				} `bson:"light,omitempty"`
				Common struct {
					Min float64 `bson:"min,omitempty"`
					Max float64 `bson:"max,omitempty"`
				} `bson:"common,omitempty"`
				Strong struct {
					Min float64 `bson:"min,omitempty"`
					Max float64 `bson:"max,omitempty"`
				} `bson:"strong,omitempty"`
				Heavy float64 `bson:"heavy,omitempty"`
			} `bson:"dose,omitempty"`
			Duration struct {
				Onset struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"onset,omitempty"`
				Comeup any `bson:"comeup,omitempty"`
				Peak   struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"peak,omitempty"`
				Offset struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"offset,omitempty"`
				Afterglow any `bson:"afterglow,omitempty"`
				Total     struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"total,omitempty"`
				Duration any `bson:"duration,omitempty"`
			} `bson:"duration,omitempty"`
			Bioavailability any `bson:"bioavailability,omitempty"`
		} `bson:"intramuscular,omitempty"`
		Intravenous *struct {
			Name string `bson:"name,omitempty"`
			Dose struct {
				Units     string  `bson:"units,omitempty"`
				Threshold float64 `bson:"threshold,omitempty"`
				Light     struct {
					Min float64 `bson:"min,omitempty"`
					Max float64 `bson:"max,omitempty"`
				} `bson:"light,omitempty"`
				Common struct {
					Min float64 `bson:"min,omitempty"`
					Max float64 `bson:"max,omitempty"`
				} `bson:"common,omitempty"`
				Strong struct {
					Min float64 `bson:"min,omitempty"`
					Max float64 `bson:"max,omitempty"`
				} `bson:"strong,omitempty"`
				Heavy float64 `bson:"heavy,omitempty"`
			} `bson:"dose,omitempty"`
			Duration struct {
				Onset struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"onset,omitempty"`
				Comeup any `bson:"comeup,omitempty"`
				Peak   struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"peak,omitempty"`
				Offset struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"offset,omitempty"`
				Afterglow any `bson:"afterglow,omitempty"`
				Total     struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"total,omitempty"`
				Duration any `bson:"duration,omitempty"`
			} `bson:"duration,omitempty"`
			Bioavailability any `bson:"bioavailability,omitempty"`
		} `bson:"intravenous,omitempty"`
		Smoked *struct {
			Name string `bson:"name,omitempty"`
			Dose struct {
				Units     string  `bson:"units,omitempty"`
				Threshold float64 `bson:"threshold,omitempty"`
				Light     struct {
					Min float64 `bson:"min,omitempty"`
					Max float64 `bson:"max,omitempty"`
				} `bson:"light,omitempty"`
				Common struct {
					Min float64 `bson:"min,omitempty"`
					Max float64 `bson:"max,omitempty"`
				} `bson:"common,omitempty"`
				Strong struct {
					Min float64 `bson:"min,omitempty"`
					Max float64 `bson:"max,omitempty"`
				} `bson:"strong,omitempty"`
				Heavy float64 `bson:"heavy,omitempty"`
			} `bson:"dose,omitempty"`
			Duration struct {
				Onset struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"onset,omitempty"`
				Comeup any `bson:"comeup,omitempty"`
				Peak   struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"peak,omitempty"`
				Offset struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"offset,omitempty"`
				Afterglow any `bson:"afterglow,omitempty"`
				Total     struct {
					Min   float64 `bson:"min,omitempty"`
					Max   float64 `bson:"max,omitempty"`
					Units string  `bson:"units,omitempty"`
				} `bson:"total,omitempty"`
				Duration any `bson:"duration,omitempty"`
			} `bson:"duration,omitempty"`
			Bioavailability any `bson:"bioavailability,omitempty"`
		} `bson:"smoked,omitempty"`
	} `bson:"roa,omitempty"`
}
