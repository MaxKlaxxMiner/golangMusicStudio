// wasm-compiler: https://wasdk.github.io/WasmFiddle/
typedef int int32;
typedef unsigned int uint32;
typedef long long int64;
typedef unsigned long long uint64;
#define SampleBits 32
#define DynamicBits 24

void volumeUpdate(int32* buf, int sampleCount, int32 vol) {
  for (int i = 0; i < sampleCount; i++) {
    buf[i] = (int32)(((int64)buf[i] * (int64)vol) >> (DynamicBits - 1));
  }
}

void volumeUpdateClamped(int32* buf, int sampleCount, int32 vol, int32 clampMax) {
  int64 min = (int64)-clampMax;
  int64 max = (int64)clampMax;
  for (int i = 0; i < sampleCount; i++) {
    int64 sample = ((int64)buf[i] * (int64)vol) >> (DynamicBits - 1);
    if (sample < min) {
      sample = min;
    } else if (sample > max) {
      sample = max;
    }
    buf[i] = (int32)sample;
  }
}
