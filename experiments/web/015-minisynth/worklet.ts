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

const watMinVersion = 10002;

declare class WorkletWat {
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

console.log("worklet: start processor");

function recMessage(event) {
    console.log("data:", event.data);
    switch (event.data.t) {
        case "watWasm": {
            const importObject = {};
            const module = new WebAssembly.Module(event.data.val);
            const instance = new WebAssembly.Instance(module, importObject);
            wg.workletWat = <any>instance.exports;
            wg.workletWatMem = new Uint8Array(wg.workletWat.mem.buffer);
            wg.workletWatSamplesLeft = new Float32Array(wg.workletWat.mem.buffer, 0, 128);
            wg.workletWatSamplesRight = new Float32Array(wg.workletWat.mem.buffer, 512, 128);
            const version = wg.workletWat.version();
            if (version < watMinVersion) {
                console.log("invalid worklet.wasm version: " + version);
                break;
            }
            console.log("worklet: run wat version: " + version);
            wg.workletWatReady = true;
            wg.workletPort.postMessage("ok: watWasmReady");
            break;
        }
        case "toneStart": {
            const code = event.data.val;
            if (wg.workletWatReady) {
                let ofs = 0;
                const wat = wg.workletWat;
                wg.fillBuffer = output => {
                    wat.noise(1024, 128, 123456789, ofs + code - 60); // C4 = 60 = mono noise
                    ofs = wat.noise(1536, 128, 123456789, ofs);

                    wat.convertIntSamplesToFloat32(0, 1024, 128);   // convert int -> float32 - left side
                    wat.convertIntSamplesToFloat32(512, 1536, 128); // convert int -> float32 - right side

                    output[0].set(wg.workletWatSamplesLeft);        // copy left-samples to output-buffer
                    output[1].set(wg.workletWatSamplesRight);       // copy right-samples to output-buffer

                    output.forEach(channel => {
                        for (let i = 0; i < channel.length; i++) {
                            channel[i] *= 0.1; // volume = 10%
                        }
                    });
                };
            } else {
                wg.fillBuffer = output => {
                    output.forEach(channel => {
                        for (let i = 0; i < channel.length; i++) {
                            channel[i] = Math.random() * 2 - 1
                            channel[i] *= 0.1;
                        }
                    });
                };
            }
            break;
        }
        case "toneEnd": {
            //const code = event.data.val;
            wg.fillBuffer = null;
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
