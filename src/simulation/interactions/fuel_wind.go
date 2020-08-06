package interactions

// // TreeImport fuel.Tree_data
// type TreeImport fuel.Tree_data

// UpdateMoisture controls the water content depending on the temperature of the surrounding temperature
func (t *TreeImport) UpdateMoisture(temperature float64) {
	temperatureDiff := temperature - 25.0 // 25ºC. This would be an equilibirum point where no water transfers occur
	diff := 0.0
	if temperatureDiff > 0 {
		diff = 0.01 * temperatureDiff
	} else {
		diff = 0
	}
	t.Dynamic.Moisture = t.Dynamic.Moisture - (diff / t.Static.BarkThickness)
}
