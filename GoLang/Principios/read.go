package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

// Definir la estructura Nombre con campos de tamaño 20
type Nombre struct {
    Fname string
    Lname string
}

func main() {
    // Pedir al usuario el nombre del archivo
    fmt.Print("Ingrese el nombre del archivo: ")
    var fileName string
    fmt.Scanln(&fileName)

    // Abrir el archivo
    file, err := os.Open(fileName)
    if err != nil {
        fmt.Println("Error al abrir el archivo:", err)
        return
    }
    defer file.Close()

    // Crear una slice para almacenar los structs
    var nombres []Nombre

    // Leer el archivo línea por línea
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        // Leer y dividir la línea en nombre y apellido
        line := scanner.Text()
        parts := strings.Fields(line)
        if len(parts) != 2 {
            fmt.Println("Línea no válida:", line)
            continue
        }

        // Crear un struct Nombre
        nombre := Nombre{
            Fname: parts[0],
            Lname: parts[1],
        }

        // Asegurar que los campos tengan un tamaño máximo de 20 caracteres
        if len(nombre.Fname) > 20 {
            nombre.Fname = nombre.Fname[:20]
        }
        if len(nombre.Lname) > 20 {
            nombre.Lname = nombre.Lname[:20]
        }

        // Añadir el struct a la slice
        nombres = append(nombres, nombre)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error al leer el archivo:", err)
        return
    }

    // Iterar sobre la slice de structs e imprimir los nombres
    fmt.Println("\nContenido de la slice de structs:")
    for _, nombre := range nombres {
        fmt.Printf("Nombre: %s, Apellido: %s\n", nombre.Fname, nombre.Lname)
    }
}
