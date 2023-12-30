import {DetailedHTMLProps, HTMLAttributes, ReactNode} from "react"
import {createPortal} from "react-dom"
import styles from "./modal.module.scss"
import {CloseButton, Title} from "@/components/ui"

type Props = {
    onClose?: () => void
    title?: ReactNode | Element | string
    opened?: boolean
    overlayProps?: {
        opacity: number
        blur: number
    }
    size?: string
} & DetailedHTMLProps<HTMLAttributes<HTMLDivElement>, HTMLDivElement>


export function Modal({onClose, title, opened, children, className, size = "30rem", style, ...props}: Props) {
    if (!opened) return null

    let classes = styles.modal
    if (className) {
        classes += " " + className
    }

    style = style ?? {}
    if (size) style.minWidth = size

    return createPortal(<div className={styles.modalBackdrop}>
        <div className={classes} {...props} style={style}>
            <div className={styles.modalHeader}>
                <Title s={3}>{title}</Title>
                <CloseButton onClick={onClose}/>
            </div>
            <div className={styles.modalBody}>
                {children}
            </div>
        </div>

    </div>, document.body)
}