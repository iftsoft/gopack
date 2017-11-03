package lla

type Param struct {
	Field	string
	Value	interface{}
}
type ParamList []Param

type ClauseType 	int8
type ClauseStat 	int8
type SortOrder		int8

const (
	Type_Statement	ClauseType	= iota
	Type_Group_AND
	Type_Group_OR
)

const (
	Stat_None	ClauseStat	= iota
	Stat_Equal
	Stat_NotEqual
	Stat_Grater
	Stat_GrEqual
	Stat_Later
	Stat_LtEqual
	Stat_IsNull
	Stat_IsNotNull
	Stat_InList
	Stat_NotInList
)

const (
	Sort_Undef	SortOrder	= iota
	Sort_Asc
	Sort_Desc
)

type Clause struct {
	Type	ClauseType
	Stat	ClauseStat
	Field	string
	Value	interface{}
	Extra	interface{}
}
type ClauseList []Clause

type Order struct {
	Field   string
	Sort	SortOrder
}
type OrderList []Order

type Conditions struct {
	Filter	*ClauseList
	Order	*OrderList
	Limit	uint32
	Offset	uint32
	Result	int64
}

// Interfaces to access object in store
//
type ObjectKeeper interface {
	Select(unit interface{}, keys ParamList) error
	Update(unit interface{}) error
	Insert(unit interface{}) error
	Delete(keys ParamList) (int64, error)
}

type ObjectFinder interface {
	Estimate(cond *Conditions) error
	Lookup(list interface{}, cond *Conditions) error
	Search(list interface{}, cond *Conditions) error
}

type StoreKeeper interface {
	Begin() error
	Commit()
	Rollback()
	GetObjectKeeper(name string) ObjectKeeper
	GetObjectFinder(name string) ObjectFinder
}

func (this ParamList) GetParamValue(name string) interface{} {
	if this == nil {
		return nil
	}
	for _, par := range this {
		if par.Field == name {
			return par.Value
		}
	}
	return nil
}

