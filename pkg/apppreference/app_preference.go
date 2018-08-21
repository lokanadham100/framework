package apppreference

type appPreference interface{
	Get() interface{}
	GetString() string
	GetBool() bool
	GetInt() int64
	GetFloat() float64
	GetSliceInt() []int64
	GetSliceString() []string	
	GetSliceBool() []bool
	GetSliceFloat() []float64
	GetMapIntInt() map[int64]int64
	GetMapStringString() map[string]string
	GetMapIntString() map[int64]string
	GetMapStringInt() map[string]int64
	GetMap
}