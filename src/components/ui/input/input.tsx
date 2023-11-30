import {HTMLAttributes} from "react"

type Props = {
    radius?: true | "sm" | "md" | "lg" | "xl"
    tooltip?: string
    error?: string
    required?: boolean
} & HTMLAttributes<HTMLInputElement>

export function Input({radius, tooltip, error, required, className, ...props}: Props) {

    let classes = className ?? ""
    switch (radius) {
        case "sm":
            classes += " rounded-sm"
            break
        case "md":
            classes += " rounded-md"
            break
        case "lg":
            classes += " rounded-lg"
            break
        case "xl":
            classes += " rounded-xl"
            break
        case true:
            classes += " rounded"
            break
        default:
            classes += ""
            break
    }

    if (tooltip) {
        classes += " tooltip"
    }

    if (error) {
        classes += " error"
    }

    if (required) {
        classes += " required"
    }


    return (
        <>
        <input {...props}/>
            {
                tooltip &&
                <span className="tooltip-text">{tooltip} - TODO</span>
            }
            {
                error &&
                <span className="error-text">{error} - TODO</span>
            }
        </>
    )
}