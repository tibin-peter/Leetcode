func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
    landTime:=landStartTime[0]+landDuration[0]
    waterTime:=waterStartTime[0]+waterDuration[0]

    for i:=0;i<len(landStartTime);i++{
        landDur:=landStartTime[i]+landDuration[i]
        if landDur<landTime{
            landTime=landDur
        }
    }

    for i:=0;i<len(waterStartTime);i++{
        waterDur:=waterStartTime[i]+waterDuration[i]
        if waterDur<waterTime{
            waterTime=waterDur
        }
    }
    bestLandFirstTime:=10000000
    for i:=0;i<len(waterStartTime);i++{
        waterStart:=landTime
        if waterStartTime[i]>waterStart{
            waterStart=waterStartTime[i]
        }
        currentFinish:=waterStart+waterDuration[i]
        if currentFinish<bestLandFirstTime{
            bestLandFirstTime=currentFinish
        }
    }
    bestWaterFirstTime:=10000000
    for i:=0;i<len(landStartTime);i++{
        landStart:=waterTime
        if landStartTime[i]>landStart{
            landStart=landStartTime[i]
        }
        currentFinish:=landStart+landDuration[i]
        if currentFinish<bestWaterFirstTime{
            bestWaterFirstTime=currentFinish
        }
    }
     if bestLandFirstTime < bestWaterFirstTime {
        return bestLandFirstTime
    }
    return bestWaterFirstTime
}