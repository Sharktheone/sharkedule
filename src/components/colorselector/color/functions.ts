import React from "react"
import styles from "@/components/colorselector/styles.module.scss"
import Color from "@/types/color/color"

export function colorDisabled(color: Color) {
    return color.isUndefined() ? styles.colorDisabled : ""
}
