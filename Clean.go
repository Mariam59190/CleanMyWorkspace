package Clean

func CleanWorkSpace(workspace *[][]*string) *[][]*string {
    for i := range *workspace {
        for j := range (*workspace)[i] {
            (*workspace)[i][j] = nil
        }
    }
    return workspace
}

