import {DetailedHTMLProps, HTMLAttributes} from "react"
import {createPortal} from "react-dom"

type Props = {
    show?: boolean
    onClose?: () => void
    title?: string
    opened?: boolean
    overlayProps?: {
        opacity: number
        blur: number
    }
} & DetailedHTMLProps<HTMLAttributes<HTMLDivElement>, HTMLDivElement>


export function Modal({show, onClose, title, opened, ...props}: Props) {

    //TODO: Render this in a portal
    return createPortal(<>

    </>, document.body)
}