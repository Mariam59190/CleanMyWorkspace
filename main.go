ackage main

import (
    "fmt"
    "github.com/Mentor-Paris/CleanMyWorkspace"
    "CleanMyWorkspace/Clean"
)

func main() {
    // Génération du souk
    workspace := CleanMyWorkspace.GenererateWorkSpace()

    fmt.Println("----- Espace de travail sale -----")
    afficherSouk(workspace)

    // Nettoyage
    workspacePropre := Clean.CleanWorkSpace(workspace)

    fmt.Println("\n----- Espace de travail propre -----")
    afficherSouk(workspacePropre)
}

// Fonction utilitaire pour afficher le souk
func afficherSouk(workspace *[][]*string) {
    for _, row := range *workspace {
        for _, item := range row {
            if item == nil {
                fmt.Print("|nil|")
            } else {
                fmt.Printf("|%s|", *item)
            }
        }
        fmt.Println()
    }
}

