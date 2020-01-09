package main

import (
	"bufio"
	"bytes"
	"log"
	"strconv"
	"strings"
)

func open() map[string]Entry {
	file := getDict()

	dictionary := make(map[string]Entry)

	scanner := bufio.NewScanner(bytes.NewReader(file))

	scanner.Scan() // skipe first line

	for scanner.Scan() {
		firstLine := strings.Split(scanner.Text(), "|")
		word := firstLine[0]
		lines, _ := strconv.Atoi(firstLine[1])

		entry := Entry{Word: word}

		for i := 1; i <= lines; i++ {
			scanner.Scan()
			args := strings.Split(scanner.Text(), "|")

			category := args[0]
			category = strings.TrimLeft(category, "(")
			category = strings.TrimRight(category, ")")

			synonyms := args[1:]

			thisWord := Word{Category: category, Synonyms: synonyms}
			entry.Entries = append(entry.Entries, thisWord)
		}

		dictionary[word] = entry
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return dictionary
}
