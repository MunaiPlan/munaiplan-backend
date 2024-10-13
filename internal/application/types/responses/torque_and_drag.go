package responses

// EffectiveTensionResponse represents the response structure for the Effective Tension prediction.
type EffectiveTensionFromMLModelResponse struct {
	Depth                           []float64 `json:"depth"`
	TowerLoadCapacity               []float64 `json:"Грузоподъёмность вышки"`
	RotaryDrilling                  []float64 `json:"Бурение ротором"`
	HelicalBucklingWithoutRotation  []float64 `json:"Спиральный изгиб(без вращения)"`
	PullUp                          []float64 `json:"Подъём"`
	SinusoidalBucklingAllOperations []float64 `json:"Синусоидальный изгиб(все операции)"`
	RunIn                           []float64 `json:"Спуск"`
	DrillingGZD                     []float64 `json:"Бурение ГЗД"`
	HelicalBucklingWithRotation     []float64 `json:"Спиральный изгиб(с вращением)"`
	TensionLimit                    []float64 `json:"Предел натяжения"`
}