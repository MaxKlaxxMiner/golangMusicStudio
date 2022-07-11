// noinspection JSUnresolvedFunction

const myRequest = new Request("main.wasm");

let wasm = null;
const myHeaders = new Headers();
myHeaders.append('pragma', 'no-cache');
myHeaders.append('cache-control', 'no-cache');
const myInit = {method: 'GET', headers: myHeaders};
WebAssembly.instantiateStreaming(fetch(myRequest, myInit)).then(obj => {
    wasm = obj.instance.exports;
});

function calc(el) {
    let tim = Date.now();
    wasm.calc();
    tim = Date.now() - tim;

    let tim2 = Date.now();
    wasm.calc2(1000000000);
    tim2 = Date.now() - tim2;

    el.innerHTML = "wasm-ok. (" + tim + " ms)<br>opti-wasm. (" + tim2 + " ms)";

    // wasm Chrome 103:  244/326 ms (go-wasm: 702 ms, js: 244 ms)
    // wasm Firefox 102: 253/343 ms (go-wasm: 734 ms, js: 408 ms)
    // wasm Vivaldi 5.3: 245/326 ms (go-wasm: 703 ms, js: 247 ms)
    // wasm IE         : xD     (js: 357 ms)
}

function run(el) {
    el.innerHTML = "...";
    console.log("start");
    setTimeout(function () {
        calc(el);
    }, 10);
}
