import {HTMLAttributes, LegacyRef} from "react"
import styles from "./text.module.scss"

type props = {
    a?: "left" | "center" | "right"
    s?: "small" | "medium" | "large" | 1 | 2 | 3 | 4 | 5 | 6
    w?: "light" | "regular" | "bold"
    c?: "primary" | "secondary" | "danger" | "success" | "warning" | "info" | "light" | "dark" | "error"
    dimmed?: boolean
    gradient?: boolean
    italic?: boolean
    underline?: boolean
    ref?: LegacyRef<HTMLParagraphElement>
} & HTMLAttributes<HTMLParagraphElement>


export function Text({a, s, w, c, children, italic, underline, className, gradient, dimmed, ref, ...props}: props) {

    let classes = className ?? ""
    classes += " " + styles.text

    switch (c) {
        case "primary":
            classes += " " + styles.primary
            break
        case "secondary":
            classes += " " + styles.secondary
            break
        case "danger":
            classes += " " + styles.danger
            break
        case "success":
            classes += " " + styles.success
            break
        case "warning":
            classes += " " + styles.warning
            break
        case "info":
            classes += " " + styles.info
            break
        case "light":
            classes += " " + styles.light
            break
        case "dark":
            classes += " " + styles.dark
            break
        default:
            classes += ""
            break
    }

    switch (a) {
        case "left":
            classes += " " + styles.left
            break
        case "right":
            classes += " " + styles.right
            break
        default:
            classes += " " + styles.center
            break
    }
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
            classes += " " + styles.medium
            break
    }

    switch (w) {
        case "bold":
            classes += " " + styles.bold
            break
        case "light":
            classes += " " + styles.light
            break
        default:
            classes += " " + styles.regular
            break
    }

    if (italic) {
        classes += " " + styles.italic
    }

    if (underline) {
        classes += " " + styles.underline
    }

    if (dimmed) {
        classes += " " + styles.dimmed
    }

    if (gradient) {
        classes += " " + styles.gradient
    }

    return (
        <p className={classes} {...props} ref={ref}> {children} </p>
    )
}