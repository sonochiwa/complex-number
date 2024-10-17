package form

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"program/operations"
	"strconv"
)

func New() {
	a := app.New()
	a.Settings().SetTheme(theme.LightTheme())
	w := a.NewWindow("Комплексные числа")

	// Создаем поля для ввода действительной и мнимой части первого числа
	real1Entry := widget.NewEntry()
	real1Entry.SetPlaceHolder("Действительная часть")
	imaginary1Entry := widget.NewEntry()
	imaginary1Entry.SetPlaceHolder("Мнимая часть")

	// Создаем поля для ввода действительной и мнимой части второго числа
	real2Entry := widget.NewEntry()
	real2Entry.SetPlaceHolder("Действительная часть")
	imaginary2Entry := widget.NewEntry()
	imaginary2Entry.SetPlaceHolder("Мнимая часть")

	// Оборачиваем поля ввода в сетку
	complex1Entry := container.NewGridWithColumns(2, real1Entry, imaginary1Entry)
	complex2Entry := container.NewGridWithColumns(2, real2Entry, imaginary2Entry)

	// Создаем текстовое поле для вывода результата
	resultLabel := widget.NewLabel("Результат:")

	// Создаем кнопки для операций
	addButton := widget.NewButton("Сложить", func() {
		c1 := parseComplexNumber(real1Entry.Text, imaginary1Entry.Text)
		c2 := parseComplexNumber(real2Entry.Text, imaginary2Entry.Text)
		resultLabel.SetText("Результат: " + c1.Add(c2).String())
	})

	subtractButton := widget.NewButton("Вычесть", func() {
		c1 := parseComplexNumber(real1Entry.Text, imaginary1Entry.Text)
		c2 := parseComplexNumber(real2Entry.Text, imaginary2Entry.Text)
		resultLabel.SetText("Результат: " + c1.Subtract(c2).String())
	})

	multiplyButton := widget.NewButton("Умножить", func() {
		c1 := parseComplexNumber(real1Entry.Text, imaginary1Entry.Text)
		c2 := parseComplexNumber(real2Entry.Text, imaginary2Entry.Text)
		resultLabel.SetText("Результат: " + c1.Multiply(c2).String())
	})

	divideButton := widget.NewButton("Разделить", func() {
		c1 := parseComplexNumber(real1Entry.Text, imaginary1Entry.Text)
		c2 := parseComplexNumber(real2Entry.Text, imaginary2Entry.Text)
		resultLabel.SetText("Результат: " + c1.Divide(c2).String())
	})

	trigonometricFormButton := widget.NewButton("Тригонометрическая форма", func() {
		c1 := parseComplexNumber(real1Entry.Text, imaginary1Entry.Text)
		//c2 := parseComplexNumber(real2Entry.Text, imaginary2Entry.Text)
		resultLabel.SetText("Результат: " + c1.ToTrigonometricForm())

	})

	ExponentialFormButton := widget.NewButton("Экспоненциальная форма", func() {
		c1 := parseComplexNumber(real1Entry.Text, imaginary1Entry.Text)
		//c2 := parseComplexNumber(real2Entry.Text, imaginary2Entry.Text)
		resultLabel.SetText("Результат: " + c1.ToExponentialForm())
	})

	// Создаем контейнер для полей ввода и кнопок
	content := container.NewVBox(
		widget.NewLabel("Первое комплексное число:"),
		complex1Entry,
		widget.NewLabel("Второе комплексное число:"),
		complex2Entry,
		widget.NewSeparator(),
		addButton,
		subtractButton,
		multiplyButton,
		divideButton,
		widget.NewLabel("Для первого числа:"),
		trigonometricFormButton,
		ExponentialFormButton,
		resultLabel,
		widget.NewSeparator(),
	)

	w.SetContent(content)
	w.Resize(fyne.NewSize(400, 400))
	w.ShowAndRun()
}

// parseComplexNumber парсит строки в действительную и мнимую части комплексного числа
func parseComplexNumber(realStr, imaginaryStr string) operations.ComplexNumber {
	realPart, _ := strconv.ParseFloat(realStr, 64)
	imaginaryPart, _ := strconv.ParseFloat(imaginaryStr, 64)
	return operations.NewComplexNumber(realPart, imaginaryPart)
}
