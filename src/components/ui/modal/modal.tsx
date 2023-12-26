import {DetailedHTMLProps, HTMLAttributes} from "react"
import {createPortal} from "react-dom"
import styles from "./modal.module.scss"
import {CloseButton, Title} from "@/components/ui"

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


export function Modal({show, onClose, title, opened, children, ...props}: Props) {

    // if (!show) return null

    return createPortal(<div className={styles.modalBackdrop}>
        <div className={styles.modal}>
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