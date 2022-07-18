package gmgen

type GenBase struct {
	Ofs           uint32
	Incr          uint32
	Vol           int32
	RemainSamples int32
}

type GenInterface interface {
	NativeSamples(buf []int32, incr, ofs uint32) uint32
	GatBase() *GenBase

	StartTone(incr uint32, vol int32) bool
	StopTone(remainSamples int32) bool
	KillTone()

	FillSamplesLQ(buf []int32) int32
	FillSamplesMQ(buf []int32) int32
	FillSamplesHQ(buf []int32) int32
	FillSamplesLQLinear(buf []int32, endIncr uint32, endVol int32) int32
	FillSamplesMQLinear(buf []int32, endIncr uint32, endVol int32) int32
	FillSamplesHQLinear(buf []int32, endIncr uint32, endVol int32) int32

	MixSamplesLQ(buf []int32) int32
	MixSamplesMQ(buf []int32) int32
	MixSamplesHQ(buf []int32) int32
	MixSamplesLQLinear(buf []int32, endIncr uint32, endVol int32) int32
	MixSamplesMQLinear(buf []int32, endIncr uint32, endVol int32) int32
	MixSamplesHQLinear(buf []int32, endIncr uint32, endVol int32) int32
}
