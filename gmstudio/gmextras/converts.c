// wasm-compiler: https://wasdk.github.io/WasmFiddle/
typedef int int32;
typedef unsigned int uint32;
typedef long long int64;
typedef unsigned long long uint64;
#define SampleBits 32
#define DynamicBits 24

void convertIntSamplesToFloat32(float* dst, int32* src, int sampleCount) {
  for (int i = 0; i < sampleCount; i++) {
    dst[i] = (float)src[i] / (float)(1 << (DynamicBits - 1));
  }
}
