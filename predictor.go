package fpc

type predictorClass uint8

const (
	fcmPredictor predictorClass = iota
	dfcmPredictor
)

type predictor interface {
	predict(actual uint64) (predicted uint64)
}

type fcm struct {
	table    []uint64
	size     uint64
	lastHash uint64
}

func newFCM(size int) *fcm {
	if !isPowerOfTwo(size) {
		panic("invalid size, must be a power of two")
	}
	return &fcm{
		table: make([]uint64, size, size),
		size:  uint64(size),
	}
}

func (f *fcm) hash(actual uint64) uint64 {
	return ((f.lastHash << 6) ^ (actual >> 48)) & (f.size - 1)
}

func (f *fcm) predict(actual uint64) uint64 {
	pred := f.table[f.lastHash]
	f.table[f.lastHash] = actual
	f.lastHash = f.hash(actual)
	return pred
}

type dfcm struct {
	table     []uint64
	size      uint64
	lastHash  uint64
	lastValue uint64
}

func newDFCM(size int) *dfcm {
	if !isPowerOfTwo(size) {
		panic("invalid size, must be a power of two")
	}
	return &dfcm{
		table: make([]uint64, size, size),
		size:  uint64(size),
	}
}

func (d *dfcm) hash(actual uint64) uint64 {
	return ((d.lastHash << 2) ^ (actual - d.lastValue>>40)) & (d.size - 1)
}

func (d *dfcm) predict(actual uint64) uint64 {
	pred := d.table[d.lastHash] + d.lastValue
	d.table[d.lastHash] = actual - d.lastValue
	d.lastHash = d.hash(actual)
	d.lastValue = actual
	return pred
}
