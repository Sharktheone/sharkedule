import {HTMLAttributes, ReactNode} from "react"
import styles from "./drawer.module.scss"
import {createPortal} from "react-dom"
import {CloseButton, Title} from "@/components/ui"
import {useClickOutside} from "@/hooks"

type Props = {
    opened: boolean
    onClose: () => void
    onClosing?: () => void
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
                           onClosing, position = "right", overlayProps, ...props
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


    let ref = useClickOutside(() => {
        close()
    })

    function close() {
        setTimeout(() => {
            console.log("closing")
            if (onClose) onClose()
        }, 500)

        ref.current?.parentElement?.classList.add(styles.closing)

        ref.current?.classList.add(styles.closing)
    }

    return createPortal(<div className={styles.drawerBackdrop}>
        <div className={classes} {...props} style={style} data-position={position} ref={ref}>
            <div className={styles.drawerHeader}>
                <Title s={3}>{title}</Title>
                <CloseButton onClick={close}/>
            </div>
            <div className={styles.drawerBody}>
                {children}
            </div>
        </div>

    </div>, document.body)
}