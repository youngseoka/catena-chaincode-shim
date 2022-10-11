package data

var (
	// Pre created btcec.PrivateKey.Serialized() for testing.
	TestingBtcecPrivKey1 = []byte{11, 247, 40, 42, 13, 67, 85, 36, 115, 14, 39, 76, 70, 183, 171, 220, 40, 204, 225, 227, 91, 151, 22, 95, 55, 216, 162, 206, 237, 123, 95, 5}
	TestingBtcecPrivKey2 = []byte{216, 58, 139, 98, 7, 249, 131, 56, 208, 12, 244, 18, 191, 246, 41, 171, 71, 94, 26, 67, 38, 158, 5, 10, 253, 124, 245, 8, 102, 200, 251, 128}
	TestingBtcecPrivKey3 = []byte{42, 42, 118, 250, 16, 115, 36, 29, 62, 52, 139, 153, 204, 36, 238, 15, 254, 4, 172, 18, 159, 134, 66, 230, 76, 222, 89, 213, 3, 153, 102, 252}
)

var (
	ExpectedBtcecPrivKey1Addr = []byte{212, 39, 203, 85, 138, 84, 103, 113, 135, 142, 177, 213, 23, 22, 164, 6, 57, 16, 226, 49}
)

var (
	ExpectedBtcecPrivKey1DID = "did:white:3xRv21hWkYpKE7gaCNRgbiiK5xye"
)
