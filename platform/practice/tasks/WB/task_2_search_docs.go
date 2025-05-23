package WB

/*
Необходимо реализовать конкурентный поиск документов на серверах.
Для этого у нас есть сторонняя библиотека с функцией

search.Search(server string, query string) ([]string, error)

которая осуществляет поиск документов на указанном сервере по указанному запросу.
У нас есть N идентичных серверов (реплики) и задача состоит в том, чтобы конкурентно
вызвать эту функцию для всех серверов и вернуть первый успешный ответ от любого из серверов
не дожидаясь ответов от других. Если какой-то сервер возвращает ошибку, то мы ее игнорируем,
дожидаясь успешного ответа от других, но если все серверы ответили с ошибкой, то наша
функция должна вернуть ошибку, что поиск не удался.

func Search(servers []string, query string) ([]string, error) {

}
*/

//func Search(servers []string, query string) ([]string, error) {
//	var wg sync.WaitGroup
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()                    // Гарантирует освобождение ресурсов
//	success := make(chan []string, 1) // Буфер на 1, чтобы не блокировать горутины
//
//	for _, server := range servers {
//		wg.Add(1)
//		go func(server string) { // передача server в go функцию избегает race condition в горутинах и
//			// позволяет получать последнее значение server из цикла.
//			defer wg.Done()
//			result, err := search.Search(ctx, server, query) // если search.Search() поддерживает отмену через контекст,
//			// это позволит остановить лишние запросы после первого успешного
//
//			if err == nil {
//				select {
//				case success <- result: // Успешная отправка
//					cancel() // Отменяем другие запросы даже если search.Search() не поддерживает контекст
//				default: // Если канал уже занят — игнорируем
//				}
//			}
//		}(server)
//	}
//
//	// Ждём первый успешный ответ или завершение всех горутин
//	go func() {
//		wg.Wait()
//		close(success) // Закрываем канал, если все горутины завершились
//	}()
//
//	if result, ok := <-success; ok {
//		return result, nil
//	}
//	return nil, errors.New("поиск не удался")
//}
