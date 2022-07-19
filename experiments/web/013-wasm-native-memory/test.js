// noinspection JSUnresolvedFunction

const myRequest = new Request("main.wasm");

let wasm = null;
let mem8 = null;
let mem32 = null;

const myHeaders = new Headers();
myHeaders.append('pragma', 'no-cache');
myHeaders.append('cache-control', 'no-cache');
const myInit = {method: 'GET', headers: myHeaders};
WebAssembly.instantiateStreaming(fetch(myRequest, myInit)).then(obj => {
    wasm = obj.instance.exports;
    mem8 = new Uint8Array(wasm.mem.buffer);
    mem32 = new Uint32Array(wasm.mem.buffer);
});

function calc(el) {
    let tim = Date.now();
    mem32[0] = 1000;
    mem32[1] = 999;
    const val = wasm.calc();
    tim = Date.now() - tim;

    el.innerHTML = "wasm-ok: " + val + "<br>(" + tim + " ms)";
}

function run(el) {
    el.innerHTML = "...";
    console.log("start");
    setTimeout(function () {
        calc(el);
    }, 10);
}
