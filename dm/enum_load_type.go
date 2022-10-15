package dm

type LoadType string

const (
	LoadTypeBulk   LoadType = "Bulk"
	LoadTypePallet LoadType = "Pallet"
	LoadTypeULD    LoadType = "ULD"
)
