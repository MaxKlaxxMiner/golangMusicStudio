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
    workletLoaded: false,
    workletStarted: false,
    workletBlockCount16: 0,
    workletReady: false,
    workletMessageTodo: [],

    workletGoInit: false,
    workletGoError: "",
    workletGoWasm: null,
    workletGoReady: false,
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

function initWorkletGo() {
    wg.workletGoInit = true;

    const loadWasm = "worklet.wasm?" + version;
    fetch(loadWasm).then(r => r.arrayBuffer().then(buffer => {
        if (r.status !== 200) {
            wg.workletGoError = loadWasm + " - " + r.statusText;
        }
        wg.workletGoWasm = new Uint8Array(buffer);

        workletSendMessage({t: "goWasm", val: wg.workletGoWasm})
    }).catch(r => {
        wg.workletGoError = loadWasm + " - " + r.toString();
    }));
}

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
    }
}

function toneWorkletJs(active) {
    workletSendMessage({t: "tone", val: active});
}

function toneWorkletGo(active) {
    workletSendMessage({t: "toneGo", val: active});
}

function updateTab(id, status, value) {
    const el = document.getElementById(id)
    if (el && (el.className !== status || el.innerText !== value)) {
        el.className = status;
        el.innerText = value;
    }
}

function updateTabMainGo() {
    if (wg.mainGoError) {
        updateTab("status_main.go", "error", "error: " + wg.mainGoError);
        return;
    }
    if (!wg.mainGoInit) {
        return;
    }
    if (wg.mainGoReady) {
        updateTab("status_main.go", "ok", "ok.");
    } else {
        updateTab("status_main.go", "", "load wasm...");
    }
}

function updateTabAudioContext() {
    if (!wg.userInput) {
        updateTab("status_audioContext", "info", "wait for user input");
        return;
    }
    if (wg.audioError) {
        updateTab("status_audioContext", "error", "error: " + wg.audioError);
        return;
    }
    if (!wg.audioInit) {
        updateTab("status_audioContext", "error", "unkown");
        return;
    }
    if (wg.audioContext == null) {
        updateTab("status_audioContext", "error", "not initialized...");
        return;
    }

    if (wg.audioContext.state === "running") {
        updateTab("status_audioContext", "ok", "ok. (running " + wg.audioContext.currentTime.toFixed(1) + "s)");
        return;
    }

    if (wg.audioContext.state === "suspended") {
        updateTab("status_audioContext", "", "suspended (wait for user input?)");
        return;
    }

    updateTab("status_audioContext", "error", "unexpected state: " + wg.audioContext.state);
}

function updateTabAudioWorklet() {
    if (wg.workletError) {
        updateTab("status_worklet", "error", "error: " + wg.workletError);
        return;
    }

    if (!wg.workletInit) {
        updateTab("status_worklet", "info", "wait for audioContext");
        return;
    }

    if (!wg.workletLoaded) {
        updateTab("status_worklet", "", "load module...");
        return;
    }

    if (!wg.workletStarted) {
        updateTab("status_worklet", "", "start module...");
        return;
    }

    updateTab("status_worklet", "ok", "ok. (running " + (wg.workletBlockCount16 * 16 * 128 / 44100).toFixed(1) + "s = " + wg.workletBlockCount16 + ")");
}

function updateTabWorkletGo() {
    let exInfo = "";
    if (wg.workletGoWasm != null) {
        exInfo = " (wasm: " + (wg.workletGoWasm.length / 1024.0).toFixed(2) + " kByte loaded)";
    }

    if (wg.workletGoError) {
        updateTab("status_workletGo", "error", "error: " + wg.workletGoError);
        return;
    }

    if (!wg.workletGoInit) {
        updateTab("status_workletGo", "", "not initialized");
        return;
    }

    if (!wg.workletInit) {
        updateTab("status_workletGo", "info", "wait for audioWorklet" + exInfo);
        return;
    }

    if (wg.workletGoWasm == null) {
        updateTab("status_workletGo", "", "load wasm...");
        return;
    }

    if (wg.workletGoReady) {
        updateTab("status_workletGo", "ok", "ok." + exInfo);
    } else {
        updateTab("status_workletGo", "", "not ready..." + exInfo);
    }

}

setInterval(() => {
    updateTab("status_main.js", "ok", "ok.");
    updateTabMainGo();
    updateTabAudioContext();
    updateTabAudioWorklet();
    updateTabWorkletGo();
}, 10);

initUserEvents();
initMainGo();
initWorkletGo();