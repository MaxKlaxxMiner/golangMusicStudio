console.log("worklet: init go wasm");

"use strict";

const wg = {
    fillBuffer: <(output: Float32Array[]) => void>null,
    bufCounter: 0,
    workletPort: <MessagePort>null,
    workletWat: <WorkletWat>null,
    workletWatMem: <Uint8Array>null,
    workletWatSamplesLeft: <Float32Array>null,
    workletWatSamplesRight: <Float32Array>null,
    workletWatReady: false,
};

const watMinVersion = 10004;

declare class WorkletWat {
    // --- 10004 ---
    // mix(buf *int, src1 *int, src2 *int, sampleCount uint)
    mix(bufPtr: number, src1Ptr: number, src2Ptr: number, sampleCount: number): void

    // --- 10003 ---
    // volumeUpdate(buf *int, sampleCount uint, vol int)
    volumeUpdate(bufPtr: number, sampleCount: number, vol: number): void

    // --- 10002 ---
    // noise(buf *int, sampleCount uint, incr uint, ofs uint) uint
    noise(bufPtr: number, sampleCount: number, incr: number, ofs: number): number

    // squareLQ(buf *int, sampleCount uint, incr uint, ofs uint) uint
    squareLQ(bufPtr: number, sampleCount: number, incr: number, ofs: number): number

    // squareHQ(buf *int, sampleCount uint, incr uint, ofs uint) uint
    squareHQ(bufPtr: number, sampleCount: number, incr: number, ofs: number): number

    // convertIntSamplesToFloat32(floats *float, ints *int, sampleCount uint)
    convertIntSamplesToFloat32(floatsPtr: number, intsPtr: number, sampleCount: number): void

    // --- 10001 ---
    version(): number

    mem: WebAssembly.Memory
}

abstract class Tone {
    abstract fill(ptrBufL: number, ptrBufR: number, sampleCount: number);
}

class Tone1 {
    midiCode: number;
    ofsL: number;
    ofsR: number;
    incrL: number;
    incrR: number;
    sampler: (bufPtr: number, sampleCount: number, incr: number, ofs: number) => number;
    volume: number;

    constructor(midiCode: number, volume: number, hq: boolean) {
        this.midiCode = midiCode;
        this.ofsL = 0;
        this.ofsR = 0;
        this.incrL = (4294967296 / (44100 / (440 * Math.pow(2, (midiCode - 69 + 0.01) / 12))) + 0.5) | 0;
        this.incrR = (4294967296 / (44100 / (440 * Math.pow(2, (midiCode - 69 - 0.01) / 12))) + 0.5) | 0;
        this.sampler = hq ? wg.workletWat.squareHQ : wg.workletWat.squareLQ;
        this.volume = volume > 2 ? volume | 0 : (8388608 * volume) | 0;
    }

    fill(ptrBufL: number, ptrBufR: number, sampleCount: number) {
        this.ofsL = this.sampler(ptrBufL, sampleCount, this.incrL, this.ofsL);
        this.ofsR = this.sampler(ptrBufR, sampleCount, this.incrR, this.ofsR);
        wg.workletWat.volumeUpdate(ptrBufL, sampleCount, this.volume);
        wg.workletWat.volumeUpdate(ptrBufR, sampleCount, this.volume);
    }
}

class Tone3 {
    midiCode: number;
    ofsL1: number;
    ofsR1: number;
    ofsL2: number;
    ofsR2: number;
    ofsL3: number;
    ofsR3: number;
    incrL1: number;
    incrR1: number;
    incrL2: number;
    incrR2: number;
    incrL3: number;
    incrR3: number;
    sampler: (bufPtr: number, sampleCount: number, incr: number, ofs: number) => number;
    volume: number;

    constructor(midiCode: number, volume: number, hq: boolean) {
        this.midiCode = midiCode;
        this.ofsL1 = 0;
        this.ofsR1 = 0;
        this.ofsL2 = 0;
        this.ofsR2 = 0;
        this.ofsL3 = 0;
        this.ofsR3 = 0;
        this.incrL1 = (4294967296 / (44100 / (440 * Math.pow(2, (midiCode - 69 + 0.01) / 12))) + 0.5) | 0;
        this.incrR1 = (4294967296 / (44100 / (440 * Math.pow(2, (midiCode - 69 - 0.01) / 12))) + 0.5) | 0;
        this.incrL2 = (4294967296 / (44100 / (440 * Math.pow(2, (midiCode - 69 + 11.95 + 0.01) / 12))) + 0.5) | 0;
        this.incrR2 = (4294967296 / (44100 / (440 * Math.pow(2, (midiCode - 69 + 11.95 - 0.01) / 12))) + 0.5) | 0;
        this.incrL3 = (4294967296 / (44100 / (440 * Math.pow(2, (midiCode - 69 + 23.90 + 0.01) / 12))) + 0.5) | 0;
        this.incrR3 = (4294967296 / (44100 / (440 * Math.pow(2, (midiCode - 69 + 23.90 - 0.01) / 12))) + 0.5) | 0;
        this.sampler = hq ? wg.workletWat.squareHQ : wg.workletWat.squareLQ;
        this.volume = volume > 2 ? volume * 0.7 | 0 : (8388608 * volume * 0.7) | 0;
    }

    fill(ptrBufL: number, ptrBufR: number, sampleCount: number) {
        this.ofsL1 = this.sampler(ptrBufL, sampleCount, this.incrL1, this.ofsL1);
        this.ofsR1 = this.sampler(ptrBufR, sampleCount, this.incrR1, this.ofsR1);

        this.ofsL2 = this.sampler(ptrInternalTmpL, sampleCount, this.incrL2, this.ofsL2);
        this.ofsR2 = this.sampler(ptrInternalTmpR, sampleCount, this.incrR2, this.ofsR2);
        wg.workletWat.mix(ptrBufL, ptrBufL, ptrInternalTmpL, sampleCount);
        wg.workletWat.mix(ptrBufR, ptrBufR, ptrInternalTmpL, sampleCount);

        this.ofsL3 = this.sampler(ptrInternalTmpL, sampleCount, this.incrL3, this.ofsL3);
        this.ofsR3 = this.sampler(ptrInternalTmpR, sampleCount, this.incrR3, this.ofsR3);
        wg.workletWat.mix(ptrBufL, ptrBufL, ptrInternalTmpL, sampleCount);
        wg.workletWat.mix(ptrBufR, ptrBufR, ptrInternalTmpR, sampleCount);

        wg.workletWat.volumeUpdate(ptrBufL, sampleCount, this.volume);
        wg.workletWat.volumeUpdate(ptrBufR, sampleCount, this.volume);
    }
}

console.log("worklet: start processor");

const sampleCount = 128;
const ptrOutputFloatL = 0;
const ptrOutputFloatR = ptrOutputFloatL + 512;
const ptrMixL = ptrOutputFloatR + 512;
const ptrMixR = ptrMixL + 512;
const ptrTmpL = ptrMixR + 512;
const ptrTmpR = ptrTmpL + 512;
const ptrInternalTmpL = ptrTmpR + 512;
const ptrInternalTmpR = ptrInternalTmpL + 512;
//const noiseVolume = (8388608 * 0.02) | 0; // 2% noise background
const noiseVolume = (8388608 * 0.00) | 0; // 0% noise background
let noiseOfs = 0;
let currentTones: { [key: number]: Tone } = {};

function watFillBuffer(output: Float32Array[]) {
    const wat = wg.workletWat;

    noiseOfs = wat.noise(ptrMixL, sampleCount, 123456789, noiseOfs);
    noiseOfs = wat.noise(ptrMixR, sampleCount, 123456789, noiseOfs);
    wat.volumeUpdate(ptrMixL, sampleCount, noiseVolume); // volume update - left side
    wat.volumeUpdate(ptrMixR, sampleCount, noiseVolume); // volume update - right side

    for (const key in currentTones) {
        const tone = currentTones[key];
        tone.fill(ptrTmpL, ptrTmpR, sampleCount);
        wat.mix(ptrMixL, ptrMixL, ptrTmpL, sampleCount);
        wat.mix(ptrMixR, ptrMixR, ptrTmpR, sampleCount);
    }

    wat.convertIntSamplesToFloat32(ptrOutputFloatL, ptrMixL, sampleCount); // convert int -> float32 - left side
    wat.convertIntSamplesToFloat32(ptrOutputFloatR, ptrMixR, sampleCount); // convert int -> float32 - right side

    output[0].set(wg.workletWatSamplesLeft);        // copy left-samples to output-buffer
    output[1].set(wg.workletWatSamplesRight);       // copy right-samples to output-buffer
}

function recMessage(event) {
    console.log("data:", event.data);
    switch (event.data.t) {
        case "watWasm": {
            const importObject = {};
            const module = new WebAssembly.Module(event.data.val);
            const instance = new WebAssembly.Instance(module, importObject);
            wg.workletWat = <any>instance.exports;
            wg.workletWatMem = new Uint8Array(wg.workletWat.mem.buffer);
            wg.workletWatSamplesLeft = new Float32Array(wg.workletWat.mem.buffer, ptrOutputFloatL, sampleCount);
            wg.workletWatSamplesRight = new Float32Array(wg.workletWat.mem.buffer, ptrOutputFloatR, sampleCount);
            const version = wg.workletWat.version();
            if (version < watMinVersion) {
                console.log("invalid worklet.wasm version: " + version);
                break;
            }
            console.log("worklet: run wat version: " + version);
            wg.workletWatReady = true;
            wg.fillBuffer = watFillBuffer;
            wg.workletPort.postMessage("ok: watWasmReady " + version);
            break;
        }
        case "toneStart": {
            const code = event.data.val;
            if (!currentTones[code]) {
                if (event.data.osc3) {
                    currentTones[code] = new Tone3(code, 0.1, event.data.hq);
                } else {
                    currentTones[code] = new Tone1(code, 0.1, event.data.hq);
                }
            }
            break;
        }
        case "toneEnd": {
            const code = event.data.val;
            delete currentTones[code];
            break;
        }
        case "toneKill": {
            currentTones = {};
            break;
        }
    }
}

declare abstract class AudioWorkletProcessor {
    port: MessagePort;

    // noinspection JSUnusedGlobalSymbols
    abstract process(inputs, outputs, parameters): boolean;
}

declare function registerProcessor(name: string, processorCtor: any);

class WatProcessor extends AudioWorkletProcessor {
    constructor() {
        super();
        this.port.onmessage = recMessage;
        wg.workletPort = this.port;
        wg.workletPort.postMessage("ok: start");
    }

    process(inputs, outputs, parameters) {
        const output = outputs[0]
        wg.bufCounter++;
        if (wg.bufCounter >= 16) {
            wg.bufCounter -= 16;
            wg.workletPort.postMessage("ok: block16");
        }
        if (wg.fillBuffer) {
            wg.fillBuffer(output);
        } else {
            output.forEach(channel => {
                for (let i = 0; i < channel.length; i++) {
                    channel[i] = 0;
                }
            })
        }
        return true
    }
}

console.log("worklet: register");

registerProcessor("worklet", WatProcessor);

console.log("worklet: ok.");
