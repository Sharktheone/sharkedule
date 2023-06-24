import styles from "./styles.module.scss"
import {CSSProperties, useEffect, useRef, useState} from "react"
import {useColors} from "./styles"
import {Entries} from "@kanban/contextmenu/entries"

type Props = {
    x: number
    y: number
    open: boolean
    close: () => void
    // type: types
}



export default function ContextMenu({x, y, open, close}: Props) {
    const {classes, cx} = useColors()
    const contextMenuRef = useRef<HTMLDivElement>(null)
    const [visible, setVisible] = useState(open)
    const [slide, setSlide] = useState(open)


    function position(): CSSProperties {
        return {
            top: `${y}px`,
            left: `${x}px`,
        }
    }

    useEffect(() => {
        if (open) {
            setVisible(true)
            setSlide(true)
            setTimeout(() => setSlide(false), 500)
        }
        else {
            setTimeout(() => setVisible(false), 490)
        }
    }, [open])


    if (!visible) return null


    return (
        <div ref={contextMenuRef}
             className={`${styles.contextmenu} ${cx(classes.contextMenu)} ${slide ? styles.open : ""} ${!open && visible ? styles.close : ""}`}
             style={position()}>
            <div>
            <button onClick={close} style={{height: "2rem", width: "3rem"}}/>
            {
                Entries.map((entry, i) => (
                    <button className={styles.entry} key={i} onClick={() => entry.handler}>
                        <div className={styles.icon}>
                            {entry.icon}
                        </div>
                        <div className={styles.name}>
                            {entry.name}
                        </div>
                    </button>

                ))
            }
            </div>
        </div>
    )
}