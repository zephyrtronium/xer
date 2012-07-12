// Copyright 2012 Branden J Brown
//
// This software is provided 'as-is', without any express or implied
// warranty. In no event will the authors be held liable for any damages
// arising from the use of this software.
//
// Permission is granted to anyone to use this software for any purpose,
// including commercial applications, and to alter it and redistribute it
// freely, subject to the following restrictions:
//
//    1. The origin of this software must not be misrepresented; you must not
//    claim that you wrote the original software. If you use this software
//    in a product, an acknowledgment in the product documentation would be
//    appreciated but is not required.
//
//    2. Altered source versions must be plainly marked as such, and must not
//    be misrepresented as being the original software.
//
//    3. This notice may not be removed or altered from any source
//    distribution.

// func int63_basicany(*xer) int63
TEXT ·int63_basicany+0(SB), $16-16
	MOVQ	x+0(FP),BP // BP = (pointer to) x
	MOVQ	(BP),BX // BX = x.sum

	// popcount_3() from http://en.wikipedia.org/wiki/Hamming_weight
	MOVQ	BX,AX // AX = v
	SHRQ	$1,AX // AX = v >> 1
	MOVQ	BX,CX // CX = v
	MOVQ	$0x5555555555555555,DX // DX = m1
	ANDQ	DX,AX // AX = (v >> 1) & m1
	SUBQ	AX,CX // CX = v - ((v >> 1) & m1)

	MOVQ	CX,AX // AX = v
	MOVQ	$0x3333333333333333,DX // DX = m2
	SHRQ	$2,CX // CX = v >> 2
	ANDQ	DX,AX // AX = v & m2
	ANDQ	DX,CX // CX = (v >> 2) & m2
	ADDQ	CX,AX // AX = (v & m2) + ((v >> 2) & m2)

	MOVQ	AX,CX // CX = v
	SHRQ	$4,AX // AX = v >> 4
	MOVQ	$0x0f0f0f0f0f0f0f0f,DX // DX = m4
	ADDQ	AX,CX // CX = v + (v >> 4)
	ANDQ	DX,CX // CX = (v + (v >> 4)) & m4

	MOVQ	$0x0101010101010101,AX // AX = h01
	IMULQ	AX,CX // CX = v * h01
	MOVQ	BX,AX // AX = x.sum
	SHRQ	$56,CX // CX = (v * h01) >> 56 = popcount(v)

	RORQ	CX,AX // AX = ror(x.sum, popcount(x.sum)), like a badass

	XORQ	AX,BX // BX = x.sum ^ s
	MOVQ	16(BP),DX // DX = x.state
	MOVL	8(BP),CX // CX = x.c
	XORQ	(DX)(CX*8),BX // BX = x.sum ^ s ^ x.state[x.c]
	MOVQ	AX,(DX)(CX*8) // x.state[x.c] = s
	MOVQ	BX,(BP) // x.sum = x.sum ^ s ^ (old) x.state[x.c]
	//MOVQ	$0x7fffffffffffffff,BX
	INCL	CX // CX = x.c + 1
	CMPL	CX,24(BP) // if CX >= len(x.state) {
	JHI	oobbasicany
	MOVL	CX,8(BP) // x.c = CX
	JMP	donebasicany // } else {
oobbasicany:
	MOVL	$0,8(BP) // x.c = 0
donebasicany: // }
	//ANDQ	BX,AX // clear top bit
	BTRQ	$63,AX
	MOVQ	AX,.noname+8(FP)
	RET

// func int63_popcntany(*xer) int64
TEXT ·int63_popcntany+0(SB), $16-16
	MOVQ	x+0(FP),BP // BP = (pointer to) x
	MOVQ	(BP),BX // BX = x.sum

	// 6a yet lacks POPCNT
	// F3 REX.W 0F B8 ModR/M
	// REX.W = 0b01001000
	// ModR/M = 0xcb = 0b11001011 = BX -> CX
	BYTE	$0xf3;	BYTE	$0x48;	BYTE	$0x0f;	BYTE	$0xb8;	BYTE	$0xcb // CX = popcount(x.sum)
	MOVQ	BX,AX
	RORQ	CX,AX // AX = ror(x.sum, popcount(x.sum))

	XORQ	AX,BX // BX = x.sum ^ s
	MOVB	8(BP),CX // CX = x.c
	XORQ	16(BP)(CX*8),BX	// BX = x.sum ^ s ^ x.state[x.c]
	MOVQ	AX,16(BP)(CX*8) // x.state[x.c] = s
	MOVQ	BX,(BP) // x.sum = x.sum ^ s ^ (old) x.state[x.c]
	//MOVQ	$0x7fffffffffffffff,BX
	INCL	CX // CX = x.c + 1
	CMPL	CX,24(BP) // if CX >= len(x.state) {
	JHI	oobpopany
	MOVL	CX,8(BP) // x.c = CX
	JMP	donepopany // } else {
oobpopany:
	MOVL	$0,8(BP) // x.c = 0
donepopany: // }
	//ANDQ	BX,AX // clear top bit
	BTRQ	$63,AX
	MOVQ	AX,.noname+8(FP)
	RET

// func int63_basic256(*xer256) int64
TEXT ·int63_basic256+0(SB), $16-16
	MOVQ	x+0(FP),BP // BP = (pointer to) x
	MOVQ	(BP),BX // BX = x.sum

	// popcount_3() from http://en.wikipedia.org/wiki/Hamming_weight
	MOVQ	BX,AX // AX = v
	SHRQ	$1,AX // AX = v >> 1
	MOVQ	BX,CX // CX = v
	MOVQ	$0x5555555555555555,DX // DX = m1
	ANDQ	DX,AX // AX = (v >> 1) & m1
	SUBQ	AX,CX // CX = v - ((v >> 1) & m1)

	MOVQ	CX,AX // AX = v
	MOVQ	$0x3333333333333333,DX // DX = m2
	SHRQ	$2,CX // CX = v >> 2
	ANDQ	DX,AX // AX = v & m2
	ANDQ	DX,CX // CX = (v >> 2) & m2
	ADDQ	CX,AX // AX = (v & m2) + ((v >> 2) & m2)

	MOVQ	AX,CX // CX = v
	SHRQ	$4,AX // AX = v >> 4
	MOVQ	$0x0f0f0f0f0f0f0f0f,DX // DX = m4
	ADDQ	AX,CX // CX = v + (v >> 4)
	ANDQ	DX,CX // CX = (v + (v >> 4)) & m4

	MOVQ	$0x0101010101010101,AX // AX = h01
	IMULQ	AX,CX // CX = v * h01
	MOVQ	BX,AX // AX = x.sum
	SHRQ	$56,CX // CX = (v * h01) >> 56 = popcount(v)

	RORQ	CX,AX // AX = ror(x.sum, popcount(x.sum)), like a badass

	XORQ	AX,BX // BX = x.sum ^ s
	MOVB	8(BP),CX // CX = x.c
	XORQ	16(BP)(CX*8),BX	// BX = x.sum ^ s ^ x.state[x.c]
	MOVQ	AX,16(BP)(CX*8) // x.state[x.c] = s
	MOVQ	BX,(BP) // x.sum = x.sum ^ s ^ (old) x.state[x.c]
	//MOVQ	$0x7fffffffffffffff,BX
	INCB	8(BP) // x.c = (x.c+1) & 255

	//ANDQ	BX,AX // clear top bit
	BTRQ	$63,AX
	MOVQ	AX,.noname+8(FP)
	RET

// func int63_popcnt256(*xer256) int64
TEXT ·int63_popcnt256+0(SB), $16-16
	MOVQ	x+0(FP),BP // BP = (pointer to) x
	MOVQ	(BP),BX // BX = x.sum

	// 6a yet lacks POPCNT
	// F3 REX.W 0F B8 ModR/M
	// REX.W = 0b01001000
	// ModR/M = 0xcb = 0b11001011 = BX -> CX
	BYTE	$0xf3;	BYTE	$0x48;	BYTE	$0x0f;	BYTE	$0xb8;	BYTE	$0xcb // CX = popcount(x.sum)
	MOVQ	BX,AX
	RORQ	CX,AX // AX = ror(x.sum, popcount(x.sum))

	XORQ	AX,BX // BX = x.sum ^ s
	MOVB	8(BP),CX // CX = x.c
	XORQ	16(BP)(CX*8),BX	// BX = x.sum ^ s ^ x.state[x.c]
	MOVQ	AX,16(BP)(CX*8) // x.state[x.c] = s
	MOVQ	BX,(BP) // x.sum = x.sum ^ s ^ (old) x.state[x.c]
	//MOVQ	$0x7fffffffffffffff,BX
	INCB	8(BP) // x.c = (x.c+1) & 255

	//ANDQ	BX,AX // clear top bit
	BTRQ	$63,AX
	MOVQ	AX,.noname+8(FP)
	RET

// func int63_basic65536(*xer65536) int64
TEXT ·int63_basic65536+0(SB), $16-16
	MOVQ	x+0(FP),BP // BP = (pointer to) x
	MOVQ	(BP),BX // BX = x.sum

	// popcount_3() from http://en.wikipedia.org/wiki/Hamming_weight
	MOVQ	BX,AX // AX = v
	SHRQ	$1,AX // AX = v >> 1
	MOVQ	BX,CX // CX = v
	MOVQ	$0x5555555555555555,DX // DX = m1
	ANDQ	DX,AX // AX = (v >> 1) & m1
	SUBQ	AX,CX // CX = v - ((v >> 1) & m1)

	MOVQ	CX,AX // AX = v
	MOVQ	$0x3333333333333333,DX // DX = m2
	SHRQ	$2,CX // CX = v >> 2
	ANDQ	DX,AX // AX = v & m2
	ANDQ	DX,CX // CX = (v >> 2) & m2
	ADDQ	CX,AX // AX = (v & m2) + ((v >> 2) & m2)

	MOVQ	AX,CX // CX = v
	SHRQ	$4,AX // AX = v >> 4
	MOVQ	$0x0f0f0f0f0f0f0f0f,DX // DX = m4
	ADDQ	AX,CX // CX = v + (v >> 4)
	ANDQ	DX,CX // CX = (v + (v >> 4)) & m4

	MOVQ	$0x0101010101010101,AX // AX = h01
	IMULQ	AX,CX // CX = v * h01
	MOVQ	BX,AX // AX = x.sum
	SHRQ	$56,CX // CX = (v * h01) >> 56 = popcount(v)

	RORQ	CX,AX // AX = ror(x.sum, popcount(x.sum)), like a badass

	XORQ	AX,BX // BX = x.sum ^ s
	MOVW	8(BP),CX // CX = x.c
	XORQ	16(BP)(CX*8),BX	// BX = x.sum ^ s ^ x.state[x.c]
	MOVQ	AX,16(BP)(CX*8) // x.state[x.c] = s
	MOVQ	BX,(BP) // x.sum = x.sum ^ s ^ (old) x.state[x.c]
	MOVQ	$0x7fffffffffffffff,BX
	INCW	8(BP) // x.c = (x.c+1) & 65535

	ANDQ	BX,AX // clear top bit
	//BTRQ	$63,AX
	MOVQ	AX,.noname+8(FP)
	RET

// func int63_popcnt65536(*xer65536) int64
TEXT ·int63_popcnt65536+0(SB), $16-16
	MOVQ	x+0(FP),BP // BP = (pointer to) x
	MOVQ	(BP),BX // BX = x.sum

	// 6a yet lacks POPCNT
	// F3 REX.W 0F B8 ModR/M
	// REX.W = 0b01001000
	// ModR/M = 0xcb = 0b11001011 = BX -> CX
	BYTE	$0xf3;	BYTE	$0x48;	BYTE	$0x0f;	BYTE	$0xb8;	BYTE	$0xcb // CX = popcount(x.sum)
	MOVQ	BX,AX
	RORQ	CX,AX // AX = ror(x.sum, popcount(x.sum))

	XORQ	AX,BX // BX = x.sum ^ s
	MOVW	8(BP),CX // CX = x.c
	XORQ	16(BP)(CX*8),BX	// BX = x.sum ^ s ^ x.state[x.c]
	MOVQ	AX,16(BP)(CX*8) // x.state[x.c] = s
	MOVQ	BX,(BP) // x.sum = x.sum ^ s ^ (old) x.state[x.c]
	MOVQ	$0x7fffffffffffffff,BX
	INCW	8(BP) // x.c = (x.c+1) & 255

	ANDQ	BX,AX // clear top bit
	//BTRQ	$63,AX
	MOVQ	AX,.noname+8(FP)
	RET

// func havePopcnt() bool
TEXT ·havePopcnt+0(SB), $0-8
	XORQ	AX,AX
	INCB	AX
	CPUID
	BTL	$23,CX
	SETCS	.noname+0(FP)
	RET
