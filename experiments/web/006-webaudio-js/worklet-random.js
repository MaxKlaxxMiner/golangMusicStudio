console.log("start")

// random-noise-processor.js
class RandomNoiseProcessor extends AudioWorkletProcessor {
    process (inputs, outputs, parameters) {
        const output = outputs[0]
        output.forEach(channel => {
            for (let i = 0; i < channel.length; i++) {
                channel[i] = Math.random() * 2 - 1
                channel[i] *= 0.05;
            }
        })
        return true
    }
}

console.log("register")

registerProcessor("worklet-random", RandomNoiseProcessor);

console.log("ok.")
