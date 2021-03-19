func canVisitAllRooms(rooms [][]int) bool {
    visited := map[int]bool{}
    canVisitAllRoomsDFS(rooms,visited,0)
    return len(visited)==len(rooms)
}

func canVisitAllRoomsDFS(rooms [][]int, visited map[int]bool, current int){
    if _, exists := visited[current]; exists{
        return
    }
    
    visited[current]=true
    
    for _, key := range rooms[current]{
        canVisitAllRoomsDFS(rooms,visited,key)
    }
}
