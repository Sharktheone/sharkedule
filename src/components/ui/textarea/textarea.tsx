import TextareaAutosize from "react-textarea-autosize"
import {HTMLAttributes, RefObject} from "react"
import styles from "./styles.module.scss"


type Props = {
    autosize?: boolean
    label?: string
    error?: string | boolean
    autosizeMinRows?: number
    autosizeMaxRows?: number
    placeholder?: string
    resize?: "none" | "both" | "horizontal" | "vertical"
    textareaRef?: RefObject<HTMLTextAreaElement>
} & Omit<HTMLAttributes<HTMLTextAreaElement>, "style">


export function Textarea({label, error, className, resize = "both", ...props}: Props) {
    let classes = styles.textarea
    if (className) {
        classes += " " + className
    }

    switch (resize) {
        case "none":
            classes += " " + styles.noResize
            break
        case "both":
            classes += " " + styles.bothResize
            break
        case "horizontal":
            classes += " " + styles.horizontalResize
            break
        case "vertical":
            classes += " " + styles.verticalResize
            break
    }

    return (
        <TA {...props} className={classes}/>
    )
}


type TAProps = {
    autosize?: boolean
    autosizeMinRows?: number
    autosizeMaxRows?: number
    textareaRef?: RefObject<HTMLTextAreaElement>
} & Omit<HTMLAttributes<HTMLTextAreaElement>, "style">

function TA({autosize, autosizeMaxRows, autosizeMinRows, textareaRef, ...props}: TAProps) {
    if (autosize) {
        return <TextareaAutosize minRows={autosizeMaxRows} maxRows={autosizeMaxRows} {...props} ref={textareaRef}/>
    }
    return <textarea {...props}/>

}