package responses

// EffectiveTensionResponse represents the response structure for the Effective Tension prediction.
type EffectiveTensionFromMLModelResponse struct {
	Depth                           []float64 `json:"Глубина"`
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

// WeightOnBitFromMlModel represents the response structure for the Weight on Bit prediction.
type WeightOnBitFromMLModelResponse struct {
	Depth                           []float64 `json:"Глубина"`
	TowerLoadCapacity               []float64 `json:"Грузоподъёмность вышки"`
	RotaryDrilling                  []float64 `json:"Бурение ротором"`
	PullUp                          []float64 `json:"Подъём"`
	RunIn                           []float64 `json:"Спуск"`
	DrillingGZD                     []float64 `json:"Бурение ГЗД"`
	MinWeightForHelicalBucklingRun  []float64 `json:"Мин. вес до спирального изгиба (спуск)"`
	MaxWeightBeforeYieldLimitPullUp []float64 `json:"Макс. вес до предела текучести (подъём)"`
}

// MomentFromMlModelResponse represents the response structure for the Moment prediction.
type MomentFromMLModelResponse struct {
	Depth          []float64 `json:"Глубина"`
	RotaryDrilling []float64 `json:"Бурение ротором"`
	PullUp         []float64 `json:"Подъём"`
	MakeUpTorque   []float64 `json:"Make-up Torque"`
	RunIn          []float64 `json:"Спуск"`
	TorqueOnMakeUp []float64 `json:"Момент свинчивания"`
}

// MinWeightFromMLModelResponse represents the response structure for the Min Weight on Bit prediction.
type MinWeightFromMLModelResponse struct {
	Depth                                             []float64 `json:"Глубина"`
	MinWeightOnBitForHelicalBucklingRotaryDrilling    []float64 `json:"Мин. вес на долоте до спирального изгиба (бурение ротором)"`
	MinWeightOnBitForSinusoidalBucklingGZDDrilling    []float64 `json:"Мин. вес на долоте до синусоидального изгиба (бурение ГЗД)"`
	MinWeightOnBitForSinusoidalBucklingRotaryDrilling []float64 `json:"Мин. вес на долоте до синусоидального изгиба (бурение ротором)"`
	MinWeightOnBitForHelicalBucklingGZDDrilling       []float64 `json:"Мин. вес на долоте до спирального изгиба (бурение ГЗД)"`
}
