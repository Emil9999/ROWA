import FarmablePlant from "./FarmablePlant"

interface ModuleFarmable{
        varietyName: string,
        plantables?: FarmablePlant,
        harvestables?: FarmablePlant
}

export default ModuleFarmable