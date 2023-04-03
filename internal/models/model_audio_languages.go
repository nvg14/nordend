// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: Nordend
//	Version: 1.0.0
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// AudioLanguages is an enum.
type AudioLanguages string

// Validate implements basic validation for this model
func (m AudioLanguages) Validate() error {
	return InKnownAudioLanguages.Validate(m)
}

var (
	AudioLanguagesAAR  AudioLanguages = "AAR"
	AudioLanguagesABK  AudioLanguages = "ABK"
	AudioLanguagesAFR  AudioLanguages = "AFR"
	AudioLanguagesAKA  AudioLanguages = "AKA"
	AudioLanguagesAMH  AudioLanguages = "AMH"
	AudioLanguagesARA9 AudioLanguages = "ARA9"
	AudioLanguagesARG  AudioLanguages = "ARG"
	AudioLanguagesASM  AudioLanguages = "ASM"
	AudioLanguagesAVA  AudioLanguages = "AVA"
	AudioLanguagesAVE  AudioLanguages = "AVE"
	AudioLanguagesAYM  AudioLanguages = "AYM"
	AudioLanguagesAZE  AudioLanguages = "AZE"
	AudioLanguagesBAK  AudioLanguages = "BAK"
	AudioLanguagesBAM  AudioLanguages = "BAM"
	AudioLanguagesBEL  AudioLanguages = "BEL"
	AudioLanguagesBEN  AudioLanguages = "BEN"
	AudioLanguagesBIS  AudioLanguages = "BIS"
	AudioLanguagesBOD  AudioLanguages = "BOD"
	AudioLanguagesBOS  AudioLanguages = "BOS"
	AudioLanguagesBRE  AudioLanguages = "BRE"
	AudioLanguagesBUL  AudioLanguages = "BUL"
	AudioLanguagesCAT  AudioLanguages = "CAT"
	AudioLanguagesCES  AudioLanguages = "CES"
	AudioLanguagesCHA  AudioLanguages = "CHA"
	AudioLanguagesCHE  AudioLanguages = "CHE"
	AudioLanguagesCHU  AudioLanguages = "CHU"
	AudioLanguagesCHV  AudioLanguages = "CHV"
	AudioLanguagesCOR  AudioLanguages = "COR"
	AudioLanguagesCOS  AudioLanguages = "COS"
	AudioLanguagesCRE  AudioLanguages = "CRE"
	AudioLanguagesCYM  AudioLanguages = "CYM"
	AudioLanguagesDAN  AudioLanguages = "DAN"
	AudioLanguagesDEU  AudioLanguages = "DEU"
	AudioLanguagesDIV  AudioLanguages = "DIV"
	AudioLanguagesDZO  AudioLanguages = "DZO"
	AudioLanguagesELL  AudioLanguages = "ELL"
	AudioLanguagesENG  AudioLanguages = "ENG"
	AudioLanguagesEPO  AudioLanguages = "EPO"
	AudioLanguagesEST  AudioLanguages = "EST"
	AudioLanguagesEUS  AudioLanguages = "EUS"
	AudioLanguagesEWE  AudioLanguages = "EWE"
	AudioLanguagesFAO  AudioLanguages = "FAO"
	AudioLanguagesFAS  AudioLanguages = "FAS"
	AudioLanguagesFIJ  AudioLanguages = "FIJ"
	AudioLanguagesFIN  AudioLanguages = "FIN"
	AudioLanguagesFRA  AudioLanguages = "FRA"
	AudioLanguagesFRY  AudioLanguages = "FRY"
	AudioLanguagesFUL  AudioLanguages = "FUL"
	AudioLanguagesGLA  AudioLanguages = "GLA"
	AudioLanguagesGLE  AudioLanguages = "GLE"
	AudioLanguagesGLG  AudioLanguages = "GLG"
	AudioLanguagesGLV  AudioLanguages = "GLV"
	AudioLanguagesGRN  AudioLanguages = "GRN"
	AudioLanguagesGUJ  AudioLanguages = "GUJ"
	AudioLanguagesHAT  AudioLanguages = "HAT"
	AudioLanguagesHAU  AudioLanguages = "HAU"
	AudioLanguagesHEB  AudioLanguages = "HEB"
	AudioLanguagesHER  AudioLanguages = "HER"
	AudioLanguagesHIN  AudioLanguages = "HIN"
	AudioLanguagesHMO  AudioLanguages = "HMO"
	AudioLanguagesHRV  AudioLanguages = "HRV"
	AudioLanguagesHUN  AudioLanguages = "HUN"
	AudioLanguagesHYE  AudioLanguages = "HYE"
	AudioLanguagesIBO  AudioLanguages = "IBO"
	AudioLanguagesIDO  AudioLanguages = "IDO"
	AudioLanguagesIII  AudioLanguages = "III"
	AudioLanguagesIKU  AudioLanguages = "IKU"
	AudioLanguagesILE  AudioLanguages = "ILE"
	AudioLanguagesINA  AudioLanguages = "INA"
	AudioLanguagesIND  AudioLanguages = "IND"
	AudioLanguagesIPK  AudioLanguages = "IPK"
	AudioLanguagesISL  AudioLanguages = "ISL"
	AudioLanguagesITA  AudioLanguages = "ITA"
	AudioLanguagesJAV  AudioLanguages = "JAV"
	AudioLanguagesJPN  AudioLanguages = "JPN"
	AudioLanguagesKAL  AudioLanguages = "KAL"
	AudioLanguagesKAN  AudioLanguages = "KAN"
	AudioLanguagesKAS  AudioLanguages = "KAS"
	AudioLanguagesKAT  AudioLanguages = "KAT"
	AudioLanguagesKAU  AudioLanguages = "KAU"
	AudioLanguagesKAZ  AudioLanguages = "KAZ"
	AudioLanguagesKHM  AudioLanguages = "KHM"
	AudioLanguagesKIK  AudioLanguages = "KIK"
	AudioLanguagesKIN  AudioLanguages = "KIN"
	AudioLanguagesKIR  AudioLanguages = "KIR"
	AudioLanguagesKOM  AudioLanguages = "KOM"
	AudioLanguagesKON  AudioLanguages = "KON"
	AudioLanguagesKOR  AudioLanguages = "KOR"
	AudioLanguagesKUA  AudioLanguages = "KUA"
	AudioLanguagesKUR  AudioLanguages = "KUR"
	AudioLanguagesLAO  AudioLanguages = "LAO"
	AudioLanguagesLAT  AudioLanguages = "LAT"
	AudioLanguagesLAV  AudioLanguages = "LAV"
	AudioLanguagesLIM  AudioLanguages = "LIM"
	AudioLanguagesLIN  AudioLanguages = "LIN"
	AudioLanguagesLIT  AudioLanguages = "LIT"
	AudioLanguagesLTZ  AudioLanguages = "LTZ"
	AudioLanguagesLUB  AudioLanguages = "LUB"
	AudioLanguagesLUG  AudioLanguages = "LUG"
	AudioLanguagesMAH  AudioLanguages = "MAH"
	AudioLanguagesMAL  AudioLanguages = "MAL"
	AudioLanguagesMAR  AudioLanguages = "MAR"
	AudioLanguagesMKD  AudioLanguages = "MKD"
	AudioLanguagesMLG1 AudioLanguages = "MLG1"
	AudioLanguagesMLT  AudioLanguages = "MLT"
	AudioLanguagesMON  AudioLanguages = "MON"
	AudioLanguagesMRI  AudioLanguages = "MRI"
	AudioLanguagesMSA6 AudioLanguages = "MSA6"
	AudioLanguagesMYA  AudioLanguages = "MYA"
	AudioLanguagesNAU  AudioLanguages = "NAU"
	AudioLanguagesNAV  AudioLanguages = "NAV"
	AudioLanguagesNBL  AudioLanguages = "NBL"
	AudioLanguagesNDE  AudioLanguages = "NDE"
	AudioLanguagesNDO  AudioLanguages = "NDO"
	AudioLanguagesNEP  AudioLanguages = "NEP"
	AudioLanguagesNLD  AudioLanguages = "NLD"
	AudioLanguagesNNO  AudioLanguages = "NNO"
	AudioLanguagesNOB  AudioLanguages = "NOB"
	AudioLanguagesNOR  AudioLanguages = "NOR"
	AudioLanguagesNYA  AudioLanguages = "NYA"
	AudioLanguagesOCI  AudioLanguages = "OCI"
	AudioLanguagesOJI  AudioLanguages = "OJI"
	AudioLanguagesORI  AudioLanguages = "ORI"
	AudioLanguagesORM  AudioLanguages = "ORM"
	AudioLanguagesOSS  AudioLanguages = "OSS"
	AudioLanguagesPAN  AudioLanguages = "PAN"
	AudioLanguagesPLI  AudioLanguages = "PLI"
	AudioLanguagesPOL  AudioLanguages = "POL"
	AudioLanguagesPOR  AudioLanguages = "POR"
	AudioLanguagesPUS  AudioLanguages = "PUS"
	AudioLanguagesQUE3 AudioLanguages = "QUE3"
	AudioLanguagesROH  AudioLanguages = "ROH"
	AudioLanguagesRON  AudioLanguages = "RON"
	AudioLanguagesRUN  AudioLanguages = "RUN"
	AudioLanguagesRUS  AudioLanguages = "RUS"
	AudioLanguagesSAG  AudioLanguages = "SAG"
	AudioLanguagesSAN  AudioLanguages = "SAN"
	AudioLanguagesSIN  AudioLanguages = "SIN"
	AudioLanguagesSLK  AudioLanguages = "SLK"
	AudioLanguagesSLV  AudioLanguages = "SLV"
	AudioLanguagesSME  AudioLanguages = "SME"
	AudioLanguagesSMO  AudioLanguages = "SMO"
	AudioLanguagesSNA  AudioLanguages = "SNA"
	AudioLanguagesSND  AudioLanguages = "SND"
	AudioLanguagesSOM  AudioLanguages = "SOM"
	AudioLanguagesSOT  AudioLanguages = "SOT"
	AudioLanguagesSPA  AudioLanguages = "SPA"
	AudioLanguagesSQI  AudioLanguages = "SQI"
	AudioLanguagesSRD  AudioLanguages = "SRD"
	AudioLanguagesSRP  AudioLanguages = "SRP"
	AudioLanguagesSSW  AudioLanguages = "SSW"
	AudioLanguagesSUN  AudioLanguages = "SUN"
	AudioLanguagesSWA  AudioLanguages = "SWA"
	AudioLanguagesSWE  AudioLanguages = "SWE"
	AudioLanguagesTAH  AudioLanguages = "TAH"
	AudioLanguagesTAM  AudioLanguages = "TAM"
	AudioLanguagesTAT  AudioLanguages = "TAT"
	AudioLanguagesTEL  AudioLanguages = "TEL"
	AudioLanguagesTGK  AudioLanguages = "TGK"
	AudioLanguagesTGL  AudioLanguages = "TGL"
	AudioLanguagesTHA  AudioLanguages = "THA"
	AudioLanguagesTIR  AudioLanguages = "TIR"
	AudioLanguagesTON  AudioLanguages = "TON"
	AudioLanguagesTSN  AudioLanguages = "TSN"
	AudioLanguagesTSO  AudioLanguages = "TSO"
	AudioLanguagesTUK  AudioLanguages = "TUK"
	AudioLanguagesTUR  AudioLanguages = "TUR"
	AudioLanguagesTWI  AudioLanguages = "TWI"
	AudioLanguagesUIG  AudioLanguages = "UIG"
	AudioLanguagesUKR  AudioLanguages = "UKR"
	AudioLanguagesURD  AudioLanguages = "URD"
	AudioLanguagesUZB  AudioLanguages = "UZB"
	AudioLanguagesVEN  AudioLanguages = "VEN"
	AudioLanguagesVIE  AudioLanguages = "VIE"
	AudioLanguagesVOL  AudioLanguages = "VOL"
	AudioLanguagesWLN  AudioLanguages = "WLN"
	AudioLanguagesWOL  AudioLanguages = "WOL"
	AudioLanguagesXHO  AudioLanguages = "XHO"
	AudioLanguagesYID  AudioLanguages = "YID"
	AudioLanguagesYOR  AudioLanguages = "YOR"
	AudioLanguagesZHA6 AudioLanguages = "ZHA6"
	AudioLanguagesZHO6 AudioLanguages = "ZHO6"
	AudioLanguagesZUL  AudioLanguages = "ZUL"

	// KnownAudioLanguages is the list of valid AudioLanguages
	KnownAudioLanguages = []AudioLanguages{
		AudioLanguagesAAR,
		AudioLanguagesABK,
		AudioLanguagesAFR,
		AudioLanguagesAKA,
		AudioLanguagesAMH,
		AudioLanguagesARA9,
		AudioLanguagesARG,
		AudioLanguagesASM,
		AudioLanguagesAVA,
		AudioLanguagesAVE,
		AudioLanguagesAYM,
		AudioLanguagesAZE,
		AudioLanguagesBAK,
		AudioLanguagesBAM,
		AudioLanguagesBEL,
		AudioLanguagesBEN,
		AudioLanguagesBIS,
		AudioLanguagesBOD,
		AudioLanguagesBOS,
		AudioLanguagesBRE,
		AudioLanguagesBUL,
		AudioLanguagesCAT,
		AudioLanguagesCES,
		AudioLanguagesCHA,
		AudioLanguagesCHE,
		AudioLanguagesCHU,
		AudioLanguagesCHV,
		AudioLanguagesCOR,
		AudioLanguagesCOS,
		AudioLanguagesCRE,
		AudioLanguagesCYM,
		AudioLanguagesDAN,
		AudioLanguagesDEU,
		AudioLanguagesDIV,
		AudioLanguagesDZO,
		AudioLanguagesELL,
		AudioLanguagesENG,
		AudioLanguagesEPO,
		AudioLanguagesEST,
		AudioLanguagesEUS,
		AudioLanguagesEWE,
		AudioLanguagesFAO,
		AudioLanguagesFAS,
		AudioLanguagesFIJ,
		AudioLanguagesFIN,
		AudioLanguagesFRA,
		AudioLanguagesFRY,
		AudioLanguagesFUL,
		AudioLanguagesGLA,
		AudioLanguagesGLE,
		AudioLanguagesGLG,
		AudioLanguagesGLV,
		AudioLanguagesGRN,
		AudioLanguagesGUJ,
		AudioLanguagesHAT,
		AudioLanguagesHAU,
		AudioLanguagesHEB,
		AudioLanguagesHER,
		AudioLanguagesHIN,
		AudioLanguagesHMO,
		AudioLanguagesHRV,
		AudioLanguagesHUN,
		AudioLanguagesHYE,
		AudioLanguagesIBO,
		AudioLanguagesIDO,
		AudioLanguagesIII,
		AudioLanguagesIKU,
		AudioLanguagesILE,
		AudioLanguagesINA,
		AudioLanguagesIND,
		AudioLanguagesIPK,
		AudioLanguagesISL,
		AudioLanguagesITA,
		AudioLanguagesJAV,
		AudioLanguagesJPN,
		AudioLanguagesKAL,
		AudioLanguagesKAN,
		AudioLanguagesKAS,
		AudioLanguagesKAT,
		AudioLanguagesKAU,
		AudioLanguagesKAZ,
		AudioLanguagesKHM,
		AudioLanguagesKIK,
		AudioLanguagesKIN,
		AudioLanguagesKIR,
		AudioLanguagesKOM,
		AudioLanguagesKON,
		AudioLanguagesKOR,
		AudioLanguagesKUA,
		AudioLanguagesKUR,
		AudioLanguagesLAO,
		AudioLanguagesLAT,
		AudioLanguagesLAV,
		AudioLanguagesLIM,
		AudioLanguagesLIN,
		AudioLanguagesLIT,
		AudioLanguagesLTZ,
		AudioLanguagesLUB,
		AudioLanguagesLUG,
		AudioLanguagesMAH,
		AudioLanguagesMAL,
		AudioLanguagesMAR,
		AudioLanguagesMKD,
		AudioLanguagesMLG1,
		AudioLanguagesMLT,
		AudioLanguagesMON,
		AudioLanguagesMRI,
		AudioLanguagesMSA6,
		AudioLanguagesMYA,
		AudioLanguagesNAU,
		AudioLanguagesNAV,
		AudioLanguagesNBL,
		AudioLanguagesNDE,
		AudioLanguagesNDO,
		AudioLanguagesNEP,
		AudioLanguagesNLD,
		AudioLanguagesNNO,
		AudioLanguagesNOB,
		AudioLanguagesNOR,
		AudioLanguagesNYA,
		AudioLanguagesOCI,
		AudioLanguagesOJI,
		AudioLanguagesORI,
		AudioLanguagesORM,
		AudioLanguagesOSS,
		AudioLanguagesPAN,
		AudioLanguagesPLI,
		AudioLanguagesPOL,
		AudioLanguagesPOR,
		AudioLanguagesPUS,
		AudioLanguagesQUE3,
		AudioLanguagesROH,
		AudioLanguagesRON,
		AudioLanguagesRUN,
		AudioLanguagesRUS,
		AudioLanguagesSAG,
		AudioLanguagesSAN,
		AudioLanguagesSIN,
		AudioLanguagesSLK,
		AudioLanguagesSLV,
		AudioLanguagesSME,
		AudioLanguagesSMO,
		AudioLanguagesSNA,
		AudioLanguagesSND,
		AudioLanguagesSOM,
		AudioLanguagesSOT,
		AudioLanguagesSPA,
		AudioLanguagesSQI,
		AudioLanguagesSRD,
		AudioLanguagesSRP,
		AudioLanguagesSSW,
		AudioLanguagesSUN,
		AudioLanguagesSWA,
		AudioLanguagesSWE,
		AudioLanguagesTAH,
		AudioLanguagesTAM,
		AudioLanguagesTAT,
		AudioLanguagesTEL,
		AudioLanguagesTGK,
		AudioLanguagesTGL,
		AudioLanguagesTHA,
		AudioLanguagesTIR,
		AudioLanguagesTON,
		AudioLanguagesTSN,
		AudioLanguagesTSO,
		AudioLanguagesTUK,
		AudioLanguagesTUR,
		AudioLanguagesTWI,
		AudioLanguagesUIG,
		AudioLanguagesUKR,
		AudioLanguagesURD,
		AudioLanguagesUZB,
		AudioLanguagesVEN,
		AudioLanguagesVIE,
		AudioLanguagesVOL,
		AudioLanguagesWLN,
		AudioLanguagesWOL,
		AudioLanguagesXHO,
		AudioLanguagesYID,
		AudioLanguagesYOR,
		AudioLanguagesZHA6,
		AudioLanguagesZHO6,
		AudioLanguagesZUL,
	}
	// KnownAudioLanguagesString is the list of valid AudioLanguages as string
	KnownAudioLanguagesString = []string{
		string(AudioLanguagesAAR),
		string(AudioLanguagesABK),
		string(AudioLanguagesAFR),
		string(AudioLanguagesAKA),
		string(AudioLanguagesAMH),
		string(AudioLanguagesARA9),
		string(AudioLanguagesARG),
		string(AudioLanguagesASM),
		string(AudioLanguagesAVA),
		string(AudioLanguagesAVE),
		string(AudioLanguagesAYM),
		string(AudioLanguagesAZE),
		string(AudioLanguagesBAK),
		string(AudioLanguagesBAM),
		string(AudioLanguagesBEL),
		string(AudioLanguagesBEN),
		string(AudioLanguagesBIS),
		string(AudioLanguagesBOD),
		string(AudioLanguagesBOS),
		string(AudioLanguagesBRE),
		string(AudioLanguagesBUL),
		string(AudioLanguagesCAT),
		string(AudioLanguagesCES),
		string(AudioLanguagesCHA),
		string(AudioLanguagesCHE),
		string(AudioLanguagesCHU),
		string(AudioLanguagesCHV),
		string(AudioLanguagesCOR),
		string(AudioLanguagesCOS),
		string(AudioLanguagesCRE),
		string(AudioLanguagesCYM),
		string(AudioLanguagesDAN),
		string(AudioLanguagesDEU),
		string(AudioLanguagesDIV),
		string(AudioLanguagesDZO),
		string(AudioLanguagesELL),
		string(AudioLanguagesENG),
		string(AudioLanguagesEPO),
		string(AudioLanguagesEST),
		string(AudioLanguagesEUS),
		string(AudioLanguagesEWE),
		string(AudioLanguagesFAO),
		string(AudioLanguagesFAS),
		string(AudioLanguagesFIJ),
		string(AudioLanguagesFIN),
		string(AudioLanguagesFRA),
		string(AudioLanguagesFRY),
		string(AudioLanguagesFUL),
		string(AudioLanguagesGLA),
		string(AudioLanguagesGLE),
		string(AudioLanguagesGLG),
		string(AudioLanguagesGLV),
		string(AudioLanguagesGRN),
		string(AudioLanguagesGUJ),
		string(AudioLanguagesHAT),
		string(AudioLanguagesHAU),
		string(AudioLanguagesHEB),
		string(AudioLanguagesHER),
		string(AudioLanguagesHIN),
		string(AudioLanguagesHMO),
		string(AudioLanguagesHRV),
		string(AudioLanguagesHUN),
		string(AudioLanguagesHYE),
		string(AudioLanguagesIBO),
		string(AudioLanguagesIDO),
		string(AudioLanguagesIII),
		string(AudioLanguagesIKU),
		string(AudioLanguagesILE),
		string(AudioLanguagesINA),
		string(AudioLanguagesIND),
		string(AudioLanguagesIPK),
		string(AudioLanguagesISL),
		string(AudioLanguagesITA),
		string(AudioLanguagesJAV),
		string(AudioLanguagesJPN),
		string(AudioLanguagesKAL),
		string(AudioLanguagesKAN),
		string(AudioLanguagesKAS),
		string(AudioLanguagesKAT),
		string(AudioLanguagesKAU),
		string(AudioLanguagesKAZ),
		string(AudioLanguagesKHM),
		string(AudioLanguagesKIK),
		string(AudioLanguagesKIN),
		string(AudioLanguagesKIR),
		string(AudioLanguagesKOM),
		string(AudioLanguagesKON),
		string(AudioLanguagesKOR),
		string(AudioLanguagesKUA),
		string(AudioLanguagesKUR),
		string(AudioLanguagesLAO),
		string(AudioLanguagesLAT),
		string(AudioLanguagesLAV),
		string(AudioLanguagesLIM),
		string(AudioLanguagesLIN),
		string(AudioLanguagesLIT),
		string(AudioLanguagesLTZ),
		string(AudioLanguagesLUB),
		string(AudioLanguagesLUG),
		string(AudioLanguagesMAH),
		string(AudioLanguagesMAL),
		string(AudioLanguagesMAR),
		string(AudioLanguagesMKD),
		string(AudioLanguagesMLG1),
		string(AudioLanguagesMLT),
		string(AudioLanguagesMON),
		string(AudioLanguagesMRI),
		string(AudioLanguagesMSA6),
		string(AudioLanguagesMYA),
		string(AudioLanguagesNAU),
		string(AudioLanguagesNAV),
		string(AudioLanguagesNBL),
		string(AudioLanguagesNDE),
		string(AudioLanguagesNDO),
		string(AudioLanguagesNEP),
		string(AudioLanguagesNLD),
		string(AudioLanguagesNNO),
		string(AudioLanguagesNOB),
		string(AudioLanguagesNOR),
		string(AudioLanguagesNYA),
		string(AudioLanguagesOCI),
		string(AudioLanguagesOJI),
		string(AudioLanguagesORI),
		string(AudioLanguagesORM),
		string(AudioLanguagesOSS),
		string(AudioLanguagesPAN),
		string(AudioLanguagesPLI),
		string(AudioLanguagesPOL),
		string(AudioLanguagesPOR),
		string(AudioLanguagesPUS),
		string(AudioLanguagesQUE3),
		string(AudioLanguagesROH),
		string(AudioLanguagesRON),
		string(AudioLanguagesRUN),
		string(AudioLanguagesRUS),
		string(AudioLanguagesSAG),
		string(AudioLanguagesSAN),
		string(AudioLanguagesSIN),
		string(AudioLanguagesSLK),
		string(AudioLanguagesSLV),
		string(AudioLanguagesSME),
		string(AudioLanguagesSMO),
		string(AudioLanguagesSNA),
		string(AudioLanguagesSND),
		string(AudioLanguagesSOM),
		string(AudioLanguagesSOT),
		string(AudioLanguagesSPA),
		string(AudioLanguagesSQI),
		string(AudioLanguagesSRD),
		string(AudioLanguagesSRP),
		string(AudioLanguagesSSW),
		string(AudioLanguagesSUN),
		string(AudioLanguagesSWA),
		string(AudioLanguagesSWE),
		string(AudioLanguagesTAH),
		string(AudioLanguagesTAM),
		string(AudioLanguagesTAT),
		string(AudioLanguagesTEL),
		string(AudioLanguagesTGK),
		string(AudioLanguagesTGL),
		string(AudioLanguagesTHA),
		string(AudioLanguagesTIR),
		string(AudioLanguagesTON),
		string(AudioLanguagesTSN),
		string(AudioLanguagesTSO),
		string(AudioLanguagesTUK),
		string(AudioLanguagesTUR),
		string(AudioLanguagesTWI),
		string(AudioLanguagesUIG),
		string(AudioLanguagesUKR),
		string(AudioLanguagesURD),
		string(AudioLanguagesUZB),
		string(AudioLanguagesVEN),
		string(AudioLanguagesVIE),
		string(AudioLanguagesVOL),
		string(AudioLanguagesWLN),
		string(AudioLanguagesWOL),
		string(AudioLanguagesXHO),
		string(AudioLanguagesYID),
		string(AudioLanguagesYOR),
		string(AudioLanguagesZHA6),
		string(AudioLanguagesZHO6),
		string(AudioLanguagesZUL),
	}

	// InKnownAudioLanguages is an ozzo-validator for AudioLanguages
	InKnownAudioLanguages = validation.In(
		AudioLanguagesAAR,
		AudioLanguagesABK,
		AudioLanguagesAFR,
		AudioLanguagesAKA,
		AudioLanguagesAMH,
		AudioLanguagesARA9,
		AudioLanguagesARG,
		AudioLanguagesASM,
		AudioLanguagesAVA,
		AudioLanguagesAVE,
		AudioLanguagesAYM,
		AudioLanguagesAZE,
		AudioLanguagesBAK,
		AudioLanguagesBAM,
		AudioLanguagesBEL,
		AudioLanguagesBEN,
		AudioLanguagesBIS,
		AudioLanguagesBOD,
		AudioLanguagesBOS,
		AudioLanguagesBRE,
		AudioLanguagesBUL,
		AudioLanguagesCAT,
		AudioLanguagesCES,
		AudioLanguagesCHA,
		AudioLanguagesCHE,
		AudioLanguagesCHU,
		AudioLanguagesCHV,
		AudioLanguagesCOR,
		AudioLanguagesCOS,
		AudioLanguagesCRE,
		AudioLanguagesCYM,
		AudioLanguagesDAN,
		AudioLanguagesDEU,
		AudioLanguagesDIV,
		AudioLanguagesDZO,
		AudioLanguagesELL,
		AudioLanguagesENG,
		AudioLanguagesEPO,
		AudioLanguagesEST,
		AudioLanguagesEUS,
		AudioLanguagesEWE,
		AudioLanguagesFAO,
		AudioLanguagesFAS,
		AudioLanguagesFIJ,
		AudioLanguagesFIN,
		AudioLanguagesFRA,
		AudioLanguagesFRY,
		AudioLanguagesFUL,
		AudioLanguagesGLA,
		AudioLanguagesGLE,
		AudioLanguagesGLG,
		AudioLanguagesGLV,
		AudioLanguagesGRN,
		AudioLanguagesGUJ,
		AudioLanguagesHAT,
		AudioLanguagesHAU,
		AudioLanguagesHEB,
		AudioLanguagesHER,
		AudioLanguagesHIN,
		AudioLanguagesHMO,
		AudioLanguagesHRV,
		AudioLanguagesHUN,
		AudioLanguagesHYE,
		AudioLanguagesIBO,
		AudioLanguagesIDO,
		AudioLanguagesIII,
		AudioLanguagesIKU,
		AudioLanguagesILE,
		AudioLanguagesINA,
		AudioLanguagesIND,
		AudioLanguagesIPK,
		AudioLanguagesISL,
		AudioLanguagesITA,
		AudioLanguagesJAV,
		AudioLanguagesJPN,
		AudioLanguagesKAL,
		AudioLanguagesKAN,
		AudioLanguagesKAS,
		AudioLanguagesKAT,
		AudioLanguagesKAU,
		AudioLanguagesKAZ,
		AudioLanguagesKHM,
		AudioLanguagesKIK,
		AudioLanguagesKIN,
		AudioLanguagesKIR,
		AudioLanguagesKOM,
		AudioLanguagesKON,
		AudioLanguagesKOR,
		AudioLanguagesKUA,
		AudioLanguagesKUR,
		AudioLanguagesLAO,
		AudioLanguagesLAT,
		AudioLanguagesLAV,
		AudioLanguagesLIM,
		AudioLanguagesLIN,
		AudioLanguagesLIT,
		AudioLanguagesLTZ,
		AudioLanguagesLUB,
		AudioLanguagesLUG,
		AudioLanguagesMAH,
		AudioLanguagesMAL,
		AudioLanguagesMAR,
		AudioLanguagesMKD,
		AudioLanguagesMLG1,
		AudioLanguagesMLT,
		AudioLanguagesMON,
		AudioLanguagesMRI,
		AudioLanguagesMSA6,
		AudioLanguagesMYA,
		AudioLanguagesNAU,
		AudioLanguagesNAV,
		AudioLanguagesNBL,
		AudioLanguagesNDE,
		AudioLanguagesNDO,
		AudioLanguagesNEP,
		AudioLanguagesNLD,
		AudioLanguagesNNO,
		AudioLanguagesNOB,
		AudioLanguagesNOR,
		AudioLanguagesNYA,
		AudioLanguagesOCI,
		AudioLanguagesOJI,
		AudioLanguagesORI,
		AudioLanguagesORM,
		AudioLanguagesOSS,
		AudioLanguagesPAN,
		AudioLanguagesPLI,
		AudioLanguagesPOL,
		AudioLanguagesPOR,
		AudioLanguagesPUS,
		AudioLanguagesQUE3,
		AudioLanguagesROH,
		AudioLanguagesRON,
		AudioLanguagesRUN,
		AudioLanguagesRUS,
		AudioLanguagesSAG,
		AudioLanguagesSAN,
		AudioLanguagesSIN,
		AudioLanguagesSLK,
		AudioLanguagesSLV,
		AudioLanguagesSME,
		AudioLanguagesSMO,
		AudioLanguagesSNA,
		AudioLanguagesSND,
		AudioLanguagesSOM,
		AudioLanguagesSOT,
		AudioLanguagesSPA,
		AudioLanguagesSQI,
		AudioLanguagesSRD,
		AudioLanguagesSRP,
		AudioLanguagesSSW,
		AudioLanguagesSUN,
		AudioLanguagesSWA,
		AudioLanguagesSWE,
		AudioLanguagesTAH,
		AudioLanguagesTAM,
		AudioLanguagesTAT,
		AudioLanguagesTEL,
		AudioLanguagesTGK,
		AudioLanguagesTGL,
		AudioLanguagesTHA,
		AudioLanguagesTIR,
		AudioLanguagesTON,
		AudioLanguagesTSN,
		AudioLanguagesTSO,
		AudioLanguagesTUK,
		AudioLanguagesTUR,
		AudioLanguagesTWI,
		AudioLanguagesUIG,
		AudioLanguagesUKR,
		AudioLanguagesURD,
		AudioLanguagesUZB,
		AudioLanguagesVEN,
		AudioLanguagesVIE,
		AudioLanguagesVOL,
		AudioLanguagesWLN,
		AudioLanguagesWOL,
		AudioLanguagesXHO,
		AudioLanguagesYID,
		AudioLanguagesYOR,
		AudioLanguagesZHA6,
		AudioLanguagesZHO6,
		AudioLanguagesZUL,
	)
)
