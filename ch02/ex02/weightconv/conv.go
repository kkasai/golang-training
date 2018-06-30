package weightconv

func PToK(p Pound) Kilogram { return Kilogram(p * 0.45359237) }
func KToP(k Kilogram) Pound{ return Pound(k / 0.45359237) }