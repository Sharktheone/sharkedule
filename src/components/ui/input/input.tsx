import {HTMLAttributes, LegacyRef} from "react"
import {Tooltip} from "@/components/ui/tooltip/tooltip"
import styles from "./input.module.scss"


type Props = {
    radius?: false | "sm" | "md" | "lg" | "xl" | "full" | "none"
    tooltip?: string
    error?: string
    required?: boolean
    ref?: LegacyRef<HTMLInputElement>
    name?: string
    type?: "text" | "password" | "email" | "number" | "tel" | "url" | "search" | "date" | "time"
    value?: string
    placeholder?: string
} & HTMLAttributes<HTMLInputElement>

export function Input({radius, tooltip, error, required, className, ...props}: Props) {

    let classes = className ?? ""
    switch (radius) {
        case "sm":
            classes += " " + styles.roundedSm
            break
        case "md":
            classes += " " + styles.roundedMd
            break
        case "lg":
            classes += " " + styles.roundedLg
            break
        case "xl":
            classes += " " + styles.roundedXl
            break
        case false:
            classes += " " + styles.roundedNone
            break
        case "full":
            classes += " " + styles.roundedFull
            break
        case "none":
            classes += " " + styles.roundedNone
            break
        default:
            classes += ""
            break
    }

    if (error) {
        classes += " " + styles.error
    }

    if (required) {
        classes += " " + styles.required
    }


    //TODO: add some fields etc
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