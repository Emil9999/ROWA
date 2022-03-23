

export function checkImage(url: string) : string{

    const availableImages: Array<string> = ["basil", "lollo bionda", "mint", "pirat"]
    if(availableImages.includes(url) ){
        return url
    }

    const checkedString = 'mint'
    return checkedString
}