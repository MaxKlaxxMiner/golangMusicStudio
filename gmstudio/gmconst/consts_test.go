package gmconst

import "testing"

func TestConsts(t *testing.T) {
	if SampleRate != 44100 && SampleRate != 48000 && SampleRate != 88200 && SampleRate != 96000 {
		t.Errorf("may not supported Samplerate: %d", SampleRate)
	}

	if SampleBits != 32 {
		t.Errorf("SampleBits != 32 not supported")
	}

	if DynamicBits < 4 || DynamicBits+1 > SampleBits {
		t.Errorf("DynamicBits out of range")
	}

	if AaBitsMQ < 1 || AaBitsMQ+SampleBits >= 64 {
		t.Errorf("AaBitsMQ out of range")
	}

	if AaBitsHQ < AaBitsMQ || AaBitsHQ+SampleBits >= 64 {
		t.Errorf("AaBitsLQ out of range")
	}
}
