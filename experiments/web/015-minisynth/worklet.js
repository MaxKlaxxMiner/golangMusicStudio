console.log("worklet: init go wasm");
"use strict";
const wg = {
    fillBuffer: null,
    bufCounter: 0,
    workletPort: null,
    workletWat: null,
    workletWatMem: null,
    workletWatMemSamples: null,
    workletWatReady: false,
};
console.log("worklet: start processor");
function recMessage(event) {
    console.log("data:", event.data);
    switch (event.data.t) {
        case "watWasm": {
            const importObject = {};
            const module = new WebAssembly.Module(event.data.val);
            const instance = new WebAssembly.Instance(module, importObject);
            wg.workletWat = instance.exports;
            wg.workletWatMem = new Uint8Array(wg.workletWat.mem.buffer);
            wg.workletWatMemSamples = new Float32Array(wg.workletWat.mem.buffer);
            const version = wg.workletWat.version();
            if (version < 10001) {
                console.log("invalid worklet.wasm version: " + version);
                break;
            }
            console.log("worklet: run wat version: " + version);
            wg.workletWatReady = true;
            wg.workletPort.postMessage("ok: watWasmReady");
            break;
        }
        case "toneStart": {
            //const code = event.data.val;
            wg.fillBuffer = output => {
                //             const bufferLeft = new Uint8Array(output[0].buffer);
                //             const bufferRight = new Uint8Array(output[1].buffer);
                //             wg.workletGoFillBuffer(bufferLeft, bufferRight);
                output.forEach(channel => {
                    for (let i = 0; i < channel.length; i++) {
                        channel[i] = Math.random() * 2 - 1;
                        channel[i] *= 0.1;
                    }
                });
            };
            break;
        }
        case "toneEnd": {
            //const code = event.data.val;
            wg.fillBuffer = null;
            break;
        }
    }
}
class WatProcessor extends AudioWorkletProcessor {
    constructor() {
        super();
        this.port.onmessage = recMessage;
        wg.workletPort = this.port;
        wg.workletPort.postMessage("ok: start");
    }
    process(inputs, outputs, parameters) {
        const output = outputs[0];
        wg.bufCounter++;
        if (wg.bufCounter >= 16) {
            wg.bufCounter -= 16;
            wg.workletPort.postMessage("ok: block16");
        }
        if (wg.fillBuffer) {
            wg.fillBuffer(output);
        }
        else {
            output.forEach(channel => {
                for (let i = 0; i < channel.length; i++) {
                    channel[i] = 0;
                }
            });
        }
        return true;
    }
}
console.log("worklet: register");
registerProcessor("worklet", WatProcessor);
console.log("worklet: ok.");
