import styles from "./styles.module.scss"
import React, {useEffect, useRef, useState} from "react"
import Color from "@/types/color/color"
import {useColors} from "./colors"
import {SegmentedControl} from "@mantine/core"
import {IconColorPicker} from "@tabler/icons-react"
import control from "./control.module.scss"
import ViewTransition from "@/components/viewTransition/viewTransition"
import CustomColor from "@/components/colorselector/color/custom"
import getColors from "@/components/colorselector/colorgen"
import Picker, {picker} from "@/components/colorselector/picker"
import SimpleColor from "@/components/colorselector/color/simple"

export function ColorSelector() {
    //TODO:  Also for the custom colors first when you define them only let them change the hsl h-value, and add a extend button for the whole spectrum
    const [selectedColor, setSelectedColor] = useState<Color>()
    const [tab, setTab] = useState("simple")
    const singleRef = useRef<HTMLButtonElement>(null)
    const controlRef = useRef<HTMLDivElement>(null)
    const {classes, cx} = useColors()

    const [picker, setPicker] = useState<picker>({} as picker)


    useEffect(() => {
        let hsl = selectedColor?.hsl()

        if (!hsl) return

        let color = new Color(hsl.h + 30, hsl.s, hsl.l)

        controlRef?.current?.style.setProperty("--gradient-color-1", selectedColor?.css() ?? "unset")

        controlRef?.current?.style.setProperty("--gradient-color-2", color?.css() ?? "unset")

    }, [selectedColor])

    useEffect(() => {
        picker.element?.classList.remove(styles.picked)
        setPicker({open: false, element: null})
    }, [tab])


    function states(color: Color) {
        if (!selectedColor) return ""

        return color.isSame(selectedColor) ? styles.selected : ""
    }


    function select(color: Color) {
        if (color.isUndefined()) return
        if (picker.open) return
        setSelectedColor(color)
    }


    function customColors() { //TODO: Get this from user context / board context
        const n = (12 - 2) * (3 * 2 - 1)

        let colors = [] as Color[]

        for (let i = 0; i < n; i++) {
            colors.push(new Color(0, 0, 0, true))
        }

        return colors
    }

    function colorContext(e: React.MouseEvent<HTMLButtonElement>) {
        e.preventDefault()
        e.stopPropagation()
        const element = e.target as HTMLButtonElement
        if (element.classList.contains(styles.picked)) {
            pickColor(element, false)
            return
        }
        pickColor(element, true)
    }

    function pickColor(element: HTMLElement | null, open = !picker.open) {
        picker.element?.classList.remove(styles.picked)
        if (open) element?.classList.add(styles.picked)
        setPicker({open: open, element: element})
    }


    return (
        <div data-view="default" className={`${styles.selector} ${cx(classes.selector)}`}>
            <SegmentedControl ref={controlRef} data={[
                {label: "Simple", value: "simple"},
                {label: "Custom", value: "custom"},
            ]} onChange={setTab} value={tab} classNames={control}/>
            <div className={styles.content}>
                <ViewTransition view={tab}>
                    <div data-id="simple" className={`${styles.custom} ${styles.tab}`}>
                        {getColors().map(shade => (
                            <div className={styles.shade}>
                                {shade.colors.map(color => (
                                    <SimpleColor color={color} select={select} states={states}
                                                 colorContext={colorContext}/>
                                ))}
                            </div>
                        ))}
                    </div>
                    <div data-id="custom" className={`${styles.custom} ${styles.tab}`}>
                        <div className={styles.customColors}>
                            {customColors().map(color => (
                                <CustomColor color={color} pickColor={pickColor} select={select} picker={picker}
                                             selectedColor={selectedColor} colorContext={colorContext}
                                             states={states}
                                />
                            ))}
                        </div>
                        <button ref={singleRef} className={`${styles.single} ${cx(classes.single)}`}
                                onClick={() => pickColor(singleRef.current)}>
                            <IconColorPicker/>
                        </button>
                    </div>
                </ViewTransition>
                <Picker data={picker} select={select}/>
            </div>
        </div>
    )
}