package form

import (
	"fmt"
	"strings"
)

type FormItems []*FormItem

// FormItem wraps item and its dependencies
type FormItem struct {
	item     item
	parent   *FormItem
	children FormItems
}

func (f *FormItem) AddItems(c ...*FormItem) *FormItem {
	for _, item := range c {
		f.AddItem(item)
	}
	return f
}

func (f *FormItem) setTextRecursive() {
	for _, i := range f.children {
		i.setText()
		if len(i.children) > 0 {
			i.setTextRecursive()
		}
	}
}

func (f *FormItem) setText() {

	if f.parent == nil {
		return
	}
	p := f.parent
	parentsCount := 0
	for {
		if p.parent != nil {
			parentsCount++
			p = p.parent
			continue
		}
		break
	}

	f.item.setPrefix(fmt.Sprintf("%s╰─", strings.Repeat("  ", parentsCount)))
}

// AddItem adds a subItem i dependant of the FormItem f
// The rules applied to display the subItem are specific to
// each FormItem
func (f *FormItem) AddItem(formItem *FormItem) *FormItem {
	formItem.parent = f
	f.children = append(f.children, formItem)
	return formItem
}

func NewFormItem(i item) *FormItem {
	return &FormItem{
		item: i,
	}
}

func (f FormItems) visibleItems() []item {
	items := make([]item, 0)
	for _, v := range f {
		items = append(items, v.item)
		if v.children != nil && v.item.displayChildren() {
			items = append(items, v.children.visibleItems()...)
		}
	}
	return items
}

// clearHiddenItemsValues range over all items and subItems and reset the value
// of the hidden ones
func (f FormItems) clearHiddenItemsValues() {
	for _, formItem := range f {
		if formItem.parent != nil && !formItem.parent.item.displayChildren() {
			formItem.item.clearValue()
		}
		if formItem.children != nil {
			formItem.children.clearHiddenItemsValues()
		}
	}
}

func (f *FormItem) Answer() interface{} {
	return f.item.answer()
}
