// +build !noasm
// Generated by PeachPy 0.2.0 from dot_product_avx.py


// func DotProductAVX(a_base *float32, a_len uint, x_cap uint, b_base *float32, b_len uint, y_cap uint) float32
TEXT ·DotProductAVX(SB),4,$0-52
	MOVQ a_base+0(FP), AX
	MOVQ b_base+24(FP), BX
	MOVQ a_len+8(FP), CX
	MOVQ b_len+32(FP), DX
	BYTE $0xC5; BYTE $0xFC; BYTE $0x57; BYTE $0xC0 // VXORPS ymm0, ymm0, ymm0
	BYTE $0xC5; BYTE $0xF4; BYTE $0x57; BYTE $0xC9 // VXORPS ymm1, ymm1, ymm1
	BYTE $0xC5; BYTE $0xF4; BYTE $0x57; BYTE $0xC9 // VXORPS ymm1, ymm1, ymm1
	BYTE $0xC5; BYTE $0xF4; BYTE $0x57; BYTE $0xC9 // VXORPS ymm1, ymm1, ymm1
	XORQ DX, DX
	JCS vector_loop_end
vector_loop_begin:
		BYTE $0xC5; BYTE $0xFC; BYTE $0x10; BYTE $0x08 // VMOVUPS ymm1, [rax]
		BYTE $0xC5; BYTE $0xFC; BYTE $0x10; BYTE $0x13 // VMOVUPS ymm2, [rbx]
		BYTE $0xC4; BYTE $0xE2; BYTE $0x75; BYTE $0xB8; BYTE $0xC2 // VFMADD231PS ymm0, ymm1, ymm2
		ADDQ $32, AX
		ADDQ $32, BX
		SUBQ $8, CX
		JCC vector_loop_begin
vector_loop_end:
	BYTE $0xC5; BYTE $0xFF; BYTE $0x7C; BYTE $0xC0 // VHADDPS ymm0, ymm0, ymm0
	BYTE $0xC5; BYTE $0xFF; BYTE $0x7C; BYTE $0xC0 // VHADDPS ymm0, ymm0, ymm0
	BYTE $0xC4; BYTE $0xE3; BYTE $0x7D; BYTE $0x19; BYTE $0xC1; BYTE $0x01 // VEXTRACTF128 xmm1, ymm0, 1
	BYTE $0xC5; BYTE $0xF8; BYTE $0x58; BYTE $0xC1 // VADDPS xmm0, xmm0, xmm1
	MOVSS X0, ret+48(FP)
	BYTE $0xC5; BYTE $0xF8; BYTE $0x77 // VZEROUPPER
	RET
