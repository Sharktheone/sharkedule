import styles from "./styles.module.scss"
import React, {useEffect, useRef, useState} from "react"
import Color from "@/types/color/color"
import {useColors} from "./colors"
import {Button, ColorPicker, SegmentedControl} from "@mantine/core"
import {IconColorPicker} from "@tabler/icons-react"
import control from "./control.module.scss"
import ViewTransition from "@/components/viewTransition/viewTransition"
import useDoubleClick from "@/hooks/useDoubleClick/useDoubleClick"
import {useClickOutside} from "@mantine/hooks"

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
    const [picker, setPicker] = useState<{open: boolean, element: HTMLButtonElement | null}>({open: false, element: null})
    const {classes, cx} = useColors()
    const controlRef = useRef<HTMLDivElement>(null)
    // const clickRef = useClickOutside<HTMLDivElement>(() => {
    //     let p = {...picker}
    //     setTimeout(() => {
    //         if (picker.x !== p.x || picker.y !== p.y) return
    //         setPicker({open: false, x: 0, y: 0})
    //     }, 100)
    // })
    const ref = useRef<HTMLDivElement>(null)
    const singleRef = useRef<HTMLButtonElement>(null)

    // clickRef.current = ref.current

    // ref?.current?.style.setProperty("--_right", "red")


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
        if (color.isUndefined()) return
        if (picker) return
        setSelectedColor(color)
    }

    function customColors() {
        const n = (num - 2) * (variants * 2 - 1)

        let colors = [] as Color[]

        for (let i = 0; i < n; i++) {
            colors.push(new Color(0, 0, 0, true))
        }

        return colors
    }

    function pickColor(element: HTMLElement, open = !picker.open) {
        picker.element?.classList.remove(styles.picked)
        if (open) element.classList.add(styles.picked)
        setPicker({open: open, element: element})
    }

    function colorDisabled(color: Color) {
        return color.isUndefined() ? styles.colorDisabled : ""
    }

    function colorContext(e: MouseEvent) {
        e.preventDefault()
        e.stopPropagation()
        const element = e.target as HTMLButtonElement
        if (element.classList.contains(styles.picked)) {
            pickColor(element, false)
            return
        }
        pickColor(element, true)
    }

    function computePickerStyles() {
        let x = picker.element?.offsetLeft ?? 0
        const y = picker.element?.offsetTop ?? 0
        const w = picker.element?.offsetParent?.clientWidth ?? 0

        const pickerWidth = ref?.current?.clientWidth

        let indicator = pickerWidth
        let bleft = "var(--_border-indicator)"
        let bright = "var(--_border-indicator)"

        if (w / 2 > x) {
            indicator = 0
            bright = "transparent"
        } else {
            x -= (pickerWidth + 56)
            bleft = "transparent"
        }

        return {
            left: "5.75rem",
            "--_left": x,
            "--_bright": bright,
            "--_bleft": bleft,
            "--_indicator": indicator,
            "--_top": y,
        } as React.CSSProperties
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
                                {shade.colors.map(color => {

                                    const r = useRef<HTMLButtonElement>(null)

                                    const {
                                        onClick,
                                        onDoubleClick
                                    } = useDoubleClick(() => select(color), () => pickColor(r.current), 100)


                                    return (
                                        <button ref={r} style={{
                                            backgroundColor: color.css()
                                        }}
                                                onClick={onClick}
                                                onDoubleClick={onDoubleClick}
                                                onContextMenu={colorContext}
                                                className={`${styles.color} ${states(color)} ${cx(classes.color)} ${colorDisabled(color)}`}/>
                                    )
                                })}
                            </div>
                        ))}
                    </div>
                    <div data-id="custom" className={`${styles.custom} ${styles.tab}`}>
                        <div className={styles.customColors}>
                            {customColors().map(color => {

                                const r = useRef<HTMLButtonElement>(null)

                                const {
                                    onClick,
                                    onDoubleClick
                                } = useDoubleClick(() => select(color), () => {pickColor(r.current)}, 100)

                                function clickHandler(e: MouseEvent<HTMLButtonElement>) {
                                    e.stopPropagation()
                                    if (picker.open && e.target !== picker.element) {
                                        pickColor(r.current, true)
                                        return
                                    }
                                    if (color.isUndefined()) pickColor(r.current)
                                    onClick()
                                }

                                function doubleClickHandler() {
                                    if (color.isUndefined()) return
                                    onDoubleClick()
                                }


                                return (
                                    <button
                                        ref={r}
                                        onClick={clickHandler}
                                        onDoubleClick={doubleClickHandler}
                                        onContextMenu={colorContext}
                                        className={`${styles.color} ${states(color)} ${cx(classes.color)}`}/>
                                )
                            })}
                        </div>
                        <button ref={singleRef} className={`${styles.single} ${cx(classes.single)}`}
                                onClick={() => pickColor(singleRef.current)}>
                            <IconColorPicker/>
                        </button>
                    </div>
                </ViewTransition>
                {
                    picker.open ? <div ref={ref} className={styles.pickerOverlay}
                                       style={computePickerStyles()}
                    >
                        <ColorPicker/>
                        <div className={styles.pickerButtons}>
                            <Button onClick={() => pickColor(singleRef.current)}>Cancel</Button>
                            <Button onClick={() => select(new Color(0, 0, 0))}>Select</Button>
                        </div>
                        {/* Hmm, I need to move this depending on the button that is pressed, I have an idea, test it later  */}
                    </div> : null
                }
            </div>
        </div>
    )
}