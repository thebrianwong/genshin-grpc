package utils

import pb_common "genshin-grpc/proto/common"

func ElementStringToEnum(element string) pb_common.Element {
	switch element {
	case "Anemo":
		return pb_common.Element_ANEMO
	case "Geo":
		return pb_common.Element_GEO
	case "Electro":
		return pb_common.Element_ELECTRO
	case "Dendro":
		return pb_common.Element_DENDRO
	case "Hydro":
		return pb_common.Element_HYDRO
	case "Pyro":
		return pb_common.Element_PYRO
	case "Cryo":
		return pb_common.Element_CRYO
	default:
		return pb_common.Element_ANEMO
	}
}
