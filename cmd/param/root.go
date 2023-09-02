package param

type ParamCmd struct {
	Create CreateCmd `cmd:"create"`
	Read   ReadCmd   `cmd:"read"`
	Sync   ReadCmd   `cmd:"sync"`
	Delete ReadCmd   `cmd:"delete"`
}
