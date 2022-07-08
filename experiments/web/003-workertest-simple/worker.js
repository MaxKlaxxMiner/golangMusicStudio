onmessage = (e) => {
    console.log("worker: recv \"" + e.data + "\"")
    if (e.data === "calc") {
        const tim = calc();
        postMessage(["calc", tim])
    }
}

function calc() {
    let tim = Date.now();
    for (let i = 0; i < 1000000000; i++) {
    }
    return Date.now() - tim;
}
