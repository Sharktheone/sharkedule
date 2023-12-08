import {DetailedHTMLProps, TextareaHTMLAttributes} from "react"


type Props = {
    autosize?: boolean
    label?: string
    error?: string
    required?: boolean

} & DetailedHTMLProps<TextareaHTMLAttributes<HTMLTextAreaElement>, HTMLTextAreaElement>


export function Textarea({autosize, ...props}: Props) {
    //TODO: Textarea with autosize
    return (
        <textarea {...props}/>
    )
}