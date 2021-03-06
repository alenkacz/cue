// Copyright 2018 The CUE Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cue

import "cuelang.org/go/cue/token"

func opIn(op op, anyOf ...op) bool {
	for _, o := range anyOf {
		if o == op {
			return true
		}
	}
	return false
}

// isCmp reports whether an op is a comparator.
func (op op) isCmp() bool {
	return opEql <= op && op <= opGeq
}

type op uint16

const (
	opUnknown op = iota

	opUnify
	opDisjunction

	opLand
	opLor
	opNot

	opEql
	opNeq
	opMat
	opNMat

	opLss
	opGtr
	opLeq
	opGeq

	opAdd
	opSub
	opMul
	opQuo
	opRem

	opIDiv
	opIMod
	opIQuo
	opIRem
)

var opStrings = []string{
	opUnknown: "??",

	opUnify:       "&",
	opDisjunction: "|",

	opLand: "&&",
	opLor:  "||",
	opNot:  "!",

	opEql:  "==",
	opNeq:  "!=",
	opMat:  "=~",
	opNMat: "!~",

	opLss: "<",
	opGtr: ">",
	opLeq: "<=",
	opGeq: ">=",

	opAdd: "+",
	opSub: "-",
	opMul: "*",
	opQuo: "/",
	opRem: "%",

	opIDiv: "div",
	opIMod: "mod",
	opIQuo: "quo",
	opIRem: "rem",
}

func (op op) String() string { return opStrings[op] }

var tokenMap = map[token.Token]op{
	token.OR:  opDisjunction, // |
	token.AND: opUnify,       // &

	token.ADD: opAdd, // +
	token.SUB: opSub, // -
	token.MUL: opMul, // *
	token.QUO: opQuo, // /
	token.REM: opRem, // %

	token.IDIV: opIDiv, // div
	token.IMOD: opIMod, // mod
	token.IQUO: opIQuo, // quo
	token.IREM: opIRem, // rem

	token.LAND: opLand, // &&
	token.LOR:  opLor,  // ||

	token.EQL: opEql, // ==
	token.LSS: opLss, // <
	token.GTR: opGtr, // >
	token.NOT: opNot, // !

	token.NEQ:  opNeq,  // !=
	token.LEQ:  opLeq,  // <=
	token.GEQ:  opGeq,  // >=
	token.MAT:  opMat,  // =~
	token.NMAT: opNMat, // !~
}

var opMap = map[op]token.Token{}

func init() {
	for t, o := range tokenMap {
		opMap[o] = t
	}
}
