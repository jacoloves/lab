// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Fill.asm

// Runs an infinite loop that listens to the keyboard input.
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel;
// the screen should remain fully black as long as the key is pressed. 
// When no key is pressed, the program clears the screen, i.e. writes
// "white" in every pixel;
// the screen should remain fully clear as long as no key is pressed.

// Put your code here.
		@R0
		M=0

		@8192
		D=A
		@number_of_pixcel
		M=D
(LOOP_KBD)
		@KBD
		D=M
		@KEY_PRESSED
		D;JNE
(KEY_NOT_PRESSED)
		@R1
		M=0
		@CHECK_STATE_CHANGE
		0;JMP
(KEY_PRESSED)
		@R1
		M=1
(CHECK_STATE_CHANGE)
		@R0
		D=M
		@R1
		D=D-M
		@LOOP_KBD
		D;JEQ

		@R1
		D=M
		@R0
		M=D

		@SET_WHITE
		D;JEQ
(SET_BLACK)
		@color
		M=-1
		@PREPARE_FILL
		0;JMP
(SET_WHITE)
		@color
		M=0
(PREPARE_FILL)
		@SCREEN
		D=A
		@position
		M=D
		@count
		M=0
(LOOP_FILL)
		@count
		D=M
		@number_of_pixcel
		D=M-D
		@FILL_END
		D;JEQ

		@color
		D=M
		@position
		A=M
		M=D

		@position
		M=M+1
		@count
		M=M+1
		@LOOP_FILL
		0;JMP
(FILL_END)
		@R1
		M=0
		@LOOP_KBD
		0;JMP
