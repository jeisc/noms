package types

import (
	. "github.com/attic-labs/noms/dbg"
	"github.com/attic-labs/noms/ref"
)

type cachedRef ref.Ref

func (cr *cachedRef) Ref(v Value) ref.Ref {
	if ref.Ref(*cr) == (ref.Ref{}) {
		Chk.NotNil(Reffer)
		*cr = cachedRef(Reffer(v))
	}
	return ref.Ref(*cr)
}