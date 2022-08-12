interface FarmablePlant {
    plant_type: string,
    module: number,
    position: number,
    group: string,
    leaf_harvest? : boolean,
    planter?: string;   

}

export default FarmablePlant