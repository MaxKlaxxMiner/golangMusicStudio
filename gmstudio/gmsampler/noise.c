// wasm-compiler: https://wasdk.github.io/WasmFiddle/
typedef int int32;
typedef unsigned int uint32;
typedef long long int64;
typedef unsigned long long uint64;
#define SampleBits 32
#define DynamicBits 24

uint32 noise(int32* buf, int sampleCount, uint32 incr, uint32 ofs) {
  for (int i = 0; i < sampleCount; i++) {
    buf[i] = (int32)(ofs << (SampleBits - DynamicBits - 4)) >> (SampleBits - DynamicBits);
    ofs = (ofs + incr) * 214013 + 2531011;
  }
  return ofs;
}
