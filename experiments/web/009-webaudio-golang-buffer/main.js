window.wg = {
    audioContext: null,
    workletInit: false,
    workletReady: false,
    workletWasm: null,
    workletNode: null
};

const version = Date.now();

const go = new Go();
WebAssembly.instantiateStreaming(fetch("main.wasm?" + version), go.importObject).then((result) => {
    go.run(result.instance);
});

fetch("worklet.wasm?" + version).then(r => r.arrayBuffer().then(buffer => {
    wg.workletWasm = new Uint8Array(buffer);
}));

function run(el) {
    el.innerHTML = "...";
    autoInitWorklet(() => {
        wg.workletNode.port.postMessage({t: "msg", val: "huhu"});
        wg.workletNode.port.postMessage({t: "wasm", val: wg.workletWasm});
        el.innerHTML = "ok.";
    });
}

function noise(active) {
    if (wg.workletNode && wg.workletNode.port) {
        wg.workletNode.port.postMessage({t: "noise", val: active});
    }
}

const initWorklet = async function () {
    wg.audioContext = new AudioContext({sampleRate: 44100, latencyHint: "interactive"} );
    await wg.audioContext.resume();
    await wg.audioContext.audioWorklet.addModule("worklet-module.js?" + version)
    wg.workletNode = new AudioWorkletNode(wg.audioContext, "worklet-module", {outputChannelCount: [2]});
    wg.workletNode.connect(wg.audioContext.destination)
}

function autoInitWorklet(okFunc) {
    if (wg.workletInit) {
        if (wg.workletReady) okFunc();
        return;
    }
    wg.workletInit = true;
    initWorklet().then(r => {
        console.log("main: worklet ok.")
        wg.workletReady = true;
        okFunc();
    }).catch(r => {
        console.log("main: worklet ERROR", r)
    });
}
