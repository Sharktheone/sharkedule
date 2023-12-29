import TextareaAutosize from "react-textarea-autosize"
import {HTMLAttributes} from "react"
import styles from "./styles.module.scss"


type Props = {
    autosize?: boolean
    label?: string
    error?: string | boolean
    autosizeMinRows?: number
    autosizeMaxRows?: number
    placeholder?: string
} & Omit<HTMLAttributes<HTMLTextAreaElement>, "style">


export function Textarea({label, error, className, ...props}: Props) {
    let classes = styles.textarea + " " + className
    return (
        <TA {...props} className={classes}/>
    )
}


type TAProps = {
    autosize?: boolean
    autosizeMinRows?: number
    autosizeMaxRows?: number
} & Omit<HTMLAttributes<HTMLTextAreaElement>, "style">

function TA({autosize, autosizeMaxRows, autosizeMinRows, ...props}: TAProps) {
    if (autosize) {
        return <TextareaAutosize minRows={autosizeMaxRows} maxRows={autosizeMaxRows} {...props}/>
    }
    return <textarea {...props}/>

}