package ozon

import (
	"fmt"
  )  

//Мы в Авито любим проводить соревнования, — недавно мы устроили чемпионат по шагам. И вот настало время подводить итоги!
//Необходимо определить userIds участников, которые прошли наибольшее количество шагов steps за все дни, не пропустив ни одного дня соревнований.

// Пример 1
//# ввод
//statistics = [
//        [{ userId: 1, steps: 1000 }, { userId: 2, steps: 1500 }],
//        [{ userId: 2, steps: 1000 }]
//]

//# вывод
//champions = { userIds: [2], steps: 2500 }

//# Пример 2
//statistics = [
//        [{ userId: 1, steps: 2000 }, { userId: 2, steps: 1500 }],
//        [{ userId: 2, steps: 4000 }, { userId: 1, steps: 3500 }]
//]

//# вывод
//champions = { userIds: [1, 2], steps: 5500 }


type UserStatistic struct {
  userId int
  steps  int
}

func findChampions(statistics [][]UserStatistic) (champions map[int]int, maxSteps int) {
  // Карта для подсчёта шагов по userId
  userSteps := make(map[int]int)
  // Множество участников, которые присутствовали каждый день
  userDays := make(map[int]int)

  totalDays := len(statistics)

  // Обрабатываем данные по каждому дню
  for _, dailyStats := range statistics {
    // Временная карта для отслеживания пользователей в этом дне
    seenToday := make(map[int]bool)
    for _, stat := range dailyStats {
      // Считаем шаги
      userSteps[stat.userId] += stat.steps
      // Отмечаем, что пользователь присутствовал сегодня
      seenToday[stat.userId] = true
    }
    // Обновляем количество дней, в которые пользователь участвовал
    for userId := range seenToday {
      userDays[userId]++
    }
  }

  // Фильтруем пользователей, чтобы оставить только тех, кто был на всех днях
  validUsers := make(map[int]int)
  for userId, days := range userDays {
    if days == totalDays {
      validUsers[userId] = userSteps[userId]
    }
  }

  // Найдём максимум шагов и определим участников с этим количеством шагов
  maxSteps = 0
  champions = make(map[int]int)
  for userId, steps := range validUsers {
    if steps > maxSteps {
      maxSteps = steps
      champions = map[int]int{userId: steps}
    } else if steps == maxSteps {
      champions[userId] = steps
    }
  }

  return champions, maxSteps
}

func T2() {
  statistics := [][]UserStatistic{
    {{userId: 1, steps: 2000}, {userId: 2, steps: 1500}},
    {{userId: 2, steps: 4000}, {userId: 1, steps: 3500}},
  }

  champions, maxSteps := findChampions(statistics)

  fmt.Printf("Champions: userIds = ")
  for userId := range champions {
    fmt.Printf("%d ", userId)
  }
  fmt.Printf(", steps = %d\n", maxSteps)
}