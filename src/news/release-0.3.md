+++
title = "Go HEP release 0.3"
date = "2017-05-11T09:50:31+02:00"
weight = 10
+++



Release `v0.3` brings read support for [ROOT](https://root.cern.ch) `TTrees` containing user-defined classes.

Not everything is supported, though.

- Changelog: https://github.com/go-hep/hep/releases/v0.3
- Sources (zip): https://github.com/go-hep/hep/archive/v0.3.zip
- Sources (tar): https://github.com/go-hep/hep/archive/v0.3.tar.gz
- DOI: https://doi.org/10.5281/zenodo.574888

## `rootio`

What is supported and tested so far (in no-split and full-split mode):

- https://github.com/go-hep/hep/blob/v0.3/rootio/gendata/gen-evnt-tree.go#L51

```c++
// edm.cxx

struct P3 {
	int32_t Px;
	double  Py;
	int32_t Pz;
};

struct Event {
	TString  Beg;
	int16_t  I16;
	int32_t  I32;
	int64_t  I64;
	uint16_t U16;
	uint32_t U32;
	uint64_t U64;
	float    F32;
	double   F64;
	TString  Str;
	P3       P3;
	int16_t  ArrayI16[ARRAYSZ];
	int32_t  ArrayI32[ARRAYSZ];
	int64_t  ArrayI64[ARRAYSZ];
	uint16_t ArrayU16[ARRAYSZ];
	uint32_t ArrayU32[ARRAYSZ];
	uint64_t ArrayU64[ARRAYSZ];
	float    ArrayF32[ARRAYSZ];
	double   ArrayF64[ARRAYSZ];
	int32_t  N;
	int16_t  *SliceI16;  //[N]
	int32_t  *SliceI32;  //[N]
	int64_t  *SliceI64;  //[N]
	uint16_t *SliceU16;  //[N]
	uint32_t *SliceU32;  //[N]
	uint64_t *SliceU64;  //[N]
	float    *SliceF32;  //[N]
	double   *SliceF64;  //[N]
	std::string StdStr;
	std::vector<int16_t> StlVecI16;
	std::vector<int32_t> StlVecI32;
	std::vector<int64_t> StlVecI64;
	std::vector<uint16_t> StlVecU16;
	std::vector<uint32_t> StlVecU32;
	std::vector<uint64_t> StlVecU64;
	std::vector<float> StlVecF32;
	std::vector<double> StlVecF64;
	std::vector<std::string> StlVecStr;
	TString End;
};
```

which can be read with a Go `struct` with the same layout, translating `std::vector<T>` into slices of `T`, `std::string` and `TString` into Go `string`, etc...
see:

https://github.com/go-hep/hep/blob/v0.3/rootio/tree_test.go#L59

```go
// event.go

type EventType struct {
	Evt EventData `rootio:"evt"`
}

type Vec3 struct {
	X int32   `rootio:"Px"`
	Y float64 `rootio:"Py"`
	Z int32   `rootio:"Pz"`
}

type EventData struct {
	Beg    string      `rootio:"Beg"`
	I16    int16       `rootio:"Int16"`
	I32    int32       `rootio:"Int32"`
	I64    int64       `rootio:"Int64"`
	U16    uint16      `rootio:"UInt16"`
	U32    uint32      `rootio:"UInt32"`
	U64    uint64      `rootio:"UInt64"`
	F32    float32     `rootio:"Float32"`
	F64    float64     `rootio:"Float64"`
	Str    string      `rootio:"Str"`
	Vec    Vec3        `rootio:"P3"`
	ArrI16 [10]int16   `rootio:"ArrayI16"`
	ArrI32 [10]int32   `rootio:"ArrayI32"`
	ArrI64 [10]int64   `rootio:"ArrayI64"`
	ArrU16 [10]uint16  `rootio:"ArrayU16"`
	ArrU32 [10]uint32  `rootio:"ArrayU32"`
	ArrU64 [10]uint64  `rootio:"ArrayU64"`
	ArrF32 [10]float32 `rootio:"ArrayF32"`
	ArrF64 [10]float64 `rootio:"ArrayF64"`
	N      int32       `rootio:"N"`
	SliI16 []int16     `rootio:"SliceI16"`
	SliI32 []int32     `rootio:"SliceI32"`
	SliI64 []int64     `rootio:"SliceI64"`
	SliU16 []uint16    `rootio:"SliceU16"`
	SliU32 []uint32    `rootio:"SliceU32"`
	SliU64 []uint64    `rootio:"SliceU64"`
	SliF32 []float32   `rootio:"SliceF32"`
	SliF64 []float64   `rootio:"SliceF64"`
	StdStr string      `rootio:"StdStr"`
	VecI16 []int16     `rootio:"StlVecI16"`
	VecI32 []int32     `rootio:"StlVecI32"`
	VecI64 []int64     `rootio:"StlVecI64"`
	VecU16 []uint16    `rootio:"StlVecU16"`
	VecU32 []uint32    `rootio:"StlVecU32"`
	VecU64 []uint64    `rootio:"StlVecU64"`
	VecF32 []float32   `rootio:"StlVecF32"`
	VecF64 []float64   `rootio:"StlVecF64"`
	VecStr []string    `rootio:"StlVecStr"`
	End    string      `rootio:"End"`
}
```

support for more use cases will come in due time.

Stay tuned!

## `hbook`

- Michael Ughetto contributed code to retrieve the bins of `hbook.H2D` histograms, thanks [@mughetto](https://github.com/mughetto)


