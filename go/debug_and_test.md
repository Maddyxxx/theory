**Для старта:** go test + Delve + Staticcheck

**Для enterprise:** GolangCI-Lint + Testify + pprof

**Для микросервисов:**  GoReplay для тестирования нагрузки

# Инструменты отладки (Debugging)
------------- 

## Delve (dlv)
------------- 
- Золотой стандарт отладки Go-программ

- лучшее применение: пошаговый анализ сложных багов

Поддержка:
- Пошагового выполнения (step-in/step-out)
- Точек останова (breakpoints)
- Анализа переменных и стека вызовов

Пример запуска:
```
bash
dlv debug ./main.go
break main.main
continue
```

## GDB (GNU Debugger)
-------------
- Работает с Go через специальный плагин
- Менее удобен, чем Delve, но полезен для низкоуровневой отладки

## pprof (профилирование)
-------------
Встроенный инструмент для анализа:
- Производительности CPU
- Использования памяти
- Блокировок горутин

- лучшее применение: оптимизация производительности

Пример:

```
import _ "net/http/pprof"
go func() { http.ListenAndServe(":6060", nil) }()
```
Анализ через:
```
bash
go tool pprof http://localhost:6060/debug/pprof/heap
```

# Инструменты тестирования
------------- 
## Встроенный testing пакет
-------------
Стандартный фреймворк для unit-тестов

Пример:
```
func TestAdd(t *testing.T) {
    if Add(2, 3) != 5 {
        t.Error("Ожидалось 5")
    }
}
```
Запуск:
```
bash
go test -v ./...
```

## Testify (stretchr/testify)
-------------
- Популярная библиотека для улучшения тестов:
- Удобные ассерты (assert.Equal(t, 5, result))
- Моки (mock.Mock)
- Наборы тестов (suites)

- лучшее применение: читаемые ассерты и моки

## GoMock (официальный мокинг)
-------------
- Генератор мок-объектов на основе интерфейсов

Пример:
```
bash
mockgen -source=repository.go > mocks/repository_mock.go
```

## Ginkgo/Gomega (BDD-стиль)
-------------
Альтернативный подход к тестированию:

```
go
Describe("Calculator", func() {
    It("should add numbers", func() {
        Expect(Add(2, 3)).To(Equal(5))
    })
})
```

#  Инструменты статического анализа
------------- 

## Staticcheck
-------------
Находит более 200 видов ошибок:
- Утечки горутин
- Неправильное использование API
- Подозрительные конструкции

- лучшее применение: поиск скрытых ошибок

Пример:

```
bash
staticcheck ./...
```

## GolangCI-Lint
-------------
- Агрегатор линтеров (объединяет 50+ инструментов)
- Поддержка CI/CD
- Конфигурация через .golangci.yml

- лучшее применение: комплексная проверка перед коммитом

## errcheck
-------------
Проверяет необработанные ошибки:
```
go
json.Unmarshal(data, &obj) // errcheck поймает отсутствие проверки ошибки
```

# Инструменты бенчмаркинга
-------------

## Встроенный бенчмаркинг
-------------
```
go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}
```
Запуск:
```
bash
go test -bench=. -benchmem
```

## pprof + benchviz
-------------
Визуализация результатов:
```
bash
go test -bench=. -cpuprofile=cpu.out
go tool pprof -svg cpu.out > cpu.svg
```

# Дополнительные инструменты
-------------

## GoReplay (тестирование нагрузки)
-------------
Запись и воспроизведение реального трафика

## Docker + Go
-------------
Отладка в контейнерах:
```
dockerfile
FROM golang:latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
```

## VS Code + Go Plugin
-------------
Интеграция всех инструментов в IDE:
- Автозавершение кода
- Встроенный дебаггер (Delve)
- Линтинг в реальном времени


### **Проектирование кода с интерфейсами для тестируемости**

**1. Основные принципы:**
- **Зависимости через интерфейсы** (не конкретные типы).
- **Инверсия зависимостей (DIP)** — модули зависят от абстракций.
- **Легкая замена реальных компонентов моками в тестах**.

---

### **2. Пример: сервис оплаты заказа**

#### **🔹 Шаг 1: Определяем интерфейс зависимости**
Допустим, у нас есть `PaymentGateway` для обработки платежей:

```go
// payment_gateway.go
type PaymentGateway interface {
    Charge(amount float64, cardToken string) (string, error)
}
```

#### **🔹 Шаг 2: Реализуем реальный и мок-сервис**

**Реальная реализация (Stripe API):**
```go
// stripe_gateway.go
type StripeGateway struct{}

func (s *StripeGateway) Charge(amount float64, cardToken string) (string, error) {
    // Реальный вызов Stripe API
    return "charge_id", nil
}
```

**Мок (для тестов):**  
Используем `testify/mock`:
```go
// mock_gateway.go
type MockGateway struct {
    mock.Mock
}

func (m *MockGateway) Charge(amount float64, cardToken string) (string, error) {
    args := m.Called(amount, cardToken) // Запись вызова
    return args.String(0), args.Error(1)
}
```

---

#### **🔹 Шаг 3: Сервис заказа с внедрением зависимости**

```go
// order_service.go
type OrderService struct {
    paymentGateway PaymentGateway // Зависимость через интерфейс
}

func (s *OrderService) ProcessOrder(amount float64, cardToken string) error {
    _, err := s.paymentGateway.Charge(amount, cardToken)
    if err != nil {
        return fmt.Errorf("payment failed: %w", err)
    }
    return nil
}
```

---

### **3. Тестирование с моками**

#### **🔹 Пример теста (testify/mock)**
```go
func TestOrderService_ProcessOrder(t *testing.T) {
    // 1. Инициализируем мок
    mockGateway := new(MockGateway)
    service := &OrderService{paymentGateway: mockGateway}

    // 2. Задаём ожидаемый вызов и результат
    mockGateway.On("Charge", 100.0, "test_token").Return("charge_123", nil)

    // 3. Вызываем тестируемый метод
    err := service.ProcessOrder(100.0, "test_token")

    // 4. Проверяем, что:
    assert.NoError(t, err)                          // Нет ошибки
    mockGateway.AssertExpectations(t)               // Метод вызван с нужными аргументами
}
```

#### **🔹 Пример с `gomock`**
1. Генерируем мок:
   ```bash
   mockgen -source=payment_gateway.go -destination=mocks/payment_gateway_mock.go
   ```
2. Используем в тесте:
   ```go
   func TestProcessOrder(t *testing.T) {
       ctrl := gomock.NewController(t)
       defer ctrl.Finish()

       mockGateway := NewMockPaymentGateway(ctrl)
       mockGateway.EXPECT().
           Charge(100.0, "test_token").
           Return("charge_123", nil)

       service := &OrderService{paymentGateway: mockGateway}
       err := service.ProcessOrder(100.0, "test_token")
       assert.NoError(t, err)
   }
   ```

---

### **4. Выгоды такого подхода**

✅ **Изоляция тестов**:
- Моки заменяют БД, API и другие внешние сервисы.
- Тесты не зависят от сети/сторонних сервисов.

✅ **Гибкость**:
- Можно подменить реализацию `PaymentGateway` без изменения `OrderService`.

✅ **Проверка взаимодействий**:
- Убеждаемся, что методы вызываются с правильными аргументами.

---

### **5. Где применять?**
- **Внешние API** (платежи, SMS, email).
- **Базы данных** (мокируем `sql.DB`).
- **Сложные вычисления** (заменяем на стабы).

**Итог**: Интерфейсы + моки = **чистый, тестируемый и гибкий код**. 🚀