func asteroidsDestroyed(mass int, asteroids []int) bool {
    sort.Slice(asteroids,func(i,j int)bool{
        return asteroids[i]<asteroids[j]
    })
    for _,v:=range asteroids{
        if v>mass{
            return false
        }
        mass+=v
    }
    return true
}