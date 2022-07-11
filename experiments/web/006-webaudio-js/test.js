function calc(el) {

    el.innerHTML = "ok.";
}

function run(el) {
    el.innerHTML = "...";
    autoInitWorklet(() => {
        console.log("start");
        setTimeout(function () {
            calc(el);
        }, 10);
    });
}

let audioContext = null;
let workletInit = false
let workletReady = false

const initWorklet = async function () {
    audioContext = new AudioContext();
    await audioContext.resume();
    await audioContext.audioWorklet.addModule("worklet-random.js")
    const randomNoiseNode = new AudioWorkletNode(audioContext, "worklet-random", {outputChannelCount: [2]});
    randomNoiseNode.connect(audioContext.destination)
}

function autoInitWorklet(okFunc) {
    if (workletInit) {
        if (workletReady) okFunc();
        return;
    }
    workletInit = true;
    initWorklet().then(r => {
        console.log("worklet: ok")
        workletReady = true;
        okFunc();
    }).catch(r => {
        console.log("worklet: ERROR", r)
    });
}
