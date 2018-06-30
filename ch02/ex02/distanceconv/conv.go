package distanceconv

func FToM(f Feet) Metre { return Metre(f * 0.3048) }
func MToF(m Metre) Feet{ return Feet(m / 0.3048) }