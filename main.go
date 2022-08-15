package main

import (
	"go_tutorial/classType"
	"go_tutorial/constant"
	"go_tutorial/defers"
	"go_tutorial/functions"
	"go_tutorial/inherit"
	"go_tutorial/interfaceType"
	mylib "go_tutorial/lib1" // 别名
	_ "go_tutorial/lib2"     // 匿名倒入
	"go_tutorial/mapType"
	"go_tutorial/pairType"
	"go_tutorial/pointers"
	"go_tutorial/polymorphism"
	"go_tutorial/slice"
	"go_tutorial/structType"
	"go_tutorial/variable"
)

func main() {
	mylib.Lib1Test()
	// lib2.Lib2Test()
	variable.TestVar()
	constant.TestConst()
	functions.TestFuncs()
	pointers.TestPointers()
	defers.TestDefers()
	slice.TestSlice()
	mapType.TestMap()
	structType.TestStruct()
	classType.TestClass()
	inherit.TestInherit()
	polymorphism.TestPolymorphism()
	interfaceType.TestInterFace()
	pairType.TestPairType()
}
