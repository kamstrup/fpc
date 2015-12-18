package fpc

type refinput struct {
	comp uint
	in   []float64
}

type reftestcase struct {
	comp         uint
	uncompressed []float64
	compressed   []byte
}

// Reference tests which have been generated by running fpc.c
var refTests = []reftestcase{
	{
		comp:         1,
		uncompressed: []float64{},
		compressed: []byte{
			0x01},
	},
	{
		comp:         3,
		uncompressed: []float64{},
		compressed: []byte{
			0x03},
	},
	{
		comp:         10,
		uncompressed: []float64{},
		compressed: []byte{
			0x0a},
	},
	{
		comp:         1,
		uncompressed: []float64{1, 1, 0.9, 0.9},
		compressed: []byte{
			0x01, 0x04, 0x00, 0x00, 0x17, 0x00, 0x00, 0x70,
			0x60, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf0,
			0x3f, 0xcd, 0xcc, 0xcc, 0xcc, 0xcc, 0xcc, 0x1c,
		},
	},
	{
		comp:         3,
		uncompressed: []float64{1, 1, 0.9, 0.9},
		compressed: []byte{
			0x03, 0x04, 0x00, 0x00, 0x17, 0x00, 0x00, 0x70,
			0x68, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf0,
			0x3f, 0xcd, 0xcc, 0xcc, 0xcc, 0xcc, 0xcc, 0x1c,
		},
	},
	{
		comp:         10,
		uncompressed: []float64{1, 1, 0.9, 0.9},
		compressed: []byte{
			0x0a, 0x04, 0x00, 0x00, 0x1f, 0x00, 0x00, 0x77,
			0x68, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf0,
			0x3f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf0,
			0x3f, 0xcd, 0xcc, 0xcc, 0xcc, 0xcc, 0xcc, 0x1c,
		},
	},
	{
		comp:         1,
		uncompressed: []float64{0, 0, 0, 0, 0},
		compressed: []byte{
			0x01, 0x05, 0x00, 0x00, 0x09, 0x00, 0x00, 0x00,
			0x00, 0x00},
	},
	{
		comp:         3,
		uncompressed: []float64{0, 0, 0, 0, 0},
		compressed: []byte{
			0x03, 0x05, 0x00, 0x00, 0x09, 0x00, 0x00, 0x00,
			0x00, 0x00},
	},
	{
		comp:         10,
		uncompressed: []float64{0, 0, 0, 0, 0},
		compressed: []byte{
			0x0a, 0x05, 0x00, 0x00, 0x09, 0x00, 0x00, 0x00,
			0x00, 0x00},
	}}
