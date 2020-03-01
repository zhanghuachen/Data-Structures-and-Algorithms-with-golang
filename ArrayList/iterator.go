package ArrayList

import "errors"

type Iterator interface {
	HasNext() bool
	Next()(interface{},error)
	Remove()
	GetIndex() int
}

type Iterable interface {
	Iterator() Iterator
}

type ArrayListIterator struct {
	list *ArrayList
	CurrentIndex int
}

func (list *ArrayList)Iterator() Iterator{
	iterator := new(ArrayListIterator)
	iterator.list = list
	iterator.CurrentIndex = 0
	return iterator
}

func (iterator *ArrayListIterator)HasNext()bool{
	return iterator.CurrentIndex < iterator.list.TheSize
}

func (iterator *ArrayListIterator)Next()(interface{},error){
	if iterator.CurrentIndex >= iterator.list.TheSize {
		return nil, errors.New("overflow")
	}
	value, err := iterator.list.Get(iterator.CurrentIndex)
	iterator.CurrentIndex++
	return value,err
}

func (iterator *ArrayListIterator)Remove(){
	iterator.list.Delete(iterator.CurrentIndex)
	iterator.CurrentIndex--
}
func (iterator *ArrayListIterator)GetIndex() int{
	return iterator.CurrentIndex
}