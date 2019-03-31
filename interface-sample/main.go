package main

// Hoge
type Hoge interface {
	getHoge() string
	getHoge2() string
}

type hoge struct {
}

func NewHoge() Hoge {
	return &hoge{}
}

func (h *hoge) getHoge() string {
	return "Hoge: hoge"
}

func (h *hoge) getHoge2() string {
	return "Hoge: hoge2"
}

// Fuga
type Fuga interface {
	getFuga() string
	getFuga2() string
}

type fuga struct {
}

func NewFuge() Fuga {
	return &fuga{}
}

func (f *fuga) getFuga() string {
	return "Fuga: fuga"
}

func (f *fuga) getFuga2() string {
	return "Fuga: fuga2"
}

// HogeFuga
// getHoge, getHoge2, getFuga, getFuga2の実装を必要とする
// 以下のように書く必要がない
// type Hogefuga interface {
//   getHoge() string
//   getHoge2() string
//   getFuga() string
//   getFuga2() string
// }
type Hogefuga interface {
	Hoge
	Fuga
}

type hogefuga struct {
}

func NewHogefuga() Hogefuga {
	return &hogefuga{}
}

func (h *hogefuga) getHoge() string {
	return "Hogefuga: hoge"
}

func (h *hogefuga) getHoge2() string {
	return "Hogefuga: hoge2"
}

func (h *hogefuga) getFuga() string {
	return "Hogefuga: fuga"
}

func (h *hogefuga) getFuga2() string {
	return "Hogefuga: fuga2"
}

type hogefuga2 struct {
	Hoge
	Fuga
}

func (hf2 *hogefuga2) getHoge() string {
	return "Hogefuga2: hoge"
}

func (hf2 *hogefuga2) getFuga2() string {
	return "Hogefuga2: fuga2"
}

// Hogefugaインターフェースに依存した関数
func printHoge(hf Hogefuga) {
	println(hf.getHoge())
}

func printHoge2(hf Hogefuga) {
	println(hf.getHoge2())
}

func printFuga(hf Hogefuga) {
	println(hf.getFuga())
}

func printFuga2(hf Hogefuga) {
	println(hf.getFuga2())
}

func main() {
	h := NewHoge()
	println(h.getHoge())
	println(h.getHoge2())

	f := NewFuge()
	println(f.getFuga())
	println(f.getFuga2())

	hf := NewHogefuga()
	printHoge(hf)
	printHoge2(hf)
	printFuga(hf)
	printFuga2(hf)

	hf2 := &hogefuga2{}
	// hogefuga2はgetFuga等を実装していない
	// 以下の２つだけをテストしたい場合などのモックであればそれでいい
	printHoge(hf2)
	printFuga2(hf2)
	// 以下はコンパイルは通るが実行するとpanicになる
	//printHoge2(hf2)
	//printFuga(hf2)
}
