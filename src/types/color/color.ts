export default class Color {
    private color: hsl
    private undefined: boolean

    constructor(h: number, s: number, l: number, undefined = false) {
        h %= 360
        this.color = {h, s, l}
        this.undefined = undefined
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

    fromHEX(hex: string, undefined = false) {
        //TODO
        this.undefined = undefined
    }

    fromRGB(r: number, g: number, b: number, undefined = false) {
        //TODO
        this.undefined = undefined
    }

    fromHSL(h: number, s: number, l: number, undefined = false) {
        this.color = {h, s, l}
        this.undefined = undefined
    }

    isSame(other: Color) {
        if (this.undefined || other.undefined) return false
        if (this.color.h != other.color.h) return false
        if (this.color.s != other.color.s) return false
        return this.color.l == other.color.l
    }

    css() {
        return `hsl(${this.color.h}deg, ${this.color.s}%, ${this.color.l}%)`
    }

    isUndefined() {
        return this.undefined
    }


    parseHSL(hsl: string) {
        //@ts-ignore TODO: it works, but typescript doesn't like it
        hsl.replace(/hsl\((\d+),\s*(\d+)%,\s*(\d+)%\)/g, (match, h, s, l) => {
            this.color = {h: +h, s: +s, l: +l}
            this.undefined = false
        })
    }


}


export type hsl = {
    h: number
    s: number
    l: number
}