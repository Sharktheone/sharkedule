import {HTMLAttributes, ReactNode} from "react"
import styles from "./drawer.module.scss"
import {createPortal} from "react-dom"
import {CloseButton, Title} from "@/components/ui"

type Props = {
    opened: boolean
    onClose: () => void
    position: "top" | "bottom" | "left" | "right"
    about?: string
    title?: ReactNode | Element | string
    overlayProps?: {
        opacity: number
        blur: number
    }
    size?: string

} & HTMLAttributes<HTMLDivElement>


export function Drawer({
                           onClose, title, opened, children, className, size = "20rem", style,
                           position = "right", overlayProps, ...props
                       }: Props) {
    if (!opened) return null

    let classes = styles.drawer
    if (className) {
        classes += " " + className
    }

    style = style ?? {}
    switch (position) {
        case "top":
            style.height = size
            break
        case "bottom":
            style.height = size
            break
        case "left":
            style.width = size
            break
        case "right":
            style.width = size
            break
    }

    return createPortal(<div className={styles.drawerBackdrop}>
        <div className={classes} {...props} style={style} data-position={position}>
            <div className={styles.drawerHeader}>
                <Title s={3}>{title}</Title>
                <CloseButton onClick={onClose}/>
            </div>
            <div className={styles.drawerBody}>
                {children}
            </div>
        </div>

    </div>, document.body)
}