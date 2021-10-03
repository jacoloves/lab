		movl $0, -4(%rdp)
		movl $1, -8(%rdp)
		jmp .L2
L3:
		movl -8(%rdp), %eax
		addl %eax, -4(%rdp)
		addl $1, -8(%rdp)
L2:
		cmpl $100, -8(%rdp)
		jle .L3
