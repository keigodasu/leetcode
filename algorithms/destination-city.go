func destCity(paths [][]string) string {
    nextPoint := paths[0][1]
    lastPoint := nextPoint
    pathsRest := paths[1:]
    
    
    for  {
        nextPoint = findRoot(pathsRest, nextPoint) 
        if nextPoint == "" {
            break
        } else {
            lastPoint = nextPoint 
        }
    } 
    
    return lastPoint
}

func findRoot(paths [][]string, target string) string {
    for _, v := range paths {
        if target == v[0] {
            return v[1]
        }
    }
    
    return ""
}
