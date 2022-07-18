package gmgen

import "github.com/MaxKlaxxMiner/golangMusicStudio/gmstudio/gmsampler"

type GenNoise struct {
	GenBase
}

func (g *GenNoise) NativeSamples(buf []int32, incr, ofs uint32) uint32 {
	return gmsampler.Noise(buf, incr, ofs)
}
func (g *GenNoise) GetBase() *GenBase { return &g.GenBase }

func (g *GenNoise) StartTone(incr uint32, vol int32) bool                               { return false }
func (g *GenNoise) StopTone(remainSamples int32) bool                                   { return false }
func (g *GenNoise) KillTone()                                                           {}
func (g *GenNoise) FillSamplesLQ(buf []int32) int32                                     { return -1 }
func (g *GenNoise) FillSamplesMQ(buf []int32) int32                                     { return -1 }
func (g *GenNoise) FillSamplesHQ(buf []int32) int32                                     { return -1 }
func (g *GenNoise) FillSamplesLQLinear(buf []int32, endIncr uint32, endVol int32) int32 { return -1 }
func (g *GenNoise) FillSamplesMQLinear(buf []int32, endIncr uint32, endVol int32) int32 { return -1 }
func (g *GenNoise) FillSamplesHQLinear(buf []int32, endIncr uint32, endVol int32) int32 { return -1 }
func (g *GenNoise) MixSamplesLQ(buf []int32) int32                                      { return -1 }
func (g *GenNoise) MixSamplesMQ(buf []int32) int32                                      { return -1 }
func (g *GenNoise) MixSamplesHQ(buf []int32) int32                                      { return -1 }
func (g *GenNoise) MixSamplesLQLinear(buf []int32, endIncr uint32, endVol int32) int32  { return -1 }
func (g *GenNoise) MixSamplesMQLinear(buf []int32, endIncr uint32, endVol int32) int32  { return -1 }
func (g *GenNoise) MixSamplesHQLinear(buf []int32, endIncr uint32, endVol int32) int32  { return -1 }
