package models

type PenerjemahReqType struct {
	Text string
}

type PenerjemahResType struct {
	RawText         string          `json:"raw_text"`
	PreproccessText string          `json:"preproccess_text"`
	Summary         string          `json:"summary"`
	BestScore       float64         `json:"best_score"`
	TreshHold       float64         `json:"tresh_hold"`
	TotAllkata      int             `json:"totallkata"`
	Calculate       []CalculateType `json:"calculate"`
}

type CalculateType struct {
	Kalimat          string               `json:"kalimat"`
	Totkata          int                  `json:"totkata"`
	ArrKata          []string             `json:"arr_kata"`
	TotKataDalamText TotKataDalamTextType `json:"tot_kata_dalam_text"`
	TFWi             TFWiMapType
	IDFWi            IDFWiMapType
	WWi              WWiMapType
	TWWi             float64 // total WWi
	NFS              float64 // total NFS
	WS               float64
}

type TotKataDalamTextType map[string]int
type TFWiMapType map[string]float64
type IDFWiMapType map[string]float64
type WWiMapType map[string]float64
type WSMapType map[string]float64
