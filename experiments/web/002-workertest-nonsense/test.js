function calc(el) {
    let tim = Date.now();
    for (let i = 0; i < 1000000000; i++) {
    }
    tim = Date.now() - tim;
    el.innerHTML = "ok. (" + tim + " ms)";

    // js Chrome 103:  244 ms
    // js Firefox 102: 408 ms
    // js Vivaldi 5.3: 247 ms
    // js IE         : 357 ms
}

function run(el) {
    el.innerHTML = "...";
    console.log("start");
    setTimeout(function(){
        calc(el);
    }, 10);
}
