window.wg = {
    calc: () => {
        console.log("error: no wasm-override")
    }
};
const go = new Go();
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
    go.run(result.instance);
});

function calc(el) {
    let tim = Date.now();
    wg.calc();
    tim = Date.now() - tim;
    el.innerHTML = "wasm-ok. (" + tim + " ms)";

    // wasm Chrome 103:  702 ms (js: 244 ms)
    // wasm Firefox 102: 734 ms (js: 408 ms)
    // wasm Vivaldi 5.3: 703 ms (js: 247 ms)
    // wasm IE         : xD     (js: 357 ms)
}

let resultEl = null;
let totalTim = 0;

function run(el) {
    el.innerHTML = "...";
    console.log("start");
    setTimeout(function () {
        calc(el);
    }, 10);
}
