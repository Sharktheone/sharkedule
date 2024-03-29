import React, {useRef} from "react"
import {colorDisabled} from "@/components/colorselector/color/functions"
import styles from "@/components/colorselector/styles.module.scss"
import Color from "@/types/color/color"


type props = {
    color: Color,
    select: (color: Color) => void
    states: (color: Color) => string
}

export default function SimpleColor({color, select, states}: props) {
    const r = useRef<HTMLButtonElement>(null)

    return (
        <button ref={r} style={{
            backgroundColor: color.css()
        }}
                onClick={() => select(color)}
                className={`${styles.color} ${states(color)} ${colorDisabled(color)}`}
        /> //TODO: Add shadow
    )
}