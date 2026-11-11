package horarios

import "time"

type DiaDelAnio struct {
	mes time.Month
	dia uint8
}

type Temporada struct {
	desde          DiaDelAnio
	hasta          DiaDelAnio
	horarioSemanal map[time.Weekday][]FranjaDeApertura
}
