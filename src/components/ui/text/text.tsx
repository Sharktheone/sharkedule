import {HTMLAttributes} from "react"
import styles from "./text.module.scss"
import {ColorModifier} from "@/types/color/color"

type props = {
    a?: "left" | "center" | "right"
    s?: "small" | "medium" | "large" | 1 | 2 | 3 | 4 | 5 | 6
    w?: "light" | "regular" | "bold"
    c?: ColorModifier
    dimmed?: boolean
    gradient?: boolean
    italic?: boolean
    underline?: boolean
} & HTMLAttributes<HTMLParagraphElement>


export function Text({a, s, w, c, children, italic, underline, className, gradient, dimmed, ...props}: props) {

    let classes = className ?? ""
    classes += " " + styles.text

    classes += " " + styles[a ?? "center"]

    switch (s) {
        case "small":
            classes += " " + styles.small
            break
        case "large":
            classes += " " + styles.large
            break
        case 1:
            classes += " " + styles.h1
            break
        case 2:
            classes += " " + styles.h2
            break
        case 3:
            classes += " " + styles.h3
            break
        case 4:
            classes += " " + styles.h4
            break
        case 5:
            classes += " " + styles.h5
            break
        case 6:
            classes += " " + styles.h6
            break
        default:
            classes += " " + styles.small
            break
    }

    classes += " " + styles[w ?? "regular"]

    if (italic) classes += " " + styles.italic

    if (underline) classes += " " + styles.underline

    if (dimmed) classes += " " + styles.dimmed

    if (gradient) classes += " " + styles.gradient

    return (
        <p className={classes} {...props} data-color={c ?? "primary"}> {children} </p>
    )
}