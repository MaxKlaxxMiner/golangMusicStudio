package gmgen

func (g *GenBase) StartTone(incr uint32, vol int32) bool                               { panic("todo") }
func (g *GenBase) StopTone(remainSamples int32) bool                                   { panic("todo") }
func (g *GenBase) KillTone()                                                           { panic("todo") }
func (g *GenBase) FillSamplesLQ(buf []int32) int32                                     { panic("todo") }
func (g *GenBase) FillSamplesMQ(buf []int32) int32                                     { panic("todo") }
func (g *GenBase) FillSamplesHQ(buf []int32) int32                                     { panic("todo") }
func (g *GenBase) FillSamplesLQLinear(buf []int32, endIncr uint32, endVol int32) int32 { panic("todo") }
func (g *GenBase) FillSamplesMQLinear(buf []int32, endIncr uint32, endVol int32) int32 { panic("todo") }
func (g *GenBase) FillSamplesHQLinear(buf []int32, endIncr uint32, endVol int32) int32 { panic("todo") }
func (g *GenBase) MixSamplesLQ(buf []int32) int32                                      { panic("todo") }
func (g *GenBase) MixSamplesMQ(buf []int32) int32                                      { panic("todo") }
func (g *GenBase) MixSamplesHQ(buf []int32) int32                                      { panic("todo") }
func (g *GenBase) MixSamplesLQLinear(buf []int32, endIncr uint32, endVol int32) int32  { panic("todo") }
func (g *GenBase) MixSamplesMQLinear(buf []int32, endIncr uint32, endVol int32) int32  { panic("todo") }
func (g *GenBase) MixSamplesHQLinear(buf []int32, endIncr uint32, endVol int32) int32  { panic("todo") }
