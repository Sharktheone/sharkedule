import {HTMLAttributes, LegacyRef} from "react"


type Pros = {
    variant?: "primary" | "secondary" | "danger" | "success" | "warning" | "info" | "light" | "dark"
    gradient?: boolean
    ref?: LegacyRef<HTMLInputElement>
    label?: string
    required?: boolean
    placeholder?: string

} & HTMLAttributes<HTMLInputElement>

export function TextInput({variant, gradient, ref, placeholder, ...props}: Pros) {

    //TODO

    return (
        <input ref={ref} placeholder={placeholder} {...props}/>
    )
}