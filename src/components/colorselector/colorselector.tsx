import styles from "./styles.module.scss"
import {Button} from "@mantine/core"
import {useState} from "react"


type ColorShades = {
    colors: Color[]
}

type Color = {
    // mainly hsl is used, but we can add rgb if we need
    h: number
    s: number
    l: number
    //
    // r: number
    // g: number
    // b: number
}

export function ColorSelector() {
    // TODO: This is ride now just a note to me...
    // Maybe add some configurable colors that are predefined and also allow user defined colors.
    // For the custom also add places to store them and also allow "single-use" colors, but maybe list them somewhere.
    // Also for the custom colors first when you define them only let them change the hsl h-value, and add a extend button for the whole spectrum

    // (use viewTransition for this one, but maybe let the option, so we can use this as a "popup" variant - user-defined?

    const [selectedColor, setSelectedColor] = useState<Color>()

    function getColors(): ColorShades[] {
        const num = 12
        const startHue = 25
        const s = 100
        const l = 50
        const variants = 3
        const lMin = 10

        let shades = [] as ColorShades[]

        for (let h = startHue; h < 360 + startHue; h += (360 / num)) {
            let colors = [] as Color[]
            for (let v = variants; v > 0; v--) {
                let color: Color = {
                    h: h % 360,
                    s: s,
                    l: l - v * lMin,
                }
                colors.push(color)
            }
            for (let v = 1; v < variants; v++) {
                let color: Color = {
                    h: h % 360,
                    s: s,
                    l: l + v * lMin,
                }
                colors.push(color)
            }

            shades.push({
                colors: colors
            })
        }
        return shades
    }

    function get(color: Color) {
        return `hsl(${color.h}deg, ${color.s}%, ${color.l}%)`
    }

    function isSame(color1: Color, color2: Color) {
        if (color1.h != color2.h) return false
        if (color1.s != color2.s) return false
        return color1.l == color2.l;

    }

    function states(color: Color) {
        function selected(color: Color) {
            if (!selectedColor) return ""

            console.log("selected??", selectedColor, color, isSame(color, selectedColor))

            return isSame(color, selectedColor) ? styles.selected : ""
        }

        function focus(color: Color) { // do we really need this?
            // return styles.focus
            return ""
        }
        return `${selected(color)} ${focus(color)}`
    }

    function select(color: Color) {
        console.log(selectedColor, "selected now", color)
        setSelectedColor(color)
    }

    return (
        <div className={styles.selector}>
            {getColors().map(shade => (
                <div className={styles.shade}>
                    {
                        shade.colors.map(color => (
                            <button style={{
                                backgroundColor: get(color)
                            }}
                                    onClick={() => select(color)}
                                 className={`${styles.color} ${states(color)}`}
                            />
                        ))
                    }
                </div>
            ))
            }
        </div>
    )
}