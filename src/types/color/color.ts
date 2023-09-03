


export class Color {
    private color: hsl

    constructor(h: number, s: number, l: number) {
        this.color = {h, s, l}
    }

    rgb() {
        //TODO
    }

    hex() {
        //TODO
    }

    hsl() {
        return this.color
    }

    fromHEX(hex: string) {
        //TODO
    }

    fromRGB(r: number, g: number, b: number) {
        //TODO
    }

    fromHSL(h: number, s: number, l: number) {
        this.color = {h, s, l}
    }


}


type hsl = {
    h: number
    s: number
    l: number
}