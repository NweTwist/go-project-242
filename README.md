# Анализатор размера диска (Go)

[![hexlet-check](https://github.com/NweTwist/go-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/NweTwist/go-project-242/actions)

Консольная утилита, которая считает размер файла или директории с флагами рекурсивного обхода, человекочитаемого формата и учёта скрытых файлов.

Учебный проект Хекслета: https://ru.hexlet.io/programs/go


## Стек

- Go

## Установка

```bash
git clone https://github.com/NweTwist/go-project-242.git
cd go-project-242
make build
```

Исполняемый файл будет создан в `bin/hexlet-path-size`.

## Использование

```bash
./bin/hexlet-path-size testdata/test.txt
./bin/hexlet-path-size -r testdata
./bin/hexlet-path-size -H testdata/test.txt
./bin/hexlet-path-size -a testdata
./bin/hexlet-path-size -r -H -a testdata
```

Флаги:

- `-r` — рекурсивно подсчитать размер директорий
- `-H` — вывести размер в человекочитаемом формате
- `-a` — учитывать скрытые файлы и директории
- `-h` — показать справку

Результат выводится в виде:

```text
2048B	testdata/test.txt
```

Установка и запуск в терминале показаны в записи [asciinema](https://asciinema.org/a/7RC8RuzGvyhmIWYS).

---

<details>
<summary>Автоматические тесты Хекслета</summary>

Тесты запускаются на каждый коммит. За запуск отвечает файл `.github/workflows/hexlet-check.yml` — не удаляйте и не переименовывайте ни его, ни репозиторий.

</details>

## О Хекслете

[Хекслет](https://ru.hexlet.io/) — школа программирования: авторские программы обучения с практикой, поддержкой наставников и реальными проектами, которые остаются в резюме. Этот репозиторий — один из таких проектов.
