package model


//种族值
type Race_Value struct{
	Hp int
	Ack int
	Def int
	S_ack int
	S_def int
	Speed int
	Total int
}

//经验
type Exp struct {
	Exp_lever int
	Init_exp int
	Exp int
}

type Base struct {
	Name string
	Pro []string
	Height int
	Weight int
	Rarity string
	Catch_rate int
	Introduce string
}


type Skill struct{
	Id int
	Name string
	Categoty string
	Pro []string
	Ways string
	Ack int
	PP int
	Init_persent int
	Target_num int
	Introduce int
}


type  Elevs  struct {
    Id int64
    Exp *Exp
    Race_value *Race_Value
    Base *Base
    Skill []*Skill
}

func NewElevs() *Elevs{
	return &Elevs{
	  Id : 0,
	  Exp : nil,
	  Race_value: nil,
	  Base: nil,
	  Skill: nil,
	}
}
