package partialsum

import (
	"github.com/hillbig/rsdic"
	"github.com/ugorji/go/codec"
)

// PartialSum stores non-negative integers V[0...N)
// and supports Sum, Find in O(1) time
// using at most (S + N) bits where S is the sum of V[0...N)
type PartialSum interface {
	// Increment add V[ind] += val
	// ind should hold ind >= Num
	IncTail(ind uint64, val uint64)

	// Num returns the number of vals
	Num() uint64

	// AllSum returns the sum of all vals
	AllSum() uint64

	// Lookup returns V[i] in O(1) time
	Lookup(ind uint64) (val uint64)

	// Sum returns V[0]+V[1]+...+V[ind-1] in O(1) time
	Sum(ind uint64) (sum uint64)

	// Lookup returns V[i] and V[0]+V[1]+...+V[i-1] in O(1) time
	LookupAndSum(ind uint64) (val uint64, sum uint64)

	// Find returns ind satisfying Sum(ind) <= val < Sum(ind+1)
	// and val - Sum(ind). If there are multiple inds
	// satisfiying this condition, return the minimum one.
	Find(val uint64) (ind uint64, offset uint64)

	// MarshalBinary encodes VecString into a binary form and returns the result.
	MarshalBinary() ([]byte, error)

	// UnmarshalBinary decodes the FixVec form a binary from generated MarshalBinary
	UnmarshalBinary([]byte) error
}

func New() PartialSum {
	dic := rsdic.New()
	dic.PushBack(true)
	return &partialSumImpl{
		dic: dic,
	}
}

type partialSumImpl struct {
	dic rsdic.RSDic
}

func (ps *partialSumImpl) IncTail(ind uint64, val uint64) {
	for i := ps.dic.OneNum(); i <= ind; i++ {
		ps.dic.PushBack(true)
	}
	for i := uint64(0); i < val; i++ {
		ps.dic.PushBack(false)
	}
}

func (ps partialSumImpl) Num() uint64 {
	return ps.dic.OneNum()
}

func (ps partialSumImpl) AllSum() uint64 {
	return ps.dic.ZeroNum()
}

func (ps partialSumImpl) Lookup(ind uint64) (val uint64) {
	return ps.dic.Select(ind+1, true) - ps.dic.Select(ind, true) - 1
}

func (ps partialSumImpl) Sum(ind uint64) (sum uint64) {
	return ps.dic.Rank(ps.dic.Select(ind, true), false)
}

func (ps partialSumImpl) LookupAndSum(ind uint64) (val uint64, sum uint64) {
	indPos := ps.dic.Select(ind, true)
	sum = ps.dic.Rank(indPos, false)
	val = ps.dic.Select(ind+1, true) - indPos - 1
	return
}

func (ps partialSumImpl) Find(val uint64) (ind uint64, offset uint64) {
	pos := ps.dic.Select(val, false)
	ind = ps.dic.Rank(pos, true) - 1
	offset = pos - ps.dic.Select(ind, true) - 1
	return
}

func (ps partialSumImpl) MarshalBinary() (out []byte, err error) {
	var mh codec.MsgpackHandle
	enc := codec.NewEncoderBytes(&out, &mh)
	err = enc.Encode(ps.dic)
	if err != nil {
		return
	}
	return
}

func (ps *partialSumImpl) UnmarshalBinary(in []byte) (err error) {
	var mh codec.MsgpackHandle
	dec := codec.NewDecoderBytes(in, &mh)
	err = dec.Decode(&ps.dic)
	if err != nil {
		return
	}
	return
}
