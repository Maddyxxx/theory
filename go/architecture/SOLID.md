SOLID — это **5 принципов объектно-ориентированного проектирования**, которые делают код гибким, понятным и поддерживаемым. Хотя Go не является чисто ООП-языком, эти принципы применимы и к нему. Разберём каждый принцип с примерами на Go.

---

## 1. **Принцип единственной ответственности (SRP)**
**«Один класс — одна ответственность»**  
В Go: **один тип/функция/пакет — одна зона ответственности**.

### 🚨 Проблема:
```go
// Плохо: UserService делает слишком много
type UserService struct{}

func (us *UserService) CreateUser() { /* ... */ }
func (us *UserService) SendEmail() { /* ... */ }  // Отправка почты — это другая ответственность!
```

### ✅ Решение:
```go
// Хорошо: разделили логику
type UserCreator struct{}
func (uc *UserCreator) Create() { /* ... */ }

type EmailSender struct{}
func (es *EmailSender) Send() { /* ... */ }
```

---

## 2. **Принцип открытости/закрытости (OCP)**
**«Программные сущности должны быть открыты для расширения, но закрыты для изменения»**.

### 🚨 Проблема:
```go
func CalculateArea(shape interface{}) float64 {
    switch s := shape.(type) {
    case Circle:    return 3.14 * s.Radius * s.Radius
    case Rectangle: return s.Width * s.Height
    // Добавление новой фигуры требует изменения функции
    }
}
```

### ✅ Решение (используем интерфейсы):
```go
type Shape interface {
    Area() float64
}

type Circle struct{ Radius float64 }
func (c Circle) Area() float64 { return 3.14 * c.Radius * c.Radius }

// Теперь можно добавлять новые фигуры без изменения CalculateArea
func TotalArea(shapes []Shape) float64 {
    total := 0.0
    for _, s := range shapes {
        total += s.Area()  // Вызов метода через интерфейс
    }
    return total
}
```

---

## 3. **Принцип подстановки Барбары Лисков (LSP)**
**«Объекты должны быть заменяемы экземплярами своих подтипов без изменения корректности программы»**.

### 🚨 Проблема:
```go
type Bird interface { Fly() }
type Duck struct{}
func (d Duck) Fly() {}

type Ostrich struct{} // Страус не летает!
func (o Ostrich) Fly() { panic("I can't fly!") }  // Нарушает LSP
```

### ✅ Решение (разделяем интерфейсы):
```go
type Bird interface { Walk() }
type FlyingBird interface {
    Bird
    Fly()
}

// Теперь Ostrich реализует только Bird, а Duck — FlyingBird
```

---

## 4. **Принцип разделения интерфейсов (ISP)**
**«Много специализированных интерфейсов лучше одного общего»**.

### 🚨 Проблема:
```go
// Плохо: один интерфейс на все случаи
type Worker interface {
    Work()
    Eat()
    Sleep()
}

// Программист должен реализовать все методы, даже если не использует Sleep()
```

### ✅ Решение:
```go
type Workable interface { Work() }
type Eatable  interface { Eat() }

type Programmer struct{}
func (p Programmer) Work() {} // Реализует только нужные интерфейсы
```

---

## 5. **Принцип инверсии зависимостей (DIP)**
**«Зависимости должны строиться на абстракциях, а не на деталях»**.

### 🚨 Проблема:
```go
// Плохо: высокоуровневый модуль зависит от низкоуровневого
type MySQLDatabase struct{}
func (db MySQLDatabase) Query() {}

type Service struct {
    db MySQLDatabase  // Жёсткая привязка к MySQL
}
```

### ✅ Решение (инъекция зависимости через интерфейс):
```go
type Database interface {
    Query()
}

type Service struct {
    db Database  // Принимает любую БД, реализующую интерфейс
}

// Теперь можно подменить MySQL на PostgreSQL без изменения Service
```

---

## 🔹 **Как применять SOLID в Go?**
1. **Интерфейсы** — основа многих принципов (OCP, ISP, DIP).
2. **Композиция вместо наследования** — Go поощряет это.
3. **Маленькие пакеты** — помогают соблюдать SRP.

Пример из стандартной библиотеки:
- `io.Reader`/`io.Writer` — идеальное соблюдение ISP и DIP.

---

## ⚖️ **Когда нарушать SOLID?**
- В **небольших проектах** — если переусложняет код.
- Для **высокой производительности** — иногда абстракции влияют на скорость.
- В **экспериментальном коде** — можно сначала сделать, потом рефакторить.

SOLID в Go — это не догма, а **гибкие правила** для создания чистого кода.