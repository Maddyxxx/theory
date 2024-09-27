***Чтание EXPLAIN***
<ln>
<li>
Начинаем читать каждую строчку сверху вниз. Смотрим на колонку type. Если индекс не используется — плохо (за исключением случаев, когда таблица очень маленькая или присутствует ключевое слово LIMIT). 
В этом случае оптимизатор намеренно предпочтет просканировать таблицу. Чем ближе значение столбца type к NULL (см. пункт о столбце type), тем лучше.
Далее стоит посмотреть на колонки rows и filtered. Чем меньше значение rows  и чем больше значение filtered,- тем лучше. 
  Однако, если значение rows слишком велико и filtered стремится к 100 %  - это очень плохо.</li>

<li>Смотрим, какой индекс был выбран из колонки key , и сравниваем со всеми ключами из possible_keys. Если индекс не оптимальный (большая селективность), 
  то стоит подумать, как изменить запрос или пробросить дополнительные данные в условие выборки, чтобы использовать наилучший индекс из possible_keys.</li>

<li>Наконец, читаем колонку Extra. Если там значение, отмеченное выше как (!!!), то, как минимум, обращаем на это вниманием. 
  Как максимум, пытаемся разобраться, почему так. В этом нам может хорошо помочь SHOW WARNINGS.</li>

<li>Переходим к следующей строке и повторяем всё заново.</li>

<li>Если не лень, то в конце перемножаем все значения в столбце rows всех строк, чтобы грубо оценить количество просматриваемых строк.</li>

<li>
  <ln>
    <li>>При чтении всегда помним о том, что:</li>

<li>EXPLAIN ничего не расскажет о триггерах и функциях (в том числе определенных пользователем), участвующих в запросе.</li>

<li>EXPLAIN не работает с хранимыми процедурами.</li>

<li>EXPLAIN не расскажет об оптимизациях, которые MySQL производит уже на этапе выполнения запроса.</li>

<li>Большинство статистической информации — всего лишь оценка, иногда очень неточная.</li>

<li>EXPLAIN не делает различий между некоторыми операциями, называя их одинаково. 
  Например, filesort может означать сортировку в памяти и на диске, а временная таблица, 
  которая создается на диске или в памяти, будет помечена как Using temporary.</li>

<li>В разных версиях MySQL EXPLAIN может выдавать совершенно разные результаты, потому что оптимизатор постоянно улучшается разработчиками, 
  поэтому не забываем обновляться.</li></ln></li></ln>