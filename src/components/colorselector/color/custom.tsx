import React, {useRef} from "react"
import useDoubleClick from "@/hooks/useDoubleClick/useDoubleClick"
import styles from "@/components/colorselector/styles.module.scss"
import Color from "@/types/color/color"
import {picker} from "@/components/colorselector/picker"


type props = {
    color: Color,
    pickColor: (element: HTMLElement | null, open?: boolean) => void,
    select: (color: Color) => void
    picker: picker
    selectedColor?: Color
    colorContext: (e: React.MouseEvent<HTMLButtonElement>) => void
    states: (color: Color) => string

}
export default function CustomColor({color, pickColor, select, picker, selectedColor, colorContext, states}: props) {
    const r = useRef<HTMLButtonElement>(null)

    const {
        onClick,
        onDoubleClick
    } = useDoubleClick(() => select(color), () => {
        pickColor(r.current)
    }, 100)

    function clickHandler(e: React.MouseEvent<HTMLButtonElement>) {
        e.stopPropagation()
        if (picker.open && e.target !== picker.element) {
            pickColor(r.current, true)
            return
        }
        if (color.isUndefined() && r.current !== null) pickColor(r.current)
        onClick()
    }

    function doubleClickHandler() {
        if (color.isUndefined()) return
        onDoubleClick()
    }

    return (
        <button
            ref={r}
            onClick={(e) => clickHandler(e)}
            onDoubleClick={doubleClickHandler}
            onContextMenu={colorContext}
            className={`${styles.color} ${states(color)}`}/>
    )
}
