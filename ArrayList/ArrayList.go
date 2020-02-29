package ArrayList

import (
	"errors"
	"fmt"
)

type List interface {
	Size() int
	Get(index int)(interface{},error)
	modify(index int,newVal interface{})error
	Inset(index int,newVal interface{})error
	Append(newVal interface{})
	Clear()
	Delete(index int)
	Print() string
}

type ArrayList struct {
	Data []interface{}
	TheSize int
}

func NewArrayList() *ArrayList{
	list := new(ArrayList)
	list.Data = make([]interface{}, 0 , 10)
	list.TheSize = 0
	return list
}

func (list *ArrayList)Size()int{
	return list.Size()
}

func (list *ArrayList)Get(index int)(interface{},error){
	if index < 0 || index >=list.TheSize {
		return nil, errors.New("overflow")
	}
	return list.Data[index], nil
}

func (list *ArrayList)modify(index int,newVal interface{})error{
	if index < 0 || index >=list.TheSize {
		return  errors.New("overflow")
	}
	list.Data[index] = newVal
	return nil
}

func (list *ArrayList)Insert(index int,newVal interface{})error{
	if index < 0 || index >= list.TheSize {
		return errors.New("overflow")
	}

	if list.TheSize == cap(list.Data){
		newData := make([]interface{},2 * list.TheSize, 2 * list.TheSize)
		for i:=0; i < len(list.Data) ;i++  {
			newData[i] = list.Data[i]
		}
		list.Data = newData
	}
	list.Data = list.Data[:list.TheSize+1]
	for i:= list.TheSize;i > index ;i--  {
		list.Data[i] = list.Data[i-1]
	}
	list.Data[index] = newVal
	list.TheSize++

	return nil
}

func (list *ArrayList)Append(newVal interface{}){
	list.Data = append(list.Data,newVal)
	list.TheSize++
}

func (list *ArrayList)Clear(){
	list.Data = make([]interface{}, 0, 10)
	list.TheSize = 0
}

func (list *ArrayList)Delete(index int){
	list.Data = append(list.Data[:index],list.Data[index+1:]...)
	list.TheSize--
}

func (list *ArrayList)Print() string{
	return fmt.Sprint(list.Data)
}














