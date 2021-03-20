type inStation struct {
    time int
    station string
}

type UndergroundSystem struct {
    idToInStation map[int]inStation
    allStations map[string]map[string][]int
}

func Constructor() UndergroundSystem {
    return UndergroundSystem{
        idToInStation: make(map[int]inStation),
        allStations: make(map[string]map[string][]int),
    }
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int)  {
    this.idToInStation[id] = inStation{t, stationName}
}


func (this *UndergroundSystem) CheckOut(id int, stationName string, t int)  {
    
    cIn := this.idToInStation[id]
    
    if _, ok := this.allStations[stationName]; !ok {
        this.allStations[stationName] = make(map[string][]int)
    }
    
    this.allStations[stationName][cIn.station] = append(this.allStations[stationName][cIn.station], t - cIn.time)
    
    return
}


func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
    
    allTimes := this.allStations[endStation][startStation]
    
    totalTime := 0
    for i := 0; i < len(allTimes); i++ {
        totalTime += allTimes[i]
    }
    
    return float64(totalTime)/ float64(len(allTimes))
}
