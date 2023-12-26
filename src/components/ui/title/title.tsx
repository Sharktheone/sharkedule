import {HTMLAttributes} from "react"
import styles from "@/components/ui/text/text.module.scss"
import {ColorModifier} from "@/types/color/color"


type props = {
    a?: "left" | "center" | "right"
    s?: 1 | 2 | 3 | 4 | 5 | 6
    w?: "light" | "regular" | "bold"
    c?: ColorModifier
    dimmed?: boolean
    gradient?: boolean
    italic?: boolean
    underline?: boolean
} & HTMLAttributes<HTMLHeadingElement>

export function Title({a, s, w, c, italic, underline, className, gradient, dimmed, ...props}: props) {

    let classes = className ?? ""

    classes += " " + styles.title

    classes += " " + styles[a ?? "center"]

    classes += " " + styles[w ?? "bold"]

    if (italic) classes += " " + styles.italic

    if (underline) classes += " " + styles.underline

    if (dimmed) classes += " " + styles.dimmed

    if (gradient) classes += " " + styles.gradient


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
        <Heading s={hX(s ?? 1)} className={classes} {...props} c={c ?? "primary"}/>
    )
}


type headingProps = {
    s: "h1" | "h2" | "h3" | "h4" | "h5" | "h6"
    c: ColorModifier
} & HTMLAttributes<HTMLHeadingElement>

function Heading({s, c, ...props}: headingProps) {
    let H = s
    return (
        <H {...props} data-color={c}/>
    )
}