

export function checkImage(checkType: string){
    
    const availableImagespng: Array<string> = ["basil", "lollo bionda", "mint", "pirat"]
    const availableImagessvg: Array<string> = ["basil", "lollo bionda", "mint", "pirat"]
    const cImage = (Givenurl: string) =>{
        const  url = Givenurl.toLowerCase()
        if (checkType == "png"){
            if(availableImagespng.includes(url) ){
                return url + '.png'
            }  else { return 'default.png'}
        } else if(checkType == "svg"){
            if(availableImagessvg.includes(url) ){
               
                return url + '.svg'
            } else { return 'default.svg'}
        }

   
    return 'default'

}
return {cImage}
}