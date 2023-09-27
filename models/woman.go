package models

type Woman struct {
	Man
}

func (w *Woman) Breathe()    { w.breathing = true }
func (w *Woman) Eat()        { w.eating = true }
func (w *Woman) Think()      { w.thinking = true }
func (w *Woman) Sex() string { return "Woman" }
