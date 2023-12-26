import {HTMLAttributes, LegacyRef} from "react"
import {ColorModifier} from "@/types/color/color"


type Pros = {
    variant?: ColorModifier
    gradient?: boolean
    ref?: LegacyRef<HTMLInputElement>
    label?: string
    required?: boolean
    placeholder?: string

} & HTMLAttributes<HTMLInputElement>

export function TextInput({variant, gradient, ref, placeholder, ...props}: Pros) {

    //TODO

    return (
        <input ref={ref} placeholder={placeholder} {...props} data-bg-color={variant ?? "primary"}/>
    )
}