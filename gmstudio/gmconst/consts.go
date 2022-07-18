package gmconst

const SampleRate = 44100 // default Samplerate (44100 = CD, 48000 = DVD)
const SampleBits = 32    // 32 bits per Sample = int32
const DynamicBits = 24   // 24 bits dynamic range (from 32, reserve: 8 bits for safe overdrive)
const AaBitsMQ = 2       // 2 bits for fast anti anti aliasing
const AaBitsHQ = 16      // 16 bits for high quality anti aliasing = 65536 subsamples

const Volume100 int32 = 1 << (DynamicBits - 1)
const VolumeLimit int32 = 1<<(SampleBits-1) - 1
