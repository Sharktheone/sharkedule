import Color from "@/types/color/color"


type ColorShades = {
    colors: Color[]
}

export const num = 12
export const variants = 3


export default function getColors(): ColorShades[] {
    const startHue = 25
    const s = 100
    const l = 50
    const lMin = 10

    let shades = [] as ColorShades[]

    for (let h = startHue; h < 360 + startHue; h += (360 / num)) {
        let colors = [] as Color[]
        for (let v = variants; v > 0; v--) {
            let color = new Color(
                h,
                s,
                l - v * lMin)
            colors.push(color)
        }
        for (let v = 1; v < variants; v++) {
            let color = new Color(
                h,
                s,
                l + v * lMin)

            colors.push(color)
        }

        shades.push({
            colors: colors
        })
    }
    return shades
}