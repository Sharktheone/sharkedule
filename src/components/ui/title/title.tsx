import {HTMLAttributes} from "react"
import styles from "@/components/ui/text/text.module.scss"


type props = {
    a?: "left" | "center" | "right"
    s?: 1 | 2 | 3 | 4 | 5 | 6
    w?: "light" | "regular" | "bold"
    c?: "primary" | "secondary" | "danger" | "success" | "warning" | "info" | "light" | "dark" | "error"
    dimmed?: boolean
    gradient?: boolean
    italic?: boolean
    underline?: boolean
} & HTMLAttributes<HTMLHeadingElement>

export function Title({a, s, w, c, children, italic, underline, className, gradient, dimmed, ...props}: props) {

    let classes = className ?? ""


    switch (c) { //TODO: This can probably be done better
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


    type headings = "h1" | "h2" | "h3" | "h4" | "h5" | "h6"

    function hX(size: number): headings {
        switch (size) {
            case 1:
                return "h1"
            case 2:
                return "h2"
            case 3:
                return "h3"
            case 4:
                return "h4"
            case 5:
                return "h5"
            case 6:
                return "h6"
            default:
                return "h1"
        }
    }

    return (
        <Heading s={hX(s ?? 1)} className={classes} {...props}/>
    )
}


type headingProps = {
    s: "h1" | "h2" | "h3" | "h4" | "h5" | "h6"
} & HTMLAttributes<HTMLHeadingElement>

function Heading({s, ...props}: headingProps) {
    let H = s
    return (
        <H {...props}/>
    )
}