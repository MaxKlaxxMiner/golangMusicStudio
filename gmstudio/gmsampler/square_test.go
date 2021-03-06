package gmsampler

import (
	"testing"
)

var squareSamplesLQA2 = []float32{
	/*   0 */ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	/* 201 */ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
}

var squareSamplesHQA2 = []float32{
	/*   0 */ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, -0.545, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 0.091,
	/* 201 */ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
}

var squareSamplesLQC4 = []float32{
	/*   0 */ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	/*  85 */ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	/* 169 */ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	/* 253 */ 1, 1, 1,
}

var squareSamplesHQC4 = []float32{
	/*   0 */ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, -0.719, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 0.438,
	/*  85 */ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, -0.158, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -0.123,
	/* 169 */ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0.404, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -0.685,
	/* 253 */ 1, 1, 1,
}

var squareSamplesLQC6 = []float32{
	/*   0 */ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	/*  22 */ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	/*  43 */ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	/*  64 */ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	/*  85 */ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	/* 106 */ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	/* 127 */ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	/* 148 */ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	/* 169 */ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	/* 190 */ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	/* 211 */ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	/* 232 */ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	/* 253 */ 1, 1, 1,
}

var squareSamplesHQC6 = []float32{
	/*   0 */ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0.07, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 0.86,
	/*  22 */ 1, 1, 1, 1, 1, 1, 1, 1, 1, 0.211, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 0.719,
	/*  43 */ 1, 1, 1, 1, 1, 1, 1, 1, 1, 0.351, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 0.579,
	/*  64 */ 1, 1, 1, 1, 1, 1, 1, 1, 1, 0.491, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 0.438,
	/*  85 */ 1, 1, 1, 1, 1, 1, 1, 1, 1, 0.632, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 0.298,
	/* 106 */ 1, 1, 1, 1, 1, 1, 1, 1, 1, 0.772, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 0.158,
	/* 127 */ 1, 1, 1, 1, 1, 1, 1, 1, 1, 0.912, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 0.017,
	/* 148 */ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, -0.947, -1, -1, -1, -1, -1, -1, -1, -1, -1, -0.123,
	/* 169 */ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, -0.807, -1, -1, -1, -1, -1, -1, -1, -1, -1, -0.263,
	/* 190 */ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, -0.666, -1, -1, -1, -1, -1, -1, -1, -1, -1, -0.404,
	/* 211 */ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, -0.526, -1, -1, -1, -1, -1, -1, -1, -1, -1, -0.544,
	/* 232 */ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, -0.386, -1, -1, -1, -1, -1, -1, -1, -1, -1, -0.685,
	/* 253 */ 1, 1, 1,
}

var squareSamplesLQC9 = []float32{
	/*   0 */ 1, 1, -1, 1, -1, -1, 1, -1, 1, 1, -1, 1, -1, -1, 1, -1, 1, 1, -1, 1, -1, -1,
	/*  22 */ 1, -1, 1, 1, -1, 1, -1, 1, 1, -1, 1, -1, -1, 1, -1, 1, 1, -1, 1, -1, -1,
	/*  43 */ 1, -1, 1, 1, -1, 1, -1, -1, 1, -1, 1, -1, -1, 1, -1, 1, 1, -1, 1, -1, -1,
	/*  64 */ 1, -1, 1, 1, -1, 1, -1, -1, 1, -1, 1, 1, -1, 1, -1, -1, 1, -1, 1, -1, -1,
	/*  85 */ 1, -1, 1, 1, -1, 1, -1, -1, 1, -1, 1, 1, -1, 1, -1, -1, 1, -1, 1, 1, -1,
	/* 106 */ 1, -1, 1, 1, -1, 1, -1, -1, 1, -1, 1, 1, -1, 1, -1, -1, 1, -1, 1, 1, -1,
	/* 127 */ 1, -1, -1, 1, -1, 1, 1, -1, 1, -1, 1, 1, -1, 1, -1, -1, 1, -1, 1, 1, -1,
	/* 148 */ 1, -1, -1, 1, -1, 1, 1, -1, 1, -1, -1, 1, -1, 1, -1, -1, 1, -1, 1, 1, -1,
	/* 169 */ 1, -1, -1, 1, -1, 1, 1, -1, 1, -1, -1, 1, -1, 1, 1, -1, 1, -1, 1, 1, -1,
	/* 190 */ 1, -1, -1, 1, -1, 1, 1, -1, 1, -1, -1, 1, -1, 1, 1, -1, 1, -1, -1, 1, -1,
	/* 211 */ 1, 1, -1, 1, -1, 1, 1, -1, 1, -1, -1, 1, -1, 1, 1, -1, 1, -1, -1, 1, -1,
	/* 232 */ 1, 1, -1, 1, -1, -1, 1, -1, 1, -1, -1, 1, -1, 1, 1, -1, 1, -1, -1, 1, -1,
	/* 253 */ 1, 1, -1,
}

var squareSamplesHQC9 = []float32{
	/*   0 */ 1, -0.366, -0.268, 0.901, -1, 0.465, 0.169, -0.803, 1, -0.564, -0.07, 0.704, -1, 0.662, -0.028, -0.605, 1, -0.761, 0.127, 0.507, -1, 0.86,
	/*  22 */ -0.226, -0.408, 1, -0.958, 0.325, 0.309, -0.943, 1, -0.423, -0.211, 0.844, -1, 0.522, 0.112, -0.746, 1, -0.621, -0.013, 0.647, -1, 0.719,
	/*  43 */ -0.085, -0.548, 1, -0.818, 0.184, 0.45, -1, 0.917, -0.283, -0.351, 0.985, -1, 0.382, 0.252, -0.886, 1, -0.48, -0.154, 0.787, -1, 0.579,
	/*  64 */ 0.055, -0.689, 1, -0.678, 0.044, 0.59, -1, 0.776, -0.142, -0.491, 1, -0.875, 0.241, 0.393, -1, 0.974, -0.34, -0.294, 0.928, -1, 0.438,
	/*  85 */ 0.195, -0.829, 1, -0.537, -0.097, 0.73, -1, 0.636, -0.002, -0.632, 1, -0.735, 0.101, 0.533, -1, 0.833, -0.199, -0.434, 1, -0.932, 0.298,
	/* 106 */ 0.336, -0.969, 1, -0.397, -0.237, 0.871, -1, 0.495, 0.138, -0.772, 1, -0.594, -0.04, 0.673, -1, 0.693, -0.059, -0.575, 1, -0.792, 0.158,
	/* 127 */ 0.476, -1, 0.89, -0.256, -0.377, 1, -0.989, 0.355, 0.279, -0.912, 1, -0.454, -0.18, 0.814, -1, 0.552, 0.081, -0.715, 1, -0.651, 0.017,
	/* 148 */ 0.616, -1, 0.75, -0.116, -0.518, 1, -0.848, 0.215, 0.419, -1, 0.947, -0.313, -0.32, 0.954, -1, 0.412, 0.222, -0.855, 1, -0.511, -0.123,
	/* 169 */ 0.757, -1, 0.609, 0.024, -0.658, 1, -0.708, 0.074, 0.559, -1, 0.807, -0.173, -0.461, 1, -0.905, 0.272, 0.362, -0.996, 1, -0.37, -0.263,
	/* 190 */ 0.897, -1, 0.469, 0.165, -0.798, 1, -0.568, -0.066, 0.7, -1, 0.666, -0.033, -0.601, 1, -0.765, 0.131, 0.502, -1, 0.864, -0.23, -0.404,
	/* 211 */ 1, -0.962, 0.329, 0.305, -0.939, 1, -0.427, -0.206, 0.84, -1, 0.526, 0.108, -0.741, 1, -0.625, -0.009, 0.643, -1, 0.723, -0.09, -0.544,
	/* 232 */ 1, -0.822, 0.188, 0.445, -1, 0.921, -0.287, -0.347, 0.981, -1, 0.386, 0.248, -0.882, 1, -0.484, -0.149, 0.783, -1, 0.583, 0.051, -0.685,
	/* 253 */ 1, -0.682, 0.048,
}

func TestSquareLQ(t *testing.T) {
	SamplerCompareTone(t, "SquareLQ A2", squareSamplesLQA2, SquareLQ, "A2")
	SamplerCompareTone(t, "SquareLQ C4", squareSamplesLQC4, SquareLQ, "C4")
	SamplerCompareTone(t, "SquareLQ C6", squareSamplesLQC6, SquareLQ, "C6")
	SamplerCompareTone(t, "SquareLQ C9", squareSamplesLQC9, SquareLQ, "C9")
}

func TestSquareHQ(t *testing.T) {
	SamplerCompareTone(t, "SquareHQ A2", squareSamplesHQA2, SquareHQ, "A2")
	SamplerCompareTone(t, "SquareHQ C4", squareSamplesHQC4, SquareHQ, "C4")
	SamplerCompareTone(t, "SquareHQ C6", squareSamplesHQC6, SquareHQ, "C6")
	SamplerCompareTone(t, "SquareHQ C9", squareSamplesHQC9, SquareHQ, "C9")
}
