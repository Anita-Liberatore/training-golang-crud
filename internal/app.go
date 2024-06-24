package main

import "fmt"

func main() {
	fmt.Print("Hello")
}

//Questo programma ha l'obiettivo di stampare la stringa "Hello" sullo schermo.
//Ecco una spiegazione dettagliata di ogni parte del codice:

//package main: Questa linea definisce il pacchetto principale del programma.
//In Go, ogni programma deve appartenere a un pacchetto,
//e main è il pacchetto speciale che definisce un'applicazione eseguibile.
//Quando un programma viene eseguito, il runtime di Go cerca la funzione main all'interno di questo pacchetto.

//import "fmt": Questa linea importa il pacchetto fmt, che sta per "format". Il pacchetto fmt
//fornisce funzioni di input/output formattato, come la stampa su console.

//func main() { ... }: Questa è la definizione della funzione main, che è il punto di
//ingresso del programma. Ogni programma Go deve avere una funzione main all'interno del
//pacchetto main. Quando il programma viene eseguito, l'esecuzione inizia da questa funzione.

//fmt.Print("Hello"): All'interno della funzione main, questa linea chiama la funzione Print del
//pacchetto fmt, che stampa la stringa "Hello" sulla console senza un carattere di nuova linea alla fine.

//In sintesi, questo programma importa il pacchetto fmt, definisce una funzione main, e all'interno di
//essa utilizza fmt.Print per stampare "Hello" sullo schermo.
