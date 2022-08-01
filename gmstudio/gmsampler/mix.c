// wasm-compiler: https://wasdk.github.io/WasmFiddle/
typedef int int32;
typedef unsigned int uint32;
typedef long long int64;
typedef unsigned long long uint64;
#define SampleBits 32
#define DynamicBits 24

void mix(int32* buf, int32* src1, int32* src2, int sampleCount) {
  for (int i = 0; i < sampleCount; i++) {
    buf[i] = src1[i] + src2[i];
  }
}

void mixClamped(int32* buf, int32* src1, int32* src2, int sampleCount, int clampMax) {
  int64 min = (int64)-clampMax;
  int64 max = (int64)clampMax;
  for (int i = 0; i < sampleCount; i++) {
    int64 sample = (int64)src1[i] + (int64)src2[i];
    if (sample < min) {
      sample = min;
    } else if (sample > max) {
      sample = max;
    }
    buf[i] = (int)sample;
  }
}
