package horarios

type FranjaDeApertura struct {
	apertura      HoraDelDia
	cierre        HoraDelDia
	ultimaEntrada *HoraDelDia
}
