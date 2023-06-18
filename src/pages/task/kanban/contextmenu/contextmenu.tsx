
import styles from "./styles.module.scss"
import {CSSProperties} from "react"
import {useColors} from "./styles"

type Props = {
    x: number
    y: number
    open: boolean
    type: types
}

export enum types {
    TASK,
    COLUMN
}

export default function ContextMenu({x, y, open}: Props)  {
    const {classes, cx} = useColors()
    function position():CSSProperties {
        return {
            top: `${y}px`,
            left: `${x}px`,
        }
    }

    if (!open) return null


    return  (
        <div className={`${styles.contextmenu} ${cx(classes.contextMenu)}`} style={position()}>
            ContextMenu
        </div>
    )
}