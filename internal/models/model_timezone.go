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

// Timezone is an enum.
type Timezone string

// Validate implements basic validation for this model
func (m Timezone) Validate() error {
	return InKnownTimezone.Validate(m)
}

var (
	TimezoneAbuDhabi                  Timezone = "Abu Dhabi"
	TimezoneAdelaide                  Timezone = "Adelaide"
	TimezoneAlaska                    Timezone = "Alaska"
	TimezoneAlmaty                    Timezone = "Almaty"
	TimezoneAmericanSamoa             Timezone = "American Samoa"
	TimezoneAmsterdam                 Timezone = "Amsterdam"
	TimezoneArizona                   Timezone = "Arizona"
	TimezoneAstana                    Timezone = "Astana"
	TimezoneAthens                    Timezone = "Athens"
	TimezoneAtlanticTimeCanada        Timezone = "Atlantic Time (Canada)"
	TimezoneAuckland                  Timezone = "Auckland"
	TimezoneAzores                    Timezone = "Azores"
	TimezoneBaghdad                   Timezone = "Baghdad"
	TimezoneBaku                      Timezone = "Baku"
	TimezoneBangkok                   Timezone = "Bangkok"
	TimezoneBeijing                   Timezone = "Beijing"
	TimezoneBelgrade                  Timezone = "Belgrade"
	TimezoneBerlin                    Timezone = "Berlin"
	TimezoneBern                      Timezone = "Bern"
	TimezoneBogota                    Timezone = "Bogota"
	TimezoneBrasilia                  Timezone = "Brasilia"
	TimezoneBratislava                Timezone = "Bratislava"
	TimezoneBrisbane                  Timezone = "Brisbane"
	TimezoneBrussels                  Timezone = "Brussels"
	TimezoneBucharest                 Timezone = "Bucharest"
	TimezoneBudapest                  Timezone = "Budapest"
	TimezoneBuenosAires               Timezone = "Buenos Aires"
	TimezoneCairo                     Timezone = "Cairo"
	TimezoneCanberra                  Timezone = "Canberra"
	TimezoneCapeVerdeIs               Timezone = "Cape Verde Is."
	TimezoneCaracas                   Timezone = "Caracas"
	TimezoneCasablanca                Timezone = "Casablanca"
	TimezoneCentralAmerica            Timezone = "Central America"
	TimezoneCentralTimeUSAmpCanada    Timezone = "Central Time (US &amp; Canada)"
	TimezoneChathamIs                 Timezone = "Chatham Is"
	TimezoneChennai                   Timezone = "Chennai"
	TimezoneChihuahua                 Timezone = "Chihuahua"
	TimezoneChongqing                 Timezone = "Chongqing"
	TimezoneCopenhagen                Timezone = "Copenhagen"
	TimezoneDarwin                    Timezone = "Darwin"
	TimezoneDhaka                     Timezone = "Dhaka"
	TimezoneDublin                    Timezone = "Dublin"
	TimezoneEasternTimeUSAmpCanada    Timezone = "Eastern Time (US &amp; Canada)"
	TimezoneEdinburgh                 Timezone = "Edinburgh"
	TimezoneEkaterinburg              Timezone = "Ekaterinburg"
	TimezoneFiji                      Timezone = "Fiji"
	TimezoneGeorgetown                Timezone = "Georgetown"
	TimezoneGreenland                 Timezone = "Greenland"
	TimezoneGuadalajara               Timezone = "Guadalajara"
	TimezoneGuam                      Timezone = "Guam"
	TimezoneHanoi                     Timezone = "Hanoi"
	TimezoneHarare                    Timezone = "Harare"
	TimezoneHawaii                    Timezone = "Hawaii"
	TimezoneHelsinki                  Timezone = "Helsinki"
	TimezoneHobart                    Timezone = "Hobart"
	TimezoneHongKong                  Timezone = "Hong Kong"
	TimezoneIndianaEast               Timezone = "Indiana (East)"
	TimezoneInternationalDateLineWest Timezone = "International Date Line West"
	TimezoneIrkutsk                   Timezone = "Irkutsk"
	TimezoneIslamabad                 Timezone = "Islamabad"
	TimezoneIstanbul                  Timezone = "Istanbul"
	TimezoneJakarta                   Timezone = "Jakarta"
	TimezoneJerusalem                 Timezone = "Jerusalem"
	TimezoneKabul                     Timezone = "Kabul"
	TimezoneKaliningrad               Timezone = "Kaliningrad"
	TimezoneKamchatka                 Timezone = "Kamchatka"
	TimezoneKarachi                   Timezone = "Karachi"
	TimezoneKathmandu                 Timezone = "Kathmandu"
	TimezoneKolkata                   Timezone = "Kolkata"
	TimezoneKrasnoyarsk               Timezone = "Krasnoyarsk"
	TimezoneKualaLumpur               Timezone = "Kuala Lumpur"
	TimezoneKuwait                    Timezone = "Kuwait"
	TimezoneKyiv                      Timezone = "Kyiv"
	TimezoneLaPaz                     Timezone = "La Paz"
	TimezoneLima                      Timezone = "Lima"
	TimezoneLisbon                    Timezone = "Lisbon"
	TimezoneLjubljana                 Timezone = "Ljubljana"
	TimezoneLondon                    Timezone = "London"
	TimezoneMadrid                    Timezone = "Madrid"
	TimezoneMagadan                   Timezone = "Magadan"
	TimezoneMarshallIs                Timezone = "Marshall Is."
	TimezoneMazatlan                  Timezone = "Mazatlan"
	TimezoneMelbourne                 Timezone = "Melbourne"
	TimezoneMexicoCity                Timezone = "Mexico City"
	TimezoneMidAtlantic               Timezone = "Mid-Atlantic"
	TimezoneMidwayIsland              Timezone = "Midway Island"
	TimezoneMinsk                     Timezone = "Minsk"
	TimezoneMonrovia                  Timezone = "Monrovia"
	TimezoneMonterrey                 Timezone = "Monterrey"
	TimezoneMontevideo                Timezone = "Montevideo"
	TimezoneMoscow                    Timezone = "Moscow"
	TimezoneMountainTimeUSAmpCanada   Timezone = "Mountain Time (US &amp; Canada)"
	TimezoneMumbai                    Timezone = "Mumbai"
	TimezoneMuscat                    Timezone = "Muscat"
	TimezoneNairobi                   Timezone = "Nairobi"
	TimezoneNewCaledonia              Timezone = "New Caledonia"
	TimezoneNewDelhi                  Timezone = "New Delhi"
	TimezoneNewfoundland              Timezone = "Newfoundland"
	TimezoneNovosibirsk               Timezone = "Novosibirsk"
	TimezoneNuku39Alofa               Timezone = "Nuku&#39;alofa"
	TimezoneOsaka                     Timezone = "Osaka"
	TimezonePacificTimeUSAmpCanada    Timezone = "Pacific Time (US &amp; Canada)"
	TimezoneParis                     Timezone = "Paris"
	TimezonePerth                     Timezone = "Perth"
	TimezonePortMoresby               Timezone = "Port Moresby"
	TimezonePrague                    Timezone = "Prague"
	TimezonePretoria                  Timezone = "Pretoria"
	TimezonePuertoRico                Timezone = "Puerto Rico"
	TimezoneQuito                     Timezone = "Quito"
	TimezoneRangoon                   Timezone = "Rangoon"
	TimezoneRiga                      Timezone = "Riga"
	TimezoneRiyadh                    Timezone = "Riyadh"
	TimezoneRome                      Timezone = "Rome"
	TimezoneSamara                    Timezone = "Samara"
	TimezoneSamoa                     Timezone = "Samoa"
	TimezoneSantiago                  Timezone = "Santiago"
	TimezoneSapporo                   Timezone = "Sapporo"
	TimezoneSarajevo                  Timezone = "Sarajevo"
	TimezoneSaskatchewan              Timezone = "Saskatchewan"
	TimezoneSeoul                     Timezone = "Seoul"
	TimezoneSingapore                 Timezone = "Singapore"
	TimezoneSkopje                    Timezone = "Skopje"
	TimezoneSofia                     Timezone = "Sofia"
	TimezoneSolomonIs                 Timezone = "Solomon Is."
	TimezoneSrednekolymsk             Timezone = "Srednekolymsk"
	TimezoneSriJayawardenepura        Timezone = "Sri Jayawardenepura"
	TimezoneStPetersburg              Timezone = "St. Petersburg"
	TimezoneStockholm                 Timezone = "Stockholm"
	TimezoneSydney                    Timezone = "Sydney"
	TimezoneTaipei                    Timezone = "Taipei"
	TimezoneTallinn                   Timezone = "Tallinn"
	TimezoneTashkent                  Timezone = "Tashkent"
	TimezoneTbilisi                   Timezone = "Tbilisi"
	TimezoneTehran                    Timezone = "Tehran"
	TimezoneTijuana                   Timezone = "Tijuana"
	TimezoneTokelauIs                 Timezone = "Tokelau Is."
	TimezoneTokyo                     Timezone = "Tokyo"
	TimezoneUTC                       Timezone = "UTC"
	TimezoneUlaanbaatar               Timezone = "Ulaanbaatar"
	TimezoneUrumqi                    Timezone = "Urumqi"
	TimezoneVienna                    Timezone = "Vienna"
	TimezoneVilnius                   Timezone = "Vilnius"
	TimezoneVladivostok               Timezone = "Vladivostok"
	TimezoneVolgograd                 Timezone = "Volgograd"
	TimezoneWarsaw                    Timezone = "Warsaw"
	TimezoneWellington                Timezone = "Wellington"
	TimezoneWestCentralAfrica         Timezone = "West Central Africa"
	TimezoneYakutsk                   Timezone = "Yakutsk"
	TimezoneYerevan                   Timezone = "Yerevan"
	TimezoneZagreb                    Timezone = "Zagreb"
	TimezoneZurich                    Timezone = "Zurich"

	// KnownTimezone is the list of valid Timezone
	KnownTimezone = []Timezone{
		TimezoneAbuDhabi,
		TimezoneAdelaide,
		TimezoneAlaska,
		TimezoneAlmaty,
		TimezoneAmericanSamoa,
		TimezoneAmsterdam,
		TimezoneArizona,
		TimezoneAstana,
		TimezoneAthens,
		TimezoneAtlanticTimeCanada,
		TimezoneAuckland,
		TimezoneAzores,
		TimezoneBaghdad,
		TimezoneBaku,
		TimezoneBangkok,
		TimezoneBeijing,
		TimezoneBelgrade,
		TimezoneBerlin,
		TimezoneBern,
		TimezoneBogota,
		TimezoneBrasilia,
		TimezoneBratislava,
		TimezoneBrisbane,
		TimezoneBrussels,
		TimezoneBucharest,
		TimezoneBudapest,
		TimezoneBuenosAires,
		TimezoneCairo,
		TimezoneCanberra,
		TimezoneCapeVerdeIs,
		TimezoneCaracas,
		TimezoneCasablanca,
		TimezoneCentralAmerica,
		TimezoneCentralTimeUSAmpCanada,
		TimezoneChathamIs,
		TimezoneChennai,
		TimezoneChihuahua,
		TimezoneChongqing,
		TimezoneCopenhagen,
		TimezoneDarwin,
		TimezoneDhaka,
		TimezoneDublin,
		TimezoneEasternTimeUSAmpCanada,
		TimezoneEdinburgh,
		TimezoneEkaterinburg,
		TimezoneFiji,
		TimezoneGeorgetown,
		TimezoneGreenland,
		TimezoneGuadalajara,
		TimezoneGuam,
		TimezoneHanoi,
		TimezoneHarare,
		TimezoneHawaii,
		TimezoneHelsinki,
		TimezoneHobart,
		TimezoneHongKong,
		TimezoneIndianaEast,
		TimezoneInternationalDateLineWest,
		TimezoneIrkutsk,
		TimezoneIslamabad,
		TimezoneIstanbul,
		TimezoneJakarta,
		TimezoneJerusalem,
		TimezoneKabul,
		TimezoneKaliningrad,
		TimezoneKamchatka,
		TimezoneKarachi,
		TimezoneKathmandu,
		TimezoneKolkata,
		TimezoneKrasnoyarsk,
		TimezoneKualaLumpur,
		TimezoneKuwait,
		TimezoneKyiv,
		TimezoneLaPaz,
		TimezoneLima,
		TimezoneLisbon,
		TimezoneLjubljana,
		TimezoneLondon,
		TimezoneMadrid,
		TimezoneMagadan,
		TimezoneMarshallIs,
		TimezoneMazatlan,
		TimezoneMelbourne,
		TimezoneMexicoCity,
		TimezoneMidAtlantic,
		TimezoneMidwayIsland,
		TimezoneMinsk,
		TimezoneMonrovia,
		TimezoneMonterrey,
		TimezoneMontevideo,
		TimezoneMoscow,
		TimezoneMountainTimeUSAmpCanada,
		TimezoneMumbai,
		TimezoneMuscat,
		TimezoneNairobi,
		TimezoneNewCaledonia,
		TimezoneNewDelhi,
		TimezoneNewfoundland,
		TimezoneNovosibirsk,
		TimezoneNuku39Alofa,
		TimezoneOsaka,
		TimezonePacificTimeUSAmpCanada,
		TimezoneParis,
		TimezonePerth,
		TimezonePortMoresby,
		TimezonePrague,
		TimezonePretoria,
		TimezonePuertoRico,
		TimezoneQuito,
		TimezoneRangoon,
		TimezoneRiga,
		TimezoneRiyadh,
		TimezoneRome,
		TimezoneSamara,
		TimezoneSamoa,
		TimezoneSantiago,
		TimezoneSapporo,
		TimezoneSarajevo,
		TimezoneSaskatchewan,
		TimezoneSeoul,
		TimezoneSingapore,
		TimezoneSkopje,
		TimezoneSofia,
		TimezoneSolomonIs,
		TimezoneSrednekolymsk,
		TimezoneSriJayawardenepura,
		TimezoneStPetersburg,
		TimezoneStockholm,
		TimezoneSydney,
		TimezoneTaipei,
		TimezoneTallinn,
		TimezoneTashkent,
		TimezoneTbilisi,
		TimezoneTehran,
		TimezoneTijuana,
		TimezoneTokelauIs,
		TimezoneTokyo,
		TimezoneUTC,
		TimezoneUlaanbaatar,
		TimezoneUrumqi,
		TimezoneVienna,
		TimezoneVilnius,
		TimezoneVladivostok,
		TimezoneVolgograd,
		TimezoneWarsaw,
		TimezoneWellington,
		TimezoneWestCentralAfrica,
		TimezoneYakutsk,
		TimezoneYerevan,
		TimezoneZagreb,
		TimezoneZurich,
	}
	// KnownTimezoneString is the list of valid Timezone as string
	KnownTimezoneString = []string{
		string(TimezoneAbuDhabi),
		string(TimezoneAdelaide),
		string(TimezoneAlaska),
		string(TimezoneAlmaty),
		string(TimezoneAmericanSamoa),
		string(TimezoneAmsterdam),
		string(TimezoneArizona),
		string(TimezoneAstana),
		string(TimezoneAthens),
		string(TimezoneAtlanticTimeCanada),
		string(TimezoneAuckland),
		string(TimezoneAzores),
		string(TimezoneBaghdad),
		string(TimezoneBaku),
		string(TimezoneBangkok),
		string(TimezoneBeijing),
		string(TimezoneBelgrade),
		string(TimezoneBerlin),
		string(TimezoneBern),
		string(TimezoneBogota),
		string(TimezoneBrasilia),
		string(TimezoneBratislava),
		string(TimezoneBrisbane),
		string(TimezoneBrussels),
		string(TimezoneBucharest),
		string(TimezoneBudapest),
		string(TimezoneBuenosAires),
		string(TimezoneCairo),
		string(TimezoneCanberra),
		string(TimezoneCapeVerdeIs),
		string(TimezoneCaracas),
		string(TimezoneCasablanca),
		string(TimezoneCentralAmerica),
		string(TimezoneCentralTimeUSAmpCanada),
		string(TimezoneChathamIs),
		string(TimezoneChennai),
		string(TimezoneChihuahua),
		string(TimezoneChongqing),
		string(TimezoneCopenhagen),
		string(TimezoneDarwin),
		string(TimezoneDhaka),
		string(TimezoneDublin),
		string(TimezoneEasternTimeUSAmpCanada),
		string(TimezoneEdinburgh),
		string(TimezoneEkaterinburg),
		string(TimezoneFiji),
		string(TimezoneGeorgetown),
		string(TimezoneGreenland),
		string(TimezoneGuadalajara),
		string(TimezoneGuam),
		string(TimezoneHanoi),
		string(TimezoneHarare),
		string(TimezoneHawaii),
		string(TimezoneHelsinki),
		string(TimezoneHobart),
		string(TimezoneHongKong),
		string(TimezoneIndianaEast),
		string(TimezoneInternationalDateLineWest),
		string(TimezoneIrkutsk),
		string(TimezoneIslamabad),
		string(TimezoneIstanbul),
		string(TimezoneJakarta),
		string(TimezoneJerusalem),
		string(TimezoneKabul),
		string(TimezoneKaliningrad),
		string(TimezoneKamchatka),
		string(TimezoneKarachi),
		string(TimezoneKathmandu),
		string(TimezoneKolkata),
		string(TimezoneKrasnoyarsk),
		string(TimezoneKualaLumpur),
		string(TimezoneKuwait),
		string(TimezoneKyiv),
		string(TimezoneLaPaz),
		string(TimezoneLima),
		string(TimezoneLisbon),
		string(TimezoneLjubljana),
		string(TimezoneLondon),
		string(TimezoneMadrid),
		string(TimezoneMagadan),
		string(TimezoneMarshallIs),
		string(TimezoneMazatlan),
		string(TimezoneMelbourne),
		string(TimezoneMexicoCity),
		string(TimezoneMidAtlantic),
		string(TimezoneMidwayIsland),
		string(TimezoneMinsk),
		string(TimezoneMonrovia),
		string(TimezoneMonterrey),
		string(TimezoneMontevideo),
		string(TimezoneMoscow),
		string(TimezoneMountainTimeUSAmpCanada),
		string(TimezoneMumbai),
		string(TimezoneMuscat),
		string(TimezoneNairobi),
		string(TimezoneNewCaledonia),
		string(TimezoneNewDelhi),
		string(TimezoneNewfoundland),
		string(TimezoneNovosibirsk),
		string(TimezoneNuku39Alofa),
		string(TimezoneOsaka),
		string(TimezonePacificTimeUSAmpCanada),
		string(TimezoneParis),
		string(TimezonePerth),
		string(TimezonePortMoresby),
		string(TimezonePrague),
		string(TimezonePretoria),
		string(TimezonePuertoRico),
		string(TimezoneQuito),
		string(TimezoneRangoon),
		string(TimezoneRiga),
		string(TimezoneRiyadh),
		string(TimezoneRome),
		string(TimezoneSamara),
		string(TimezoneSamoa),
		string(TimezoneSantiago),
		string(TimezoneSapporo),
		string(TimezoneSarajevo),
		string(TimezoneSaskatchewan),
		string(TimezoneSeoul),
		string(TimezoneSingapore),
		string(TimezoneSkopje),
		string(TimezoneSofia),
		string(TimezoneSolomonIs),
		string(TimezoneSrednekolymsk),
		string(TimezoneSriJayawardenepura),
		string(TimezoneStPetersburg),
		string(TimezoneStockholm),
		string(TimezoneSydney),
		string(TimezoneTaipei),
		string(TimezoneTallinn),
		string(TimezoneTashkent),
		string(TimezoneTbilisi),
		string(TimezoneTehran),
		string(TimezoneTijuana),
		string(TimezoneTokelauIs),
		string(TimezoneTokyo),
		string(TimezoneUTC),
		string(TimezoneUlaanbaatar),
		string(TimezoneUrumqi),
		string(TimezoneVienna),
		string(TimezoneVilnius),
		string(TimezoneVladivostok),
		string(TimezoneVolgograd),
		string(TimezoneWarsaw),
		string(TimezoneWellington),
		string(TimezoneWestCentralAfrica),
		string(TimezoneYakutsk),
		string(TimezoneYerevan),
		string(TimezoneZagreb),
		string(TimezoneZurich),
	}

	// InKnownTimezone is an ozzo-validator for Timezone
	InKnownTimezone = validation.In(
		TimezoneAbuDhabi,
		TimezoneAdelaide,
		TimezoneAlaska,
		TimezoneAlmaty,
		TimezoneAmericanSamoa,
		TimezoneAmsterdam,
		TimezoneArizona,
		TimezoneAstana,
		TimezoneAthens,
		TimezoneAtlanticTimeCanada,
		TimezoneAuckland,
		TimezoneAzores,
		TimezoneBaghdad,
		TimezoneBaku,
		TimezoneBangkok,
		TimezoneBeijing,
		TimezoneBelgrade,
		TimezoneBerlin,
		TimezoneBern,
		TimezoneBogota,
		TimezoneBrasilia,
		TimezoneBratislava,
		TimezoneBrisbane,
		TimezoneBrussels,
		TimezoneBucharest,
		TimezoneBudapest,
		TimezoneBuenosAires,
		TimezoneCairo,
		TimezoneCanberra,
		TimezoneCapeVerdeIs,
		TimezoneCaracas,
		TimezoneCasablanca,
		TimezoneCentralAmerica,
		TimezoneCentralTimeUSAmpCanada,
		TimezoneChathamIs,
		TimezoneChennai,
		TimezoneChihuahua,
		TimezoneChongqing,
		TimezoneCopenhagen,
		TimezoneDarwin,
		TimezoneDhaka,
		TimezoneDublin,
		TimezoneEasternTimeUSAmpCanada,
		TimezoneEdinburgh,
		TimezoneEkaterinburg,
		TimezoneFiji,
		TimezoneGeorgetown,
		TimezoneGreenland,
		TimezoneGuadalajara,
		TimezoneGuam,
		TimezoneHanoi,
		TimezoneHarare,
		TimezoneHawaii,
		TimezoneHelsinki,
		TimezoneHobart,
		TimezoneHongKong,
		TimezoneIndianaEast,
		TimezoneInternationalDateLineWest,
		TimezoneIrkutsk,
		TimezoneIslamabad,
		TimezoneIstanbul,
		TimezoneJakarta,
		TimezoneJerusalem,
		TimezoneKabul,
		TimezoneKaliningrad,
		TimezoneKamchatka,
		TimezoneKarachi,
		TimezoneKathmandu,
		TimezoneKolkata,
		TimezoneKrasnoyarsk,
		TimezoneKualaLumpur,
		TimezoneKuwait,
		TimezoneKyiv,
		TimezoneLaPaz,
		TimezoneLima,
		TimezoneLisbon,
		TimezoneLjubljana,
		TimezoneLondon,
		TimezoneMadrid,
		TimezoneMagadan,
		TimezoneMarshallIs,
		TimezoneMazatlan,
		TimezoneMelbourne,
		TimezoneMexicoCity,
		TimezoneMidAtlantic,
		TimezoneMidwayIsland,
		TimezoneMinsk,
		TimezoneMonrovia,
		TimezoneMonterrey,
		TimezoneMontevideo,
		TimezoneMoscow,
		TimezoneMountainTimeUSAmpCanada,
		TimezoneMumbai,
		TimezoneMuscat,
		TimezoneNairobi,
		TimezoneNewCaledonia,
		TimezoneNewDelhi,
		TimezoneNewfoundland,
		TimezoneNovosibirsk,
		TimezoneNuku39Alofa,
		TimezoneOsaka,
		TimezonePacificTimeUSAmpCanada,
		TimezoneParis,
		TimezonePerth,
		TimezonePortMoresby,
		TimezonePrague,
		TimezonePretoria,
		TimezonePuertoRico,
		TimezoneQuito,
		TimezoneRangoon,
		TimezoneRiga,
		TimezoneRiyadh,
		TimezoneRome,
		TimezoneSamara,
		TimezoneSamoa,
		TimezoneSantiago,
		TimezoneSapporo,
		TimezoneSarajevo,
		TimezoneSaskatchewan,
		TimezoneSeoul,
		TimezoneSingapore,
		TimezoneSkopje,
		TimezoneSofia,
		TimezoneSolomonIs,
		TimezoneSrednekolymsk,
		TimezoneSriJayawardenepura,
		TimezoneStPetersburg,
		TimezoneStockholm,
		TimezoneSydney,
		TimezoneTaipei,
		TimezoneTallinn,
		TimezoneTashkent,
		TimezoneTbilisi,
		TimezoneTehran,
		TimezoneTijuana,
		TimezoneTokelauIs,
		TimezoneTokyo,
		TimezoneUTC,
		TimezoneUlaanbaatar,
		TimezoneUrumqi,
		TimezoneVienna,
		TimezoneVilnius,
		TimezoneVladivostok,
		TimezoneVolgograd,
		TimezoneWarsaw,
		TimezoneWellington,
		TimezoneWestCentralAfrica,
		TimezoneYakutsk,
		TimezoneYerevan,
		TimezoneZagreb,
		TimezoneZurich,
	)
)
