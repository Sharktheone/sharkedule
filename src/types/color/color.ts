


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

    isSame(other: Color) {
        if (this.color.h != other.color.h) return false
        if (this.color.s != other.color.s) return false
        return this.color.l == other.color.l;
    }


}


type hsl = {
    h: number
    s: number
    l: number
}