import styles from "./styles.module.scss"
import {useEffect, useRef, useState} from "react"
import Color from "@/types/color/color"
import {useColors} from "./colors"
import {SegmentedControl} from "@mantine/core"
import {IconColorPicker} from "@tabler/icons-react"
import control from "./control.module.scss"
import ViewTransition from "@/components/viewTransition/viewTransition"

type ColorShades = {
    colors: Color[]
}

const num = 12
const variants = 3

export function ColorSelector() {
    // TODO: This is ride now just a note to me...
    // Maybe add some configurable colors that are predefined and also allow user defined colors.
    // For the custom also add places to store them and also allow "single-use" colors, but maybe list them somewhere.
    // Also for the custom colors first when you define them only let them change the hsl h-value, and add a extend button for the whole spectrum
    // Selector for switching between normal and custom colors
    // Add button to hide 1/2 of the color shades

    // (use viewTransition for this one, but maybe let the option, so we can use this as a "popup" variant - user-defined?

    const [selectedColor, setSelectedColor] = useState<Color>()
    const [tab, setTab] = useState("simple")
    const {classes, cx} = useColors()
    const controlRef = useRef<HTMLDivElement>(null)


    useEffect(() => {
        let hsl = selectedColor?.hsl()

        if (!hsl) return

        let color = new Color(hsl.h + 30, hsl.s, hsl.l)

        console.log(controlRef)

        controlRef?.current?.style.setProperty("--gradient-color-1", selectedColor?.css() ?? "unset")

        controlRef?.current?.style.setProperty("--gradient-color-2", color?.css() ?? "unset")

    }, [selectedColor])

    function getColors(): ColorShades[] {
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

    function states(color: Color) {
        if (!selectedColor) return ""

        return color.isSame(selectedColor) ? styles.selected : ""
    }

    function select(color: Color) {
        setSelectedColor(color)
    }

    function customColors() {
        const n = (num - 2) * (variants * 2 - 1)

        let colors = [] as Color[]

        for (let i = 0; i < n; i++) {
            colors.push(new Color(0, 0, 0))
        }

        return colors
    }


    return (
        <div data-view="default" className={`${styles.selector} ${cx(classes.selector)}`}>
            <SegmentedControl ref={controlRef} data={[
                {label: 'Simple', value: 'simple'},
                {label: 'Custom', value: 'custom'},
            ]} onChange={setTab} value={tab} classNames={control}/>

            <div className={styles.content}>
                <ViewTransition view={tab}>
                    <div data-id="custom" className={`${styles.custom} ${styles.tab}`}>
                        <div className={styles.customColors}>
                            {customColors().map(color => (
                                <button
                                    onClick={() => select(color)}
                                    className={`${styles.color} ${states(color)} ${cx(classes.color)}`}/>
                            ))}
                        </div>
                        <button className={`${styles.single} ${cx(classes.single)}`}>
                            <IconColorPicker/>
                        </button>
                    </div>
                    <div data-id="simple" className={`${styles.colors} ${styles.tab}`}>
                        {getColors().map(shade => (
                            <div className={styles.shade}>
                                {shade.colors.map(color => (
                                    <button style={{
                                        backgroundColor: color.css()
                                    }}
                                            onClick={() => select(color)}
                                            className={`${styles.color} ${states(color)} ${cx(classes.color)}`}/>
                                ))}
                            </div>
                        ))}
                    </div>
                </ViewTransition>
            </div>
        </div>
    )
}