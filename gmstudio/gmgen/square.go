package gmgen

import "github.com/MaxKlaxxMiner/golangMusicStudio/gmstudio/gmsampler"

type GenSquare struct {
	GenBase
}

func (g *GenSquare) NativeSamples(buf []int32, incr, ofs uint32) uint32 {
	return gmsampler.SquareHQ(buf, incr, ofs)
}
func (g *GenSquare) GetBase() *GenBase { return &g.GenBase }

func (g *GenSquare) StartTone(incr uint32, vol int32) bool                               { return false }
func (g *GenSquare) StopTone(remainSamples int32) bool                                   { return false }
func (g *GenSquare) KillTone()                                                           {}
func (g *GenSquare) FillSamplesLQ(buf []int32) int32                                     { return -1 }
func (g *GenSquare) FillSamplesMQ(buf []int32) int32                                     { return -1 }
func (g *GenSquare) FillSamplesHQ(buf []int32) int32                                     { return -1 }
func (g *GenSquare) FillSamplesLQLinear(buf []int32, endIncr uint32, endVol int32) int32 { return -1 }
func (g *GenSquare) FillSamplesMQLinear(buf []int32, endIncr uint32, endVol int32) int32 { return -1 }
func (g *GenSquare) FillSamplesHQLinear(buf []int32, endIncr uint32, endVol int32) int32 { return -1 }
func (g *GenSquare) MixSamplesLQ(buf []int32) int32                                      { return -1 }
func (g *GenSquare) MixSamplesMQ(buf []int32) int32                                      { return -1 }
func (g *GenSquare) MixSamplesHQ(buf []int32) int32                                      { return -1 }
func (g *GenSquare) MixSamplesLQLinear(buf []int32, endIncr uint32, endVol int32) int32  { return -1 }
func (g *GenSquare) MixSamplesMQLinear(buf []int32, endIncr uint32, endVol int32) int32  { return -1 }
func (g *GenSquare) MixSamplesHQLinear(buf []int32, endIncr uint32, endVol int32) int32  { return -1 }
