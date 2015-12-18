package fpc

import (
	"reflect"
	"testing"
)

func TestDecodePrefix(t *testing.T) {
	type output struct {
		n1, n2 uint8
		p1, p2 predictorClass
	}
	testcases := []struct {
		in   byte
		want output
	}{
		{
			in: binstr2byte("01110111"),
			want: output{
				n1: 8,
				n2: 8,
				p1: 0,
				p2: 0,
			},
		},
		{
			in: binstr2byte("11110111"),
			want: output{
				n1: 8,
				n2: 8,
				p1: 1,
				p2: 0,
			},
		},
		{
			in: binstr2byte("00101111"),
			want: output{
				n1: 2,
				n2: 8,
				p1: 0,
				p2: 1,
			},
		},
	}
	for i, tc := range testcases {
		var have output
		have.n1, have.n2, have.p1, have.p2 = decodePrefix(tc.in)
		if !reflect.DeepEqual(have, tc.want) {
			t.Errorf("decodePrefix test=%d  have=%+v  want=%+v", i, have, tc.want)
		}
	}
}

func TestDecodeOne(t *testing.T) {
	type output struct {
		n      int
		v1, v2 uint64
	}
	testcases := []struct {
		in   string
		want output
	}{
		{
			in: "01110111",
			want: output{
				n:  1,
				v1: 0,
				v2: 0,
			},
		},
		{
			in: "01110110 00000001",
			want: output{
				n:  2,
				v1: 0,
				v2: 1,
			},
		},
		{
			in: "01100110 00000001 00000001",
			want: output{
				n:  3,
				v1: 1,
				v2: 1,
			},
		},
		{
			in: "00110110 11111111 11111111 00000000  00000000 00000000 00000001",
			want: output{
				n:  7,
				v1: 65535,
				v2: 1,
			},
		},
	}
	for i, tc := range testcases {
		var have output
		have.n, have.v1, have.v2 = decodeOne(binstr2bytes(tc.in))
		if !reflect.DeepEqual(have, tc.want) {
			t.Errorf("decodeOne test=%d  have=%+v  want=%+v", i, have, tc.want)
		}
	}
}

func TestDecodeBlockHeader(t *testing.T) {
	testcases := []struct {
		in   []byte
		want blockHeader
	}{
		{
			in: []byte{0x00, 0x80, 0x00, 0xb6, 0x35, 0x02},
			want: blockHeader{
				nRecords: 32768,
				nBytes:   144822,
			},
		},
		{
			in: []byte{0x00, 0x80, 0x00, 0xc2, 0x43, 0x00},
			want: blockHeader{
				nRecords: 32768,
				nBytes:   17346,
			},
		},
	}
	for i, tc := range testcases {
		have := decodeBlockHeader(tc.in)
		if !reflect.DeepEqual(have, tc.want) {
			t.Errorf("decodeBlockHeader test=%d  have=%+v  want=%+v", i, have, tc.want)
		}
	}
}
