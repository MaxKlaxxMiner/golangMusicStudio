// noinspection DuplicatedCode

console.log("worklet: init go wasm");

"use strict";

globalThis.wg = {
    fillBuffer: null,
    bufCounter: 0,
    workletPort: null,
    // workletWat: null,
    // workletWatMem: null,
    // workletWatMemSamples: null,
    // workletWatReady: false,
};

console.log("worklet: start processor");

function recMessage(event) {
    console.log("data:", event.data);
    switch (event.data.t) {
        // case "watWasm": {
        //     const importObject = {};
        //     const module = new WebAssembly.Module(event.data.val);
        //     const instance = new WebAssembly.Instance(module, importObject);
        //     wg.workletWat = instance.exports;
        //     wg.workletWatMem = new Uint8Array(wg.workletWat.mem.buffer);
        //     wg.workletWatMemSamples = new Float32Array(wg.workletWat.mem.buffer);
        //     if (wg.workletWat.active() === 1) {
        //         wg.workletWatReady = true;
        //         wg.workletPort.postMessage("ok: watWasmReady");
        //     }
        //     console.log(wg.workletWatMemSamples);
        //     break;
        // }
        // case "tone": {
        //     if (event.data.val) {
        //         wg.fillBuffer = output => {
        //             output.forEach(channel => {
        //                 for (let i = 0; i < channel.length; i++) {
        //                     channel[i] = Math.random() * 2 - 1
        //                     channel[i] *= 0.1;
        //                 }
        //             });
        //         }
        //     } else {
        //         wg.fillBuffer = null;
        //     }
        //     break;
        // }
        // case "toneGo": {
        //     if (event.data.val) {
        //         wg.fillBuffer = output => {
        //             const bufferLeft = new Uint8Array(output[0].buffer);
        //             const bufferRight = new Uint8Array(output[1].buffer);
        //             wg.workletGoFillBuffer(bufferLeft, bufferRight);
        //         };
        //     } else {
        //         wg.fillBuffer = null;
        //     }
        //     break;
        // }
    }
}

// random-noise-processor.js
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
                    //channel[i] = i / 128 / 32;
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
