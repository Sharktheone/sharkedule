import styles from "@/components/colorselector/styles.module.scss"
import {Button, ColorPicker} from "@mantine/core"
import Color from "@/types/color/color"
import React, {useEffect, useRef, useState} from "react"
import {Simulate} from "react-dom/test-utils"
import select = Simulate.select

export type picker = {
    open: boolean,
    element: HTMLElement | null
}

type props = {
    data: picker,
    select: (color: Color) => void
}

export default function Picker({data, select}: props) {
    const [pickerValue, setPickerValue] = useState<Color>(new Color(0, 0, 0))
    const ref = useRef<HTMLDivElement>(null)

    useEffect(() => {
        if (!data.element) return
        data.element.style.backgroundColor = pickerValue.css()
    }, [pickerValue])

    function pickerChange(string: string) {
        let col = new Color(0, 0, 0)
        col.parseHSL(string)

        setPickerValue(col)
    }


    function computePickerStyles() {
        let x = data.element?.offsetLeft ?? 0
        const y = data.element?.offsetTop ?? 0
        const w = data.element?.offsetParent?.clientWidth ?? 0

        const pickerWidth = ref?.current?.clientWidth ?? 0

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

    if (!data.open) return null

    return (
        <div ref={ref} className={styles.pickerOverlay}
             style={computePickerStyles()}
        >
            <ColorPicker format="hsl" onChange={pickerChange}/>
            <div className={styles.pickerButtons}>
                <Button onClick={() => pickColor(singleRef.current)}>Cancel</Button>
                <Button onClick={() => select(new Color(0, 0, 0))}>Select</Button>
            </div>
        </div>
    )
}