// Code generated by execgen; DO NOT EDIT.

// Copyright 2019 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package colexecprojconst

import (
	"bytes"
	"regexp"

	"github.com/cockroachdb/cockroach/pkg/col/coldata"
)

type projPrefixBytesBytesConstOp struct {
	projConstOpBase
	constArg        []byte
	negate          bool
	caseInsensitive bool
}

func (p projPrefixBytesBytesConstOp) Next() coldata.Batch {
	_negate := p.negate
	_caseInsensitive := p.caseInsensitive
	batch := p.Input.Next()
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}
	vec := batch.ColVec(p.colIdx)
	var col *coldata.Bytes
	col = vec.Bytes()
	projVec := batch.ColVec(p.outputIdx)
	p.allocator.PerformOperation([]coldata.Vec{projVec}, func() {
		// Capture col to force bounds check to work. See
		// https://github.com/golang/go/issues/39756
		col := col
		projCol := projVec.Bool()
		_outNulls := projVec.Nulls()

		hasNullsAndNotCalledOnNullInput := vec.Nulls().MaybeHasNulls()
		if hasNullsAndNotCalledOnNullInput {
			colNulls := vec.Nulls()
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					if !colNulls.NullAt(i) {
						// We only want to perform the projection operation if the value is not null.
						arg := col.Get(i)
						if _caseInsensitive {
							arg = bytes.ToUpper(arg)
						}
						projCol[i] = bytes.HasPrefix(arg, p.constArg) != _negate
					}
				}
			} else {
				_ = projCol.Get(n - 1)
				_ = col.Get(n - 1)
				for i := 0; i < n; i++ {
					if !colNulls.NullAt(i) {
						// We only want to perform the projection operation if the value is not null.
						arg := col.Get(i)
						if _caseInsensitive {
							arg = bytes.ToUpper(arg)
						}
						projCol[i] = bytes.HasPrefix(arg, p.constArg) != _negate
					}
				}
			}
			projVec.SetNulls(_outNulls.Or(*colNulls))
		} else {
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					arg := col.Get(i)
					if _caseInsensitive {
						arg = bytes.ToUpper(arg)
					}
					projCol[i] = bytes.HasPrefix(arg, p.constArg) != _negate
				}
			} else {
				_ = projCol.Get(n - 1)
				_ = col.Get(n - 1)
				for i := 0; i < n; i++ {
					arg := col.Get(i)
					if _caseInsensitive {
						arg = bytes.ToUpper(arg)
					}
					projCol[i] = bytes.HasPrefix(arg, p.constArg) != _negate
				}
			}
		}
	})
	return batch
}

type projSuffixBytesBytesConstOp struct {
	projConstOpBase
	constArg        []byte
	negate          bool
	caseInsensitive bool
}

func (p projSuffixBytesBytesConstOp) Next() coldata.Batch {
	_negate := p.negate
	_caseInsensitive := p.caseInsensitive
	batch := p.Input.Next()
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}
	vec := batch.ColVec(p.colIdx)
	var col *coldata.Bytes
	col = vec.Bytes()
	projVec := batch.ColVec(p.outputIdx)
	p.allocator.PerformOperation([]coldata.Vec{projVec}, func() {
		// Capture col to force bounds check to work. See
		// https://github.com/golang/go/issues/39756
		col := col
		projCol := projVec.Bool()
		_outNulls := projVec.Nulls()

		hasNullsAndNotCalledOnNullInput := vec.Nulls().MaybeHasNulls()
		if hasNullsAndNotCalledOnNullInput {
			colNulls := vec.Nulls()
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					if !colNulls.NullAt(i) {
						// We only want to perform the projection operation if the value is not null.
						arg := col.Get(i)
						if _caseInsensitive {
							arg = bytes.ToUpper(arg)
						}
						projCol[i] = bytes.HasSuffix(arg, p.constArg) != _negate
					}
				}
			} else {
				_ = projCol.Get(n - 1)
				_ = col.Get(n - 1)
				for i := 0; i < n; i++ {
					if !colNulls.NullAt(i) {
						// We only want to perform the projection operation if the value is not null.
						arg := col.Get(i)
						if _caseInsensitive {
							arg = bytes.ToUpper(arg)
						}
						projCol[i] = bytes.HasSuffix(arg, p.constArg) != _negate
					}
				}
			}
			projVec.SetNulls(_outNulls.Or(*colNulls))
		} else {
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					arg := col.Get(i)
					if _caseInsensitive {
						arg = bytes.ToUpper(arg)
					}
					projCol[i] = bytes.HasSuffix(arg, p.constArg) != _negate
				}
			} else {
				_ = projCol.Get(n - 1)
				_ = col.Get(n - 1)
				for i := 0; i < n; i++ {
					arg := col.Get(i)
					if _caseInsensitive {
						arg = bytes.ToUpper(arg)
					}
					projCol[i] = bytes.HasSuffix(arg, p.constArg) != _negate
				}
			}
		}
	})
	return batch
}

type projContainsBytesBytesConstOp struct {
	projConstOpBase
	constArg        []byte
	negate          bool
	caseInsensitive bool
}

func (p projContainsBytesBytesConstOp) Next() coldata.Batch {
	_negate := p.negate
	_caseInsensitive := p.caseInsensitive
	batch := p.Input.Next()
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}
	vec := batch.ColVec(p.colIdx)
	var col *coldata.Bytes
	col = vec.Bytes()
	projVec := batch.ColVec(p.outputIdx)
	p.allocator.PerformOperation([]coldata.Vec{projVec}, func() {
		// Capture col to force bounds check to work. See
		// https://github.com/golang/go/issues/39756
		col := col
		projCol := projVec.Bool()
		_outNulls := projVec.Nulls()

		hasNullsAndNotCalledOnNullInput := vec.Nulls().MaybeHasNulls()
		if hasNullsAndNotCalledOnNullInput {
			colNulls := vec.Nulls()
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					if !colNulls.NullAt(i) {
						// We only want to perform the projection operation if the value is not null.
						arg := col.Get(i)
						if _caseInsensitive {
							arg = bytes.ToUpper(arg)
						}
						projCol[i] = bytes.Contains(arg, p.constArg) != _negate
					}
				}
			} else {
				_ = projCol.Get(n - 1)
				_ = col.Get(n - 1)
				for i := 0; i < n; i++ {
					if !colNulls.NullAt(i) {
						// We only want to perform the projection operation if the value is not null.
						arg := col.Get(i)
						if _caseInsensitive {
							arg = bytes.ToUpper(arg)
						}
						projCol[i] = bytes.Contains(arg, p.constArg) != _negate
					}
				}
			}
			projVec.SetNulls(_outNulls.Or(*colNulls))
		} else {
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					arg := col.Get(i)
					if _caseInsensitive {
						arg = bytes.ToUpper(arg)
					}
					projCol[i] = bytes.Contains(arg, p.constArg) != _negate
				}
			} else {
				_ = projCol.Get(n - 1)
				_ = col.Get(n - 1)
				for i := 0; i < n; i++ {
					arg := col.Get(i)
					if _caseInsensitive {
						arg = bytes.ToUpper(arg)
					}
					projCol[i] = bytes.Contains(arg, p.constArg) != _negate
				}
			}
		}
	})
	return batch
}

type projSkeletonBytesBytesConstOp struct {
	projConstOpBase
	constArg        [][]byte
	negate          bool
	caseInsensitive bool
}

func (p projSkeletonBytesBytesConstOp) Next() coldata.Batch {
	_negate := p.negate
	_caseInsensitive := p.caseInsensitive
	batch := p.Input.Next()
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}
	vec := batch.ColVec(p.colIdx)
	var col *coldata.Bytes
	col = vec.Bytes()
	projVec := batch.ColVec(p.outputIdx)
	p.allocator.PerformOperation([]coldata.Vec{projVec}, func() {
		// Capture col to force bounds check to work. See
		// https://github.com/golang/go/issues/39756
		col := col
		projCol := projVec.Bool()
		_outNulls := projVec.Nulls()

		hasNullsAndNotCalledOnNullInput := vec.Nulls().MaybeHasNulls()
		if hasNullsAndNotCalledOnNullInput {
			colNulls := vec.Nulls()
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					if !colNulls.NullAt(i) {
						// We only want to perform the projection operation if the value is not null.
						arg := col.Get(i)

						{
							if _caseInsensitive {
								arg = bytes.ToUpper(arg)
							}
							var idx, skeletonIdx int
							for skeletonIdx < len(p.constArg) {
								idx = bytes.Index(arg, p.constArg[skeletonIdx])
								if idx < 0 {
									break
								}
								arg = arg[idx+len(p.constArg[skeletonIdx]):]
								skeletonIdx++
							}
							projCol[i] = skeletonIdx == len(p.constArg) != _negate
						}
					}
				}
			} else {
				_ = projCol.Get(n - 1)
				_ = col.Get(n - 1)
				for i := 0; i < n; i++ {
					if !colNulls.NullAt(i) {
						// We only want to perform the projection operation if the value is not null.
						arg := col.Get(i)

						{
							if _caseInsensitive {
								arg = bytes.ToUpper(arg)
							}
							var idx, skeletonIdx int
							for skeletonIdx < len(p.constArg) {
								idx = bytes.Index(arg, p.constArg[skeletonIdx])
								if idx < 0 {
									break
								}
								arg = arg[idx+len(p.constArg[skeletonIdx]):]
								skeletonIdx++
							}
							projCol[i] = skeletonIdx == len(p.constArg) != _negate
						}
					}
				}
			}
			projVec.SetNulls(_outNulls.Or(*colNulls))
		} else {
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					arg := col.Get(i)

					{
						if _caseInsensitive {
							arg = bytes.ToUpper(arg)
						}
						var idx, skeletonIdx int
						for skeletonIdx < len(p.constArg) {
							idx = bytes.Index(arg, p.constArg[skeletonIdx])
							if idx < 0 {
								break
							}
							arg = arg[idx+len(p.constArg[skeletonIdx]):]
							skeletonIdx++
						}
						projCol[i] = skeletonIdx == len(p.constArg) != _negate
					}
				}
			} else {
				_ = projCol.Get(n - 1)
				_ = col.Get(n - 1)
				for i := 0; i < n; i++ {
					arg := col.Get(i)

					{
						if _caseInsensitive {
							arg = bytes.ToUpper(arg)
						}
						var idx, skeletonIdx int
						for skeletonIdx < len(p.constArg) {
							idx = bytes.Index(arg, p.constArg[skeletonIdx])
							if idx < 0 {
								break
							}
							arg = arg[idx+len(p.constArg[skeletonIdx]):]
							skeletonIdx++
						}
						projCol[i] = skeletonIdx == len(p.constArg) != _negate
					}
				}
			}
		}
	})
	return batch
}

type projRegexpBytesBytesConstOp struct {
	projConstOpBase
	constArg *regexp.Regexp
	negate   bool
}

func (p projRegexpBytesBytesConstOp) Next() coldata.Batch {
	_negate := p.negate
	batch := p.Input.Next()
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}
	vec := batch.ColVec(p.colIdx)
	var col *coldata.Bytes
	col = vec.Bytes()
	projVec := batch.ColVec(p.outputIdx)
	p.allocator.PerformOperation([]coldata.Vec{projVec}, func() {
		// Capture col to force bounds check to work. See
		// https://github.com/golang/go/issues/39756
		col := col
		projCol := projVec.Bool()
		_outNulls := projVec.Nulls()

		hasNullsAndNotCalledOnNullInput := vec.Nulls().MaybeHasNulls()
		if hasNullsAndNotCalledOnNullInput {
			colNulls := vec.Nulls()
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					if !colNulls.NullAt(i) {
						// We only want to perform the projection operation if the value is not null.
						arg := col.Get(i)
						projCol[i] = p.constArg.Match(arg) != _negate
					}
				}
			} else {
				_ = projCol.Get(n - 1)
				_ = col.Get(n - 1)
				for i := 0; i < n; i++ {
					if !colNulls.NullAt(i) {
						// We only want to perform the projection operation if the value is not null.
						arg := col.Get(i)
						projCol[i] = p.constArg.Match(arg) != _negate
					}
				}
			}
			projVec.SetNulls(_outNulls.Or(*colNulls))
		} else {
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					arg := col.Get(i)
					projCol[i] = p.constArg.Match(arg) != _negate
				}
			} else {
				_ = projCol.Get(n - 1)
				_ = col.Get(n - 1)
				for i := 0; i < n; i++ {
					arg := col.Get(i)
					projCol[i] = p.constArg.Match(arg) != _negate
				}
			}
		}
	})
	return batch
}
