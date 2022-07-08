function calc() {
    let tim = Date.now();
    for (let i = 0; i < 1000000000; i++) {
    }
    tim = Date.now() - tim;
    asyncOutput("main-ok. (" + tim + " ms)");
}

let resultEl = null;
let totalTim = 0;

function asyncOutput(line) {
    setTimeout(() => {
        if (resultEl.innerHTML === "...") {
            resultEl.innerHTML = line;
        } else {
            resultEl.innerHTML += "<br>" + line;
            const tim = Date.now() - totalTim;
            resultEl.innerHTML += "<br>total-time: " + tim + " ms";
        }
    }, 1)
}

function run(el) {
    el.innerHTML = "...";
    resultEl = el;
    console.log("start");
    setTimeout(() => {
        totalTim = Date.now()
        calc();
        calc();
    }, 10);
}

function run2(el) {
    el.innerHTML = "...";
    resultEl = el;
    console.log("start");
    setTimeout(() => {
        totalTim = Date.now()
        console.log("main: send -> worker \"calc\"")
        myWorker.postMessage("calc");
        calc();
    }, 10);
}

let myWorker = null;

if (window.Worker) {
    myWorker = new Worker("worker.js");

    myWorker.onmessage = (e) => {
        if (e.data.length && e.data[0] === "calc") {
            asyncOutput("worker-ok. (" + e.data[1] + " ms)");
        }
    }
} else {
    console.error("Worker not found!")
}
