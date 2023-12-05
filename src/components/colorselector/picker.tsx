import styles from "@/components/colorselector/styles.module.scss"
import {Button, ColorPicker} from "@/components/ui"
import Color from "@/types/color/color"
import React, {CSSProperties, useEffect, useRef, useState} from "react"


export type picker = {
    open: boolean,
    element: HTMLElement | null
    index?: number
}

type props = {
    data: picker,
    setData: (data: picker) => void
    select: (color: Color) => void
    finish: () => void
}

export default function Picker({data, select, setData, finish}: props) {
    const [pickerValue, setPickerValue] = useState<Color>(new Color(0, 0, 0))
    const ref = useRef<HTMLDivElement>(null)
    const [pickerStyles, setPickerStyles] = useState<CSSProperties>(computePickerStyles())

    useEffect(() => {
        if (!data.element) return
        data.element.style.backgroundColor = pickerValue.css()

        if (data.index === -1337) setData({open: false, element: null, index: -2})
    }, [pickerValue])

    useEffect(() => {
        setPickerStyles(computePickerStyles())
    }, [data])

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
        } as CSSProperties
    }

    function cancel() {
        let col = new Color(0, 0, 0, true)
        setPickerValue(col)
        console.log("cancel")
        data.element?.classList.remove(styles.picked)
        setData({open: false, element: data.element, index: -1337})
    }

    function handleSelect() {
        //TODO: We need to pass custom colors from the user / board context
        select(pickerValue)
        finish()
    }

    if (!data.open) return null

    return (
        <div ref={ref} className={styles.pickerOverlay}
             style={pickerStyles}
        >
            <ColorPicker format="hsl" onChange={pickerChange}/>
            <div className={styles.pickerButtons}>
                <Button onClick={cancel}>Cancel</Button>
                <Button onClick={handleSelect}>Select</Button>
            </div>
        </div>
    )
}