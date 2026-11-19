func reverseBits(n int) int {
	num := uint32(n)
	var res uint32 = 0
	for i := 0; i < 32; i++ {
		bit := (num >> i) & 1
		res |= (bit << (31 - i))
	}
	return int(res)
}
