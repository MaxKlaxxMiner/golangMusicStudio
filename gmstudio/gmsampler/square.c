// wasm-compiler: https://wasdk.github.io/WasmFiddle/
typedef int int32;
typedef unsigned int uint32;
typedef long long int64;
typedef unsigned long long uint64;
#define SampleBits 32
#define DynamicBits 24

uint32 squareLQ(int32* buf, int sampleCount, uint32 incr, uint32 ofs) {
  for (int i = 0; i < sampleCount; i++) {
    buf[i] = ((int32)1 << (DynamicBits - 1)) - (int32)(ofs >> (SampleBits - 1) << DynamicBits);
    ofs += incr;
  }
  return ofs;
}

uint32 squareHQ(int32* buf, int sampleCount, uint32 incr, uint32 ofs) {
  int32 curVal = (int32)(((int64)1 << (DynamicBits - 1)) - (int64)(ofs >> (SampleBits - 1) << DynamicBits));
  uint64 incr1 = ((uint64)1 << (DynamicBits + DynamicBits)) / (uint64)incr;
  for (int i = 0; i < sampleCount; i++) {
    if ((ofs >> (SampleBits - 1)) == ((ofs + incr) >> (SampleBits - 1))) {
      buf[i] = curVal;
      ofs += incr;
      continue;
    }
    uint32 ofsA2 = (ofs + incr) & (uint32)(((uint64)1 << SampleBits) - 1);
    uint32 f = ((uint32)1 << (SampleBits - 1)) - (ofs & (((uint32)1 << (SampleBits - 1)) - 1));
    curVal = (int32)(((uint64)f * incr1) >> DynamicBits);
    if ((ofs >> (SampleBits - 1)) != 0) {
      curVal = ((int32)1 << (DynamicBits - 1)) - curVal;
    } else {
      curVal = curVal - ((int32)1 << (DynamicBits - 1));
    }

    buf[i] = curVal;
    ofs = ofsA2;
    curVal = (int32)(((int64)1 << (DynamicBits - 1)) - (int64)(ofs >> (SampleBits - 1) << DynamicBits));
  }
  return ofs;
}

uint32 squareHQ2(int32* buf, int sampleCount, uint32 incr, uint32 ofs) {
  int i;
  int32 curVal = (int32)8388608 - (int32)(ofs >> 31 << 24);
  uint64 incr1 = (uint64)281474976710656 / (uint64)incr;
  for (int i = 0; i < sampleCount; i++) {
    if ((ofs>>31) == ((ofs+incr)>>31)) {
      buf[i] = curVal;
      ofs += incr;
    } else {
      curVal = (int32)(((uint64)((uint32)2147483648 - (ofs & 2147483647)) * incr1) >> 24);
      if (ofs >> 31) {
        curVal = 8388608 - curVal;
      } else {
        curVal = curVal - 8388608;
      }
      buf[i] = curVal;
      ofs += incr;
      curVal = (int32)8388608 - (int32)(ofs >> 31 << 24);
    }
  }
  return ofs;
}
