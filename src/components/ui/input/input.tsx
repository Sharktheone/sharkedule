import {HTMLAttributes} from "react"
import {Tooltip} from "@/components/ui/tooltip/tooltip"

type Props = {
    radius?: true | "sm" | "md" | "lg" | "xl" | "full"
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

    if (error) {
        classes += " error"
    }

    if (required) {
        classes += " required"
    }


    return (
        <>
            <input className={classes} {...props}/>
            {
                tooltip &&
                <Tooltip text={tooltip}/>
            }
            {
                error &&
                <span className="error-text">{error} - TODO
                    {/* This text needs to be probably over the input field. */}
                </span>

            }
        </>
    )
}