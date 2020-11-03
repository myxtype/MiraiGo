package tlv

import "github.com/myxtype/MiraiGo/binary"

func T108(arr []byte) []byte {
	return binary.NewWriterF(func(w *binary.Writer) {
		w.WriteUInt16(0x108)
		w.WriteTlv(arr)
	})
}
