import React, {useRef} from "react"
import {useDoubleClick} from "@/hooks"
import styles from "@/components/colorselector/styles.module.scss"
import Color from "@/types/color/color"
import {picker} from "@/components/colorselector/picker"


type props = {
    color: Color,
    pickColor: (element: HTMLElement | null, index: number, open?: boolean) => void,
    select: (color: Color) => void
    picker: picker
    colorContext: (e: React.MouseEvent<HTMLButtonElement>, index: number) => void
    states: (color: Color) => string
    index: number,

}
export default function CustomColor({color, pickColor, select, picker, colorContext, states, index}: props) {
    const r = useRef<HTMLButtonElement>(null)

    const {
        onClick,
        onDoubleClick
    } = useDoubleClick(() => select(color), () => {
        pickColor(r.current, index)
    }, 100)

    function clickHandler(e: React.MouseEvent<HTMLButtonElement>) {
        e.stopPropagation()
        if (picker.open && e.target !== picker.element) {
            pickColor(r.current, index, true)
            return
        }
        if (color.isUndefined() && r.current !== null) pickColor(r.current, index)
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
            onContextMenu={e => colorContext(e, index)}
            className={`${styles.color} ${states(color)}`}/>
    )
}
