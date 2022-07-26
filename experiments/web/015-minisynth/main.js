window.wg = {
    userInput: false,

    mainGoInit: false,
    mainGoError: "",
    mainGoReady: false,

    audioInit: false,
    audioError: "",
    audioContext: null,

    workletInit: false,
    workletError: "",
    workletNode: null,
    workletLoaded: false,
    workletStarted: false,
    workletBlockCount16: 0,
    workletReady: false,
    workletMessageTodo: [],
}

const version = Date.now();

function initUserEvents() {
    onclick = (e) => {
        wg.userInput = true;
        initAudio();
    }
    onmousedown = (e) => {
        wg.userInput = true;
        initAudio();
    }
    onkeydown = (e) => {
        switch (e.key) {
            case "Alt":
            case "Control":
            case "Shift":
            case "CapsLock":
            case "Escape":
            case "ScrollLock":
            case "NumLock":
                return
        }
        wg.userInput = true;
        initAudio();
    }
}

function initMainGo() {
    try {
        if (typeof Go !== "function") {
            throw "Go-Handler not found: maybe fail wasm_exec.js?";
        }
        const go = new Go();
        const loadWasm = "main.wasm?" + version;
        wg.mainGoInit = true;
        WebAssembly.instantiateStreaming(fetch(loadWasm), go.importObject).then(r => {
            go.run(r.instance).catch(r => {
                wg.mainGoError = loadWasm + " - " + r.toString();
            });
        }).catch(r => {
            wg.mainGoError = loadWasm + " - " + r.toString();
        });
    } catch (e) {
        wg.mainGoError = e.toString()
    }
}

function initAudio() {
    if (wg.audioInit) return;
    if (!wg.userInput) return;
    wg.audioInit = true;

    try {
        if (wg.audioContext == null) {
            wg.audioContext = new AudioContext({sampleRate: 44100, latencyHint: "interactive"});
            const interval = setInterval(() => {
                if (wg.audioContext.state === "running") {
                    clearInterval(interval);
                    if (!wg.audioInit) {
                        wg.userInput = true;
                        initAudio();
                    }
                }
            }, 10);
        } else {
            wg.audioContext.resume();
        }
    } catch (e) {
        wg.audioError = e.toString();
        return
    }

    if (wg.audioContext.state === "suspended") {
        console.log("no user gesture");
        wg.audioInit = false;
        wg.userInput = false;
        return;
    }

    wg.workletInit = true;
    const workletUrl = "worklet.js?" + version;
    wg.audioContext.audioWorklet.addModule(workletUrl).then(r => {
        wg.workletNode = new AudioWorkletNode(wg.audioContext, "worklet", {outputChannelCount: [2]});
        wg.workletNode.port.onmessage = workletReceiveMessage;
        wg.workletNode.connect(wg.audioContext.destination);
        wg.workletLoaded = true;
    }).catch(r => {
        wg.workletError = workletUrl + " - " + r.toString();
    })
}

// function initWorkletGo() {
//     wg.workletGoInit = true;
//
//     const loadWasm = "worklet.wasm?" + version;
//     fetch(loadWasm).then(r => r.arrayBuffer().then(buffer => {
//         if (r.status !== 200) {
//             wg.workletGoError = loadWasm + " - " + r.statusText;
//         }
//         wg.workletGoWasm = new Uint8Array(buffer);
//
//         workletSendMessage({t: "goWasm", val: wg.workletGoWasm})
//     }).catch(r => {
//         wg.workletGoError = loadWasm + " - " + r.toString();
//     }));
// }

// function initWorkletWat() {
//     wg.workletWatInit = true;
//
//     const loadWasm = "wat.wasm?" + version;
//     fetch(loadWasm).then(r => r.arrayBuffer().then(buffer => {
//         if (r.status !== 200) {
//             wg.workletWatError = loadWasm + " - " + r.statusText;
//         }
//         wg.workletWatWasm = new Uint8Array(buffer);
//
//         workletSendMessage({t: "watWasm", val: wg.workletWatWasm})
//
//         const importObject = {};
//         const module = new WebAssembly.Module(wg.workletWatWasm);
//         const instance = new WebAssembly.Instance(module, importObject);
//         wg.workletWat = instance.exports;
//         wg.workletWatMem = new Uint8Array(wg.workletWat.mem.buffer);
//         wg.workletWatMemSamples = new Float32Array(wg.workletWat.mem.buffer);
//         wg.workletWatMemInts = new Int32Array(wg.workletWat.mem.buffer);
//     }).catch(r => {
//         wg.workletWatError = loadWasm + " - " + r.toString();
//     }));
// }

function workletSendMessage(msg) {
    if (!wg.workletReady) {
        if (msg != null) wg.workletMessageTodo.push(msg);
        return;
    }
    if (wg.workletMessageTodo.length > 0) {
        for (let i = 0; i < wg.workletMessageTodo.length; i++) {
            wg.workletNode.port.postMessage(wg.workletMessageTodo[i]);
        }
        wg.workletMessageTodo = [];
    }
    if (msg != null) wg.workletNode.port.postMessage(msg);
}

function workletReceiveMessage(msg) {
    switch (msg.data) {
        case "ok: start": {
            wg.workletStarted = true;
            break;
        }
        case "ok: block16": {
            wg.workletBlockCount16++;
            wg.workletReady = true;
            workletSendMessage(null);
            break;
        }
        case "ok: goWasmReady": {
            wg.workletGoReady = true;
            break;
        }
        case "ok: watWasmReady": {
            wg.workletWatReady = true;
            break;
        }
    }
}

// function toneWorkletJs(active) {
//     workletSendMessage({t: "tone", val: active});
// }
//
// function toneWorkletGo(active) {
//     workletSendMessage({t: "toneGo", val: active});
// }

function toneStart(midiCode) {
}

function toneStop(midiCode) {
}

window.addEventListener('load', (event) => {
    initUserEvents();
    initMainGo();
    // initWorkletGo();
    // initWorkletWat();
});
