package gmconst

import "testing"

func TestConsts(t *testing.T) {
	if SampleBits != 32 {
		t.Errorf("SampleBits != 32 not supported")
	}

	if DynamicBits < 4 || DynamicBits+1 > SampleBits {
		t.Errorf("DynamicBits out of range")
	}

	if AaBitsLQ < 1 || AaBitsLQ+SampleBits >= 64 {
		t.Errorf("AaBitsLQ out of range")
	}

	if AaBitsHQ < AaBitsLQ || AaBitsHQ+SampleBits >= 64 {
		t.Errorf("AaBitsLQ out of range")
	}
}
