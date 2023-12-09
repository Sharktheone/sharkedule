import {HTMLAttributes} from "react"

type Props = {
    opened: boolean
    onClose: () => void
    position: "top" | "bottom" | "left" | "right"
    about?: string
    title?: string
    overlayProps?: {
        opacity: number
        blur: number
    }

} & HTMLAttributes<HTMLDivElement>


export function Drawer({...props}: Props) {

    //TODO: render in portal
    return (
        <div/>
    )
}