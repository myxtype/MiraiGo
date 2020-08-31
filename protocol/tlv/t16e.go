package tlv

import "github.com/myxtype/MiraiGo/binary"

func T16E(buildModel []byte) []byte {
	return binary.NewWriterF(func(w *binary.Writer) {
		w.WriteUInt16(0x16e)
		w.WriteTlv(binary.NewWriterF(func(w *binary.Writer) {
			w.Write(buildModel)
		}))
	})
}
